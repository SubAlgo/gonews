package model

import (
	"fmt"
	"unicode/utf8"

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
