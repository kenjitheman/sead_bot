package tg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	"Допомога":       "help",
	"Підтримка":      "support",
	"Контакти":       "contacts",
	"Заява на вступ": "application_form",
	"Питання":        "questions",
	"Ми в мережі":    "socials",
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
			tgbotapi.NewKeyboardButton("Питання"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Заява на вступ"),
			tgbotapi.NewKeyboardButton("Ми в мережі"),
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

	socialsKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Канал"),
			tgbotapi.NewKeyboardButton("Вебсайт"),
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
						msg.Text = okEmoji + " Вже працюю"
						msg.ReplyMarkup = generalKeyboard
					} else {
						msg.Text = "Бот вже запущений\nСтоп - зупинити бота"
					}

				case "help":
					if isBotRunning {
						infoEmoji := emoji.Sprintf("%v", emoji.Information)
						msg.Text = infoEmoji + " Підказки\n\n Допомога - отримати всі команди\n Старт - запустити бота\n Стоп - зупинити бота\n Контакти - отримати контактну інформацію ключових членів клубу\n Заява на вступ - отримати посилання на форму (онлайн-заявку на вступ до клубу)\n Канал - отримати посилання на telegram kанал клубу\n Вебсайт - отримати посилання на вебсайт клубу\n Підтримка - повідомити про знайдені помилки\n\n Питання - задати питання і отримати відповідь від адміністратора"
						msg.ReplyMarkup = generalKeyboard
					}

				case "contacts":
					if isBotRunning {
						msg.Text = "Президент: @kenjitheman\nВіце-президент: [contact info]\nСекретар: [contact info]\nСкарбник: [contact info]"
						msg.ReplyMarkup = generalKeyboard
					}

				case "application_form":
					if isBotRunning {
						applicationFormUrl := os.Getenv("GOOGLE_FORM_URL")
						infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
						msg.Text = infinityEmoji + " " + applicationFormUrl
						msg.ReplyMarkup = generalKeyboard
					}

				case "channel":
					if isBotRunning {
						channelUrl = os.Getenv("CHANNEL_URL")
						infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)
						msg.Text = infinityEmoji + " " + channelUrl
						msg.ReplyMarkup = generalKeyboard
					}

				case "socials":
					if isBotRunning {
						msg.ReplyMarkup = socialsKeyboard
						msg.Text = "Будь ласка, виберіть варіант:"
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Printf("[ERROR] error sending message: %v\n", err)
							return
						}

						select {
						case response := <-updates:
							if response.Message != nil &&
								response.Message.Chat.ID == update.Message.Chat.ID {
								take := response.Message.Text
								infinityEmoji := emoji.Sprintf("%v", emoji.Infinity)

								switch take {
								case "Канал":
									channelURL := os.Getenv("CHANNEL_URL")
									msg.Text = infinityEmoji + " " + channelURL
								case "Вебсайт":
									websiteURL := os.Getenv("WEBSITE_URL")
									msg.Text = infinityEmoji + " " + websiteURL
								default:
									idkEmoji := emoji.Sprintf("%v", emoji.OpenHands)
									msg.Text = idkEmoji + " Вибачте, але я вас не розумію\nДопомога - отримати всі команди"
								}

								msg.ReplyMarkup = generalKeyboard
							}
						case <-time.After(30 * time.Second): // Set a timeout for user response
							msg.Text = "Відповіді не отримано.\nСпробуйте пізніше."
							msg.ReplyMarkup = generalKeyboard
						}
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

				case "questions":
					if isBotRunning {
						cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
						creatorChatIDStr := os.Getenv("CREATOR_CHAT_ID")
						creatorChatID, err = strconv.ParseInt(creatorChatIDStr, 10, 64)
						if err != nil {
							log.Panic(err)
						}
						msg.Text = cactusEmoji + " Будь ласка, введіть ваше запитання:"
						bot.Send(msg)

						response := <-updates
						if response.Message != nil {
							if response.Message.Chat.ID != update.Message.Chat.ID {
								continue
							}
							description := response.Message.Text
							GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
							msg.Text = GreenHeartEmoji + " Ми надамо відповідь якнайшвидше!"
							supportMsg := tgbotapi.NewMessage(
								creatorChatID,
								fmt.Sprintf(
									"Запитання від користувача %s:\n%s",
									update.Message.From.UserName,
									description,
								),
							)
							bot.Send(supportMsg)
						}
					} else {
						msg.Text = "Бот вже зупинений.\nСтарт - запустити бота."
					}
				case "support":
					if isBotRunning {
						cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
						creatorChatIDStr = os.Getenv("CREATOR_CHAT_ID")
						creatorChatID, err = strconv.ParseInt(creatorChatIDStr, 10, 64)
						if err != nil {
							log.Panic(err)
						}
						msg.Text = cactusEmoji + " Будь ласка, опишіть проблему:"
						bot.Send(msg)

						response := <-updates
						if response.Message != nil {
							if response.Message.Chat.ID != update.Message.Chat.ID {
								continue
							}
							description := response.Message.Text
							GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
							msg.Text = GreenHeartEmoji + " Дякую за звіт про помилку!"
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
