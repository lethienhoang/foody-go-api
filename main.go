package main

import (
	"github.com/foody-go-api/context"
	"log"
)


func main() {

	appContext := context.NewAppContext()
	if err := appContext.RunService(); err != nil {
		log.Panicln(err)
	}
}
