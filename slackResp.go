package main

import "encoding/json"

type slackResp struct {
	RespType    string   `json:"response_type"`
	Text        string   `json:"text"`
	Attachments []attach `json:"attachments"`
}

type attach struct {
	Fallback    string  `json:"fallback"`
	Color       string  `json:"color"`
	Pretext     string  `json:"pretext"`
	Author_name string  `json:"author_name"`
	Author_link string  `json:"author_link"`
	Author_icon string  `json:"author_icon"`
	Title       string  `json:"title"`
	Title_link  string  `json:"title_link"`
	text        string  `json:"text"`
	Fields      []field `json:"fields"`
	Image_url   string  `json:"image_url"`
	Thumb_url   string  `json:"thumb_url"`
	Footer      string  `json:"footer"`
	Footer_icon string  `json:"footer_icon"`
	Ts          string  `json:"ts"`
}

type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short string `json:"short"`
}

func (s *slackResp) String() string {

	b, err := json.Marshal(s)
	if err != nil {
		panic("Unable to convert to string")
	}

	return string(b)
}
