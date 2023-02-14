package main

import (
	"log"

	kafka "github.com/Mateuscristofori/praticing_golang/application/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()

	kafka.Publish("Olá", "readtest", producer)

}
