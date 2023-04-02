package cluster

import (
	"fmt"
	"github.com/PandaSekh/otterdb/internal/database"
	"github.com/PandaSekh/otterdb/internal/fnvHash"
	"github.com/PandaSekh/otterdb/internal/utils"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	defaultSize = 3
)

type Cluster struct {
	nodes []Node
	size  int
	host  string
	port  string
}

// NewDefault generates a Cluster with the default buckets size
func NewDefault() *Cluster {
	c := &Cluster{
		size:  defaultSize,
		nodes: make([]Node, defaultSize),
		host:  "localhost",
		port:  strconv.Itoa(1111 + rand.Intn(9999-1111)),
	}

	for i := range c.nodes {
		c.nodes[i].instance = database.NewDefaultRemoteDatabase()
	}

	go c.startClusterServer()

	return c
}

func (c *Cluster) startClusterServer() {
	fmt.Printf("otterdb cluster server started on: %s:%s\n", c.host, c.port)
	server, err := net.Listen("tcp", c.host+":"+c.port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	utils.PrintLocalIp()
	defer func(server net.Listener) {
		_ = server.Close()
	}(server)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go c.processClusterClientRequest(connection)
	}
}

func (c *Cluster) processClusterClientRequest(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	request := string(buffer[:mLen])

	// get correct key
	key := strings.Split(request, " ")[1] // todo not hardcoded -> use encoder/decoder

	// get server to send message to
	nodeIndex := c.getNodeIndex(key)
	node := c.nodes[nodeIndex]

	// send request to correct server
	response := sendRequestToNode(node.instance.GetHost(), node.instance.GetPort(), request)

	// return response
	_, err = connection.Write([]byte(response))
	_ = connection.Close()
}

func sendRequestToNode(host string, port string, request string) string {
	//establish connection
	connection, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}

	// Send data
	_, err = connection.Write([]byte(request))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)

	return string(buffer[:mLen])
}

func (c *Cluster) getNodeIndex(key string) int {
	return int(fnvHash.NewDefault().Hash(key) % uint64(c.size))
}

func (c *Cluster) GetPort() string {
	return c.port
}

func (c *Cluster) GetHost() string {
	return c.host
}

type Node struct {
	instance *database.RemoteDatabase
}

func (c *Cluster) String() string {
	return fmt.Sprintf("Nodes: %v, Size: %d", c.nodes, c.size)
}

func (n *Node) String() string {
	return fmt.Sprintf("Database: %v", n.instance)
}
