package main

import (
	"flag"
	"log"
	"os"

	"github.com/minya/telegram"
)

func main() {
	botToken, port := initialize()

	listenErr := telegram.StartListen(botToken, port, handleUpdate)
	if nil != listenErr {
		log.Printf("Unable to start listen: %v\n", listenErr)
	}
}

func handleUpdate(update telegram.Update) interface{} {
	log.Printf("[INFO] Update received: %v\n", update)
	var replyMessage string
	if update.Message.Text == "" {
		replyMessage = "I don't understand you"
	} else {
		replyMessage = "You said: " + update.Message.Text
	}

	replyMsg := telegram.ReplyMessage{
		ChatId: update.Message.Chat.Id,
		Text:   replyMessage,
	}
	log.Printf("[INFO] Reply message: %v\n", replyMsg)
	return replyMsg
}

func initialize() (string, int) {
	var logPath string
	var botToken string
	var port int
	flag.StringVar(&logPath, "logpath", "echobot", "Path to write logs")
	flag.StringVar(&botToken, "token", "", "Bot Token (secret)")
	flag.IntVar(&port, "port", -1, "Http port to listen")
	flag.Parse()

	if botToken == "" {
		log.Fatal("Bot token is empty")
	}

	if port == -1 {
		log.Fatal("Port is empty")
	}

	if port < 0 || port > 65535 {
		log.Fatal("Port is out of range")
	}

	setUpLogger(logPath)

	return botToken, port
}

func setUpLogger(logPath string) {
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
}
