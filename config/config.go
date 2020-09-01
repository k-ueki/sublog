package config

import (
	"log"
	"os"

	"github.com/k-ueki/sublog/blogs"

	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	BlogCompanyList []string
	ParentBlogURL   map[string]string
	SlackURL        string
	DBDial          string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("/home/ec2-user/sublog/config.ini")
	//cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: config.ini, err: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		BlogCompanyList: blogs.CompanyList,
		ParentBlogURL:   blogs.CompanyBlogURL,
		SlackURL:        cfg.Section("slack").Key("url").String(),
		DBDial:          cfg.Section("db").Key("dial").String(),
	}
}
