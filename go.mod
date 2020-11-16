module github.com/shupkg/wtool

go 1.15

replace (
	github.com/shupkg/wbin => ../wbin
	github.com/shupkg/wproto => ../wproto
)

require (
	github.com/shupkg/wbin v0.0.0-00010101000000-000000000000
	github.com/shupkg/wproto v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.1
	google.golang.org/protobuf v1.25.0
)
