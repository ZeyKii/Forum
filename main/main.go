package main

import (
	"fmt"
	"forum/views"
)

func main() {
	fmt.Println("Server web lauch on http://localhost:8082/")
	forum.StartServer()
}
