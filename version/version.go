package version

var (
	// Version defines version of the package
	Version string
	// Revision is current revision
	Revision string
	// BuildDate is a latest build date
	BuildDate string
	// GoVersion defines version of the Go language where it was built
	GoVersion string
)

// Info returns map of version info
var Info = map[string]string{
	"version":   Version,
	"revision":  Revision,
	"buildDate": BuildDate,
	"goVersion": GoVersion,
}
