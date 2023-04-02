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
	db   LocalDatabase
	host string
	port string
}

func (d *RemoteDatabase) Get(key string, c chan interface{}) {
	d.db.mu.Lock()
	defer d.db.mu.Unlock()
	v, found := d.db.table.Get(key)

	if !found {
		close(c)
	} else {
		c <- v
	}
}

func (d *RemoteDatabase) Set(key string, value interface{}, c chan bool) {
	d.db.mu.Lock()
	defer d.db.mu.Unlock()

	d.db.table.Set(key, value)
	c <- true
}

func (d *RemoteDatabase) Remove(key string, c chan bool) {
	d.db.mu.Lock()
	defer d.db.mu.Unlock()
	res := d.db.table.Remove(key)

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
	return fmt.Sprintf("%v", d.db.table)
}

func (d *RemoteDatabase) GetTable() ds.HashTable {
	return d.db.table
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
	res := []byte("ok")

	inputs := strings.Split(req, " ")

	switch inputs[0] {
	case "get":
		c := make(chan interface{}, 1)

		go d.Get(inputs[1], c)

		val, open := <-c
		if open && val != nil {
			res = []byte(val.(string))
		} else {
			res = []byte("Not Found")
		}
	case "set":
		c := make(chan bool, 1)
		go d.Set(inputs[1], inputs[2], c)

		done, open := <-c
		if !open || !done {
			res = []byte("Not ok")
		}
	case "rem":
		c := make(chan bool, 1)
		go d.Remove(inputs[1], c)
		done, open := <-c
		if !open || !done {
			res = []byte("Not ok")
		}
	}

	_, err = connection.Write(res)
	_ = connection.Close()
}

func NewDefaultRemoteDatabase() *RemoteDatabase {
	return NewRemoteDatabase("localhost", strconv.Itoa(1111+rand.Intn(9999-1111)))
}

func NewRemoteDatabase(host string, port string) *RemoteDatabase {
	dkv := &RemoteDatabase{
		db: LocalDatabase{
			table: *ds.NewSized(4000),
			mu:    &sync.Mutex{},
		},
		host: host,
		port: port,
	}

	go dkv.startServer()

	return dkv
}
