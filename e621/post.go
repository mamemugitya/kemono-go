package e621

import (
	"strconv"
	"strings"
)

const (
	NO_PARENT   int = -1
	NO_CHILDREN int = -1
)

type Rating string

const (
	Safe         Rating = "s"
	Questionable        = "q"
	Explicit            = "e"
)

type FileExt string

const (
	JPG  FileExt = "jpg"
	PNG          = "png"
	GIF          = "gif"
	SWF          = "swf"
	WEBM         = "webm"
)

type Status int

const (
	Active Status = iota
	Flagged
	Pending
	Deleted
)

//Note: the following fields are not available for deleted posts: source, sources, md5, file_size, file_ext, preview_width, preview_height, sample_url, sample_width, sample_height, has_children, children
type Post struct {
	ID               int    `json:"id"`
	TagsString       string `json:"tags"`
	LockedTagsString string `json:"locked_tags"`
	Tags             []string
	LockedTags       []string
	Description      string    `json:"description"`
	CreatedAt        CreatedAt `json:"created_at"`
	CreatorID        int       `json:"creator_id"`
	Author           string    `json:"author"`
	Change           int       `json:"change"`
	Source           string    `json:"source"`
	Score            int       `json:"score"`
	FavCount         int       `json:"fav_count"`
	Md5              string    `json:"md5"`
	FileSize         int       `json:"file_size"`
	FileURL          string    `json:"file_url"`
	FileExt          FileExt   `json:"file_ext"`
	PreviewURL       string    `json:"preview_url"`
	PreviewWidth     int       `json:"preview_width"`
	PreviewHeight    int       `json:"preview_height"`
	SampleURL        string    `json:"sample_url"`
	SampleWidth      int       `json:"sample_width"`
	SampleHeight     int       `json:"sample_height"`
	Rating           Rating    `json:"rating"`
	Status           Status    `json:"status"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	HasComments      bool      `json:"has_comments"`
	HasNotes         bool      `json:"has_notes"`
	HasChildren      bool      `json:"has_children"`
	ChildrenPtr      string    `json:"children"`  //null or id
	ParentIDPtr      string    `json:"parent_id"` //null or id
	Children         int
	ParentID         int
	Artist           []string `json:"artist"`
	Sources          []string `json:"sources"`
	Delreason        string   `json:"delreason"`
}

type TypedTagsPost struct {
	ID               int       `json:"id"`
	LockedTagsString string    `json:"locked_tags"`
	Tags             TypedTags `json:"tags"`
	LockedTags       []string
	Description      string    `json:"description"`
	CreatedAt        CreatedAt `json:"created_at"`
	CreatorID        int       `json:"creator_id"`
	Author           string    `json:"author"`
	Change           int       `json:"change"`
	Source           string    `json:"source"`
	Score            int       `json:"score"`
	FavCount         int       `json:"fav_count"`
	Md5              string    `json:"md5"`
	FileSize         int       `json:"file_size"`
	FileURL          string    `json:"file_url"`
	FileExt          FileExt   `json:"file_ext"`
	PreviewURL       string    `json:"preview_url"`
	PreviewWidth     int       `json:"preview_width"`
	PreviewHeight    int       `json:"preview_height"`
	SampleURL        string    `json:"sample_url"`
	SampleWidth      int       `json:"sample_width"`
	SampleHeight     int       `json:"sample_height"`
	Rating           Rating    `json:"rating"`
	Status           Status    `json:"status"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	HasComments      bool      `json:"has_comments"`
	HasNotes         bool      `json:"has_notes"`
	HasChildren      bool      `json:"has_children"`
	ChildrenPtr      string    `json:"children"`  //null or id
	ParentIDPtr      string    `json:"parent_id"` //null or id
	Children         int
	ParentID         int
	Artist           []string `json:"artist"`
	Sources          []string `json:"sources"`
	Delreason        string   `json:"delreason"`
}

type CreatedAt struct {
	JSONClass string `json:"json_class"`
	S         int    `json:"s"`
	N         int    `json:"n"`
}

type TypedTags struct {
	General   []string `json:"general"`
	Artist    []string `json:"artist"`
	CopyRight []string `json:"copyright"`
	Character []string `json:"character"`
	Species   []string `json:"species"`
}

func (f FileExt) String() string {
	switch f {
	case JPG:
		return "JPG"
	case PNG:
		return "PNG"
	case GIF:
		return "GIF"
	case SWF:
		return "SWF"
	case WEBM:
		return "WEBM"
	default:
		return "Unknown"
	}
}

func (r Rating) String() string {
	switch r {
	case Safe:
		return "Safe"
	case Questionable:
		return "Questionable"
	case Explicit:
		return "Explicit"
	default:
		return "Unknown"
	}
}

func (s Status) String() string {
	switch s {
	case Active:
		return "Active"
	case Flagged:
		return "Flagged"
	case Pending:
		return "Pending"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func (p *Post) initializeFields() {
	p.Tags = strings.Split(p.TagsString, " ")
	p.LockedTags = strings.Split(p.LockedTagsString, " ")

	if p.ChildrenPtr != "" {
		p.Children, _ = strconv.Atoi(p.ChildrenPtr)
	} else {
		p.Children = NO_CHILDREN
	}

	if p.ParentIDPtr != "" {
		p.ParentID, _ = strconv.Atoi(p.ParentIDPtr)
	} else {
		p.ParentID = NO_PARENT
	}
}

func (p *TypedTagsPost) initializeFields() {
	p.LockedTags = strings.Split(p.LockedTagsString, " ")

	if p.ChildrenPtr != "" {
		p.Children, _ = strconv.Atoi(p.ChildrenPtr)
	} else {
		p.Children = NO_CHILDREN
	}

	if p.ParentIDPtr != "" {
		p.ParentID, _ = strconv.Atoi(p.ParentIDPtr)
	} else {
		p.ParentID = NO_PARENT
	}
}
