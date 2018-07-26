package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"misteraladin.com/jasmine/rate-structure/config"
	gorm2 "github.com/jinzhu/gorm"
)

var (
	dbConfig  = config.Config.DB
	dbHotel  = config.Config.DBHotel
	mysqlConn *gorm.DB
	mysqlConn1 *gorm2.DB
	err       error
	err1       error
)

// initialize database
func init() {
	if dbConfig.Driver == "mysql" {
		setupMysqlConn()
		setupMysqlConn1()
	}
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)
	// mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

func setupMysqlConn1() {
	connectionString1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbHotel.User, dbHotel.Password, dbHotel.Host, dbHotel.Port, dbHotel.Name)
	mysqlConn1, err1 = gorm2.Open("mysql", connectionString1)
	if err1 != nil {
		panic(err)
	}
	err1 = mysqlConn1.DB().Ping()
	if err1 != nil {
		panic(err1)
	}
	mysqlConn1.LogMode(true)
	// mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn1() *gorm2.DB {
	return mysqlConn1
}
