package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type AWS struct {
	Name string
	URL  string
}

func NewAWS(blogMap map[string]string) *AWS {
	name := "AWS"
	return &AWS{
		Name: name,
		URL:  blogMap[name],
	}
}

func (a *AWS) GetTableName() string {
	return "AWS_blog"
}

func (a *AWS) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(a.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find("main > article").Each(func(i int, s *goquery.Selection) {
		datetime, _ := s.Find("footer.blog-post-meta > time").Attr("datetime")
		date, _ := time.Parse("2006-01-02T15:04:05-07:00", datetime)
		if date.After(lastDate) {
			url, _ := s.Find("h2.blog-post-title > a").Attr("href")
			title := s.Find("h2.blog-post-title > a").Text()

			blog := NewBlog(title, url, a.Name, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return &blogList, nil
}
