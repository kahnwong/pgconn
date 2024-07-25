package config

type Connection struct {
	Hostname  string
	ProxyKind string
	ProxyHost string
	Port      int
	Username  string
	Password  string
	Dbname    string
}

var ConnMap = createConnMap(readConfig())
