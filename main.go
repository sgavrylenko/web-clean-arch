package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sgavrylenko/web-clean-arch/route"
	"log"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()
	log.Fatal(route.RunAPI(":8080"))
}
