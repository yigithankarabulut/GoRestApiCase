package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/yigithankarabulut/vatansoftgocase/src/apiserver"
	"os"
)

func main() {
	if err := apiserver.New(
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
		apiserver.WithServerEnv(os.Getenv("SERVER_ENV")),
	); err != nil {
		log.Fatal(err)
	}
}
