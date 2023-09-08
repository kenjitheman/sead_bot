package tg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	isBotRunning  bool
	creatorChatID int64
)

var ukrainianCommands = map[string]string{
	"Допомога":       "help",
	"Підтримка":      "support",
	"Контакти":       "contacts",
	"Заява на вступ": "application_form",
	"Стоп":           "stop",
	"Старт":          "start",
}

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file")
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning = false

	generalKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Допомога"),
			tgbotapi.NewKeyboardButton("Підтримка"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Контакти"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Заява на вступ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Стоп"),
		),
	)

	startKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Старт"),
		),
	)

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			userInput := update.Message.Text

			if command, ok := ukrainianCommands[userInput]; ok {
				switch command {
				case "start":
					if !isBotRunning {
						isBotRunning = true
						okEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
						msg.Text = okEmoji + " вже працюю"
						msg.ReplyMarkup = generalKeyboard
					} else {
						msg.Text = "бот вже запущений\nСтоп - зупинити бота"
					}

				case "help":
					if isBotRunning {
						infoEmoji := emoji.Sprintf("%v", emoji.Information)
						msg.Text = infoEmoji + " Підказки\n\n Допомога - щоб отримати всі команди\n Старт - запустити бота\n Стоп - зупинити бота\n Контакти - отримати контактну інформацію ключових членів клубу\n Заява на вступ - отримати посилання на форму (онлайн-заявку на вступ до клубу)\n Підтримка - щоб повідомити про знайдені помилки"
						msg.ReplyMarkup = generalKeyboard
					}

				case "contacts":
					if isBotRunning {
						infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
						contactInfo := `
              Президент: @kenjitheman
              Віце-президент: [contact info]
              Секретар: [contact info]
              Скарбник: [contact info]
                    `
						msg.Text = infinityEmoji + " " + contactInfo
						msg.ReplyMarkup = generalKeyboard
					}

				case "application_form":
					if isBotRunning {
						applicationFormUrl := os.Getenv("GOOGLE_FORM_URL")
						infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
						msg.Text = infinityEmoji + " " + applicationFormUrl
						msg.ReplyMarkup = generalKeyboard
					}

				case "stop":
					if isBotRunning {
						isBotRunning = false
						stopEmoji := emoji.Sprintf("%v", emoji.RedCircle)
						msg.Text = stopEmoji + " зупинився"
						msg.ReplyMarkup = startKeyboard
					} else {
						msg.Text = "бот вже зупинений\nСтарт - запустити бота"
					}

				case "support":
					if isBotRunning {
						cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
						creatorChatIDStr := os.Getenv("CREATOR_CHAT_ID")
						creatorChatID, err = strconv.ParseInt(creatorChatIDStr, 10, 64)
						if err != nil {
							log.Panic(err)
						}
						msg.Text = cactusEmoji + " будь ласка, опишіть проблему:"
						bot.Send(msg)

						response := <-updates
						if response.Message != nil {
							if response.Message.Chat.ID != update.Message.Chat.ID {
								continue
							}
							description := response.Message.Text
							GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
							msg.Text = GreenHeartEmoji + " дякую за звіт про помилку!"
							supportMsg := tgbotapi.NewMessage(
								creatorChatID,
								fmt.Sprintf(
									" bug report from user %s:\n%s",
									update.Message.From.UserName,
									description,
								),
							)
							bot.Send(supportMsg)
						}
					} else {
						msg.Text = "бот вже зупинений\nСтарт - запустити бота"
					}

				default:
					if isBotRunning {
						idkEmoji := emoji.Sprintf("%v", emoji.OpenHands)
						msg.Text = idkEmoji + " вибачте, але я вас не розумію\n/help"
					}
				}

				if _, err := bot.Send(msg); err != nil {
					fmt.Printf("[ERROR] error sending message")
				}
			}
		}
	}
}
