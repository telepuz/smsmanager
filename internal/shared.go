package internal

type Message struct {
	Index   int    `xml:"Index"`
	Phone   string `xml:"Phone"`
	Content string `xml:"Content"`
	Date    string `xml:"Date"`
}
