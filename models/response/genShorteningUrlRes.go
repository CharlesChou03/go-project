package models

// GenShorteningUrlRes is the request of API POST /api/url-shortening-service/v1/generate
type GenShorteningUrlRes struct {
	OriginalUrl   string `json:"originalUrl"`
	ShorteningUrl string `json:"shorteningUrl"`
	ExpiredAt     int64  `json:"expiredAt"`
}
