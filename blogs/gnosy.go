package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type Gnosy struct {
	Name string
	URL  string
}

func NewGnosy(blogMap map[string]string) *Gnosy {
	name := "Gnosy"
	return &Gnosy{
		Name: name,
		URL:  blogMap[name],
	}
}

func (g *Gnosy) GetTableName() string {
	return "Gnosy_blog"
}

func (g *Gnosy) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(g.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find("div.archive-entries > section").Each(func(i int, s *goquery.Selection) {
		datetime, _ := s.Find(".archive-entry-header > .archive-date > a > time").Attr("datetime")
		date, _ := time.Parse("2006-01-02", datetime)
		if date.After(lastDate) {
			url, _ := s.Find(".archive-entry-header > .archive-date > a").Attr("href")
			title := s.Find(".archive-entry-header > .entry-title > a").Text()

			blog := NewBlog(title, url, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return &blogList, nil
}
