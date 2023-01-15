package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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

	result := buildResult(lowercased)

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

func buildResult(inputQuery string) string {
	result := ""

	for i, letter := range inputQuery {
		repeatTimes := 1

		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			repeatTimes = random()
		}

		for j := 1; j <= repeatTimes; j++ {
			if i%2 != 0 {
				result += strings.ToUpper(string(letter))
			} else {
				result += string(letter)
			}
		}
	}

	return result
}

func random() int {
	rand.Seed(time.Now().UnixNano())

	switch randomNumber := rand.Intn(100) + 1; {
	case randomNumber > 40:
		return 1
	case randomNumber > 30:
		return 2
	case randomNumber > 20:
		return 3
	case randomNumber > 10:
		return 4
	default:
		return 5
	}
}
