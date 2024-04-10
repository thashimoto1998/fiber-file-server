package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Post("/", func(c fiber.Ctx) error {
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			if token := form.Value["token"]; len(token) > 0 {
				// Get key value:
				fmt.Println(token[0])
			}

			// Get all files from "documents" key:
			files := form.File["documents"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("./uploaded-%s", file.Filename)); err != nil {
					return err
				}
			}
		}

		return nil
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
