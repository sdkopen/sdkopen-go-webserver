package http

type ContentType int

const (
	ContentTypeJSON ContentType = iota
	ContentTypeTextPlain
	ContentTypePDF
	ContentTypeOctetStream

	ContentTypeHeader string = "Content-Type"
)

func (ct ContentType) String() string {
	return [...]string{
		"application/json",
		"text/plain",
		"application/pdf",
		"application/octet-stream",
	}[ct]
}
