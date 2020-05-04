package cafa

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) baseRequestWithBody(conf *Configuration, headers map[string]string, requestPath, httpMethod string) (*Response, error) {
	requestUrl := c.baseUrl + requestPath
	body, err := json.Marshal(conf)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling configuration failed")
	}
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(httpMethod, requestUrl, bodyReader)
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
	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		res.Message = "unmarshaling response failed"
		res.StatusCode = strconv.Itoa(resp.StatusCode)
	}

	return res, nil
}

func (c *Client) baseRequestWithQueryParameters(conf *Configuration, headers map[string]string, requestPath, httpMethod string) (*Response, error) {
	requestUrl := c.baseUrl + requestPath

	params := url.Values{}
	params.Add("application", conf.Application)
	params.Add("level", conf.Level)
	params.Add("level_id", conf.LevelID)
	params.Add("type", conf.Type)
	params.Add("name", conf.Name)

	req, err := http.NewRequest(httpMethod, requestUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build HTTP request")
	}
	req.URL.RawQuery = params.Encode()

	for name, value := range headers {
		req.Header.Add(name, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		res.Message = "unmarshaling response failed"
		res.StatusCode = strconv.Itoa(resp.StatusCode)
	}

	return res, nil
}

func (c *Client) baseRequestByID(headers map[string]string, requestPath, httpMethod string) (*Response, error) {
	requestUrl := c.baseUrl + requestPath
	req, err := http.NewRequest(httpMethod, requestUrl, nil)
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

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		res.Message = "unmarshaling response failed"
		res.StatusCode = strconv.Itoa(resp.StatusCode)
	}

	return res, nil
}

func (c *Client) baseRequestByIDWithBody(headers map[string]string, requestPath, httpMethod string, body io.Reader) (*Response, error) {
	requestUrl := c.baseUrl + requestPath
	req, err := http.NewRequest(httpMethod, requestUrl, body)
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

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		res.Message = "unmarshaling response failed"
		res.StatusCode = strconv.Itoa(resp.StatusCode)
	}

	return res, nil
}
