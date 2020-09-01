package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/k-ueki/sublog/blogs"

	"github.com/k-ueki/sublog/config"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.Config.DBDial), &gorm.Config{})
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return db, nil
}

func init() {
	db, err := DBConnection()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
		os.Exit(1)
	}

	for idx := range config.Config.BlogCompanyList {
		sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s_blog (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(111) NOT NULL,
			url VARCHAR(512) NOT NULL,
			created_at DATETIME NOT NULL)`, config.Config.BlogCompanyList[idx])
		if err := db.Exec(sql).Error; err != nil {
			log.Fatal("cannot create table:company", err)
			os.Exit(1)
		}
	}
}

func GetLastDate(db *gorm.DB, tableName string) (time.Time, error) {
	var blog blogs.Blog
	if err := db.Table(tableName).Order("created_at DESC").First(&blog).Error; err != nil {
		return time.Time{}, err
	}
	return blog.CreatedAt, nil
}
