package domain

type ViewTemplate struct {
	SignatureCount int
	Signatures     []string
}

type Guestbook struct {
	Signatures []Signature `json:"signatures"`
}

type Signature struct {
	Signature string `json:"signature"`
}
