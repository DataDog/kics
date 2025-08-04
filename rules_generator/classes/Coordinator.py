from classes.CodeProcessor import CodeProcessor
from classes.RulesGenerator import RulesGenerator
import sys

from alive_progress import alive_bar
from pathlib import Path
from collections.abc import Callable
from typing import List, Dict, Optional, Any


class Coordinator:
    def __init__(self) -> None:
        self.codeProcessor = CodeProcessor()
        self.rulesGenerator = RulesGenerator()

    def __generate_new_rule(self, rule: str):
        message = self.codeProcessor.read_snippet(rule)
        received = self.rulesGenerator.send_rule_request(message)
        result = self.codeProcessor.write_rule_snippet(rule, received)
        print(f"Generated module support support for {rule.name}!")
        return result

    def __generate_new_terraforms(self, rule: str, update: Dict[str, Any]):
        message = self.codeProcessor.read_snippet(rule)
        message += "\n@@@@@"
        for source in update:
            message += f"\nsource: {source}"
            for resource in update[source]:
                message += f"\nresource: {resource}"
            message += "\n"
        received = self.rulesGenerator.send_terraform_request(message)
        result = self.codeProcessor.write_terraform_snippets(rule, received)
        print(f"Generated terraform files support for {rule.name}!")
        return result

    def __generate_new_files(
        self,
        path_str: str,
        skip: int,
    ) -> None:
        path = Path(path_str)
        common = self.codeProcessor.load_common()
        with alive_bar(len(list(path.glob("*")))) as bar:
            rules_list = sorted(path.iterdir())[skip:]
            try:
                for rule in rules_list:
                    update = self.__generate_new_rule(
                        rule,
                    )
                    if update == {}:
                        continue
                    common = self.codeProcessor.update_common(common, update)
                    self.__generate_new_terraforms(rule, update)
                    bar()
                self.codeProcessor.write_common(common)
            except KeyboardInterrupt:
                print("Ctrl-C pressed!")
                if update != None:
                    common = self.codeProcessor.update_common(common, update)
                self.codeProcessor.write_common(common)
                sys.exit(0)
            except Exception as e:
                if update != None:
                    common = self.codeProcessor.update_common(common, update)
                self.codeProcessor.write_common(common)
                sys.exit(
                    f"The following error has occured during the rule generation: {e}"
                )

    def generate_module_support(self, path_str: str, skip: int) -> None:
        self.__generate_new_files(
            path_str,
            skip,
        )
