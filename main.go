package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/http"

)


type Item struct {
	Title string
	Body string
}
type API int

var database []Item


func (a *API) GetDB(title string,reply *[]Item) error {
	
	*reply = database
	return nil
}

func (a *API) AddItem(item Item,reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) GetItemByTitle(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title{
			 getItem = val
			 break
		}
	}

	*reply = getItem
	return nil
}

func (a *API) EditItem(edited Item, reply *Item) error {
	var editedItem Item

	for idx, val := range database {
		if val.Title == edited.Title{
			database[idx] = edited
			editedItem = edited
		}
	}

	*reply = editedItem
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {

	var deleted Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			deleted = item
			break
		}
	}
	*reply = deleted
	return nil
}


func main(){

	fmt.Println("Hello")

	var api = new(API)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error registering API",err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp",":4040")

	if err != nil {
		log.Fatal("Listener Error",err)
	}
	log.Printf("Serving RPC On Port %d",4040)

	err = http.Serve(listener,nil)
	if err != nil {
		log.Fatal("Server Error",err)
	}

	/*
	fmt.Println("Initial Database State", database)
	a := Item{"first","This Is The First Item"}
	b := Item{"second","This Is The Second Item"}
	c := Item{"third","This Is The Fourth Item"}
	d := Item{"fourth","This Is The Fourth Item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	AddItem(d)

	fmt.Println("Database State After Adding Items", database)

	e := GetItemByTitle("second")

	fmt.Println("Get Item : ", e)

	EditItem("third",Item{"third","This Is The Third Item"})

	fmt.Println("Database Before Deleting Item", database)
	DeleteItem(d)
	fmt.Println("Database After Deleting Item", database)
*/
}