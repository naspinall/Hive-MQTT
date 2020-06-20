package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func subscribeHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(msg.Payload())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:8080").SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetClientID("TestClientID")
	channel := make(chan int)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("hello", 0, subscribeHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := c.Publish("hello", 0, false, "First Message"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := c.Publish("hello", 0, false, "Second Message"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	<-channel
}
