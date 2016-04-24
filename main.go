package main

import (
  "os"
  "log"
  "time"
  "github.com/tucnak/telebot"
)

var bot *telebot.Bot

func main() {

  token := os.Getenv("BOT_API_TOKEN")
  if len(token) > 0 {
    log.Printf("Telegram Bot Token: %v\n", token)
  } else {
    log.Fatal("Please set 'BOT_API_TOKEN' from environment variable")
  }

  if newBot, err := telebot.NewBot(token); err != nil {
    log.Fatal(err);
  } else {
    bot = newBot
  }

  bot.Messages = make(chan telebot.Message, 1000)
  bot.Queries = make(chan telebot.Query, 1000)

  go messages()
  go queries()

  bot.Start(1 * time.Second)
}

func messages() {
  for message := range bot.Messages {
    switch {
    case message.Text == "/start":
      bot.SendMessage(message.Chat, "Hello, " + message.Sender.FirstName + "!", nil)
    }
  }
}

func queries() {
  for query := range bot.Queries {
    log.Println("--- new query ---")
    log.Println("from:", query.From)
    log.Println("text:", query.Text)

    // There you build a slice of let's say, article results:
    results := []telebot.Result{/*...*/}

    // And finally respond to the query:
    if err := bot.Respond(query, results); err != nil {
      log.Println("ouch:", err)
    }
  }
}
