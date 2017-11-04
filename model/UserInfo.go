package model

type User struct {
	Id          int64
	Name        string
	UserImage   string
	HomeImage   string
	MoodMessage string
	TwitterId   int64
	MyPoint     int64
}

var TestUser = User{
	Id: 1,
	Name: "中野 梓",
	UserImage: "http://dic.nicovideo.jp/oekaki/185531.png",
	HomeImage: "https://pbs.twimg.com/media/Cw9MWd7VQAA7HIF.jpg",
	MoodMessage: "あずにゃんペロペロとか辞めてくださいっ！",
	TwitterId: 3112106569,
	MyPoint: 114514,
}
