package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker/v2"
)

func main() {
	apptoken, err := getEnv("SLACK_APP_TOKEN")

	if err != nil {
		log.Fatal(err.Error())
	}

	botToken, err := getEnv("SLACK_BOT_TOKEN")

	if err != nil {
		log.Fatal(err.Error())
	}

	api := slacker.NewClient(botToken, apptoken)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	api.AddCommand(&slacker.CommandDefinition{
		Command:  "my yop is <year>",
		Examples: []string{"Echoes the word"},
		Handler: func(cc *slacker.CommandContext) {
			year := cc.Request().IntegerParam("year", 1990)

			cc.Response().Reply(fmt.Sprintf("You are %d years old", (time.Now().Year() - year)))

		},
	})

	api.Listen(ctx)
}

func getEnv(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	val := os.Getenv(key)

	return val, nil
}
