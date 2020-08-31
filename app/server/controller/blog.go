package controller

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/k-ueki/sublog/blogs"
	"github.com/k-ueki/sublog/config"
	database "github.com/k-ueki/sublog/db"
)

type BlogController struct {
	DB *gorm.DB
}

func NewBlogContoller() *BlogController {
	db, err := database.DBConnection()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
		os.Exit(1)
	}
	return &BlogController{DB: db}
}

func (c *BlogController) GetCookpadBlog() (*blogs.BlogList, error) {
	cp := blogs.NewCookpad(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, cp.GetTableName())

	blogs, err := cp.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, cp.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetCyberAgentBlog() (*blogs.BlogList, error) {
	ca := blogs.NewCyberAgent(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, ca.GetTableName())

	blogs, err := ca.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, ca.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetDeNABlog() (*blogs.BlogList, error) {
	d := blogs.NewDeNA(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, d.GetTableName())

	blogs, err := d.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, d.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetEurekaBlog() (*blogs.BlogList, error) {
	e := blogs.NewEureka(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, e.GetTableName())

	blogs, err := e.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, e.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetGmoBlog() (*blogs.BlogList, error) {
	gmo := blogs.NewGMO(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, gmo.GetTableName())

	blogs, err := gmo.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, gmo.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetGnosyBlog() (*blogs.BlogList, error) {
	gnosy := blogs.NewGnosy(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, gnosy.GetTableName())

	blogs, err := gnosy.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, gnosy.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetMercariBlog() (*blogs.BlogList, error) {
	mer := blogs.NewMercari(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, mer.GetTableName())

	blogs, err := mer.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, mer.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}
