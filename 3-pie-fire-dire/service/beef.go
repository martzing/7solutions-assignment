package service

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/martzing/7solutions-assignment/3-pie-fire-dire/configs"
)

func GetBeefSummary() (map[string]int32, error) {
	resp, err := http.Get(*configs.BeefUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	body = bytes.ReplaceAll(body, []byte("\n"), []byte(" "))
	data := bytes.Split(body, []byte(" "))
	beef := make(map[string]int32)
	for _, v := range data {
		b := bytes.ReplaceAll(v, []byte(","), []byte(""))
		b = bytes.ReplaceAll(b, []byte("."), []byte(""))
		b = bytes.ReplaceAll(b, []byte(" "), []byte(""))
		if string(b) == "" {
			continue
		}
		if beef[string(b)] == 0 {
			beef[string(b)] = 1
		} else {
			beef[string(b)]++
		}
	}
	return beef, nil
}
