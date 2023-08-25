import os

import yaml

home_directory = os.path.expanduser("~")


# read config
file_path = f"{home_directory}/.config/pg-conn-cli/db.yaml"
if os.path.exists(file_path):
    with open(file_path, "r") as yaml_file:
        yaml_data = yaml.safe_load(yaml_file)
elif not os.path.exists(file_path):
    print(f"Config file at {file_path} does not exist")
    raise FileNotFoundError


def _read_db_config():
    # convert database list to dict
    db_config = {item["name"]: item for item in yaml_data}

    # confict role list to dict
    for database in db_config:
        db_config[database]["roles"] = {
            item["username"]: item for item in db_config[database]["roles"]
        }

    return db_config


def _get_database_roles(database: str, db_config):
    if database in databases:
        return db_config[database]["roles"]
    else:
        print(f"Database {database} is not specified in config")
        return []


db_config = _read_db_config()

databases = db_config.keys()
db_role_mapping = {
    database: _get_database_roles(database=database, db_config=db_config)
    for database in databases
}
