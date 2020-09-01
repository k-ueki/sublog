package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type Cookpad struct {
	Name string
	URL  string
}

func NewCookpad(blogMap map[string]string) *Cookpad {
	name := "Cookpad"
	return &Cookpad{
		Name: name,
		URL:  blogMap[name],
	}
}

func (c *Cookpad) GetTableName() string {
	return "Cookpad_blog"
}

func (c *Cookpad) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(c.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find(".archive-entries > section").Each(func(i int, s *goquery.Selection) {
		datetime, _ := s.Find(".archive-entry-header > .date > a > time").Attr("datetime")
		date, _ := time.Parse("2006-01-02", datetime)
		if date.After(lastDate) {
			url, _ := s.Find(".archive-entry-header > .entry-title > a").Attr("href")
			title := s.Find(".archive-entry-header > .entry-title > a").Text()

			blog := NewBlog(title, url, c.Name, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return &blogList, nil
}
