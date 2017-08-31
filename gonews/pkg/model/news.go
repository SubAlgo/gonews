package model

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News type
type News struct {
	ID        bson.ObjectId
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
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
func CreateNews(news News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	// รายละเอียด คลิป4 [52.00]
	muteNews.Lock()         //ทำการ lock เพื่อป้องกันการส่ง Requres
	defer muteNews.Unlock() //เพื่อให้แน่ใจว่า unlock แล้ว
	newsStorage = append(newsStorage, news)
}

// ListNews fff
func ListNews() []*News {
	muteNews.Lock()         //ทำการ lock เพื่อป้องกันการส่ง Requres
	defer muteNews.Unlock() //เพื่อให้แน่ใจว่า unlock แล้ว
	r := make([]*News, len(newsStorage))
	for i := range newsStorage {
		n := newsStorage[i]
		r[i] = &n
	}
	return r
}

// GetNews fff
func GetNews(id string) *News {
	muteNews.Lock()         //ทำการ lock เพื่อป้องกันการส่ง Requres
	defer muteNews.Unlock() //เพื่อให้แน่ใจว่า unlock แล้ว

	for _, news := range newsStorage {
		if news.ID == id {
			n := news
			return &n
		}
	}
	return nil
}

// DeleteNews fff
func DeleteNews(id string) {
	// รายละเอียด คลิป4 [52.00]
	muteNews.Lock()         //ทำการ lock เพื่อป้องกันการส่ง Requres
	defer muteNews.Unlock() //เพื่อให้แน่ใจว่า unlock แล้ว
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
			return
		}
	}
}
