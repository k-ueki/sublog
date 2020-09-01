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

	cpBlogs, err := blogContoller.GetAndSaveCookpadBlog()
	if err != nil {
		log.Fatal(err)
	}
	cpBlogs.Append(&list)

	caBlogs, err := blogContoller.GetAndSaveCyberAgentBlog()
	if err != nil {
		log.Fatal(err)
	}
	caBlogs.Append(&list)

	dBlogs, err := blogContoller.GetAndSaveDeNABlog()
	if err != nil {
		log.Fatal(err)
	}
	dBlogs.Append(&list)

	eBlogs, err := blogContoller.GetAndSaveEurekaBlog()
	if err != nil {
		log.Fatal(err)
	}
	eBlogs.Append(&list)

	gmoBlogs, err := blogContoller.GetAndSaveGmoBlog()
	if err != nil {
		log.Fatal(err)
	}
	gmoBlogs.Append(&list)

	gnosyBlogs, err := blogContoller.GetAndSaveGnosyBlog()
	if err != nil {
		log.Fatal(err)
	}
	gnosyBlogs.Append(&list)

	merBlogs, err := blogContoller.GetAndSaveMercariBlog()
	if err != nil {
		log.Fatal(err)
	}
	merBlogs.Append(&list)

	mfBlogs, err := blogContoller.GetAndSaveMoneyForwardBlog()
	if err != nil {
		log.Fatal(err)
	}
	mfBlogs.Append(&list)

	vgBlogs, err := blogContoller.GetAndSaveVoyageGroupBlog()
	if err != nil {
		log.Fatal(err)
	}
	vgBlogs.Append(&list)

	ztBlogs, err := blogContoller.GetAndSaveZozoTownBlog()
	if err != nil {
		log.Fatal(err)
	}
	ztBlogs.Append(&list)

	slack.Send(&list)
}
