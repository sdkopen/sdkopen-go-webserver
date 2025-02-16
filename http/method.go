package http

type Method int

const (
	MethodGet Method = iota
	MethodHead
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
	MethodConnect
	MethodOptions
	MethodTrace
)

func (hm Method) String() string {
	return [...]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}[hm]
}
