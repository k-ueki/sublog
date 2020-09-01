package blogs

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/k-ueki/sublog/util"
)

type VoyageGroup struct {
	Name string
	URL  string
}

func NewVoyageGroup(blogMap map[string]string) *VoyageGroup {
	name := "VOYAGEGROUP"
	return &VoyageGroup{
		Name: name,
		URL:  blogMap[name],
	}
}

func (v *VoyageGroup) GetTableName() string {
	return "VOYAGEGROUP_blog"
}

func (v *VoyageGroup) Get(lastDate time.Time) (*BlogList, error) {
	res, err := util.HttpGet(v.URL)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var blogList BlogList
	doc.Find("div.archive-entries > section").Each(func(i int, s *goquery.Selection) {
		datetime, _ := s.Find(".archive-entry-header > .date > a > time").Attr("datetime")
		date, _ := time.Parse("2006-01-02", datetime)
		if date.After(lastDate) {
			url, _ := s.Find(".archive-entry-header > .entry-title > a").Attr("href")
			title := s.Find(".archive-entry-header > .entry-title > a").Text()

			blog := NewBlog(title, url, v.Name, date)
			blogList.Blogs = append(blogList.Blogs, blog)
		}
	})

	return &blogList, nil
}
