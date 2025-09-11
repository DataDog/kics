from openai import OpenAI
from classes.models.Output import Output


class RulesGenerator:

    def __init__(self):
        self.client = OpenAI()

    def send_rule_request(self, code_snippet: str) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure add support for the modules.
Know that the rego get_module_equivalent_key function takes 4 arguments which are the provider of the module, the source of the module, the resource and the key to look for. Also, be careful of nested attributes, separate the tests to get the key first and the the nested attribute.
Answer with only the refactored code, one new line, "@@@@@", a second new line and a json with the source of the module as a key holding two attributes, resources which is a list of the supported resources and inputs which holds the mapping between the module variable and the resource variable."""
        response = self.client.responses.create(
            model="ft:gpt-4.1-2025-04-14:datadog-staging:module-support:BzKPAmRd",
            input=[
                {"role": "system", "content": system_message},
                {"role": "developer", "content": code_snippet},
            ],
        )
        return response.output_text

    def send_terraform_request(self, code_snippet: str) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. 
        The User will provide you with a module source, one or multiple terraform resources and a rego file. Your job is to emit terraform code examples.
    Answer with a positive example that will trigger the rule, one new line, "@@@@@", a second new line and a negative example that won't trigger the rule."""
        response = self.client.responses.create(
            model="ft:gpt-4.1-2025-04-14:datadog-staging:terraform-examples:BzIuIlyu",
            input=[
                {"role": "system", "content": system_message},
                {"role": "developer", "content": code_snippet},
            ],
        )
        return response.output_text

    def send_rule_doc_formatting_request(self, metadata: dict) -> Output:
        system_message = """You are to format some text. You will be provided with a rule name, line break, @@@@@, line break and a rule description.
Your job is to ensure that the case of the rule's name is good: except the first word, proper nouns and acronyms, no word should start with a capital case.
Your job is to reformulate the description to make it as clear as possible."""

        queryName = metadata["queryName"]
        descriptionText = metadata["descriptionText"]
        response = self.client.responses.parse(
            model="gpt-5-mini-2025-08-07",
            input=[
                {"role": "system", "content": system_message},
                {
                    "role": "user",
                    "content": f"queryName: {queryName}\n\descriptionText: {descriptionText}",
                },
            ],
            text_format=Output,
        )
        return response.output_parsed
