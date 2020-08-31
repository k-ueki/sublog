package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type DeNA struct {
	Name string
	URL  string
}

func NewDeNA(blogMap map[string]string) *DeNA {
	name := "DeNA"
	return &DeNA{
		Name: name,
		URL:  blogMap[name],
	}
}

func (d *DeNA) GetTableName() string {
	return "DeNA_blog"
}

func (d *DeNA) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(d.URL)
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
