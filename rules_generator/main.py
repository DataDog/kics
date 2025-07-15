from classes.CodeProcessor import CodeProcessor
from classes.RulesGenerator import RulesGenerator
from classes.JsonlGenerator import JsonlGenerator
from classes.Coordinator import Coordinator


def generate_new_rules(path: str):
    coordinator = Coordinator()

    coordinator.generate_new_rules(path)


def get_new_rule():
    rulesGenerator = RulesGenerator()
    codeProcessor = CodeProcessor()

    snippet = codeProcessor.__read_snippet(
        "assets/queries/terraform/aws/alb_is_not_integrated_with_waf"
    )
    new_rule = rulesGenerator.send_request(snippet)

    return new_rule


def write_example_jsonl():
    jsonlGenerator = JsonlGenerator()

    snippets = []
    couple_1 = {
        "no_modules": """Line 1.1.1 with "double quoted string"
        Line 1.1.2""",
        "modules": """Line 1.2.1
        Line 1.2.2""",
    }
    couple_2 = {
        "no_modules": """Line 2.1.1
        Line 2.1.2""",
        "modules": """Line 2.2.1
        Line 2.2.2""",
    }
    snippets.append(couple_1)
    snippets.append(couple_2)

    s = jsonlGenerator.__write_jsonl(snippets)
    print(s)


def write_definitive_jsonl():
    jsonlGenerator = JsonlGenerator("")
    return jsonlGenerator.write_jsonl_from_path()


if __name__ == "__main__":
    generate_new_rules("assets/queries/terraform/aws")
