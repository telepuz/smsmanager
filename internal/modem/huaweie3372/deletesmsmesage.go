package huaweie3372

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func (h *HuaweiE3372) DeleteSMSMessage(messageID int) error {
	sestok, err := h.getSesTokInfo()
	if err != nil {
		return fmt.Errorf(
			"DeleteSMSMessage: %v",
			err,
		)
	}

	xmlbody := fmt.Sprintf(`
<request>
	<Index>%v</Index>
</request>`,
		messageID,
	)

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"http://%s/api/sms/delete-sms",
			h.modemURL,
		),
		bytes.NewBuffer([]byte(xmlbody)),
	)
	if err != nil {
		return fmt.Errorf(
			"DeleteSMSMessage: %v",
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
		return fmt.Errorf(
			"DeleteSMSMessage: %v",
			err,
		)
	}
	defer resp.Body.Close()

	slog.Info(fmt.Sprintf(
		"DeleteSMSMessage: Deleted message %v",
		messageID,
	))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(
			"DeleteSMSMessage: %v",
			err,
		)
	}
	v := smsList{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf(
			"DeleteSMSMessage: %v",
			err,
		)
	}

	fmt.Print(v)
	return nil
}
