package main

import (
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/services"
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
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
