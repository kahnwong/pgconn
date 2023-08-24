import typer
from pg_conn_cli.utils.prereq import check_if_pgcli_in_path

def entrypoint():
    check_if_pgcli_in_path()


if __name__ == "__main__":
    entrypoint()
