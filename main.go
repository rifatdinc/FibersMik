package main

import (
	"rifatdinc/signalspeaker/macadres"
	"rifatdinc/signalspeaker/routeros"
	"rifatdinc/signalspeaker/telegram"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var username string = "admin"
var password string = "mc4152"

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/macvendor", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"vendor": macadres.RequestMac("C4:AD:34:E4:3D:F9"),
		})
	})

	// app.Get("/Fiber", func(c *fiber.Ctx) error {
	// 	var Ipadresses = []string{"10.50.254.253"}
	// 	// commnad := []string{"/interface/wireless/scan", "=number=wlan1", "=save-file=rifatDinc.txt", "=duration=5"}
	// 	command := []string{"/interface/wireless/print"}
	// 	admin := "admin"
	// 	password := "mc4152"
	// 	result := routeros.Loops(command, &admin, &password, Ipadresses)
	// 	return c.JSON(map[string][]string{"Rahatol": result})

	// })

	app.Get("/Connect_Dev", func(c *fiber.Ctx) error {
		go func(ip string) {
			routeros.Nas(ip)
		}("31.145.83.206:8728")
		// routeros.Client("10.50.254.253:8728", "admin", "mc4152")

		return c.JSON(map[string]string{"result": "xx"})
	})

	app.Get("/TelegramApi", func(c *fiber.Ctx) error {
		telegram.Telegram("/interface/wireless/print")
		return c.SendString("x")
	})

	app.Listen(":3011")
}
