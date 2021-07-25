package models

// GetOriginalUrlReq is the request of API GET /api/url-shortening-service/v1/getoriginalurl
type GetOriginalUrlReq struct {
	ShorteningUrl string `json:"shorteningUrl"`
}

func (r GetOriginalUrlReq) Validate() bool {
	if r.ShorteningUrl == "" {
		return false
	}
	return true
}
