from openai import OpenAI


class RulesGenerator:

    #    system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure add support for the modules.
    #    Answer with only the refactored code, two new lines, a "#####" separator and an example terraform file to test on."""

    system_message = """You're an expert writing Rego and also in Terraform and its module support. When provided with a Rego file about a Terraform infrastructure add support for the modules. 
    Answer with only the refactored code."""

    def __init__(self):
        self.client = OpenAI()

    def send_request(self, code_snippet: str) -> str:
        response = self.client.responses.create(
            model="ft:gpt-4.1-2025-04-14:datadog-staging:module-support:Bs7l0H5K",
            input=[
                {"role": "system", "content": self.system_message},
                {"role": "developer", "content": code_snippet},
            ],
        )
        return response.output_text
