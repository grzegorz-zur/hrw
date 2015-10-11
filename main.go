package main

import (
	"log"
	"time"
)

func main() {
	config, err := ReadConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		time.Sleep(time.Duration(config.Interval) * time.Second)
		if ok := Check(config.Ping); !ok {
			log.Println("router reset")
			err := Reset(config.Router)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
