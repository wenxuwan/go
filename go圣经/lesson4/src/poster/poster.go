package poster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const URL = "https://omdbapi.com/?t="
const KEY = "&apikey=ee961568"

type picUrl struct {
	PICTUREURL string `json:"Poster"`
}

func GetPicUrl1(movi_name string) (picurl *picUrl, err error) {
	resp, err := http.Get(URL + movi_name + KEY)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&picurl); err != nil {
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	fmt.Println(picurl.PICTUREURL, movi_name)
	SaveImage(picurl.PICTUREURL, movi_name)
	return
}

func GetPicUrl(movi_name string) (*picUrl, error) {
	resp, err := http.Get(URL + movi_name + KEY)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result picUrl
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	fmt.Println(result.PICTUREURL, movi_name)
	SaveImage(result.PICTUREURL, movi_name)
	return &result, nil
}

func SaveImage(picUrl string, moviName string) {
	picName := moviName + ".jpg"
	image, err := os.Create(picName)
	if err != nil {
		log.Println("create file failed:", picName, err)
		return
	}
	response, err := http.Get(picUrl)
	if err != nil {
		log.Println("Get picture failed:", picUrl, err)
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read file failed")
	}
	image.Write(data)
	image.Close()
}
