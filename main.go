package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/joho/godotenv"
	alert "github.com/kaustubhkapatral/monitoring-bot/alert"
	config "github.com/kaustubhkapatral/monitoring-bot/config"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file not found")
	}
	config.SetConfig()
}
func main() {
	fmt.Println("Started ")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			if err := alert.HexSend(); err != nil {
				fmt.Println("Unable to start Hex send", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			if err := alert.JailSend(); err != nil {
				fmt.Println("Unable to start jail send", err)
			}
			time.Sleep(7200 * time.Second)
		}
	}()
	wg.Wait()
}
