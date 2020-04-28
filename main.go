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

func GetItemByTitle(title string)Item {
	var getItem Item

	for _, val := range database {
		if val.title == title{
			 getItem = val
			 break
		}
	}

	return getItem
}

func EditItem(title string, edited Item)Item {
	var editedItem Item

	for idx, val := range database {
		if val.title == title{
			database[idx] = edited
			editedItem = edited
		}
	}
	return editedItem
}

func DeleteItem(item Item) Item{

	var deleted Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			deleted = item
			break
		}
	}
	
	return deleted
}


func main(){

	fmt.Println("Hello")

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

}