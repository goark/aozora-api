module github.com/spiegel-im-spiegel/aozora-api

go 1.16

require (
	github.com/spf13/cobra latest
	github.com/spiegel-im-spiegel/errs latest
	github.com/spiegel-im-spiegel/gocli latest
)

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
