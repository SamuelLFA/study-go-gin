package main

import (
	"github.com/samuellfa/study-go-gin/database"
	"github.com/samuellfa/study-go-gin/routes"
)

func main() {
	database.ConectWithDatabase()
	routes.HandleRequest()
}
