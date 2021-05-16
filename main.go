package main

import (
	"kafkagoproxy/kafka"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("weclome use kafka api!")
	})
	app.Post("/api/Kafka/UploadLogData", func(c *fiber.Ctx) error {
		kafkacnn := new(kafka.Kafkaconn)
		if err := c.BodyParser(kafkacnn); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    0,
				"data":    nil,
				"message": err.Error(),
			})
		}
		//log.Fatal(kafkacnn.Partion)
		k := kafkacnn
		if partition, offset, err := kafka.SendMessage(*k); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    0,
				"data":    nil,
				"message": err.Error(),
			})
		} else {
			kfakaResult := kafka.KafkaResult{
				Topics:  k.Topics,
				Partion: partition,
				Offset:  offset,
			}
			return c.Status(200).JSON(fiber.Map{
				"code":    1,
				"data":    kfakaResult,
				"message": "sucess",
			})
		}

	})

	app.Post("/api/Kafka/AscUploadLogData", func(c *fiber.Ctx) error {
		kafkacnn := new(kafka.Kafkaconn)
		if err := c.BodyParser(kafkacnn); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    0,
				"data":    nil,
				"message": err.Error(),
			})
		}
		k := kafkacnn

		if err := kafka.AsycSendMessage(*k); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    0,
				"data":    "",
				"message": err.Error(),
			})
		} else {
			return c.Status(200).JSON(fiber.Map{
				"code":    1,
				"data":    nil,
				"message": "sucess",
			})
		}

	})
	//log.Fatal(app.Listen(3000))
	app.Listen(":3000")
}
