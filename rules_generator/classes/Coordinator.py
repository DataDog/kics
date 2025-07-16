from classes.CodeProcessor import CodeProcessor
from classes.RulesGenerator import RulesGenerator

from alive_progress import alive_bar
from pathlib import Path


class Coordinator:
    def __init__(self) -> None:
        self.codeProcessor = CodeProcessor()
        self.rulesGenerator = RulesGenerator()

    def generate_new_rules(self, path_str: str):
        path = Path(path_str)
        with alive_bar(len(list(path.glob("*")))) as bar:
            for rule in path.iterdir():
                old_snippet = self.codeProcessor.read_snippet(rule)
                new_snippet = self.rulesGenerator.send_rule_request(old_snippet)
                self.codeProcessor.write_snippet(rule, new_snippet)
                print(f"Generated module support for {rule.name}!")
                bar()

    def generate_new_terraforms(self, path_str: str):
        path = Path(path_str)
        with alive_bar(len(list(path.glob("*")))) as bar:
            for rule in path.iterdir():
                snippet = self.codeProcessor.read_snippet(rule)
                terraforms = self.rulesGenerator.send_terraform_request(snippet)
                self.codeProcessor.write_terraform_files(rule, terraforms)
                print(f"Generated terraform files for {rule.name}!")
                bar()
