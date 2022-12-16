package main

import (
	"encoding/json"
	"fmt"
)

var authors = []Author{
	{Id: "author-1", Firstname: "Jake", Lastname: "Roghetti", Username: "jeti", Password: "sunshine"},
	{Id: "author-2", Firstname: "Cathy", Lastname: "mcgothen", Password: "arby's"},
}

var articles = []Article{
	{Id: "atr-1", Author: "author-1", Title: "crazyTitle", Content: "some content"},
	{Id: "atr-219", Author: "author-2", Title: "another day", Content: "there was this time that I did something crazy."},
}

func main() {
	fmt.Println("Starting application...")
	data, _ := json.Marshal(authors)

	fmt.Println(string(data))
}
