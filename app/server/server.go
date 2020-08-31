package main

import (
	"fmt"
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
	blogURLMap := config.Config.ParentBlogURL

	mer := blogs.NewMercari(blogURLMap)
	merLatest, _ := database.GetLastDate(db, mer.GetTableName())

	merBlogs, err := mer.Get(merLatest)
	if err != nil {
		log.Fatal(err)
	}
	if err := merBlogs.Save(db, mer.GetTableName()); err != nil {
		log.Fatal(err)
	}
	merBlogs.Append(&list)

	ca := blogs.NewCyberAgent(blogURLMap)
	caLatest, _ := database.GetLastDate(db, ca.GetTableName())
	fmt.Println(caLatest)

	caBlogs, err := ca.Get(caLatest)
	if err != nil {
		log.Fatal(err)
	}
	if err := caBlogs.Save(db, ca.GetTableName()); err != nil {
		log.Fatal(err)
	}
	caBlogs.Append(&list)

	slack.Send(&list)
}
