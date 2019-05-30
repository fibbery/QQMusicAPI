package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
获取歌曲的media id
*/
func searchMediaId(sing string, author string) string {
	params := url.Values{
		"key": {sing,},
	}
	data := ""
	for key, value := range params {
		data += key + "=" + value[0] + "&"
	}
	request, _ := http.NewRequest("GET", basePath+searchMediaUrl+"?"+data, nil)
	client := &http.Client{}
	response, e := client.Do(request)
	if e != nil {
		fmt.Printf("get media id has occur error , %v", e)
		return ""
	}
	bytes, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Printf("this response is %s", string(bytes))
	//TODO 解析获取的数据
	return ""
}

/**
添加到我的歌单
*/
func addToMyPlayList(dirId int, mediaId string) {
	//todo 请求需要cookie
	params := url.Values{
		"format":      {"json",},
		"inCharset":   {"utf8",},
		"outCharset":  {"utf-8",},
		"notice":      {"0",},
		"platform":    {"yqq.post",},
		"needNewCode": {"0",},
		"uin":         {"1152921504755912068",},
		"midlist":     {"003meFNe1QWeUA",},
		"typelist":    {"13",},
		"dirid":       {"1",},
		"addtype":     {},
		"formsender":  {"4",},
		"source":      {"153",},
		"r2":          {"0",},
		"r3":          {"1",},
		"utf8":        {"1",},
		"g_tk":        {"5381",},
	}
	response, e := http.PostForm(basePath+addPlaylistUrl, params)
	if e != nil {
		fmt.Printf("get media id has occur error , %v", e)
	}
	bytes, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Printf("this response is %s", string(bytes))
}
