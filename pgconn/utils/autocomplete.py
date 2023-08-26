import typer

from pgconn.utils.config import db_role_mapping


def complete_db():  # noqa
    return list(db_role_mapping.keys())


def complete_role(ctx: typer.Context):  # noqa
    database = ctx.params.get("database")

    return db_role_mapping[database]
