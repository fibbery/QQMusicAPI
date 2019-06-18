package main

import (
	"github.com/fibbery/qqmusic/music"
	"encoding/json"
	"fmt"
)

func main() {
	response := music.SearchResponse{
		CommonResponse: music.CommonResponse{1, 1},
		Data: struct {
			Singers *music.MediaCollection `json:"singer,omitempty"`
			Mvs     *music.MediaCollection `json:"mv,omitempty"`
			Albums  *music.MediaCollection `json:"album,omitempty"`
			Songs   *music.MediaCollection `json:"song,omitempty"`
		}{
			Singers: &music.MediaCollection{
				Type: int(music.SONG),
				Items: &[]music.Media{
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
