package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sangeeth518/learning-routes/database"
	"github.com/sangeeth518/learning-routes/router"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("failed to lod env")
		return

	}
	database.ConnectTODb()
}
func main() {

	router.Routing()
}
