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
	var final_database []Item
	fmt.Println("Client")
	

	client,err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection ", err)
	}

	a := Item{"first","This Is The First Item"}
	b := Item{"second","This Is The Second Item"}
	c := Item{"third","This Is The Fourth Item"}
	d := Item{"fourth","This Is The Fourth Item"}

	client.Call("API.AddItem",a,&reply)
	client.Call("API.AddItem",b,&reply)
	client.Call("API.AddItem",c,&reply)
	client.Call("API.AddItem",d,&reply)

	fmt.Println("Initial DB State")
	client.Call("API.GetDB","GetDB",&database)
	fmt.Println(database)

	fmt.Println("Get Item By Title first")
	e := Item{}
	client.Call("API.GetItemByTitle","first",&e)
	fmt.Println(e)

	fmt.Println("Edit Item 3")
	f := Item{"third","This Is The Third Item"}
	client.Call("API.EditItem",Item{"third","This Is The Third Item"},&f)
	fmt.Println(f)

	fmt.Println("Final DB State")
	client.Call("API.GetDB","GetDB",&final_database)
	fmt.Println(final_database)

}
