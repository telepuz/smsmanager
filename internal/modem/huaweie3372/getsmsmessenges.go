package huaweie3372

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/telepuz/smsmanager/internal"
)

type smsList struct {
	Count    int         `xml:"Count"`
	Messages messageList `xml:"Messages"`
}

type messageList struct {
	MessageList []internal.Message `xml:"Message"`
}

func (h *HuaweiE3372) GetSMSMessenges() ([]internal.Message, error) {
	sestok, err := h.getSesTokInfo()
	if err != nil {
		return nil, fmt.Errorf(
			"GetSMSMessenges: %v",
			err,
		)
	}

	xmlbody := `
<request>
	<PageIndex>1</PageIndex>
	<ReadCount>10</ReadCount>
	<BoxType>1</BoxType>
	<SortType>0</SortType>
	<Ascending>0</Ascending>
	<UnreadPreferred>1</UnreadPreferred>
</request>`

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"http://%s/api/sms/sms-list",
			h.modemURL,
		),
		bytes.NewBuffer([]byte(xmlbody)),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"GetSMSMessenges: %v",
			err,
		)
	}
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("__RequestVerificationToken", sestok.TokInfo)
	req.AddCookie(&http.Cookie{
		Name:  "SessionID",
		Value: sestok.SesInfo,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"GetSMSMessenges: %v",
			err,
		)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"GetSMSMessenges: %v",
			err,
		)
	}
	v := smsList{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf(
			"GetSMSMessenges: %v",
			err,
		)
	}

	return v.Messages.MessageList, nil
}
