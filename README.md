<h3 align="center">software engineering and development club's tg bot</h3>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

## project structure

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

## installation

```sh
git clone https://github.com/kenjitheman/seadclub-bot
```

## usage

- run it using docker:
    - you need to paste your api keys in dockerfile:

```ENV
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- run it:

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- run it without docker:
    - you need to **create .env file** with env variables
        - and you need to **UNCOMMENT** the following lines in bot.tg:

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

- run it:

```sh
go run main.go
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit)
