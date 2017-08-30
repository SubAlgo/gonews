package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

// News type
type News struct {
	ID        string
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	newsStroage []*News
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
	//return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews create News struct
func CreateNews(news *News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	newsStroage = append(newsStroage, news)
}

// ListNews fff
func ListNews() []*News {
	return newsStroage
}

// GetNews fff
func GetNews(id string) *News {
	for _, news := range newsStroage {
		if news.ID == id {
			return news
		}
	}
	return nil
}

// DeleteNews fff
func DeleteNews(id string) {
	for i, news := range newsStroage {
		if news.ID == id {
			newsStroage = append(newsStroage[:i], newsStroage[i+1:]...)
			return
		}
	}
}
