package controller

import (
	"log"
	"os"

	"github.com/k-ueki/sublog/blogs"
	"github.com/k-ueki/sublog/config"
	database "github.com/k-ueki/sublog/db"
	"gorm.io/gorm"
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

func (c *BlogController) GetAndSaveCookpadBlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveCyberAgentBlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveDeNABlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveEurekaBlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveGnosyBlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveMercariBlog() (*blogs.BlogList, error) {
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

func (c *BlogController) GetAndSaveMoneyForwardBlog() (*blogs.BlogList, error) {
	mf := blogs.NewMoneyForward(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, mf.GetTableName())

	blogs, err := mf.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, mf.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetAndSaveVoyageGroupBlog() (*blogs.BlogList, error) {
	vg := blogs.NewVoyageGroup(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, vg.GetTableName())

	blogs, err := vg.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, vg.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (c *BlogController) GetAndSaveZozoTownBlog() (*blogs.BlogList, error) {
	zt := blogs.NewZozoTown(config.Config.ParentBlogURL)
	latest, _ := database.GetLastDate(c.DB, zt.GetTableName())

	blogs, err := zt.Get(latest)
	if err != nil {
		return nil, err
	}
	if err := blogs.Save(c.DB, zt.GetTableName()); err != nil {
		return nil, err
	}
	return blogs, nil
}
