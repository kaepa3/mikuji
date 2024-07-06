package uranai

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/kaepa3/mikuji/mikuji/seiza"
)

const (
	baseUrL = "https://api.jugemkey.jp/api/horoscope/free/"
)

func Today() (*HoroScope, error) {

	url, err := createUrL()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := parseJson(resp)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

type UranaiResult struct {
	Content string      `json:"content"`
	Item    string      `json:"item"`
	Money   int         `json:"money"`
	Total   int         `json:"total"`
	Job     int         `json:"job"`
	Color   string      `json:"color"`
	Day     interface{} `json:"day"`
	Love    int         `json:"love"`
	Rank    int         `json:"rank"`
	Sign    string      `json:"sign"`
}

func (u *UranaiResult) GetArrayResult() []interface{} {
	rv := reflect.ValueOf(*u)
	num := rv.NumField()
	list := make([]interface{}, num)
	for idx := 0; idx < num; idx++ {
		val := rv.Field(idx)
		list[idx] = val.Interface()
	}
	return list
}

type HoroScope struct {
	HoroScope map[string][]UranaiResult `json:"horoscope"`
}

func (h *HoroScope) GetRecord(sz seiza.Seiza) *UranaiResult {
	for _, v := range h.HoroScope {
		for _, val := range v {
			name := seiza.IndexToString(sz)
			if val.Sign == name {
				return &val
			}
		}
	}
	return nil
}

func getData(val []byte) (*HoroScope, error) {
	var data HoroScope // nil slice
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func parseJson(r *http.Response) (*HoroScope, error) {
	// data := make([]Item, 0) のように要素数0の slice としても良い
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	val, err := getData(body)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func createUrL() (string, error) {
	now := time.Now()
	yearStr := strconv.Itoa(now.Year())
	dayStr := strconv.Itoa(now.Day())
	callURL, err := url.JoinPath(baseUrL, yearStr, dayStr)
	if err != nil {
		return "", err
	}
	return callURL, nil
}
