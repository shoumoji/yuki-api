package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var db *sqlx.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("info: ", err)
	}

	// MYSQL接続
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDRESS")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlAddr, mysqlPort, mysqlDBName)
	log.Println("info: DSN ->", dsn)
	for {
		var err error
		db, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Println("SQL Connect Error\n", err)
			time.Sleep(time.Second * 10)
			continue
		}
		break
	}

	schema := `
	CREATE TABLE IF NOT EXISTS yuki-data (
		device_id CHAR(40) NOT NULL,
		points    BIGINT   NOT NULL,
		date      DATE     NOT NULL
	);
	`
	db.MustExec(schema)

	// timezone
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	e := echo.New()
	// hard - api
	e.POST("/post", handlePOST)

	//	e.GET("/data/each", handleEachData)
	//	e.GET("/data/total", handleTotalData)

	// テスト用
	e.GET("/test/each", handleTestEach)
	e.GET("/test/total", handleTestTotal)

	e.Logger.Fatal(e.Start(":3000"))
}
