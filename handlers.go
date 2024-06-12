package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
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
		return c.SendStatus(400)
	}
	log.Infof("processing %s", payload)
	expression, err := evaluateExpression(payload.Input)
	if err != nil {
		log.Errorf("Error evaluating %s", payload.Input)
		return c.SendStatus(500)
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
