package scoredatabase

import (
	"context"
	"hr/configs/models/square"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTopicList(start, end int64, collection *mongo.Collection) ([]square.Topic, error) {
	filter := bson.D{}
	options := options.Find().SetSort(bson.D{{"created_at", -1}}).SetSkip(start).SetLimit(end - start + 1)

	// 执行查询
	cursor, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// 解码结果
	var topics []square.Topic
	if err := cursor.All(context.TODO(), &topics); err != nil {
		return nil, err
	}

	return topics, nil
}
