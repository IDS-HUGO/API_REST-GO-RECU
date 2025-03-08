package messaging

import (
	"encoding/json"
	"log"

	"demo/src/products/domain/entities"
	"demo/src/products/domain/messaging"
)

type PublishProductCreated struct {
	messageBroker messaging.IMessageBroker
}

func NewPublishProductCreated(mb messaging.IMessageBroker) PublishProductCreated {
	return PublishProductCreated{messageBroker: mb}
}

func (p *PublishProductCreated) Execute(product *entities.Product) error {
	message, err := json.Marshal(product)
	if err != nil {
		log.Println("[Messaging] Error al serializar producto:", err)
		return err
	}

	return p.messageBroker.PublishMessage("INVENTARIO_ACTUALIZADO", string(message))
}
