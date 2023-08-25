from time import sleep

import typer
from rich import print
from typing_extensions import Annotated

from pg_conn_cli.utils.autocomplete import complete_db
from pg_conn_cli.utils.autocomplete import complete_role
from pg_conn_cli.utils.config import db_config
from pg_conn_cli.utils.config import db_role_mapping
from pg_conn_cli.utils.prereq import is_binary_in_path
from pg_conn_cli.utils.run import connect_db
from pg_conn_cli.utils.run import start_proxy


# --------------- init app --------------- #
app = typer.Typer()
list_app = typer.Typer()
app.add_typer(list_app, name="list")


# --------------- command: list --------------- #
@list_app.command()
def database():
    database_bullets = "\n".join([f"- {i}" for i in db_role_mapping.keys()])
    r = f"[green]Database list:[/green]\n{database_bullets}"

    print(r)


@list_app.command()
def role(
    database: Annotated[str, typer.Argument(autocompletion=complete_db)],
):
    if roles := db_role_mapping.get(database):
        role_bullets = "\n".join([f"- {i}" for i in roles])

        r = f"[green]Dbname:[/green] {database} \n[blue]Roles:[/blue]\n{role_bullets}"

        print(r)


# --------------- command: connect --------------- #
@app.command()
def connect(
    database: Annotated[str, typer.Argument(autocompletion=complete_db)],
    role: Annotated[str, typer.Argument(autocompletion=complete_role)],
):
    # -------------- db credentials -------------- #
    hostname = db_config[database]["hostname"]
    dbname = db_config[database]["roles"][role]["dbname"]
    username = db_config[database]["roles"][role]["username"]
    password = db_config[database]["roles"][role]["password"]

    if proxy := db_config[database].get("proxy"):
        if proxy["kind"] == "cloud-sql-proxy":
            proxy_proc = start_proxy(
                kind=proxy["kind"],
                host=proxy["host"],
                db_hostname=db_config[database]["hostname"],
            )

            sleep(0.50)  # important, so proxy has some time to start up

            connect_db(
                hostname=hostname,
                dbname=dbname,
                username=username,
                password=password,
            )

            proxy_proc.kill()
            proxy_proc.wait()
        elif proxy["kind"] == "ssh":
            connect_db(
                hostname=hostname,
                dbname=dbname,
                username=username,
                password=password,
                ssh_tunnel=proxy["host"],
            )
    elif not proxy:
        connect_db(
            hostname=hostname,
            dbname=dbname,
            username=username,
            password=password,
        )


# --------------- entrypoint --------------- #
def entrypoint():
    if is_binary_in_path("pgcli"):
        app()


if __name__ == "__main__":
    entrypoint()
