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
	doc.Find(".list-content > article").Each(func(i int, s *goquery.Selection) {
		title := s.Find("div > h2 > a").Text()

		datetime := time.Now().Format("2006-01-02")
		date, _ := time.Parse("2006-01-02", datetime)

		url, _ := s.Find("div > h2 > a").Attr("href")
		url = d.URL + url

		blog := NewBlog(title, url, d.Name, date)
		//fmt.Println(blog)
		blogList.Blogs = append(blogList.Blogs, blog)
	})

	return &blogList, nil
}
