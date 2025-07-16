import argparse
from pathlib import Path

from classes.CodeProcessor import CodeProcessor
from classes.RulesGenerator import RulesGenerator
from classes.JsonlGenerator import JsonlGenerator
from classes.Coordinator import Coordinator


def parse_args():
    parser = argparse.ArgumentParser(
        description="Generate module support and terraform examples from existing rules"
    )
    parser.add_argument("generation", type=str, help="Which function to call")
    parser.add_argument("path", type=str, help="Path")
    return parser.parse_args()


def generate_new_rules(path: str):
    coordinator = Coordinator()

    coordinator.generate_new_rules(path)


def get_new_rule(path: str) -> str:
    rulesGenerator = RulesGenerator()
    codeProcessor = CodeProcessor()

    snippet = codeProcessor.read_snippet(path)
    new_rule = rulesGenerator.send_rule_request(snippet)

    return new_rule


def write_definitive_module_jsonl():
    jsonlGenerator = JsonlGenerator("")
    return jsonlGenerator.write_module_jsonl()


def write_definitive_terraforms_jsonl():
    jsonlGenerator = JsonlGenerator("")
    return jsonlGenerator.write_terraform_jsonl()


if __name__ == "__main__":
    args = parse_args()
    asked = args.generation
    path = args.path
    if asked == "generate_new_rules":
        generate_new_rules(path)
    elif asked == "get_new_rule":
        print(get_new_rule(path))
    elif asked == "write_definitive_module_jsonl":
        print(write_definitive_module_jsonl())
    elif asked == "write_definitive_terraforms_jsonl":
        print(write_definitive_terraforms_jsonl())
