package alert

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	config "github.com/kaustubhkapatral/monitoring-bot/config"
	valcheck "github.com/kaustubhkapatral/monitoring-bot/valcheck"
)

func HexSend() error {
	botToken := config.NewApp.Token
	chatIDConfig := config.NewApp.ChatId
	chatID, err := strconv.ParseInt(chatIDConfig, 10, 64)
	if err != nil {
		fmt.Println("Unable to convert int64", err)
		return err
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	bot.Debug = true
	msg := tgbotapi.NewMessage(chatID, "")

	height, err := valcheck.HexCheck()
	if err != nil {
		fmt.Println("Error checking hex ", err)
		return err
	}
	if height == "" {
		return nil
	}

	msg.Text = "Validator missed block at height \t" + height

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func JailSend() error {
	botToken := config.NewApp.Token
	chatIDConfig := config.NewApp.ChatId
	chatID, err := strconv.ParseInt(chatIDConfig, 10, 64)
	if err != nil {
		fmt.Println("Unable to convert int64", err)
		return err
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	bot.Debug = true
	msg := tgbotapi.NewMessage(chatID, "")
	resp, err := valcheck.JailCheck()
	if err != nil {
		fmt.Println("Error checking jailed ", err)
		return err
	}

	if resp == true {
		msg.Text = "Validator is not in acive set"
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	} else {
		msg.Text = "Validator is in active set"
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
