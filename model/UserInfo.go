package model

type User struct {
	Id          int64
	UserName    string
	UserImage   string
	HomeImage   string
	MoodMessage string
	TwitterId   int64
	MyPoint     int64
}

var TestUser = User{
	Id:          1,
	UserName:    "中野 梓",
	UserImage:   "http://dic.nicovideo.jp/oekaki/185531.png",
	HomeImage:   "https://pbs.twimg.com/media/Cw9MWd7VQAA7HIF.jpg",
	MoodMessage: "あずにゃんペロペロとか辞めてくださいっ！",
	TwitterId:   3112106569,
	MyPoint:     114514,
}

var DefaultUser = User{
	UserName:    "デフォ・ルト",
	UserImage:   "https://pbs.twimg.com/profile_images/748345537074651137/WnyzqKAU.jpg",
	HomeImage:   "http://haginoryokkou.com/wp-content/uploads/2016/09/noimage.png",
	MoodMessage: "",
	MyPoint:     1000,
}