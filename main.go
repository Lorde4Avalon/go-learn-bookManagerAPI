package main

import (
	"fmt"

	"github.com/Rioa-Avalon/go-demo/models"
	"github.com/Rioa-Avalon/go-demo/routers"
)

func main() {
	models.Init()
	router := routers.RouterRegist()
	fmt.Println("Server is running at http://localhost:3303")
	router.Run(":3303")
}
