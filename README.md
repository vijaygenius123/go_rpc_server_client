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

```