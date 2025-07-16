from openai import OpenAI


class RulesGenerator:

    def __init__(self):
        self.client = OpenAI()

    def send_rule_request(self, code_snippet: str) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure add support for the modules. 
    Answer with only the refactored code."""
        response = self.client.responses.create(
            model="ft:gpt-4.1-2025-04-14:datadog-staging:module-support:Bs7l0H5K",
            input=[
                {"role": "system", "content": system_message},
                {"role": "developer", "content": code_snippet},
            ],
        )
        return response.output_text

    def send_terraform_request(self, code_snippet: str) -> str:
        system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure, create examples tf files using the related module.
    Answer with a positive example that will trigger the rule, one new line, "#####", a second new line and a negative example that won't trigger the rule."""
        response = self.client.responses.create(
            model="ft:gpt-4.1-2025-04-14:datadog-staging:module-support:Bs7l0H5K",
            input=[
                {"role": "system", "content": system_message},
                {"role": "developer", "content": code_snippet},
            ],
        )
        return response.output_text
