package music

type MediaType int

const (
	SONG MediaType = iota + 1
	SINGER
	ALBUM
	MV
)

type CommonResponse struct {
	Code    int `json:"code"`
	Subcode int `json:"subcode"`
}

type Media struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//资源Id，实际起搜索作用的
	MediaId string `json:"mid"`
	Singer  string `json:"singer"`
	DocId   string `json:"docid"`
	PicUrl  string `json:"pic"`
}

type MediaCollection struct {
	Type  int      `json:"type"`
	Name  string   `json:"name"`
	Order int      `json:"order"`
	Count int      `json:"count"`
	Items *[]Media `json:"itemlist"`
}

type SearchResponse struct {
	CommonResponse
	Data struct {
		Singers *MediaCollection `json:"singer,omitempty"`
		Mvs     *MediaCollection `json:"mv,omitempty"`
		Albums  *MediaCollection `json:"album,omitempty"`
		Songs   *MediaCollection `json:"song,omitempty"`
	} `json:"data,omitempty"`
}
