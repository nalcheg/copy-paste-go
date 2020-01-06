package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedStr  string
	}{
		{
			name:         "unauthorized",
			url:          "http://127.0.0.1:9011",
			expectedCode: 401,
			expectedStr:  "",
		}, {
			name:         "greetings",
			url:          "http://127.0.0.1:9011/welcome?pass=true&iam=nal",
			expectedCode: 200,
			expectedStr:  "Hi, nal !",
		},
	}

	go func() {
		main()
	}()

	time.Sleep(1000 * time.Millisecond)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := http.Client{}
			resp, err := client.Get(test.url)
			if err != nil {
				t.Error(err)
			}
			defer func() {
				if err := resp.Body.Close(); err != nil {
					t.Error(err)
				}
			}()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}

			if string(body) != test.expectedStr {
				t.Error("unexpected body")
			}
			if resp.StatusCode != test.expectedCode {
				t.Error("unexpected response http code")
			}
		})
	}
}
