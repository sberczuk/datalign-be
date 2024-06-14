package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func Test_sendRequest(t *testing.T) {

	value := makeBody()
	payload := struct {
		Input string
	}{Input: value}

	marshal, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	requestURL := fmt.Sprintf("http://localhost:%d/eval", 3000)
	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(marshal))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(res.Status)
	resBody, err := io.ReadAll(res.Body)
	fmt.Println(string(resBody))
}

func makeBody() string {
	atom := "42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+"
	var b []byte
	for i := 0; i < 3200; i++ {
		b = append(b, atom...)
	}
	b = append(b, []byte("3")...)
	return string(b)
}
