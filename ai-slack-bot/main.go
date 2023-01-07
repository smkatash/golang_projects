package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"

)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}

func main() {
	err := godotenv.Load(".env") // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	b_token := os.Getenv("SLACK_BOT_TOKEN")
	app_token := os.Getenv("SLACK_APP_TOKEN")
	bot := slacker.NewClient(b_token, app_token)
	go printCommandEvents(bot.CommandEvents())

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfram",
		Examples: []string{"who is the president of Germany?"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			query := request.Param("message")
			fmt.Println(query)
			//client.Parse
			response.Reply("received")
		},
	})

	cnt, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(cnt)
	if err != nil {
		log.Fatal(err)
	}
}