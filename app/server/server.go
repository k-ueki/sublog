package main

import (
	"log"

	"github.com/k-ueki/sublog/app/server/controller"

	"github.com/k-ueki/sublog/slack"

	"github.com/k-ueki/sublog/blogs"
)

func main() {
	blogContoller := controller.NewBlogContoller()
	list := blogs.BlogList{}

	cpBlogs, err := blogContoller.GetCookpadBlog()
	if err != nil {
		log.Fatal(err)
	}
	cpBlogs.Append(&list)

	caBlogs, err := blogContoller.GetCyberAgentBlog()
	if err != nil {
		log.Fatal(err)
	}
	caBlogs.Append(&list)

	merBlogs, err := blogContoller.GetMercariBlog()
	if err != nil {
		log.Fatal(err)
	}
	merBlogs.Append(&list)

	slack.Send(&list)
}
