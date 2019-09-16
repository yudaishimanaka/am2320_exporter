package main

import (
	"log"

	"github.com/oltoko/go-am2320"
)

func main() {
	am2320 := am2320.Create(am2320.DefaultI2CAddr)

	res, err := am2320.Read()
	if err != nil {
		log.Fatalln("Failed to read from AM2320", err)
	}

	log.Printf("%.2f", res.Temperature)
	log.Printf("%.2f", res.Humidity)
}
