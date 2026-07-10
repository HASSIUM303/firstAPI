package main

import (
	"fmt"
	"todoapi/http"
	"todoapi/todo"
)

func main() {
	var todoList = todo.NewList()
	var httpHandlers = http.NewHandlers(todoList)
	var server = http.NewServer(httpHandlers)

	if err := server.Start(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
