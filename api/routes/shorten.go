package routes

import (
	"time"

	"main.go/api/helpers"
)

type request struct {
	URL    string        `json:"url"`
	Custom string        `json:"custom"`
	Expiry time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	Custom         string        `json:"custom"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	//url validation
	if !govaildator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Remove domain from URL"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)
}
