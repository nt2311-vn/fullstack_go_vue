package logger

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/kataras/golog"
)

var glog = golog.New()

type remote struct{}

func (r *remote) Write(data []byte) (n int, err error) {
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
