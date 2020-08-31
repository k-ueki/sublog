package blogs

import (
	"time"

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

func (m *Mercari) GetTableName() string {
	return "mercari_blog"
}

func (m *Mercari) Get(lastDate time.Time) (BlogList, error) {
	res, err := util.HttpRequest(m.URL + "/blog")
	if err != nil {
		return BlogList{nil}, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return BlogList{nil}, err
	}

	var blogList BlogList
	doc.Find(".post-list__item").Each(func(i int, s *goquery.Selection) {
		datetime, _ := s.Find("time").Attr("datetime")
		date, _ := time.Parse("2006-01-02", datetime)
		if date.After(lastDate) {
			url, _ := s.Find("a").Attr("href")
			url = m.URL + url
			title := s.Find(".post__title").Text()

			blog := NewBlog(title, url, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return blogList, nil
}
