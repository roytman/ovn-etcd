package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Listener int
type Reply struct  {
	Data string
}

var cli *clientv3.Client

func (l *Listener) Abc(line string, reply *string) error {
	fmt.Printf("Receive: %v\n", line)
	*reply = "yay"
	return nil
}

//func (l *Listener) Efg(line string, reply *string) error {
//	fmt.Printf("Receive: %v\n", line)
//	timeout := 5 * time.Second
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	_, err := cli.Put(ctx, "sample_key", "sample_value")
//	//resp, err := cli.Put(ctx, "sample_key", "sample_value")
//	cancel()
//	if err != nil {
//	    // handle error!
//	}
//	// use the response
//	*reply = "yaya efg"
//	return nil
//}
//this server is used as the client for the etcd server.(wrapper)
func main() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cli.Close()

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.RegisterName("", listener)
	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
