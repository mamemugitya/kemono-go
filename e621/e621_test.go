package e621

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetPostByID(t *testing.T) {
	expect := []*Post{
		&Post{
			ID:               1,
			TagsString:       "test tags",
			LockedTagsString: "locked_tags",
			Tags:             []string{"test", "tags"},
			LockedTags:       []string{"locked_tags"},
			Description:      "This is a test case.",
			CreatedAt: CreatedAt{
				JSONClass: "Time",
				S:         12345,
				N:         67890,
			},
			CreatorID:     3,
			Author:        "OTY",
			Change:        12345,
			Source:        "https://test.net/12345.png",
			Score:         50,
			FavCount:      135,
			Md5:           "098f6bcd4621d373cade4e832627b4f6",
			FileSize:      70000,
			FileURL:       "https://testfile.net/12345.png",
			FileExt:       PNG,
			PreviewURL:    "https://testpreview.net/12345.png",
			PreviewWidth:  150,
			PreviewHeight: 100,
			SampleURL:     "https://testsample.net/12345.png",
			SampleWidth:   800,
			SampleHeight:  500,
			Rating:        Explicit,
			Status:        Active,
			Width:         4500,
			Height:        3000,
			HasComments:   false,
			HasNotes:      false,
			HasChildren:   true,
			ChildrenPtr:   "13578",
			Children:      13578,
			ParentIDPtr:   "",
			ParentID:      NO_PARENT,
			Artist:        []string{"OTY"},
			Sources: []string{"https://sources1.net/source.png",
				"https://sources2.net/source.png",
				"https://sources3.net/source.png"},
			Delreason: "",
		},
		&Post{
			ID:               2,
			TagsString:       "test tags",
			LockedTagsString: "",
			Tags:             []string{"test", "tags"},
			LockedTags:       []string{""},
			Description:      "This is a test case.",
			CreatedAt: CreatedAt{
				JSONClass: "Time",
				S:         12345,
				N:         67890,
			},
			CreatorID:     3,
			Author:        "OTY",
			Change:        12345,
			Source:        "https://test.net/12345.png",
			Score:         50,
			FavCount:      135,
			Md5:           "098f6bcd4621d373cade4e832627b4f6",
			FileSize:      70000,
			FileURL:       "https://testfile.net/12345.png",
			FileExt:       PNG,
			PreviewURL:    "https://testpreview.net/12345.png",
			PreviewWidth:  150,
			PreviewHeight: 100,
			SampleURL:     "https://testsample.net/12345.png",
			SampleWidth:   800,
			SampleHeight:  500,
			Rating:        Explicit,
			Status:        Active,
			Width:         4500,
			Height:        3000,
			HasComments:   false,
			HasNotes:      false,
			HasChildren:   true,
			ChildrenPtr:   "13578",
			Children:      13578,
			ParentIDPtr:   "",
			ParentID:      NO_PARENT,
			Artist:        []string{"OTY"},
			Sources: []string{"https://sources1.net/source.png",
				"https://sources2.net/source.png",
				"https://sources3.net/source.png"},
			Delreason: "",
		},
	}

	files := []string{"../testdata/getpostid-1.json", "../testdata/getpostid-2.json"}

	for idx, file := range files {
		mux, mockServerURL := NewMockServer()
		client := NewTestClient(mockServerURL)
		handlePath := fmt.Sprintf("/post/show.json")
		mux.HandleFunc(handlePath, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, file)
		})

		const testNum = 1

		post, _ := client.GetPostByID(context.Background(), testNum)

		if !reflect.DeepEqual(post, expect[idx]) {
			t.Errorf("\n got=%+v, \n want=%+v", post, expect[idx])
		}
	}

}

func TestCheckMD5(t *testing.T) {
	expect := []*CheckMD5{
		&CheckMD5{
			Md5:    "098f6bcd4621d373cade4e832627b4f6",
			Exists: true,
			PostID: 12345,
		},
	}

	files := []string{"../testdata/checkmd5-1.json"}

	for idx, file := range files {

		mux, mockServerURL := NewMockServer()
		client := NewTestClient(mockServerURL)
		handlePath := fmt.Sprintf("/post/check_md5.json")
		mux.HandleFunc(handlePath, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, file)
		})

		const testMD5 = "098f6bcd4621d373cade4e832627b4f6"

		result, _ := client.CheckMD5(context.Background(), testMD5)

		if !reflect.DeepEqual(result, expect[idx]) {
			t.Errorf("\n got=%+v, \n want=%+v", result, expect[idx])
		}
	}
}
