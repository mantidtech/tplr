package tplr

var version = "snapshot"

// Version returns the version number of this app
// Use the build.sh script to have this populated meaningfully
func Version() string {
	return version
}
