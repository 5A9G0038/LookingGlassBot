package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func mtr(IPAddress string) string {
	if net.ParseIP(IPAddress) == nil {
		Command := ("mtr -G 2 -c 5 -erwbz " + IPAddress)
		mtrCommand := exec.Command("bash", "-c", Command)
		out, err := mtrCommand.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return "An error occurred"
		}
		return string(out)
	} else {
		return ("IP Address: '" + IPAddress + "' - Invalid ")
	}
}

func ping(IPAddress string) string {
	if net.ParseIP(IPAddress) == nil {
		Command := ("ping -c 5 " + IPAddress)
		mtrCommand := exec.Command("bash", "-c", Command)
		out, err := mtrCommand.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return "An error occurred"
		}
		return string(out)
	} else {
		return ("IP Address: '" + IPAddress + "' - Invalid ")
	}
}

func traceroute(IPAddress string) string {
	if net.ParseIP(IPAddress) == nil {
		Command := ("traceroute " + IPAddress)
		mtrCommand := exec.Command("bash", "-c", Command)
		out, err := mtrCommand.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return "An error occurred"
		}
		return string(out)
	} else {
		return ("IP Address: '" + IPAddress + "' - Invalid ")
	}
}

func main() {
	godotenv.Load()

	TGBotToken := os.Getenv("Telegram_Bot_Token")

	bot, err := tgbotapi.NewBotAPI(TGBotToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Command() {
		case "mtr":
			IPAddress := strings.ReplaceAll(update.Message.Text, "/mtr", "")
			msg.Text = "`" + mtr(IPAddress) + "`"
		case "ping":
			IPAddress := strings.ReplaceAll(update.Message.Text, "/ping", "")
			msg.Text = "`" + ping(IPAddress) + "`"
		case "traceroute":
			IPAddress := strings.ReplaceAll(update.Message.Text, "/traceroute", "")
			msg.Text = "`" + traceroute(IPAddress) + "`"
		default:
			msg.Text = "Oh... Sorry, I don't know this command."
		}
		msg.ParseMode = "markdown"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}
