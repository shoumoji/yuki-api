package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type RawData struct {
	DeviceID   string `json:"device-id"`
	HuitPoints string `json:"huit-points"`
}

type Data struct {
	DeviceID string `json:"device-id"`
	Points   uint64 `json:"points"`
	Date     string `json:"date"`
}

func handlePOST(c echo.Context) error {
	postData := new(RawData)
	if err := c.Bind(postData); err != nil {
		return err
	}
	// sqlに入れる
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
