package main

import (
	"fmt"
	"rifatdinc/signalspeaker/macadres"
	"rifatdinc/signalspeaker/routeros"

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

	app.Get("/Fiber", func(c *fiber.Ctx) error {
		var Ipadresses = []string{"10.50.254.253"}
		// commnad := []string{"/interface/wireless/scan", "=number=wlan1", "=save-file=rifatDinc.txt", "=duration=5"}
		command := "/interface/wireless/print"
		admin := "admin"
		password := "mc4152"
		result := routeros.Loops(&command, &admin, &password, Ipadresses)
		fmt.Println(result)
		return c.JSON(map[string][]string{"Rahatol": result})

	})

	app.Get("/Connect_Dev", func(c *fiber.Ctx) error {
		var Ipadresses = []string{"10.50.254.253"}
		dot := routeros.Mikregistration(&username, &password, Ipadresses)
		fmt.Println(dot)
		return c.JSON(dot)
	})

	app.Post("/UploadFile", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":3011")
}
