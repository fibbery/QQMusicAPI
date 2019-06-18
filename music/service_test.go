package music

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		mediaType MediaType
		key       string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "搜索音乐名", args: args{SONG, "林俊杰"}},
		{name: "搜索人名", args: args{SINGER, "林俊杰"}},
		{name: "搜索MV", args: args{MV, "林俊杰"}},
		{name: "搜索专辑", args: args{ALBUM, "林俊杰"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := Search(tt.args.mediaType, tt.args.key)
			data, _ :=json.MarshalIndent(item,"","    ")
			fmt.Printf("%s\n", data)
		})
	}
}
