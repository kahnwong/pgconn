# pgconn

`pgcli` wrapper to connect to PostgreSQL database specified in `db.yaml`. Proxy/tunnel connection is automatically created and killed when pgcli is exited.

## Pre-requisites

- Install [pgcli](https://www.pgcli.com/install)
- If you are connecting to GCP's Cloud SQL: install [cloud-sql-proxy](https://github.com/GoogleCloudPlatform/cloud-sql-proxy)

## Setup

1. `pipx install git+https://github.com/kahnwong/pgconn.git`
2. `pgconn --install-completion`
3. create a config file in `~/.config/pgconn/db.yaml`

```yaml
- name: sample-db
  hostname: localhost
  proxy: # this block is optional
    kind: cloud-sql-proxy
    host: $GCP_PROJECT:$GCP_REGION:$INSTANCE_IDENTIFIER
  roles:
    - username: postgres
      password: postgrespassword
      dbname: sample_db

# if using ssh tunnelling
proxy:
  kind: ssh
  host: $SSH_CONFIG_HOST
```

## Available commands

```bash
connect DATABASE ROLE
list
    databases
    roles DATABASE
```

## Examples

`list databases`

```bash
❯ pgconn list databases
Available databases:
    nuc-postgres
    local-postgres
```

`list roles`

```bash
❯ pgconn list roles nuc-map
Database: nuc-postgres
Roles:
    postgres
```

`connect`

```bash
❯ pgconn connect nuc-map postgres
Server: PostgreSQL 15.3
Version: 3.5.0
Home: http://pgcli.com
postgres@192:map>
```
