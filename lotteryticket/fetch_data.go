package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func FetchData() {

	for i := 1; i < 10; i++ {
		data, err := getRequestData(i)
		if err != nil {
			fmt.Printf("i %v, err %v\n", i, err)
			//break
		}
		fmt.Print(data)
	}
}

func getRequestData(page int) (da *Data, err error) {

	baseUrl := "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice"
	u, err := url.Parse(baseUrl)
	params := url.Values{}
	params.Add("name", "ssq")
	params.Add("pageNo", strconv.Itoa(page))
	params.Add("pageSize", strconv.Itoa(30))
	params.Add("systemType", "PC")
	u.RawQuery = params.Encode()
	data := &Data{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Printf("create request failed: %v", err)
		return data, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "HMF_CI=34242d4476bad37784f4140c41b8b97f69b294e4d8dcb71ecd3bfc4d198e4c63b2e8430f727496262967582ff49c75419e2a7d347941f7668dd2b65e4d9bf1b205; 21_vq=28")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		_ = fmt.Errorf("request failed, page %v", err)
		return data, err
	}

	if resp.StatusCode != http.StatusOK {
		v := fmt.Sprintf("request failed, code %v", resp.StatusCode)
		return data, errors.New(v)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating gzip reader: %v\n", err)
			return
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Errorf("body parse failed")
		return data, err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Errorf("json parse failed")
		return data, err
	}
	return data, nil

}
