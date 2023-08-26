# pgconn

`pgcli` wrapper to connect to PostgreSQL database specified in `db.yaml`. Proxy/tunnel connections are automatically created and killed when pgcli is exited.

## Pre-requisites

- Install [pgcli](https://www.pgcli.com/install)
- If you are connecting to GCP's Cloud SQL: install [cloud-sql-proxy](https://github.com/GoogleCloudPlatform/cloud-sql-proxy)

## Setup

1. `pipx install -e .`
2. `pgconn --install-completion`
3. create a config file in `~/.config/pgconn/db.yaml`

```yaml
- name: sammple-db
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
    database
    role DATABASE
```

## Examples

`list database`

```bash
❯ pgconn list database
Database list:
- nuc-map
- local-postgres
```

`list role`

```bash
❯ pgconn list role nuc-map
Dbname: nuc-map
Roles:
- postgres
- postgres2
```

`connect`

```bash
❯ pgconn connect nuc-map postgres
Server: PostgreSQL 15.3
Version: 3.5.0
Home: http://pgcli.com
postgres@192:map>
```
