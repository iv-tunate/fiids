package models

type Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description *string `xml:"description"`
		Language string `xml:"language"`
		Item []RSSItem `xml:"item"`
}

type RSSFeed struct{
	Channel Channel `xml:"channel"`
}

type RSSItem struct{
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

