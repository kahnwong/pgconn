package config

func createConnMap(config Config) map[string]map[string]map[string]Connection {
	configMap := make(map[string]map[string]map[string]Connection)

	for _, a := range config.Pgconn {
		configMap[a.Account] = map[string]map[string]Connection{}

		for _, db := range a.Dbs {
			configMap[a.Account][db.Name] = map[string]Connection{}
			hostname := db.Hostname
			proxyKind := db.Proxy.Kind
			proxyHost := db.Proxy.Host

			for _, role := range db.Roles {
				configMap[a.Account][db.Name][role.Username] = Connection{
					Hostname:  hostname,
					ProxyKind: proxyKind,
					ProxyHost: proxyHost,
					Port:      5432,
					Username:  role.Username,
					Password:  role.Password,
					Dbname:    role.Dbname,
				}
			}
		}
	}

	return configMap
}
