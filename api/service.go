package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func Search(mediaType MediaType, key string) *MediaCollection {
	params := url.Values{
		"key": {key},
	}
	result := new(SearchResponse)
	http_get(searchMediaUrl, params, result)
	if result.Code != 0 {
		log.Printf("execute %s fail, the code is %d, the subcode is %d", searchMediaUrl, result.Code, result.Subcode)
		return nil
	}
	switch mediaType {
	case SONG:
		return result.Data.Songs
	case SINGER:
		return result.Data.Singers
	case ALBUM:
		return result.Data.Albums
	case MV:
		return result.Data.Mvs
	default:
		return nil
	}
}

/**
封装get请求
*/
func http_get(method string, params url.Values, rep interface{}) {
	var data string
	for key, value := range params {
		data += key + "=" + value[0] + "&"
	}
	request, _ := http.NewRequest("GET", basePath+searchMediaUrl+"?"+data, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("get media id has occur error , %v", err)
		rep = nil
		return
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(rep); err != nil {
		log.Printf("unmarshalling response fail : %s \n", err)
	}
}

/**
封装post请求
*/
func http_post(method string, params url.Values, rep interface{}) {
	response, e := http.PostForm(basePath+method, params)
	if e != nil {
		fmt.Printf("get media id has occur error , %v", e)
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(rep); err != nil {
		log.Printf("unmarshalling response fail : %s \n", e)
	}
}
