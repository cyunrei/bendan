package commands

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	collect "github.com/sxyazi/go-collection"
)

// 部分句子来自某 Mastodon bot，作者不希望链接出现在象外，故不添加引用链接。
// 感谢Ta的创意，及这些富有力量的文字！
var seDontworry = []string{
	"没关系，努力过了，很厉害了",
	"没关系，努力假装是人类，很厉害了",
	"没关系，相信自己的感受，你的感觉是对的，很厉害了",
	"没关系，鼓起勇气去做，一小步也很厉害了",
	"没关系，没进步也没关系，很厉害了",
	"没关系，还活着已经很厉害了",
	"没关系，一分钟也很厉害了",
	"没关系，能发泄出来就很厉害了",
	"没关系，已经很努力了，很厉害了",
	"没关系，生不动气，保持清醒，也很厉害了",
	"没关系，有勇气say hi已经很厉害了",
	"没关系，写不出来但还在坚持写，很厉害了",
	"没关系，写下TODO已经很厉害了",
	"没关系，建好文件夹已经很厉害了",
	"没关系，感到沮丧也没关系，很厉害了",
	"没关系，用简单模式通关也很厉害了",
	"没关系，好好休息吧，已经很厉害了",
	"没关系，能和同事沟通已经很厉害了",
	"没关系，能打开VSCode已经很厉害了",
	"没关系，去面过了，很厉害了",
	"没关系，疯狂焦虑也没有放弃，很厉害了",
	"没关系，迈出了第一步，很厉害了",
	"没关系，都会过去的",
	"没关系，不想坚持下去了也没关系，辛苦了",
	"没关系，辛苦了",
	"没关系，不热爱生活也可以的，你已经很厉害了",
	"没关系，逃避虽可耻但有用",
	"没关系，就算没结果，但努力过了，很厉害了",
	"没关系，和花花草草说说话，已经很厉害了",
	"没关系，能知道自己喜欢什么，很厉害了",
	"没关系，能说出想说的，很厉害了",
	"没关系，下定决心躺平已经很厉害了",
	"没关系，没成功也很尝试过了，很厉害了",
	"没关系，不努力也可以的，很厉害了",
	"没关系，不热爱生活也可以的，很厉害了",
	"没关系，做不到还在坚持，已经很厉害了",
	"没关系，有时间能玩游戏，很厉害了",
	"没关系，不要因为无法改变的事惩罚自己",
	"没关系，工作真的很累，真的辛苦了",
	"没关系，带着面具活下去也已经很厉害了",
	"没关系，明天再学吧，很厉害了",
	"没关系，学不动了也很厉害了",
	"没关系，忘了就忘了，重要的东西一定会再次出现",
	"没关系，会写hello world已经很厉害了",
	"没关系，坚持到现在了，很厉害了",
	"没关系，心情很容易被影响不是你的错，很厉害了",
	"没关系，还在对这个世界有所期待，很厉害了",
	"没关系，认识到自己也有做不到的事情，很厉害了",
	"没关系，做了比没做好，很厉害了",
	"没关系，学习真的很难，真的辛苦了",
	"没关系，意识到自己累了，很厉害了",
	"没关系，你努力的样子真的很酷",
	"紧紧握住彼此的手吧，没关系，都很厉害了",
	"没关系，一分钟也很厉害了",
	"没关系，对人性仍抱有期待，已经很厉害了",
}

func Dontworry(msg *tgbotapi.Message) bool {
	if !collect.Contains([]string{"/没关系", "/mgx", "/dontworry"}, msg.Text) {
		return false
	}

	sel := func(userID int64) string {
		_, min, _ := time.Now().Clock()
		sum := sha1.Sum([]byte(fmt.Sprintf("%d %d", userID, min)))
		return seDontworry[int(binary.BigEndian.Uint16(sum[:]))%len(seDontworry)]
	}

	if msg.ReplyToMessage == nil {
		SendText(msg.Chat.ID, sel(msg.From.ID))
	} else {
		ReplyText(msg, sel(msg.ReplyToMessage.From.ID))
	}
	return true
}