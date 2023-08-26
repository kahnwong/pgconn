from setuptools import find_packages
from setuptools import setup

if __name__ == "__main__":
    setup(
        name="pgconn",
        version="0.1",
        description="pgcli wrapper to connect to PostgreSQL database specified in `db.yaml`. Proxy/tunnel connections are automatically created and killed when pgcli is exited.",
        long_description=open("README.md").read(),
        long_description_content_type="text/markdown",
        url="https://github.com/kahnwong/pgconn",
        author="Karn Wong",
        author_email="karn@karnwong.me",
        license="MIT",
        project_urls={
            "Documentation": "https://github.com/kahnwong/pgconn",
            "Source": "https://github.com/kahnwong/pgconn",
        },
        packages=find_packages(exclude=["tests"]),
        install_requires=[
            "typer[all]>=0.9.0",
            "pyyaml",
        ],
        python_requires=">=3.9",
        entry_points={
            "console_scripts": [
                "pgconn=pgconn.app:entrypoint",
            ],
        },
    )
