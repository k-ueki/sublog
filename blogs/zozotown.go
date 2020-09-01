package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type ZozoTown struct {
	Name string
	URL  string
}

func NewZozoTown(blogMap map[string]string) *ZozoTown {
	name := "ZOZOTOWN"
	return &ZozoTown{
		Name: name,
		URL:  blogMap[name],
	}
}

func (z *ZozoTown) GetTableName() string {
	return "ZOZOTOWN_blog"
}

func (z *ZozoTown) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(z.URL)
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

			blog := NewBlog(title, url, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return &blogList, nil
}
