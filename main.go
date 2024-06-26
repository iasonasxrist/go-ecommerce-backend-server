package main

import (
	"ecommerce.com/config"
	"ecommerce.com/internal/api"
	"fmt"	
	"log"
)

func main() {
	fmt.Println("Hello, Go!")

	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded properly %v \n", err)
	}

	api.StartServer(cfg)

}
