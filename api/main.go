package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//Creo las constanstes para las rutas de la api

const (
	PhotoApi = "https://api.pexels.com/v1/"
	VideoApi = "https://api.pexels.com/videos/"
)

// defino la struct de client
type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

// definimos la funcion NewClient
func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

//definimos la structura de los resultados
type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int         `json:"id"`
	Width           int         `json:"width"`
	Height          int         `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	PhotographerId  int         `json:"photographer_id"`
	AvgColor        string      `json:"avg-color"`
	Src             PhotoSource `json:"src"`
	Alt             string      `json:"alt"`
}

type CuratedResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photo    []Photo `json:"photos"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

//creamos el método SearchPhotos

func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoApi+"search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)

	return &result, err
}

func (c *Client) CuratedPhotos(perPage, page int) (*CuratedResult, error) {
	url := fmt.Sprint(PhotoApi+"/curated?per_page=%d&page=%d", perPage, page)

	resp, err := c.requestDoWithAuth("GET", url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var result CuratedResult

	err = json.Unmarshal(data, &result)

	return &result, err
}

// método para validar la autenticación
func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	resp, err := c.hc.Do(req)

	if err != nil {
		return resp, err
	}

	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))

	if err != nil {
		return resp, nil
	} else {
		c.RemainingTimes = int32(times)
	}

	return resp, nil
}

func main() {

	//Esto debería estar en un archivo .env
	os.Setenv("PexelsToken", "563492ad6f9170000100000144d00e693ffd439c9820763ad87580c0")
	var TOKEN = os.Getenv("PexelsToken")

	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves", 15, 1)

	if err != nil {
		fmt.Errorf("Search error:%v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("Search result wrong")
	}
	fmt.Println(result)

}
