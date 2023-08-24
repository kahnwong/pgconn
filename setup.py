from setuptools import find_packages
from setuptools import setup

if __name__ == "__main__":
    setup(
        name="pg_conn_cli",
        version="0.1",
        packages=find_packages(exclude=["tests"]),
        install_requires=[
            "typer[all]>=0.9.0",
            "pyyaml",
        ],
        entry_points={
            "console_scripts": [
                "pg-conn-cli=pg_conn_cli.app:entrypoint",
            ],
        },
    )
