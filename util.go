package rabbitmq

import (
	"log"
)

func Must(err error, msg string) {
	if err != nil {
		log.Printf("Error encountered in rabbit-go: %s", err.Error())
	}
}
