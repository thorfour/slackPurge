package main

import "encoding/json"

type slackResp struct {
	RespType    string   `json:"response_type"`
	Text        string   `json:"text"`
	Attachments []attach `json:"attachments"`
}

type attach struct {
	Fallback    string   `json:"fallback"`
	Color       string   `json:"color"`
	Pretext     string   `json:"pretext"`
	Author_name string   `json:"author_name"`
	Author_link string   `json:"author_link"`
	Author_icon string   `json:"author_icon"`
	Title       string   `json:"title"`
	Title_link  string   `json:"title_link"`
	Text        string   `json:"text"`
	Fields      []field  `json:"fields"`
	Image_url   string   `json:"image_url"`
	Thumb_url   string   `json:"thumb_url"`
	Footer      string   `json:"footer"`
	Footer_icon string   `json:"footer_icon"`
	Ts          string   `json:"ts"`
	Actions     []action `json:"actions"`
}

type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short string `json:"short"`
}

type action struct {
	Name    string  `json:"name"`
	Text    string  `json:"text"`
	Style   string  `json:"style"`
	Type    string  `json:"type"`
	Value   string  `json:"value"`
	Confirm confirm `json:"confirm"`
}

type confirm struct {
	Title        string `json:"title"`
	Text         string `json:"text"`
	Ok_text      string `json:"ok_text"`
	Dismiss_text string `json:"dismiss_text"`
}

func (s *slackResp) String() string {

	b, err := json.Marshal(s)
	if err != nil {
		panic("Unable to convert to string")
	}

	return string(b)
}
