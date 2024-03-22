package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/sxyazi/bendan/commands/enhance"
	"github.com/sxyazi/bendan/utils"
)

func EnhanceViaQuery(query *tgbotapi.InlineQuery) any {
	text := query.Query

	if et := enhance.Do(text); et[0] != text {
		text = et[0] // TODO: return enhanced multi-url
	} else {
		return false
	}

	return tgbotapi.InlineQueryResultArticle{
		Type:  "article",
		ID:    uuid.New().String(),
		Title: utils.TruncateUTF8(text, 64),
		InputMessageContent: tgbotapi.InputTextMessageContent{
			Text: text,
		},
	}
}
