package blogs

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type Mercari struct {
	Name string
	URL  string
}

func NewMercari(blogMap map[string]string) *Mercari {
	name := "mercari"
	return &Mercari{
		Name: name,
		URL:  blogMap[name],
	}
}

func (m *Mercari) Get() error {
	res, err := util.HttpRequest(m.URL + "/blog")
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".post-list__item").Each(func(i int, s *goquery.Selection) {
		date, _ := s.Find("time").Attr("datetime")
		url, _ := s.Find("a").Attr("href")
		url = m.URL + url
		title := s.Find(".post__title").Text()

		blog := NewBlog(title, url, date)
		fmt.Println(*blog)
	})
	return nil
}
