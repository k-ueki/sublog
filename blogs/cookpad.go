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
	doc.Find("div#main-inner").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i)
		//datetime := s.Find("article.js-entry-article").Text()
		//fmt.Println(datetime)
		//	date, _ := time.Parse("2006-01-02T15:04:05-07:00", datetime)
		//	if date.After(lastDate) {
		//		url, _ := s.Find(".card-caeng__title-link").Attr("href")
		//title := s.Find("div.entry-inner").Text()
		//fmt.Println(title)
		//
		//		blog := NewBlog(title, url, date)
		//		blogList.Blogs = append(blogList.Blogs, blog)
		//	}
	})

	return &blogList, nil
}
