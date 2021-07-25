package models

// GetOriginalUrlRes is the request of API GET /api/url-shortening-service/v1/getoriginalurl
type GetOriginalUrlRes struct {
	OriginalUrl   string `json:"originalUrl"`
	ShorteningUrl string `json:"shorteningUrl"`
	ExpiredAt     int64  `json:"expiredAt"`
	CreatedAt     int64  `json:"createdAt"`
}
