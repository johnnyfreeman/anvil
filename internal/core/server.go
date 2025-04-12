package core

type Server struct {
	Host         string // SSH hostname or IP
	Hostname     string // `hostname`
	OS           string // from `/etc/os-release` or `lsb_release -d`
	OSVersion    string // extracted from the same as above
	Kernel       string // `uname -r`
	Architecture string // `uname -m`
}
