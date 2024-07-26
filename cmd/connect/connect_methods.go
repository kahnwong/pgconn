package connect

import (
	"math/rand"
	"time"

	"github.com/kahnwong/pgconn/config"
)

var connMap = config.ConnMap

type connection struct {
	config.Connection
}

func (c connection) SetProxyPort() int {
	// prevent port conflict in case
	// simultaneously connecting to proxied db
	if c.ProxyKind == "" {
		return c.Port
	} else {
		minPort := 5432
		maxPort := 8000

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		port := r.Intn(maxPort-minPort+1) + minPort

		return port
	}
}
