package config

import "github.com/k-ueki/sublog/blogs"

type Config struct {
	BlogCompanyList []string
	ParentBlogURL   map[string]string
}

func NewConfig() *Config {
	return &Config{
		BlogCompanyList: blogs.CompanyList,
		ParentBlogURL:   blogs.CompanyBlogURL,
	}
}
