package main

import (
	"character-manager/database"
	"character-manager/internal/routes"
	"fmt"
	"github.com/BrandonEchols/common-go-utils/common"
	"net/http"
)

func main() {
	cfg := common.GetConfigGetter("internal/config/config.json")

	_, err := database.NewDb(cfg)
	if err != nil {
		fmt.Printf("Error connecting to database. Err: %s\n", err.Error)
		return
	}

	server := routes.SetupRouter()

	port := "localhost:" + cfg.MustGetConfigVar("port")
	fmt.Println("Starting server on " + port)
	err = http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("Error returned from http.ListenAndServe. Err", err.Error())
	}
}
