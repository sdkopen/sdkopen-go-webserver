package http

type ContentType int

const (
	ContentTypeJSON ContentType = iota
	ContentTypeCSV
	ContentTypeXML
	ContentTypeTextXML
	ContentTypeTextPlain
	ContentTypePDF
	ContentTypeOctetStream

	ContentTypeHeader string = "Content-Type"
)

func (ct ContentType) String() string {
	return [...]string{
		"application/json",
		"text/csv",
		"application/xml",
		"text/xml",
		"text/plain",
		"application/pdf",
		"application/octet-stream",
	}[ct]
}
