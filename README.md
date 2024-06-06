# pgconn

`pgcli` wrapper to connect to PostgreSQL database specified in `db.yaml`. Proxy/tunnel connection is automatically created and killed when pgcli is exited.

## Pre-requisites

- Install [pgcli](https://www.pgcli.com/install)
- If you are connecting to GCP's Cloud SQL: install [cloud-sql-proxy](https://github.com/GoogleCloudPlatform/cloud-sql-proxy)

## Setup

1. Install

```bash
go install github.com/kahnwong/pgconn@latest
```

2. create a config file with SOPs in `~/.config/pgconn/pgconn.sops.yaml`

```yaml
pgconn:
  - account: personal
    dbs:
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
connect ACCOUNT DATABASE ROLE
list
    accounts
    databases ACCOUNT
    roles ACCOUNT DATABASE
```

## Examples

`list accounts`

```bash
❯ pgconn list accounts
Available accounts:
  - personal
  - foo
```

`list databases`

```bash
❯ pgconn list databases personal
Account: personal
Databases:
  - nuc-postgres
  - local-postgres
```

`list roles`

```bash
❯ pgconn list roles personal nuc-map
Account: personal
Database: nuc-map
Roles:
  - a
  - b
```

`connect`

```bash
❯ pgconn connect personal nuc-map postgres
Server: PostgreSQL 15.3
Version: 3.5.0
Home: http://pgcli.com
postgres@192:map>
```
