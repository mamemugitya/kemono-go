package e621

import "strings"

type PostTagHistory struct {
	ID         int    `json:"id"`
	PostID     int    `json:"post_id"`
	CreatedAt  int    `json:"created_at"`
	TagsString string `json:"tags"`
	Tags       []string
	Source     string `json:"source"`
}

func (p *PostTagHistory) initializeFields() {
	p.Tags = strings.Split(p.TagsString, " ")
}
