package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
	"regexp"
	"strconv"
)

const MAX_REQUEST_LENGTH = 320000

// post the payload
func eval(c fiber.Ctx) error {
	payload := struct {
		Input string
	}{}
	body := c.Body()

	requestEval := string(body)
	requestLength := len(requestEval)
	log.Infof("reqeust length = %d", requestLength)
	if requestLength > MAX_REQUEST_LENGTH {
		return c.Status(http.StatusRequestEntityTooLarge).SendString("request is too large")
	}

	log.Infof("Got a request %V", requestEval)
	err := json.Unmarshal(body, &payload)
	if err != nil {
		log.Errorf("bad request for payload %v", requestEval)
		return c.Status(400).SendString(err.Error())
	}

	err = validateExpression(payload.Input)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	log.Infof("processing %s", payload)
	expression, err := evaluateExpression(payload.Input)

	// Since the expression language doesn't seem to differentiate kinds of errors
	// this error could be a 422 (unprossible entity) or a 500 (server error)
	// 500 seems a safer defauly
	if err != nil {
		log.Errorf("Error evaluating %s", payload.Input)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	respValue := strconv.FormatFloat(expression, 'f', -1, 32)

	log.Infof("returning answer %s for ", respValue, payload.Input)
	c.Status(http.StatusOK)
	err = c.JSON(fiber.Map{
		"answer": respValue,
	},
	)
	if err != nil {
		return c.SendStatus(500)
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	return err
}

func validateExpression(input string) error {

	// ideally I'd use go Playground validator, but since this is simple and in the interests of time I'll
	// use a regex
	if len(input) == 0 {
		return fmt.Errorf("empty expression")
	}

	// TODO: we want this to end with number followed by optional parentheses or spaces
	// but that doesn't handle checking for balanced parens.
	// We could add that check without a regex by writing a func that checks for balanced parens or just
	// not support parens (and put that check back in the front end
	// Optimization: use Compile to initialize this at startup
	re, err := regexp.Compile(`^[\d\s\\+\-*\/\(\)\.]+[\s\d\)]$`)
	if err != nil {
		return err
	}
	matched := re.MatchString(input)
	//matched, err := regexp.Match(`^[\d\s\\+\-*\/ \(\)\.]+[\d\)]$`, []byte(input))

	if !matched {
		return fmt.Errorf("invalid chars")
	}

	return nil
}
