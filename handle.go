package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type RawData struct {
	deviceID   string `json:device-id`
	huitPoints string `json:huit-points`
}

type Data struct {
	deviceID string `json:device-id`
	points   uint64 `json:points`
	date     string `json:date`
}

func handlePOST(c echo.Context) error {
	postData := new(RawData)
	if err := c.Bind(postData); err != nil {
		return err
	}
	// sqlに入れる
}

func handleEachData(c echo.Context) error {

}

func handleTotalData(c echo.Context) error {

}

func handleTestEach(c echo.Context) error {
	data := []Data{
		{
			deviceID: "XX9021",
			points:   1234,
			date:     "2000-01-01T01:00:00+09:00",
		},
		{
			deviceID: "XX0832",
			points:   5839,
			date:     "2000-01-01T01:00:00+09:00",
		},
		{
			deviceID: "XX0932",
			points:   1111333,
			date:     "2000-01-01T01:00:00+09:00",
		},
	}
	c.JSON(http.StatusOK, data)
}

func handleTestTotal(c echo.Context) error {
	data := []Data{
		{
			deviceID: "XX0921testtotal",
			points:   28308,
			date:     "2000-01-01T01:00:00+09:00",
		},
		{
			deviceID: "XX0832testtotal",
			points:   58393928,
			date:     "2020-08-015T13:00:00+09:00",
		},
		{
			deviceID: "XX0832testtotal",
			points:   58393928,
			date:     "2040-05-015T18:32:00+09:00",
		},
	}
	return c.JSON(http.StatusOK, data)
}
