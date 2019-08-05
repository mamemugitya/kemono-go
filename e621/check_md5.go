package e621

const (
	NO_POST int = -1
)

type CheckMD5 struct {
	Md5    string `json:"md5"`
	Exists bool   `json:"exists"`
	PostID int    `json:"post_id"` // This field is valied only if the post exists.
}

func (c *CheckMD5) initializeFields() {
	if c.Exists == false {
		c.PostID = NO_POST
	}
}
