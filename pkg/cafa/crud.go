package cafa

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

//put needed headers in a map, conf to save. in response you will get the response
func (c *Client) Save(conf *Configuration, headers map[string]string) (*Response, error) {
	return c.baseRequestWithBody(conf, headers, requestPath, httpPost)
}

func (c *Client) Update(conf *Configuration, headers map[string]string) (*Response, error) {
	return c.baseRequestWithBody(conf, headers, requestPath, httpPut)
}

func (c *Client) Get(filters *Configuration, headers map[string]string) ([]ConfigurationResponse, error) {
	requestUrl := c.baseUrl + requestPath

	params := url.Values{}
	params.Add("application", filters.Application)
	params.Add("level", filters.Level)
	params.Add("level_id", filters.LevelID)
	params.Add("type", filters.Type)
	params.Add("name", filters.Name)

	req, err := http.NewRequest(httpGet, requestUrl, nil)
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

	if resp.StatusCode != http.StatusOK {
		var res Response

		err := json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling response failed with status %d", resp.StatusCode))
		}

		return nil, fmt.Errorf("error performing GET request: %s (status code %s)", res.Message, res.StatusCode)
	}

	var res []ConfigurationResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling response failed with status %d", resp.StatusCode))
	}

	return res, nil

}

func (c *Client) Delete(filters *Configuration, headers map[string]string) (*Response, error) {
	return c.baseRequestWithQueryParameters(filters, headers, requestPath, httpDelete)
}
