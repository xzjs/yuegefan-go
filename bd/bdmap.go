package bd

import (
	"encoding/json"
	"errors"
	"fmt"

	"yuegefan/conf"

	"github.com/go-resty/resty/v2"
)

type BDMap struct {
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type DetailInfo struct {
	Price         string `json:"price"`
	OverallRating string `json:"overall_rating"`
}

type Result struct {
	Name       string     `json:"name"`
	Location   Location   `json:"location"`
	DetailInfo DetailInfo `json:"detail_info"`
}

type Response struct {
	Status  int64    `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"results"`
}

func (m *BDMap) Search(query string) (location Location, err error) {
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"query":  query,
		"region": "北京",
		"ak":     conf.GetConf().Map.AK,
		"output": "json",
	}).Get("https://api.map.baidu.com/place/v2/search")
	if err != nil {
		return
	}
	response := Response{}
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return
	}
	if response.Status != 0 {
		err = errors.New(response.Message)
		return
	}
	if len(response.Results) == 0 {
		err = errors.New("未查询到相关结果")
		return
	}
	location = response.Results[0].Location
	return
}

func (m *BDMap) SearchCircle(location Location) (results []Result, err error) {
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"query":    "美食",
		"tag":      "美食",
		"location": fmt.Sprintf("%f,%f", location.Lat, location.Lng),
		"ak":       conf.GetConf().Map.AK,
		"output":   "json",
		"scope":    "2",
		"filter":   "industry_type:cater|sort_name:overall_rating",
	}).Get("https://api.map.baidu.com/place/v2/search")
	if err != nil {
		return
	}
	response := Response{}
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return
	}
	if response.Status != 0 {
		err = errors.New(response.Message)
		return
	}
	if len(response.Results) == 0 {
		err = errors.New("未查询到相关结果")
		return
	}
	results = response.Results
	return
}
