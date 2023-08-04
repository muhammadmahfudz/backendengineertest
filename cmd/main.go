package main

import (
	"Backend_Engineer_Interview_Assignment/common/app"
	ct "Backend_Engineer_Interview_Assignment/common/config"
	"log"
)

func main() {
	cfg, err := ct.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)

}
