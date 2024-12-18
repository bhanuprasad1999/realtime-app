package services

import (
	"github.com/google/uuid"
)

type Message struct {
	MessageID   string    `json:"message_id"`
	ChannelName string `json:"channel_name"`
	Username string `json:"username"`
	TextMessage string `json:"text_message"`
}

func generateMessageID() string{
	return uuid.New().String()
}

func AddMessage(text_message string, username string, channel_name string){
	var all_messages []Message
	// message object 
	var message Message
	message.ChannelName = channel_name
	message.Username = username
	message.TextMessage = text_message
	message.MessageID = generateMessageID()
	
	readJSONFile(&all_messages, "messages")
	all_messages = append(all_messages, message)
	writeJSONFile(all_messages, "messages")
}