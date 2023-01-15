package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery != nil {
			handleInline(bot, update.InlineQuery)
		}
	}
}

func handleInline(bot *tgbotapi.BotAPI, inlineQuery *tgbotapi.InlineQuery) {
	lowercased := strings.ToLower(inlineQuery.Query)

	result := ""

	for i, letter := range lowercased {
		if i%2 != 0 {
			result += strings.ToUpper(string(letter))
		} else {
			result += string(letter)
		}
	}

	article := tgbotapi.NewInlineQueryResultArticle(inlineQuery.ID, "✨ Chonizado ✨", result)
	article.Description = result

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{article},
	}

	if _, err := bot.Request(inlineConf); err != nil {
		log.Println(err)
	}
}
