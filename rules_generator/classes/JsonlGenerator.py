import json
from typing import List, Dict
from pathlib import Path


class JsonlGenerator:

    def __init__(self, prefix: str) -> None:
        self.prefix = Path(prefix)

    def __write_json(
        self, user_message: str, assistant_message: str, system_message: str
    ) -> str:
        json_string = {
            "messages": [
                {"role": "system", "content": system_message},
                {"role": "user", "content": user_message},
                {"role": "assistant", "content": assistant_message},
            ]
        }

        return json.dumps(json_string)

    def __write_jsonl_from_couples(
        self, couples: List[Dict[str, str]], key1: str, key2: str, system_message: str
    ) -> str:
        s = ""
        for couple in couples:
            s += self.__write_json(couple[key1], couple[key2], system_message) + "\n"
        return s

    def __create_no_module_and_module_structure(self) -> List[Dict[str, str]]:
        snippets = []
        standard_path = "queries/terraform/aws"
        if not (standard_path.is_dir()):
            return snippets
        for rule_path in (Path(self.prefix) / standard_path).iterdir():
            with (
                open(rule_path / "query.rego", "r") as f_no_modules,
                open(
                    (rule_path / "query.rego", "r")
                    .absolute()
                    .as_posix()
                    .replace(self.prefix, "assets")
                ) as f_modules,
            ):
                snippets.append(
                    {"no_modules": f_no_modules.read(), "modules": f_modules.read()}
                )
        return snippets

    def __create_snippet_and_terraforms_structure(self):
        couples = []
        standard_path = Path("module_support/queries/terraform/aws")
        if not (standard_path.is_dir()):
            return couples
        for rule_path in (standard_path).iterdir():
            test_path = rule_path / "test"
            query_path = rule_path / "query.rego"
            if not (test_path.is_dir() and query_path.is_file()):
                continue
            positive_path = ""
            negative_path = ""
            for test in test_path.iterdir():
                if "positive" in test.name:
                    positive_path = test
                elif "negative" in test.name:
                    negative_path = test
            if positive_path == "" or negative_path == "":
                continue
            with (
                open(positive_path, "r") as f_positive,
                open(negative_path, "r") as f_negative,
                open(query_path, "r") as f_rule,
            ):
                couples.append(
                    {
                        "rule": f_rule.read(),
                        "terraforms": f"{f_positive.read()}\n#####\n{f_negative.read()}",
                    }
                )
        return couples

    def write_module_jsonl(self) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure, add support for the modules. 
    Answer with only the refactored code."""
        return self.__write_jsonl_from_couples(
            self.__create_no_module_and_module_structure(),
            "no_modules",
            "modules",
            system_message,
        )

    def write_terraform_jsonl(self):
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure, create examples tf files using the related module.
    Answer with a positive example that will trigger the rule, one new line, "#####", a second new line and a negative example that won't trigger the rule."""
        return self.__write_jsonl_from_couples(
            self.__create_snippet_and_terraforms_structure(),
            "rule",
            "terraforms",
            system_message,
        )
