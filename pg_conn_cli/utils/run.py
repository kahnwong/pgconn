import subprocess

from pg_conn_cli.utils.prereq import is_binary_in_path


def start_proxy(kind: str, host: str, db_hostname: str = False):
    if kind == "cloud-sql-proxy":
        binary_name = "cloud-sql-proxy"
        if is_binary_in_path(binary_name) is True:
            background_command = [
                "cloud-sql-proxy",
                host,
                "--quiet",
            ]
        else:
            exit()

    elif kind == "ssh":  # only used when running proxy only
        background_command = [
            "ssh",
            "-N",
            "-L",
            f"5432:{db_hostname}:5432",
            host,
        ]

    return subprocess.Popen(
        background_command,
        shell=False,
        stdin=None,
        stdout=None,
        close_fds=True,
    )


def connect_db(
    hostname: str,
    dbname: str,
    username: str,
    password: str,
    port: str = "5432",
    ssh_tunnel: str = None,
):
    pgcli_command = "pgcli"

    connection_uri = (
        f"postgresql://{username}:{password}@{hostname}:{port}/{dbname}?sslmode=disable"
    )

    command = [
        pgcli_command,
        connection_uri,
    ]

    if ssh_tunnel:
        command.extend(["--ssh-tunnel", ssh_tunnel])

    try:
        subprocess.run(command, check=True)

    except subprocess.CalledProcessError as e:
        print(f"Error running pgcli: {e}")
