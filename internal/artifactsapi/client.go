package artifactsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
)

const (
	MAX_RETRY_COUNT = 3
)

type Client struct {
	client *http.Client
	config *Config
}

func NewClient(config *Config) *Client {
	client := &http.Client{}

	return &Client{
		client: client,
		config: config,
	}
}

func (c *Client) Get(path string, queries map[string]string, response dto.Response) (dto.StatusCode, error) {
	return c.doRequest(
		http.MethodGet,
		path,
		queries,
		nil,
		response,
	)
}

func (c *Client) Post(path string, body dto.Request, response dto.Response) (dto.StatusCode, error) {
	return c.doRequest(
		http.MethodPost,
		path,
		map[string]string{},
		body,
		response,
	)
}

func (c *Client) constructURL(path string) url.URL {
	return url.URL{
		Scheme: "https",
		Host:   c.config.Host(),
		Path:   path,
	}
}

func (c *Client) prepareHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.Token()))
}

func (c *Client) doRequest(
	method string,
	path string,
	queries map[string]string,
	body dto.Request,
	response dto.Response,
) (int, error) {
	url := c.constructURL(path)
	query := url.Query()
	for key, value := range queries {
		query.Set(key, value)
	}
	url.RawQuery = query.Encode()
	bodyBytes, err := json.Marshal(body)

	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return 0, err
	}
	c.prepareHeaders(req)

	statusCode := 0
	round := 0

	for {
		res, err := c.client.Do(req)
		if err != nil {
			if round < MAX_RETRY_COUNT {
				time.Sleep(500 * time.Millisecond)
				round++
				continue
			} else {
				return 0, err
			}
		}
		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}

		if err := json.Unmarshal(resBody, response); err != nil {
			return 0, err
		}

		statusCode = res.StatusCode
		break
	}

	return statusCode, nil
}
