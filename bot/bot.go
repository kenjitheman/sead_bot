package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning := false

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	lastUserMessageTime := time.Now()

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			userInput := update.Message.Text

			if autoOff != nil {
				autoOff.Stop()
			}

			if time.Since(lastUserMessageTime) > 5*time.Minute {
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}

			lastUserMessageTime = time.Now()

			autoOff := time.NewTimer(5 * time.Minute)
			go func() {
				<-autoOff.C
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}()

			switch userInput {
			case "/start", "start", "Старт", "/Старт":
				if !isBotRunning {
					isBotRunning = true
					msg.Text = startMsg
					msg.ReplyMarkup = GeneralKeyboard
				} else {
					msg.Text = alreadyStartedMsg
				}

			case "/help", "help", "Допомога", "/Допомога":
				if isBotRunning {
					msg.Text = helpMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "/contacts", "contacts", "Контакти", "/Контакти":
				if isBotRunning {
					msg.Text = contactsMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "/form", "form", "Заява на вступ", "/Заява на вступ":
				if isBotRunning {
					msg.Text = formMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "/socials", "socials", "Ми в мережі", "/Ми в мережі", "/we_on_the_web", "we_on_the_web", "We on the web", "/We on the web":
				if isBotRunning {
					msg.Text = weOnTheWebMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "/stop", "stop", "Стоп", "/Стоп":
				if isBotRunning {
					isBotRunning = false
					msg.Text = stopMsg
					msg.ReplyMarkup = StartKeyboard
				} else {
					msg.Text = areadyStoppedMsg
				}

			case "/ask", "/bug_report", "ask", "bug_report", "Питання", "Повідомити про помилку", "/Питання", "/Повідомити про помилку":
				if isBotRunning {
					msg.ReplyMarkup = BackKeyboard
					switch userInput {
					case "/ask", "ask", "Питання", "/Питання":
						initialMessage = askForQuestionMsg
						afterMessage = thxForQuestionMsg

					case "/bug_report", "bug_report", "Повідомити про помилку", "/Повідомити про помилку":
						initialMessage = askForBugReportMsg
						afterMessage = thxForBugReportMsg

					default:
						msg.Text = idkMsg
					}
					msg.Text = initialMessage
					bot.Send(msg)

					response := <-updates

					if response.Message != nil {
						if response.Message.Chat.ID != update.Message.Chat.ID {
							continue
						}
						description := response.Message.Text

						if description == "Назад" || description == "/Назад" || description == "back" || description == "/back" {
							msg.Text = backToMenuMsg
							msg.ReplyMarkup = GeneralKeyboard
						} else {
							var supportMsg tgbotapi.MessageConfig
							switch userInput {
							case "/ask", "ask", "Питання", "/Питання":
								supportMsg = tgbotapi.NewMessage(
									creatorChatID,
									fmt.Sprintf(
										askQuestionMsg,
										update.Message.From.UserName,
										description,
									),
								)
							case "/bug_report", "bug_report", "Повідомити про помилку", "/Повідомити про помилку":
								supportMsg = tgbotapi.NewMessage(
									creatorChatID,
									fmt.Sprintf(
										reportMsg,
										update.Message.From.UserName,
										description,
									),
								)
							}

							msg.Text = afterMessage
							bot.Send(supportMsg)
							msg.ReplyMarkup = GeneralKeyboard
						}

					}
				} else {
					msg.Text = alreadyStartedMsg
				}
			default:
				if isBotRunning {
					msg.Text = idkMsg
				}
			}
			lastUserMessageTime = time.Now()
			if _, err := bot.Send(msg); err != nil {
				fmt.Println("[ERROR] error sending message")
			}
		}
	}

	if autoOff != nil {
		autoOff.Stop()
	}
}
