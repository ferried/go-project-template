package main


import (
	"log"
	"sailor/cmd/apis/app"
)

func main() {
	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
