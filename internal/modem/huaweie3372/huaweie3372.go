package huaweie3372

import "log/slog"

type HuaweiE3372 struct {
	modemURL string
}

func New(modemURL string) *HuaweiE3372 {
	slog.Debug("Create new Huawei modem")
	return &HuaweiE3372{
		modemURL: modemURL,
	}
}
