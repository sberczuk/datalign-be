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

// post the payload
func eval(c fiber.Ctx) error {
	payload := struct {
		Input string
	}{}
	body := c.Body()
	log.Infof("Got a request %V", string(body))
	err := json.Unmarshal(body, &payload)
	if err != nil {
		log.Errorf("bad request for payload %v", string(body))
		return c.Status(400).SendString(err.Error())
	}

	err = validateExpression(payload.Input)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	log.Infof("processing %s", payload)
	expression, err := evaluateExpression(payload.Input)
	if err != nil {
		log.Errorf("Error evaluating %s", payload.Input)
		return c.Status(500).SendString(err.Error())
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

	matched, err := regexp.Match(`^[\d\\+\-*\/ \(\)\.]+$`, []byte(input))

	if !matched || err != nil {
		return fmt.Errorf("invalid chars")

	}

	return nil
}
