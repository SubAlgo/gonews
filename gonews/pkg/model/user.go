package model

import (
	"fmt"
	"log"
	"unicode/utf8"

	"gopkg.in/mgo.v2"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register Function Resigter Member
func Register(username, password string) error {
	if utf8.RuneCountInString(username) < 4 {
		return fmt.Errorf("username must >= 4 chars")
	}
	if utf8.RuneCountInString(password) < 6 {
		return fmt.Errorf("password must >= 6 chars")
	}
	hpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s := mongoSession.Copy()
	defer s.Close()
	err = s.DB(database).C("users").Insert(bson.M{
		"username": username,
		"password": string(hpwd),
	})

	if err != nil {
		return err
	}
	return nil
}

// Login Function Login
func Login(username, password string) (string, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var user bson.M
	err := s.
		DB(database).
		C("users").
		Find(bson.M{"username": username}).
		One(&user)
		//DB("gonews").C("users").Find(bson.M{"username": username}).One(&user) คือ เชื่อมต่อ Database แล้วตรวจดู usernameว่ามีหรือไม่

	log.Println("Inputusername: ", username)
	log.Println("InputPassword: ", password)

	if err == mgo.ErrNotFound {
		return "", fmt.Errorf("username or password wrong1")
	}

	if err != nil {
		return "", err
	}

	hpwd, _ := user["password"].(string)
	log.Println(hpwd)
	err = bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(password))
	if err != nil {
		//return "", fmt.Errorf("username or password wrong2")
		return "", err
	}
	userID, _ := user["_id"].(bson.ObjectId)
	return userID.Hex(), nil
}

// CheckUserID fff
func CheckUserID(id string) (bool, error) {
	s := mongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return false, fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	n, err := s.
		DB(database).
		C("users").
		FindId(objectID).
		Count()

	if err != nil {
		return false, err
	}
	if n <= 0 {
		return false, nil
	}
	return true, nil
}
