package slack

import (
	"log"
	"sync"

	"github.com/k-ueki/sublog/blogs"
	"github.com/k-ueki/sublog/config"
	"github.com/k-ueki/sublog/util"
)

func Send(blogs *blogs.BlogList) {
	ch := make(chan string)
	for _, blog := range blogs.Blogs {
		go blog.GenerateJson(ch)
	}
	wg := &sync.WaitGroup{}
	for _ = range blogs.Blogs {
		wg.Add(1)
		go func() {
			if err := util.HttpPost(config.Config.SlackURL, <-ch); err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
