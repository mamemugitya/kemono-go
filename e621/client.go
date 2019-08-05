package e621

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	UserName   string
	Logger     *log.Logger
}

func NewClient(baseURL, userName string, logger *log.Logger) (*Client, error) {
	CheckedBaseURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    CheckedBaseURL,
		HTTPClient: http.DefaultClient,
		UserName:   userName,
		Logger:     logger,
	}, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	var p []byte
	res.Body.Read(p)
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

//post/show
func (c *Client) GetPostByID(ctx context.Context, id uint) (*Post, error) {
	subURL := "/post/show.json"
	query := map[string]string{
		"id": fmt.Sprintf("%d", id),
	}
	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var out Post
	err = decodeBody(res, &out)
	out.initializeFields()
	return &out, err
}

func (c *Client) GetPostByMD5(ctx context.Context, md5 string) (*Post, error) {
	subURL := "/post/show.json"
	query := map[string]string{
		"md5": fmt.Sprintf("%s", md5),
	}
	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var out Post
	err = decodeBody(res, &out)
	out.initializeFields()
	return &out, err
}

func (c *Client) CheckMD5(ctx context.Context, md5 string) (*CheckMD5, error) {
	subURL := "/post/check_md5.json"
	query := map[string]string{
		"md5": fmt.Sprintf("%s", md5),
	}
	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var out CheckMD5
	err = decodeBody(res, &out)
	out.initializeFields()
	return &out, err
}

func (c *Client) List(ctx context.Context, p ListParams) (*[]Post, error) {
	subURL := "/post/index.json"
	query := make(map[string]string)

	if p.Limit != 0 {
		query["limit"] = fmt.Sprintf("%d", p.Limit)
	}

	if p.BeforeId != 0 {
		query["before_id"] = fmt.Sprintf("%d", p.BeforeId)
	}

	if p.Page != 0 {
		query["page"] = fmt.Sprintf("%d", p.Page)
	}

	if p.Tags != "" {
		query["tags"] = p.Tags
	}

	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var out []Post

	err = decodeBody(res, &out)
	for idx, _ := range out {
		out[idx].initializeFields()
	}
	return &out, err
}

func (c *Client) ListWithTypedTags(ctx context.Context, p ListParams) (*[]TypedTagsPost, error) {
	subURL := "/post/index.json"
	query := make(map[string]string)
	query["typed_tags"] = "true"

	if p.Limit != 0 {
		query["limit"] = fmt.Sprintf("%d", p.Limit)
	}

	if p.BeforeId != 0 {
		query["before_id"] = fmt.Sprintf("%d", p.BeforeId)
	}

	if p.Page != 0 {
		query["page"] = fmt.Sprintf("%d", p.Page)
	}

	if p.Tags != "" {
		query["tags"] = p.Tags
	}

	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var out []TypedTagsPost

	err = decodeBody(res, &out)
	for idx, _ := range out {
		out[idx].initializeFields()
	}
	return &out, err
}

func (c *Client) Popular(ctx context.Context, interval string) (*[]Post, error) {
	subURL := "/post/"
	if interval == "day" || interval == "d" {
		subURL = subURL + "popular_by_day.json"
	} else if interval == "week" || interval == "w" {
		subURL = subURL + "popular_by_week.json"
	} else if interval == "month" || interval == "m" {
		subURL = subURL + "popular_by_month.json"
	} else {
		return nil, errors.New("No matched interval.")
	}

	req, err := c.NewRequest(context.Background(), subURL, nil, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var out []Post
	err = decodeBody(res, &out)
	for idx, _ := range out {
		out[idx].initializeFields()
	}
	return &out, err
}

func (c *Client) PostTagHistory(ctx context.Context, p PostTagHistoryParams) (*[]PostTagHistory, error) {
	subURL := "post_tag_history/index.json"
	query := make(map[string]string)

	if p.PostId != 0 {
		query["post_id"] = fmt.Sprintf("%d", p.PostId)
	}

	if p.DateStart != "" {
		query["date_start"] = p.DateStart
	}

	if p.DateEnd != "" {
		query["date_end"] = p.DateEnd
	}

	if p.UserId != 0 {
		query["user_id"] = fmt.Sprintf("%d", p.UserId)
	}

	if p.UserName != "" {
		query["user_name"] = p.UserName
	}

	if p.Source != "" {
		query["source"] = p.Source
	}

	if p.Tags != "" {
		query["tags"] = p.Tags
	}

	if p.Reason != "" {
		query["reason"] = p.Reason
	}

	if p.Description != "" {
		query["description"] = p.Description
	}

	if p.Before != "" {
		query["before"] = p.Before
	}

	if p.After != "" {
		query["after"] = p.After
	}

	req, err := c.NewRequest(context.Background(), subURL, query, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var out []PostTagHistory

	err = decodeBody(res, &out)
	for idx, _ := range out {
		out[idx].initializeFields()
	}
	return &out, err
}

func (client *Client) NewRequest(ctx context.Context, subURL string, query map[string]string, method string, body io.Reader) (*http.Request, error) {
	endpointURL := *client.BaseURL
	endpointURL.Path = path.Join(client.BaseURL.Path, subURL)
	if query != nil {
		q := endpointURL.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		endpointURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.WithContext(ctx)
	userAgent := fmt.Sprintf("kemo-go/1.0 by otyaken (Username:%s)", client.UserName)
	req.Header.Set("User-Agent", userAgent)

	return req, nil

}
