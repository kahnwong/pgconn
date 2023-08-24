import shutil
from typing import Union


def check_if_pgcli_in_path() -> Union[None, RuntimeError]:
    binary_name = "pgcli"
    binary_path = shutil.which(binary_name)

    if binary_path is not None:
        return None
    else:
        raise RuntimeError(f"{binary_name} not found in PATH")
