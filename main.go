package main

import (
	"github.com/spaceNottis123/api-go-gin/database"
	"github.com/spaceNottis123/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
