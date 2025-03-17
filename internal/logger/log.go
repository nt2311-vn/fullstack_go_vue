package logger

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kataras/golog"
)

var Logger = golog.New()

type remote struct{}

func (r remote) Write(data []byte) (n int, err error) {
	go func() {
		req, err := http.NewRequest("POST", "http://localhost:8010/log", bytes.NewBuffer(data))

		if err == nil {
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("error in making request", err)
			}
			defer resp.Body.Close()
		}
	}()

	return len(data), nil
}

func SetLoggingOutput(localStdout bool) {
	if localStdout {
		configureLocal()
		return
	}

	configureRemote()
}

func configureLocal() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	Logger.SetOutput(os.Stdout)
	Logger.SetLevel("debug")
	Logger.SetLevelOutput("info", file)
}

func configureRemote() {
	r := remote{}

	Logger.SetLevelFormat("info", "json")
	Logger.SetLevelOutput("info", r)
}
