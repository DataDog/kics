from pathlib import Path
from typing import List


class CodeProcessor:
    def __read_file(self, path: Path) -> str:
        try:
            with open(path, "r") as f:
                return f.read()
        except Exception as e:
            print(e)

    def read_snippet(self, path: Path) -> str:
        return self.__read_file(path / "query.rego")

    def __write_file(self, path: Path, content: str) -> None:
        try:
            with open(path, "w") as f:
                f.write(content)
        except Exception as e:
            print(e)

    def write_snippet(self, path: Path, snippet: str) -> None:
        self.__write_file(path / "query.rego", snippet)
