package main

import (
	"bufio"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "0.0.0.0:12345")
	//Only change this
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply *string
		err = client.Call("abc", "test", &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v", *reply)
	}
}
