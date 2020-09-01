package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type Eureka struct {
	Name string
	URL  string
}

func NewEureka(blogMap map[string]string) *Eureka {
	name := "Eureka"
	return &Eureka{
		Name: name,
		URL:  blogMap[name],
	}
}

func (e *Eureka) GetTableName() string {
	return "Eureka_blog"
}

func (e *Eureka) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(e.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find(".js-collectionStream > div.streamItem").Each(func(i int, s *goquery.Selection) {
		s.Find("section > div.row > div.col").Each(func(j int, ss *goquery.Selection) {
			datetime, _ := s.Find("div.ui-caption > time").Attr("datetime")
			date, _ := time.Parse("2006-01-02T15:04:05.999Z", datetime)
			if date.After(lastDate) {
				url, _ := s.Find("a").Attr("href")
				title := ss.Find("a > h3").Text()

				if title != "" && url != "" {
					blog := NewBlog(title, url, date)
					blogList.Blogs = append(blogList.Blogs, blog)
				}
			}
		})
	})

	return &blogList, nil
}
