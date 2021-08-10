package parameter

const (
	Platform = "platform"
	Asset    = "asset"
	Target   = "target"
	Document = "document"
	Sheet    = "sheet"
	Package  = "package"
)

type Parameter struct {
	Platform string
	Asset    string
	Target   string
	Document string
	Sheet    string
	Package  string
}
