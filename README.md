# SEAD Club's helper tg bot

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

## Project structure

```go
.
├── bot
│   ├── bot.go
│   ├── keyboards.go
│   └── vars.go
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```

## Installation

```sh
git clone https://github.com/seadclub/seadclub-bot
```

## Usage

- Run it using docker:
    - You need to paste your api keys in Dockerfile:

```env
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- Run it:

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- Run it without docker:
    - You need to **create .env file** with env variables
        - And you need to **UNCOMMENT** the following lines in bot.tg:

```go
// "github.com/joho/godotenv"
```

```go
// err := godotenv.Load("../.env")
// if err != nil {
// 	fmt.Println("[ERROR] error loading .env file")
// 	log.Panic(err)
// }
```

- Run it:

```sh
go run main.go
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](https://choosealicense.com/licenses/mit)
