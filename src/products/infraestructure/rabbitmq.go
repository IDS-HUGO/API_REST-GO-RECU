package infraestructure

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal("[RabbitMQ] Error al conectar:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("[RabbitMQ] Error al abrir canal:", err)
	}

	log.Println("[RabbitMQ] Conexión exitosa")
	return &RabbitMQ{Conn: conn, Channel: ch}
}

func (r *RabbitMQ) PublishMessage(queue string, body string) error {
	q, err := r.Channel.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		log.Println("[RabbitMQ] Error al declarar cola:", err)
		return err
	}

	err = r.Channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Println("[RabbitMQ] Error al publicar mensaje:", err)
		return err
	}

	log.Println("[RabbitMQ] Mensaje publicado en", queue, ":", body)
	return nil
}

func (r *RabbitMQ) Close() {
	r.Conn.Close()
	r.Channel.Close()
}
