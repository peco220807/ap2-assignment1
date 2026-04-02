package main
import (
	"log"
	"appointment-service/internal/app"
)

func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}