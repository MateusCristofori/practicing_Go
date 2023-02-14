package kafka

import (
	"log"
	"os"

	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Irá criar outro producer para mandar mensagens no Kafka. Abstração para retornar uma "instância" de um Producer Kafka.
func NewKafkaProducer() *ckafka.Producer {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(&configMap)

	if err != nil {
		log.Println(p)
	}

	return p
}

// Função que irá ser usada para publicar mensagens no Kafka através de um producer.
func Publish(topic string, msg string, producer *ckafka.Producer) error {
	message := ckafka.Message{
		// Estamos passando qual o tópico (vem por parâmetro), a partição é o próprio Kafka que define e o valor é um array de bytes que foi transformado de string para bytes usando "[]byte(msg)".
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	// No final, a mensagem será produzida com o método "produce()".
	err := producer.Produce(&message, nil)

	if err != nil {
		return err
	}
	return nil
}
