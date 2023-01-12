package db

import (
	"github.com/sxyazi/bendan/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func AddPinned(pm *types.PinnedMessage) error {
	pm.CreatedAt = time.Now()
	_, err := Db().Collection("pinned").InsertOne(ctx, pm)
	if err != nil {
		log.Println("AddPinned error:", err)
		return err
	}

	return nil
}

func RemovePinned(id int, chatId int64) (int64, error) {
	one, err := Db().Collection("pinned").DeleteOne(ctx, bson.M{"id": id, "chatId": chatId})
	if err != nil {
		log.Println("RemovePinned error:", err)
		return 0, err
	}

	return one.DeletedCount, nil
}

func GetPinned(chatId int64) ([]*types.PinnedMessage, error) {
	opt := options.Find().SetSort(bson.M{"$natural": -1})
	cur, err := Db().Collection("pinned").Find(ctx, bson.M{"chatId": chatId}, opt)
	if err != nil {
		log.Println("GetPinned error:", err)
		return nil, err
	}

	var records []*types.PinnedMessage
	if err = cur.All(ctx, &records); err != nil {
		log.Println("GetPinned error:", err)
		return nil, err
	}
	return records, nil
}
