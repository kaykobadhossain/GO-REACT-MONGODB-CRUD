package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error handling dotenv file")
	}
	PORT := os.Getenv("PORT")
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MONGODB CONNECTED SUCCESSFULLY")

	collection = client.Database("golangDB").Collection("todos")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})
	app.Get("/api/todos", getAllTodos)
	app.Post("/api/todos", createATodo)
	app.Patch("/api/todos/:id", updateATodo)
	app.Delete("/api/todos/:id", deleteATodo)

	log.Fatal(app.Listen(":" + PORT))
}

func getAllTodos(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"msg": "All todo List"})
}

func createATodo(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"msg": "create A Todo"})
}
func updateATodo(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"msg": "update a todo"})
}
func deleteATodo(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"msg": "delete a todo"})
}
