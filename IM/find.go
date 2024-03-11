package IM

import (
	"4-TikTok/configs"
	"4-TikTok/model/ws"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"sort"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint) (err error) {
	collection := configs.MongoDBClient.Database(database).Collection(id)
	if collection == nil {
		fmt.Println("集合为空")
	}
	comment := ws.Message{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix(),
		Read:      read,
	}
	_, err = collection.InsertOne(context.Background(), comment)
	if err != nil {
		fmt.Println("这表示集合出错了！")
	}
	return
}

func FindMany(database string, sendId string, id string, pageSize int) (results []ws.Result, err error) {
	var resultsMe []ws.Message
	var resultsYou []ws.Message
	sendIdCollection := configs.MongoDBClient.Database(database).Collection(sendId)
	idCollection := configs.MongoDBClient.Database(database).Collection(id)
	// 如果不知道该使用什么context，可以通过context.TODO() 产生context
	sendIdTimeCursor, _ := sendIdCollection.Find(context.Background(),
		options.Find(), options.Find().SetLimit(int64(pageSize)))
	idTimeCursor, _ := idCollection.Find(context.Background(),
		options.Find(), options.Find().SetLimit(int64(pageSize)))
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return //这是go语言的语法特性,当返回值被命名之后可以直接将其返回
}

func FirstFindtMsg(database string, sendId string, id string) (results []ws.Result, err error) {
	// 首次查询(把对方发来的所有未读都取出来)
	var resultsMe []ws.Message
	sendIdCollection := configs.MongoDBClient.Database(database).Collection(sendId)
	filter := bson.M{"read": 0}
	sendIdCursor, _ := sendIdCollection.Find(context.Background(), filter)
	if sendIdCursor.Next(context.TODO()) == false { //ToDo:用于判断在mongodb下的查询文档是否为空，而不能单纯的使用游标是否为0进行判断结果集是否为空
		fmt.Println("空")
		return
	}
	fmt.Println("执行第一步")
	var unReads []ws.Message
	err = sendIdCursor.All(context.TODO(), &unReads) //用于将查询结果中的所有文档解码并存储到指定的切片变量中
	if err != nil {
		log.Println("sendIdCursor err", err)
	}
	results, err = AppendAndSort(resultsMe, unReads)
	fmt.Println("直行到删除部分")
	// 将所有的维度设置为已读
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"$set": bson.M{"read": 1},
	})
	for _, res := range results {
		fmt.Println(res)
	}
	fmt.Println("执行到最后一步")
	return results, nil
}

func AppendAndSort(resultsMe, resultsYou []ws.Message) (results []ws.Result, err error) {
	//这是一个切片操作
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}
