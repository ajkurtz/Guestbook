package domain

type ViewTemplate struct {
	SignatureCount int
	Signatures     []string
}

type Guestbook struct {
	Signatures     []Signature
}

type Signature struct {
	Signature     string
}