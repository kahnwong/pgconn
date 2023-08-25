import shutil
from typing import Union


def is_binary_in_path(binary_name: str) -> Union[None, RuntimeError]:
    if shutil.which(binary_name):
        return True
    else:
        print(f"{binary_name} not found in PATH")
        return False
