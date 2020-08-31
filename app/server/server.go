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

	dBlogs, err := blogContoller.GetDeNABlog()
	if err != nil {
		log.Fatal(err)
	}
	dBlogs.Append(&list)

	eBlogs, err := blogContoller.GetEurekaBlog()
	if err != nil {
		log.Fatal(err)
	}
	eBlogs.Append(&list)

	gmoBlogs, err := blogContoller.GetGmoBlog()
	if err != nil {
		log.Fatal(err)
	}
	gmoBlogs.Append(&list)

	gnosyBlogs, err := blogContoller.GetGmoBlog()
	if err != nil {
		log.Fatal(err)
	}
	gnosyBlogs.Append(&list)

	merBlogs, err := blogContoller.GetMercariBlog()
	if err != nil {
		log.Fatal(err)
	}
	merBlogs.Append(&list)

	slack.Send(&list)
}
