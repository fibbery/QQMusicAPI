package api

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/**
	处理从网站 https://www.tunemymusic.com/zh/Spotify-to-File.php#step3 导出的 txt 歌单
	map 中key是歌名  value是作者
*/
func parseTxt(filename string) map[string]string {
	isTxt := strings.HasSuffix(filename, ".txt")
	result := make(map[string]string)
	if !isTxt {
		return result
	}
	file, e := os.Open(filename)
	if e != nil {
		fmt.Printf("read file %s error, %v", filename, e)
		return result
	}
	reader := bufio.NewReader(file)
	//需要替换（）
	reg, _ := regexp.Compile("\\([\\s\\S]*\\)")
	for {
		item, e := reader.ReadString('\n')
		if e == io.EOF {
			break
		}
		item = strings.TrimSuffix(item, "\r\n")
		if item = strings.TrimSpace(item); item == "" {
			continue
		}
		temp := strings.Split(item, " - ")
		sing := reg.ReplaceAllString(temp[1], "")
		author := reg.ReplaceAllString(temp[0], "")
		result[sing] = author
	}
	return result
}