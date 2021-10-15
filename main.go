package main

import (
	"github.com/foody-go-api/context"
)


func main() {

	appContext := context.NewAppContext()
	appContext.Run()
}
