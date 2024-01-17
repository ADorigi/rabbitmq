package rabbitmq

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestPublishCTXByte(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	rabbit := NewRabbit()
	rabbit.Configure()
	rabbit.ConnectSocket()
	rabbit.ConnectChannel()
	rabbit.DeclareQueue("hello")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Sedning to queue from 'go test'"
	rabbit.PublishCTXByte(ctx, []byte(body))

	rabbit.Close()
}
