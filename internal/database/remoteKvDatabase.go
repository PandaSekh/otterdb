package database

import (
	"fmt"
	"github.com/PandaSekh/otterdb/internal/ds"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type RemoteDatabase struct {
	LocalDatabase
	host string
	port string
}

func (d *RemoteDatabase) Get(key string, c chan interface{}) {
	d.mu.Lock()
	defer d.mu.Unlock()
	v, found := d.table.Get(key)

	if !found {
		close(c)
	} else {
		c <- v
	}
}

func (d *RemoteDatabase) Set(key string, value interface{}, c chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.table.Set(key, value)

	c <- true
}

func (d *RemoteDatabase) Remove(key string, c chan bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res := d.table.Remove(key)

	c <- res
}

func (d *RemoteDatabase) Contains(key string, c chan bool) {
	channel := make(chan interface{}, 1)
	d.Get(key, channel)

	val, open := <-channel
	if !open && val == nil {
		c <- false
	}

	c <- true
}

func (d *RemoteDatabase) String() string {
	return fmt.Sprintf("%v", d.table)
}

func (d *RemoteDatabase) GetTable() ds.HashTable {
	return d.table
}

func (d *RemoteDatabase) GetHost() string {
	return d.host
}

func (d *RemoteDatabase) GetPort() string {
	return d.port
}

func (d *RemoteDatabase) startServer() {
	fmt.Printf("Server started on: %s:%s\n", d.host, d.port)
	server, err := net.Listen("tcp", d.host+":"+d.port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer func(server net.Listener) {
		_ = server.Close()
	}(server)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go d.onClientRequest(connection)
	}
}

func (d *RemoteDatabase) onClientRequest(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	req := string(buffer[:mLen])
	var res string

	inputs := strings.Split(req, " ")
	switch inputs[0] {
	case "get":
		var c chan interface{}

		d.Get(inputs[1], c)

		val, open := <-c
		if !open && val == nil {
			fmt.Println("Key not found")
		} else {
			fmt.Sprintf("%v", val)
		}
	case "set":
		var c chan bool
		d.Set(inputs[1], inputs[2], c)
		done, open := <-c
		if !open || !done {
			fmt.Println("Not done")
		} else {
			fmt.Println("Done")
		}
	case "rem":
		var c chan bool
		d.Remove(inputs[1], c)
		done, open := <-c
		if !open || !done {
			fmt.Println("Not done")
		} else {
			fmt.Println("Done")
		}
	}

	_, err = connection.Write([]byte(res))
	_ = connection.Close()
}

func NewDefaultRemoteDatabase() *RemoteDatabase {
	return NewRemoteDatabase("localhost", strconv.Itoa(1111+rand.Intn(9999-1111)))
}

func NewRemoteDatabase(host string, port string) *RemoteDatabase {
	dkv := &RemoteDatabase{
		LocalDatabase: LocalDatabase{
			table: *ds.NewSized(4000),
			mu:    &sync.Mutex{},
		},
		host: host,
		port: port,
	}

	go dkv.startServer()

	return dkv
}
