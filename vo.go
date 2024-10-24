package curl2

// @enumGenerated
type ContentType string

func (g ContentType) Values() []string {
	return []string{ApplicationJson.String(), ApplicationXml.String(), ApplicationFormUrlEncoded.String()}
}
func (g ContentType) String() string {
	return string(g)
}

const (
	ApplicationJson           ContentType = "application/json"
	ApplicationXml            ContentType = "application/xml"
	ApplicationFormUrlEncoded ContentType = "application/x-www-form-urlencoded"
)

// @enumGenerated
type Method string

func (g Method) Values() []string {
	return []string{MethodGet.String(), MethodHead.String(), MethodPost.String(), MethodPut.String(), MethodPatch.String(), MethodDelete.String(), MethodConnect.String(), MethodOptions.String(), MethodTrace.String()}
}
func (g Method) String() string {
	return string(g)
}

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)
