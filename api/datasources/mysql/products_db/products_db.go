package products_db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBTYPE = "mysql"
	SCHEMA = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	Client   *gorm.DB
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbName   = os.Getenv("MYSQL_DATABASE")

	datasourceName = fmt.Sprintf(SCHEMA, username, password, dbName)
)

func init() {
	var err error
	Client, err = gorm.Open(DBTYPE, datasourceName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database successfully configure")
}
