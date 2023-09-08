FROM golang:alpine
WORKDIR /app
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
ENV CREATOR_CHAT_ID=YOUR_CHAT_ID
ENV GOOGLE_FORM_URL=YOUR_URL
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 8080
CMD ["./main"]
