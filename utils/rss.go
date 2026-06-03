package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/iv-tunate/fiids/models"
)

func UrlToFeed(url string) (models.RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	
	rssFeed := models.RSSFeed{}
	if err != nil{
		return rssFeed, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return rssFeed, err
	}
	err = xml.Unmarshal(data, &rssFeed)
	 if err != nil{
		return rssFeed, err
	}

	return rssFeed, nil
}