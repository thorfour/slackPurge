package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSlackRespStringer(t *testing.T) {

	s := new(slackResp)
	s.Text = "Hello Dave"
	s.Attachments = make([]attach, 1)
	s.Attachments[0].Color = "blue"
	s.Attachments[0].text = "I can't do that Dave"

	resp := fmt.Sprint(s)

	exp := "{\"response_type\":\"\",\"text\":\"Hello Dave\",\"attachments\":[{\"fallback\":\"\",\"color\":\"blue\",\"pretext\":\"\",\"author_name\":\"\",\"author_link\":\"\",\"author_icon\":\"\",\"title\":\"\",\"title_link\":\"\",\"fields\":null,\"image_url\":\"\",\"thumb_url\":\"\",\"footer\":\"\",\"footer_icon\":\"\",\"ts\":\"\"}]}"

	if strings.Compare(resp, exp) != 0 {
		t.Fail()
	}
}
