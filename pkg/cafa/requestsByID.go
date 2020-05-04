package cafa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type MoveRequest struct {
	LevelID string `json:"level_id"`
}

type UpdateByIDRequest struct {
	Value interface{}
}

func (c *Client) MoveByID(id int, headers map[string]string, req *MoveRequest) (*Response, error) {
	reqPath := fmt.Sprintf("%s/%d", requestPath, id)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling move request failed")
	}
	bodyReader := bytes.NewReader(body)
	return c.baseRequestByIDWithBody(headers, reqPath, httpPut, bodyReader)
}

func (c *Client) GetByID(id int, headers map[string]string) (*ConfigurationResponse, error) {
	reqPath := fmt.Sprintf("%s/%d", requestPathV3, id)
	requestUrl := c.baseUrl + reqPath

	req, err := http.NewRequest(httpGet, requestUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build HTTP request")
	}

	for name, value := range headers {
		req.Header.Add(name, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	res := &ConfigurationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling response failed with status %d", resp.StatusCode))
	}

	return res, nil
}
func (c *Client) DeleteByID(id int, headers map[string]string) (*Response, error) {
	reqPath := fmt.Sprintf("%s/%d", requestPathV3, id)
	return c.baseRequestByID(headers, reqPath, httpDelete)
}

func (c *Client) UpdateByID(id int, headers map[string]string, req *UpdateByIDRequest) (*Response, error) {
	reqPath := fmt.Sprintf("%s/%d", requestPathV3, id)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling move request failed")
	}
	bodyReader := bytes.NewReader(body)
	return c.baseRequestByIDWithBody(headers, reqPath, httpPut, bodyReader)
}
