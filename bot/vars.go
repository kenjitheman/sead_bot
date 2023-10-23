package bot

import (
	"github.com/enescakir/emoji"
	"strconv"
	"time"
)

const (
	creatorChatIDStr string = "5785150199"
	websiteUrl       string = "https://seadclub.online"
	googleFormUrl    string = "http://join.seadclub.online"
)

var (
	creatorChatID, err  = strconv.ParseInt(creatorChatIDStr, 10, 64)
	isBotRunning        bool
	initialMessage      string
	afterMessage        string
	lastUserMessageTime time.Time
	autoOff             *time.Timer
)

var (
	okEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	infoEmoji       = emoji.Sprintf("%v", emoji.Information)
	infinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	stopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	cactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	greenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
	idkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
)

var (
	idkMsg             = idkEmoji + " Вибачте, але я вас не розумію\n/help - отримати всі команди"
	startMsg           = okEmoji + " Привіт, я бот клубу Інженерії Програмного Забезпечення та Цифрових Технологій\n/help - отримати всі команди"
	alreadyStartedMsg  = okEmoji + " Бот вже запущений\n/stop - зупинити бота"
	stopMsg            = stopEmoji + " Зупинився"
	helpMsg            = infoEmoji + " Підказки\n\n+ /help - отримати всі команди\n+ /start - запустити бота\n+ /stop - зупинити бота\n+ /contacts - отримати контактну інформацію ключових членів клубу\n+ /form - отримати посилання на форму (онлайн-заявку на вступ до клубу)\n+ /we_on_the_web - отримати посилання на нас в мережі\n+ /bug_report - повідомити про знайдені помилки\n+ /ask - задати питання і отримати відповідь від адміністратора"
	contactsMsg        = "Президент: @kenjitheman\nВіце-президент: @ya_code"
	bugReportMsg       = okEmoji + " Дякую за звіт про помилку!"
	askMsg             = okEmoji + " Ми надамо відповідь якнайшвидше!"
	weOnTheWebMsg      = infinityEmoji + " " + websiteUrl
	formMsg            = infinityEmoji + " " + googleFormUrl
	reportMsg          = " Звіт про помилку від користувача @%s:\n%s"
	askQuestionMsg     = " Запитання від користувача @%s:\n%s"
	areadyStoppedMsg   = stopEmoji + " Бот вже зупинений\n/start - запустити бота"
	backToMenuMsg      = okEmoji + " Повернувся до меню"
	askForQuestionMsg  = cactusEmoji + " Будь ласка, введіть ваше запитання:"
	askForBugReportMsg = cactusEmoji + " Будь ласка, опишіть проблему:"
	thxForQuestionMsg  = greenHeartEmoji + " Ми надамо відповідь якнайшвидше!"
	thxForBugReportMsg = greenHeartEmoji + " Дякую за звіт про помилку!"
	autoOffMsg         = stopEmoji + " Бот вимкнувся автоматично через 5 хвилин бездіяльності\n/start - запустити бота"
)
