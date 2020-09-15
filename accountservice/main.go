package main

import (
	"fmt"
	"github.com/ed201971/godaddy/accountservice/services"
	"github.com/ed201971/godaddy/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	services.StartWebServer("6767")
}

func initializeBoltClient() {
	services.DBClient = &dbclient.BoltClient{}
	services.DBClient.OpenBoltDb()
	services.DBClient.Seed()
}
