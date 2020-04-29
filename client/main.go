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
	var database []Item
	fmt.Println("Client")
	

	client,err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection ", err)
	}

	a := Item{"first","This Is The First Item"}
	b := Item{"second","This Is The Second Item"}

	client.Call("API.AddItem",a,&reply)
	fmt.Println(reply)
	client.Call("API.AddItem",b,&reply)

	client.Call("API.GetDB","GetDB",&database)
	fmt.Println(database)

	c := Item{}
	client.Call("API.GetItemByTitle","first",&c)
	fmt.Println(c)
}
