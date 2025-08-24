package huaweie3372

type HuaweiE3372 struct {
	modemURL string
}

func New(modemURL string) *HuaweiE3372 {
	return &HuaweiE3372{
		modemURL: modemURL,
	}
}
