package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type RawData struct {
	DeviceID   string `json:"device_id"`
	HuitPoints int64  `json:"huit_points"`
}

type Data struct {
	DeviceID string `json:"device_id" db:"device_id"`
	Points   uint64 `json:"points" db:"points"`
	Date     string `json:"date" db:"date"`
}

func handlePOST(c echo.Context) (err error) {
	postData := new(RawData)
	if err = c.Bind(postData); err != nil {
		return err
	}

	if len(postData.DeviceID) >= 40 {
		// デバイスIDデカすぎ
		return err
	}

	if postData.HuitPoints <= 0 {
		// 入力されたデータが0以下なら、当該device_idのデータを削除(テスト用機能)
		q, err := db.Prepare("DELETE FROM yuki_data WHERE device_id = ?")
		if err != nil {
			return err
		}
		q.Exec(postData.DeviceID)
		return c.String(http.StatusNoContent, "Delete done")
	}

	// sqlにつっこむ
	now := time.Now()
	nowStr := fmt.Sprintf("%s", now.Format("2006-01-02 15:04:05"))

	ins, err := db.Prepare("INSERT INTO yuki_data(device_id, points, date) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(postData.DeviceID, postData.HuitPoints, nowStr)
	if err != nil {
		return err
	}
	ins.Close()

	return c.String(http.StatusCreated, "OK")
}

// dailyの上位device10人分
func handleEachData(c echo.Context) error {
	var dataList []Data

	rawToday := time.Now()
	today := fmt.Sprintf("%s", rawToday.Format("2006-01-02"))
	err := db.Select(&dataList, "SELECT device_id, SUM(points) AS points, DATE(date) AS date FROM yuki_data WHERE DATE(date)= ? GROUP BY device_id ORDER BY SUM(points) DESC LIMIT 10", today)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dataList)
}

// deviceごとの今までの総計
func handleTotalData(c echo.Context) error {
	type Total struct {
		DeviceID string `json:"device_id" db:"device_id"`
		Points   uint64 `json:"points" db:"points"`
	}
	var totalDataList []Total

	err := db.Select(&totalDataList, "SELECT device_id, SUM(points) AS points FROM yuki_data GROUP BY device_id ORDER BY SUM(points) DESC LIMIT 10")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, totalDataList)
}

// dailyの全てのデバイスの上位40
func handleAllData(c echo.Context) error {
	var dataList []Data

	rawToday := time.Now()
	today := fmt.Sprintf("%s", rawToday.Format("2006-01-02"))
	err := db.Select(&dataList, "SELECT device_id, SUM(points) AS points, DATE(date) AS date FROM yuki_data WHERE DATE(date)= ? GROUP BY device_id ORDER BY SUM(points) DESC LIMIT 40", today)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dataList)
}

func handleTestEach(c echo.Context) error {
	data := []Data{
		{
			DeviceID: "XX9021",
			Points:   1234,
			Date:     "2000-01-01T01:00:00+09:00",
		},
		{
			DeviceID: "XX0832",
			Points:   5839,
			Date:     "2000-01-01T01:00:00+09:00",
		},
		{
			DeviceID: "XX0932",
			Points:   1111333,
			Date:     "2000-01-01T01:00:00+09:00",
		},
	}
	return c.JSON(http.StatusOK, data)
}

func handleTestTotal(c echo.Context) error {
	data := []Data{
		{
			DeviceID: "XX0921testtotal",
			Points:   28308,
			Date:     "2000-01-01T01:00:00+09:00",
		},
		{
			DeviceID: "XX0832testtotal",
			Points:   58393928,
			Date:     "2020-08-015T13:00:00+09:00",
		},
		{
			DeviceID: "XX0832testtotal",
			Points:   58393928,
			Date:     "2040-05-015T18:32:00+09:00",
		},
	}
	return c.JSON(http.StatusOK, data)
}
