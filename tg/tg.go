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
	isBotRunning     bool
	creatorChatID    int64
	creatorChatIDStr string
	websiteUrl       string
	channelUrl       string
)

var ukrainianCommands = map[string]string{
	"Допомога":               "help",
	"Повідомити про помилку": "support",
	"Контакти":               "contacts",
	"Заява на вступ":         "application_form",
	"Питання":                "questions",
	"Ми в мережі":            "socials",
	"Стоп":                   "stop",
	"Старт":                  "start",
	"/start":                 "/start",
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
			tgbotapi.NewKeyboardButton("Заява на вступ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Контакти"),
			tgbotapi.NewKeyboardButton("Питання"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ми в мережі"),
			tgbotapi.NewKeyboardButton("Повідомити про помилку"),
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

	backKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
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
				case "start", "/start":
					if !isBotRunning {
						isBotRunning = true
						okEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
						msg.Text = okEmoji + " Вже працюю"
						msg.ReplyMarkup = generalKeyboard
					} else {
						okEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
						msg.Text = okEmoji + " Бот вже запущений\nСтоп - зупинити бота"
					}

				case "help":
					if isBotRunning {
						infoEmoji := emoji.Sprintf("%v", emoji.Information)
						msg.Text = infoEmoji + " Підказки\n\n+ Допомога - отримати всі команди\n+ Старт - запустити бота\n+ Стоп - зупинити бота\n+ Контакти - отримати контактну інформацію ключових членів клубу\n+ Заява на вступ - отримати посилання на форму (онлайн-заявку на вступ до клубу)\n+ Ми в мережі - отримати посилання на нас в мережі\n+ Повідомити про помилку - повідомити про знайдені помилки\n+ Питання - задати питання і отримати відповідь від адміністратора"
						msg.ReplyMarkup = generalKeyboard
					}

				case "contacts":
					if isBotRunning {
						msg.Text = "Президент: @kenjitheman\nВіце-президент: @ya_code"
						msg.ReplyMarkup = generalKeyboard
					}

				case "application_form":
					if isBotRunning {
						applicationFormUrl := os.Getenv("GOOGLE_FORM_URL")
						infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
						msg.Text = infinityEmoji + " " + applicationFormUrl
						msg.ReplyMarkup = generalKeyboard
					}

				case "socials":
					if isBotRunning {
						websiteUrl := os.Getenv("WEBSITE_URL")
						channelUrl := os.Getenv("CHANNEL_URL")
						msg.Text = "Вебсайт: " + websiteUrl + "\nКанал: " + channelUrl
						msg.ReplyMarkup = generalKeyboard
					}

				case "stop":
					if isBotRunning {
						isBotRunning = false
						stopEmoji := emoji.Sprintf("%v", emoji.RedCircle)
						msg.Text = stopEmoji + " Зупинився"
						msg.ReplyMarkup = startKeyboard
					} else {
						msg.Text = "Бот вже зупинений\nСтарт - запустити бота"
					}

				case "questions", "support":
					if isBotRunning {
						cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
						greenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
						creatorChatIDStr := os.Getenv("CREATOR_CHAT_ID")
						creatorChatID, err = strconv.ParseInt(creatorChatIDStr, 10, 64)
						if err != nil {
							log.Panic(err)
						}
						msg.ReplyMarkup = backKeyboard

						initialMessage := cactusEmoji
						afterMessage := greenHeartEmoji
						if command == "questions" {
							initialMessage += " Будь ласка, введіть ваше запитання:"
							afterMessage += " Ми надамо відповідь якнайшвидше!"
						} else if command == "support" {
							initialMessage += " Будь ласка, опишіть проблему:"
							afterMessage += " Дякую за звіт про помилку!"
						}
						msg.Text = initialMessage
						bot.Send(msg)

						response := <-updates

						if response.Message != nil {
							if response.Message.Chat.ID != update.Message.Chat.ID {
								continue
							}
							description := response.Message.Text
							if description == "Назад" {
								msg.ReplyMarkup = generalKeyboard
							} else {
								var supportMsg tgbotapi.MessageConfig
								if command == "questions" {
									supportMsg = tgbotapi.NewMessage(
										creatorChatID,
										fmt.Sprintf(
											"Запитання від користувача @%s:\n%s",
											update.Message.From.UserName,
											description,
										),
									)
								} else if command == "support" {
									supportMsg = tgbotapi.NewMessage(
										creatorChatID,
										fmt.Sprintf(
											"Звіт про помилку від користувача @%s:\n%s",
											update.Message.From.UserName,
											description,
										),
									)
								}
								msg.Text = afterMessage
								bot.Send(supportMsg)
								msg.ReplyMarkup = generalKeyboard
							}
						}
					} else {
						msg.Text = "Бот вже зупинений\nСтарт - запустити бота"
					}

				default:
					if isBotRunning {
						idkEmoji := emoji.Sprintf("%v", emoji.OpenHands)
						msg.Text = idkEmoji + " Вибачте, але я вас не розумію\nДопомога - отримати всі команди"
					}
				}

				if _, err := bot.Send(msg); err != nil {
					fmt.Printf("[ERROR] error sending message")
				}
			}
		}
	}
}
