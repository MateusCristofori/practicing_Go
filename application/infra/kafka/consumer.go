// Esse "cosumer" será responsável por pegar os tópicos do Kafka.
package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Um struct para representar o consumer do Kafka. Ele irá ter o "MsgChan" como atributo que será um canal usando como referência o "kafka.Message". Com esse canal, podemos usar processamentos simultâneos em threads diferentes. É, basicamente, uma programação assíncrona. Esse atributo representa as mensagens do Kafka.
type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// Basicamente um método como forma de abstração onde iremos ter acesso a "instância" de um novo KafkaConsumer.
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

// Método de consumo que irá ficar lendo os tópicos do Kafka.
func (k *KafkaConsumer) Consume() {
	// Configurações necessárias para o Kafka. Iremos colocar todas elas dentro do "ConfigMap{}".
	configMap := ckafka.ConfigMap{
		// "os.Getenv()" é como pegamos variáveis de ambiente para o código. É basicamente um "process.env..." do Node.
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(&configMap)
	if err != nil {
		log.Fatalf("Error consuming Kafka message: " + err.Error())
	}
	// Após criarmos o consumer, precisamos dizer de quais tópicos queremos consumir os dados. Com o "string{}" iremos colocar todas as informações diretamente dentro do array (todas as informações da variável de ambiente "KafkaReadTopics").
	topics := []string{
		os.Getenv("KafkaReadTopic"),
	}
	// Inscrição do consumer no tópicos.
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	// Esse loop irá permitir que o consumer fique consumindo infinitamente as mensagens do Kafka.
	for {
		// Variável "mensagem" que irá ler as mensagens do kafka. O argumento "-1" da função representa o "timeOut" para a leitura de mensagens, nesse caso, o consumer não irá esperar para ler as mensagens.
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// Caso o "err" seja "vazio" (caso não ocorra nenhum erro), a variável mensagem será passada para o atributo "MsgChan" do struct KafkaConsumer.
			k.MsgChan <- msg
		}
	}

}
