package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ipApi struct {
	Client *http.Client
}

// MyIP return public ip address of current machine
func (ia *ipApi) MyIP() (ip string, err error) {

	resp, err := ia.Client.Get("https://api.test.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	infos := make(map[string]string)
	err = json.Unmarshal(body, &infos)
	if err != nil {
		return "", err
	}

	ip, ok := infos["ip"]
	if !ok {
		return "", fmt.Errorf("invalid response result")
	}
	return ip, nil
}

func TestMyIP(t *testing.T) {
	tests := []struct {
		code     int
		text     string
		ip       string
		hasError bool
	}{
		{code: 200, text: "{\"ip\":\"1.2.3.4\"}", ip: "1.2.3.4", hasError: false},
		{code: 403, text: "", ip: "", hasError: true},
		{code: 200, text: "abcd", ip: "", hasError: true},
	}

	for row, test := range tests {
		client := NewTestHttpClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: test.code,
				Body:       ioutil.NopCloser(bytes.NewBufferString(test.text)),
				Header:     make(http.Header),
			}
		})
		api := &ipApi{Client: client}

		ip, err := api.MyIP()
		if test.hasError {
			assert.Error(t, err, "row %d", row)
		} else {
			assert.NoError(t, err, "row %d", row)
		}
		assert.Equal(t, test.ip, ip, "ip should equal, row %d", row)
	}
}
