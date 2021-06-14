package server

import (
	"fmt"
	"net/http"

	"ex/handlers"
)

func RunServer() {
	fmt.Println("start")
	defer func() {
		fmt.Println("end")
	}()
	http.HandleFunc("/calculator", handlers.Calculator)
	http.ListenAndServe(":8090", nil)
}
