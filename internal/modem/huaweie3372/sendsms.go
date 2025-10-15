package huaweie3372

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
)

func (h *HuaweiE3372) SendSMS(phoneNumber, text string) error {
	slog.Debug("SendSMS: Send SMS from modem")
	sestok, err := h.getSesTokInfo()
	if err != nil {
		return fmt.Errorf(
			"SendSMS: %v",
			err,
		)
	}

	xmlbody := fmt.Sprintf(`
<request>
	<Index>-1</Index>
	<Phones><Phone>%s</Phone></Phones>
	<Sca></Sca>
	<Content>%s</Content>
	<Length>-1</Length>
	<Reserved>1</Reserved>
	<Date>-1</Date>
</request>`,
		phoneNumber,
		text)

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"http://%s/api/sms/send-sms",
			h.modemURL,
		),
		bytes.NewBuffer([]byte(xmlbody)),
	)
	if err != nil {
		return fmt.Errorf(
			"SendSMS: %v",
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
			"SendSMS: %v",
			err,
		)
	}
	defer resp.Body.Close()
	return nil
}
