package main

import (
	"fmt"
	"net/rpc"
	"log"
)

type Item struct {
	Title string
	Body string
}

func main() {

	var reply Item
	fmt.Println("Client")
	

	client,err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection ", err)
	}

	a := Item{"first","This Is The First Item"}

	client.Call("API.AddItem",a,&reply)
	fmt.Println(reply)
}
