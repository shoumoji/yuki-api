package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type RawData struct {
	DeviceID   string `json:"device_id"`
	HuitPoints uint64 `json:"huit_points"`
}

type Data struct {
	DeviceID string `json:"device_id"`
	Points   uint64 `json:"points"`
	Date     string `json:"date"`
}

func handlePOST(c echo.Context) (err error) {
	postData := new(RawData)
	if err = c.Bind(postData); err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}

	if len(postData.DeviceID) >= 40 {
		// デバイスIDデカすぎ
		return c.String(http.StatusServiceUnavailable, "")
	}

	// sqlにつっこむ
	now := time.Now()
	nowStr := fmt.Sprintf("%s", now.Format("2006-01-02T15:04:05+07:00"))

	ins, err := db.Prepare("INSERT INTO yuki_data(device_id, points, date) VALUES(?,?,?)")
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	ins.Exec(postData.DeviceID, postData.HuitPoints, nowStr)
	return c.String(http.StatusCreated, "")
}

// func handleEachData(c echo.Context) error {
// 	return
// }
//
// func handleTotalData(c echo.Context) error {
//
// }

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
