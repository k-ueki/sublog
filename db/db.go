package db

import (
	"log"
	"os"

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

	if err := db.Exec(`CREATE TABLE IF NOT EXISTS company (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(55) NOT NULL
		)`).Error; err != nil {
		log.Fatal("cannot create table:company", err)
		os.Exit(1)
	}
	if err := db.Exec(`CREATE TABLE IF NOT EXISTS blog (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(77) NOT NULL,
			url VARCHAR(85) NOT NULL,
			created_at DATETIME NOT NULL
		)`).Error; err != nil {
		log.Fatal("cannot create table:blog", err)
		os.Exit(1)
	}
	if err := db.Exec(`CREATE TABLE IF NOT EXISTS company_blog (
			company_id INT NOT NULL,
			blog_id    INT NOT NULL,
			FOREIGN KEY company_blog_company_id	(company_id) REFERENCES company(id),
			FOREIGN KEY company_blog_blog_id(blog_id) REFERENCES blog(id)
		)`).Error; err != nil {
		log.Fatal("cannot create table:company_blog", err)
		os.Exit(1)
	}

	//for _, com := range blogs.CompanyList {
	//	if err := db.Debug().Exec("INSERT INTO company (`name`) VALUES (?)", com).Error; err != nil {
	//		log.Fatalf("cannot insert company:%s err:%v", com, err.Error)
	//	}
	//}
}
