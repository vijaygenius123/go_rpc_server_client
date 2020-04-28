package main

import (
	"fmt"
)


type Item struct {
	title string
	body string
}

var database []Item

func AddItem(item Item) Item{
	database = append(database, item)
	return item
}

func main(){

	fmt.Println("Hello")

	fmt.Println("Initial Database State", database)
	a := Item{"first","This Is The First Item"}
	b := Item{"second","This Is The Second Item"}
	c := Item{"third","This Is The Third Item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)

	fmt.Println("Database State After Adding Items", database)

}