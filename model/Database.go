package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var MyDB *sql.DB

func DataBaseInit() {
	dataSource := os.Getenv("HEROKU_DATABASE_URL")
	var err error
	MyDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
}

func InsertNewUserByTwitter(twitterId int64) error {
	_, err := MyDB.Exec("INSERT INTO users (userName, userImage, homeImage, moodMessage, twitterId, myPorint)" +
		" VALUES(?, ?, ?, ?, ?, ?)", DefaultUser.UserName, DefaultUser.UserImage, DefaultUser.HomeImage,
			DefaultUser.MoodMessage, twitterId, DefaultUser.MyPoint)
	if err != nil {
		return err
	}
	return nil
}

func IsUserExistByTwitter(twitterId int64)(bool){
	var count int64
	if err := MyDB.QueryRow("select count(id) from users where twitterId = ?;", twitterId).Scan(&count); err != nil {
		return false
	}
	if count == 1{
		return true
	}else{
		return false
	}
}

func SelectUserDataById(id string)(User, error){
	var userData User
	if err := MyDB.QueryRow("select id, userName, userImage, homeImage, moodMessage, " +
		"myPoint from users where id = ?;", id).Scan(&userData.Id, &userData.UserName,
			&userData.UserImage, &userData.HomeImage, &userData.MoodMessage, &userData.MyPoint); err != nil {
				return User{}, err
	}
	return userData, nil
}

