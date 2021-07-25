package models

// GenShorteningUrlReq is the request of API POST /api/url-shortening-service/v1/generate
type GenShorteningUrlReq struct {
	UserId      string `json:"userId"`
	OriginalUrl string `json:"originalUrl"`
	ExpiredAt   int64  `json:"expiredAt"`
}

func (r GenShorteningUrlReq) Validate() bool {
	if r.UserId == "" || r.OriginalUrl == "" {
		return false
	}
	if r.ExpiredAt < 0 {
		return false
	}
	return true
}
