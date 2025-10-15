package huaweie3372

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type sesTokInfo struct {
	TokInfo string `xml:"TokInfo"`
	SesInfo string `xml:"SesInfo"`
}

func (h *HuaweiE3372) getSesTokInfo() (*sesTokInfo, error) {
	slog.Debug("getSesTokInfo: Get token for modem")
	resp, err := http.Get(
		fmt.Sprintf(
			"http://%s/api/webserver/SesTokInfo",
			h.modemURL,
		),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"getSesTokInfo: %v",
			err,
		)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"getSesTokInfo: %v",
			err,
		)
	}
	v := sesTokInfo{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf(
			"getSesTokInfo: %v",
			err,
		)
	}
	return &v, nil
}
