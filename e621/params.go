package e621

type ListParams struct {
	Limit    int
	BeforeId int
	Page     int
	Tags     string
}

type PostTagHistoryParams struct {
	PostId      uint
	DateStart   string
	DateEnd     string
	UserId      uint
	UserName    string
	Source      string
	Tags        string
	Reason      string
	Description string
	Before      string
	After       string
}
