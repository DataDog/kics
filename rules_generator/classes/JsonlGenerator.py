import json
from typing import List, Dict
from pathlib import Path


class JsonlGenerator:

    def __init__(self, prefix: str) -> None:
        self.prefix = Path(prefix)

    def __write_json(
        self, code_snippet_no_module: str, code_snippet_module: str
    ) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure add support for the modules. 
    Answer with only the refactored code."""
        json_string = {
            "messages": [
                {"role": "system", "content": system_message},
                {"role": "user", "content": code_snippet_no_module},
                {"role": "assistant", "content": code_snippet_module},
            ]
        }

        return json.dumps(json_string)

    def __write_jsonl_from_snippets(self, snippets: List[Dict[str, str]]) -> str:
        s = ""
        for couple in snippets:
            s += self.__write_json(couple["no_modules"], couple["modules"]) + "\n"
        return s

    def __create_no_module_and_module_structure(self) -> List[Dict[str, str]]:
        snippets = []
        standard_path = "queries/terraform/aws"
        for rule_path in (Path(self.prefix) / standard_path).iterdir():
            with (
                open(rule_path / "query.rego", "r") as f_no_modules,
                open(
                    (rule_path / "query.rego")
                    .absolute()
                    .as_posix()
                    .replace(self.prefix, "assets")
                ) as f_modules,
            ):
                snippets.append(
                    {"no_modules": f_no_modules.read(), "modules": f_modules.read()}
                )
        return snippets

    def write_jsonl(self) -> str:
        return self.__write_jsonl_from_snippets(
            self.__create_no_module_and_module_structure()
        )
