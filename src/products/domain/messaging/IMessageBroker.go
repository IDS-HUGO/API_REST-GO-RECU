package messaging

type IMessageBroker interface {
	PublishMessage(queue string, message string) error
}
