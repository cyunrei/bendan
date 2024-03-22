package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	. "github.com/sxyazi/bendan/utils"
)

var queryChain = []func(*tgbotapi.InlineQuery) any{
	PurifyViaQuery,
	EnhanceViaQuery,
}

var viaMessage = []func(*tgbotapi.Message) bool{
	ForwardMark,
	Pin,
	Whoami,
	Me,
	Eval,
	Dontworry,
	Call,
	Mark,
	Forward,
	Purify,
	YesRight,
	YesIs,
	YesCan,
	YesLook,
}

var Bot *tgbotapi.BotAPI

func Handle(update *tgbotapi.Update) {
	if update.InlineQuery != nil {
		HandleQuery(update.InlineQuery)
		return
	}

	var message *tgbotapi.Message
	if update.Message != nil {
		message = update.Message
	} else if update.ChannelPost != nil {
		message = update.ChannelPost
	} else {
		return
	}

	if NeedToIgnore(Bot, message.Text) {
		return
	}

	//log.Printf("[%s] says: %s", message.From.UserName, message.Text)
	for _, f := range viaMessage {
		if f(message) {
			break
		}
	}
}

func HandleQuery(query *tgbotapi.InlineQuery) {
	var results []any
	for _, f := range queryChain {
		result := f(query)
		if _, ok := result.(bool); !ok {
			results = append(results, result)
		}
	}
	RespondInlineQuery(query.ID, results)
}
