package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-microservice/internal/controller"
	"go-microservice/internal/di"
	"go-microservice/internal/shared"
	"log"
	"time"
)

func main() {
	err := di.Container.Invoke(func(deps shared.Deps, ch controller.Holder) error {
		var (
			app = echo.New()
		)

		ch.SetupRoutes(app)

		// Get the current time and timezone
		currentTime := time.Now()
		localTimeZone := currentTime.Location()
		fmt.Printf("\nLocal Timezone: %s at %s\n", localTimeZone, currentTime.String())

		// start app
		log.Fatal(app.Start(":8080"))

		return nil
	})

	if err != nil {
		panic(err)
	}
}
