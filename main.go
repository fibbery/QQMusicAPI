package main

import (
	"QQMusicAPI/api"
	"encoding/json"
	"fmt"
)

func main() {
	response := api.SearchResponse{
		CommonResponse: api.CommonResponse{1, 1},
		Data: struct {
			Singers *api.MediaCollection `json:"singer,omitempty"`
			Mvs     *api.MediaCollection `json:"mv,omitempty"`
			Albums  *api.MediaCollection `json:"album,omitempty"`
			Songs   *api.MediaCollection `json:"song,omitempty"`
		}{
			Singers: &api.MediaCollection{
				Type: int(api.SONG),
				Items: &[]api.Media{
					{
						Id:      "62342",
						Name:    "他说",
						DocId:   "62342",
						Singer:  "linjunjie",
						MediaId: "000TuW8h0AH2n4",
						PicUrl:  "http://www.baidu.com",
					},
				},
			},
		},
	}

	data, _ := json.Marshal(response)
	fmt.Printf("%s", data)

}
