package main

import (
	"log"
	"os"

	"github.com/k-ueki/sublog/slack"

	"github.com/k-ueki/sublog/config"

	database "github.com/k-ueki/sublog/db"

	"github.com/k-ueki/sublog/blogs"
)

func main() {
	db, err := database.DBConnection()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
		os.Exit(1)
	}

	list := blogs.BlogList{}

	mer := blogs.NewMercari(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(db, mer.GetTableName())

	merBlogs, err := mer.Get(latest)
	if err != nil {
		log.Fatal(err)
	}
	if err := merBlogs.Save(db, mer.GetTableName()); err != nil {
		log.Fatal(err)
	}
	merBlogs.Append(&list)

	slack.Send(&list)
}
