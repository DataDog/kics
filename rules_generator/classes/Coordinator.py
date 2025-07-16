from classes.CodeProcessor import CodeProcessor
from classes.RulesGenerator import RulesGenerator

from alive_progress import alive_bar
from pathlib import Path
from collections.abc import Callable


class Coordinator:
    def __init__(self) -> None:
        self.codeProcessor = CodeProcessor()
        self.rulesGenerator = RulesGenerator()

    def __generate_new_file(
        self,
        path_str: str,
        send_message: Callable[[str], None],
        write: Callable[[Path, str], None],
        object_type: str,
    ) -> None:
        path = Path(path_str)
        with alive_bar(len(list(path.glob("*")))) as bar:
            for rule in path.iterdir():
                message = self.codeProcessor.read_snippet(rule)
                received = send_message(message)
                write(rule, received)
                print(f"Generated {object_type} support for {rule.name}!")
                bar()

    def generate_new_rules(self, path_str: str) -> None:
        self.__generate_new_file(
            path_str,
            self.rulesGenerator.send_rule_request,
            self.codeProcessor.write_snippet,
            "module support",
        )

    def generate_new_terraforms(self, path_str: str) -> None:
        self.__generate_new_file(
            path_str,
            self.rulesGenerator.send_terraform_request,
            self.codeProcessor.write_terraform_snippets,
            "terraform files",
        )
