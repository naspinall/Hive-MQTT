package main

import (
	"github.com/naspinall/Hive-MQTT/pkg/server"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hive"
	dbname   = "hive"
)

func main() {
	mqtt := server.NewMQTTBroker()
	mqtt.Listen("localhost", "8080")
}
