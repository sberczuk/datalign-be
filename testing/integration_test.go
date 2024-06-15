//go:build integration

package testing

// run this test when developing
// not integrated with go test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestSendRequest(t *testing.T) {

	value := makeBody(3200)
	payload := struct {
		Input string
	}{Input: value}

	marshal, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	post(marshal)
}

func Test_sendHugeRequest(t *testing.T) {

	value := makeBody(32000)
	payload := struct {
		Input string
	}{Input: value}

	marshal, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	post(marshal)
}

func Test_sendVeryHugeRequest(t *testing.T) {

	value := makeBody(320000)
	payload := struct {
		Input string
	}{Input: value}

	marshal, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	post(marshal)
}

func Test_sendBadRequest(t *testing.T) {

	value := makeErrorBody()
	payload := struct {
		Input string
	}{Input: value}

	marshal, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	post(marshal)
}

func post(marshal []byte) error {
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
	if err != nil {
		return err
	}
	fmt.Println(string(resBody))
	return nil
}

func makeBody(n int) string {
	atom := "42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+"
	var b []byte
	for i := 0; i < n; i++ {
		b = append(b, atom...)
	}
	b = append(b, []byte("3")...)
	return string(b)
}

func makeErrorBody() string {
	atom := "42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+42-42+16-16+1+"
	var b []byte
	for i := 0; i < 3200; i++ {
		b = append(b, atom...)
	}
	b = append(b, []byte("3+")...)
	return string(b)
}
