package core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sniping/types"
)

func SetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

func GetResponse(resp *http.Response) (map[string]any, error) {
    var response map[string]any
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				log.Fatal(err)
			}
    return response, nil
}

func DoRequest(
    client *http.Client,
    method string,
    url string,
    payload any,
) (types.ResponseData, error) {

    var body io.Reader

    if payload != nil {
        jsonData, err := json.Marshal(payload)
        if err != nil {
            return types.ResponseData{}, err
        }
        body = bytes.NewBuffer(jsonData)
    }

    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return types.ResponseData{}, err
    }

    // headers
	SetHeaders(req)

    resp, err := client.Do(req)
	if err != nil {
		return types.ResponseData{}, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

    res := types.ResponseData{
        Status: resp.StatusCode,
        Length: len(data),
        Body: data,
    }
    return res, nil
}