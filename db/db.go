package db

import (
	"fmt"
	"log"
	"os"

	"github.com/k-ueki/sublog/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@/sublog?parseTime=true")
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

	con := config.NewConfig()
	for idx := range con.BlogCompanyList {
		sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s_blog (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(111) NOT NULL,
			url VARCHAR(111) NOT NULL,
			created_at DATETIME NOT NULL)`, con.BlogCompanyList[idx])
		if err := db.Exec(sql).Error; err != nil {
			log.Fatal("cannot create table:company", err)
			os.Exit(1)
		}
	}
}
