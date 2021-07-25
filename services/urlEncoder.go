package services

import (
	"github.com/CharlesChou03/url-shortening-service.git/utils"
)

func UrlShorten(url string, num int) string {
	return utils.GetRandomString(num)
}
