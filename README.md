# go_rpc_server_client


## Create A Basic CRUD

Create a simple in memory db
```go
type Item struct{
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

```