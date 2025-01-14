package parseflag

import "flag"

var (
	TgToken       string
	Dev           bool
	WebHook       bool
	Port          int
	OwnerID       int64
	DSN           string
	ImgMaxSize    int
	MaxFree       int
	ImageSavePath string
	HOST          string
	WEBSITE       string
	I18nExtraPath string
)

func init() {
	flag.BoolVar(&Dev, "dev", false, "development mode")
	flag.IntVar(&ImgMaxSize, "imgMax", 1638400, "image maximum resolution")
	flag.Int64Var(&OwnerID, "owner", 2143676086, "owner telegram id")
	flag.IntVar(&Port, "p", 8082, "port")
	flag.IntVar(&MaxFree, "mf", 0, "free user max free time")
	flag.StringVar(&HOST, "host", ``, "webhook and api host")
	flag.StringVar(&WEBSITE, "web", `pyhdxy.top`, "website site host")
	flag.BoolVar(&WebHook, "webhook", false, "enable telegram bot webhook")
	flag.StringVar(&ImageSavePath, "isp", `/mnt/stable-diffusion-webui-bot`, "Image Save Path")
	flag.StringVar(&DSN, "dsn", `./stable-diffusion-webui-bot.db`, "database, postgres|sqlite")
	flag.StringVar(&TgToken, "t", ``, "telegram bot token")
	flag.StringVar(&I18nExtraPath, "i18np", `./i18n-extra`, "i18n extra translated")
}
