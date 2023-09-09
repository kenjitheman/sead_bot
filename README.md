<h3 align="center">software engineering and development club's tg bot</h3>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

## installation

```
git clone https://github.com/kenjitheman/sead_bot
```

## usage

- run it using docker:
    - you need to paste your api keys in dockerfile:

```
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
ENV CREATOR_CHAT_ID=YOUR_CHAT_ID
ENV GOOGLE_FORM_URL=YOUR_URL
ENV SITE_URL=YOUR_URL
ENV CHANNEL_URL=YOUR_URL
```

- run it:

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- run it withot docker:
    - create .env file with env variables
        - you need to uncomment the following lines in tg module:

```
// "github.com/joho/godotenv"
```

```
// err := godotenv.Load("../.env")
// if err != nil {
// 	fmt.Println("[ERROR] error loading .env file")
// 	log.Panic(err)
// }
```

- run it:

```
go run cmd/main.go
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit)
