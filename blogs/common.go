package blogs

import (
	"time"
)

type (
	Company struct {
		ID   int
		Name string
	}

	Blog struct {
		ID        int
		Title     string
		URL       string
		CreatedAt time.Time
	}

	CompanyBlog struct {
		CompanyID int
		BLogID    int
	}
)

type BlogInterface interface {
	Get() error
}

var (
	CompanyList = []string{
		"CyberAgent",
		"Cookpad",
		"DeNA",
		"Eureka",
		"GMO",
		"Gnosy",
		"mercari",
		"MoneyForward",
		"VOYAGEGROUP",
		"ZOZOTOWN",
	}

	CompanyBlogURL = map[string]string{
		"CyberAgent":   "https://developers.cyberagent.co.jp/blog/archives/category/engineer/",
		"Cookpad":      "https://techlife.cookpad.com/archive",
		"DeNA":         "https://engineer.dena.com/",
		"Eureka":       "https://medium.com/eureka-engineering",
		"GMO":          "https://blog.gmo.media/",
		"Gnosy":        "https://tech.gunosy.io/",
		"mercari":      "https://engineering.mercari.com",
		"MoneyForward": "https://moneyforward.com/engineers_blog/",
		"VOYAGE GROUP": "https://techlog.voyagegroup.com/archive",
		"ZOZOTOWN":     "https://techblog.zozo.com/",
	}
)

func NewBlog(title, url, date string) *Blog {
	t, _ := time.Parse("2006-01-02", date)
	return &Blog{
		Title:     title,
		URL:       url,
		CreatedAt: t,
	}
}

func (b *Blog) Save() error {
	return nil
}
