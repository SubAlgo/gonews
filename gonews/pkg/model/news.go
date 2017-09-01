package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News type
type News struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time `bson:"CreatedAt"`
	UpdatedAt time.Time `bson:"UpdatedAt"`
}

var (
	newsStorage []News
	muteNews    sync.Mutex //เพื่อป้องกัน การ Create data พร้อมๆ ที่อาจมีปัญหา
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
	//return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews create News struct
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	//news.ID.Hex()

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil
}

// ListNews fff
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

// GetNews fff
func GetNews(id string) (*News, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid ID")
	}
	objectID := bson.ObjectIdHex(id)
	if !objectID.Valid() {
		return nil, fmt.Errorf("invalid id")
	}
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// DeleteNews fff
func DeleteNews(id string) error {
	/*objectID := bson.ObjectId(id)
	if !objectID.Valid() { // ถ้าไม่เจอ id ให้ทำ
		return fmt.Errorf("invalid id")
	}
	*/
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid ID")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNews update|Edit news
func UpdateNews(news *News) error {
	if news.ID == "" {
		return fmt.Errorf("required to update")
	}

	news.UpdatedAt = time.Now()
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").UpdateId(news.ID, news)
	if err != nil {
		return err
	}
	return nil
}
