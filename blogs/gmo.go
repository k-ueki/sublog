package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type GMO struct {
	Name string
	URL  string
}

func NewGMO(blogMap map[string]string) *GMO {
	name := "GMO"
	return &GMO{
		Name: name,
		URL:  blogMap[name],
	}
}

func (g *GMO) GetTableName() string {
	return "GMO_blog"
}

func (g *GMO) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(g.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find(".article-list > .col-md-12").Each(func(i int, s *goquery.Selection) {
		//	datetime, _ := s.Find(".card-caeng__content > .card-caeng__meta > time").Attr("datetime")
		//	date, _ := time.Parse("2006-01-02T15:04:05-07:00", datetime)
		//	if date.After(lastDate) {
		//		url, _ := s.Find(".card-caeng__title-link").Attr("href")
		//		title := s.Find(".card-caeng__title-link").Text()
		//
		//		blog := NewBlog(title, url, date)
		//		blogList.Blogs = append(blogList.Blogs, blog)
		//	}
	})

	return &blogList, nil
}
