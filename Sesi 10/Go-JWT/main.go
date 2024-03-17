package main

import (
	"Go-JWT/database"
	"Go-JWT/router"
	"fmt"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
	fmt.Println("Server Running In : LocalHost:8080")
}
