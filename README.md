# pg-conn-cli

## Pre-requisites

- Install [pgcli](https://www.pgcli.com/install)
- If you are connecting to GCP's Cloud SQL: install [cloud-sql-proxy](https://github.com/GoogleCloudPlatform/cloud-sql-proxy)

## Setup

```bash
pipx install -e .
pg-conn-cli --install-completion
```

Then create a config file in `~/.config/pg-conn-cli/db.yaml`

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
❯ pg-conn-cli list database
Database list:
- nuc-map
- local-postgres
```

`list role`

```bash
❯ pg-conn-cli list role nuc-map
Dbname: nuc-map
Roles:
- postgres
- postgres2
```

`connect`

```bash
❯ pg-conn-cli connect nuc-map postgres
Server: PostgreSQL 15.3
Version: 3.5.0
Home: http://pgcli.com
postgres@192:map>
```
