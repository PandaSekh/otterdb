package repl

import (
	"bufio"
	"fmt"
	"github.com/PandaSekh/otterdb/internal/cluster"
	"net"
	"os"
	"strings"
)

var clusterHost string
var clusterPort string

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)
	remoteCluster := cluster.NewDefault()
	clusterHost = remoteCluster.GetHost()
	clusterPort = remoteCluster.GetPort()
	fmt.Println("otterdb started.")
	for {
		fmt.Print("# ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.Replace(input, "\n", "", -1)
		processInput(input)
	}
}

func processInput(input string) {
	channel := make(chan string, 1)
	go callServer(input, channel)
	result := <-channel
	fmt.Printf("Result: %v\n", result)
}

func callServer(input string, channel chan string) {
	//establish connection
	connection, err := net.Dial("tcp", clusterHost+":"+clusterPort)
	if err != nil {
		panic(err)
	}
	///send some data
	_, err = connection.Write([]byte(input))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	channel <- string(buffer[:mLen])
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)
}
