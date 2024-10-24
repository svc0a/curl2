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
