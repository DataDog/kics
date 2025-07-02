---
title: "Lambda Permission Misconfigured"
meta:
  name: "aws/lambda_permission_misconfigured"
  id: "75ec6890-83af-4bf1-9f16-e83726df0bd0"
  display_name: "Lambda Permission Misconfigured"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata

**Id:** `75ec6890-83af-4bf1-9f16-e83726df0bd0`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Low

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

### Description

 This check examines the action field in the aws_lambda_permission resource to ensure it is set to "lambda:InvokeFunction". When this field is misconfigured with broader or unintended actions, such as "lambda:DeleteFunction", it grants unnecessary or overly permissive access to the Lambda function. This can allow third-party AWS services or principals to perform destructive or unintended operations on the function, increasing the risk of unauthorized deletion, modification, or misuse. If left unaddressed, this misconfiguration could result in loss of critical business logic, disruption of service, or escalation of privileges within your cloud environment.


## Compliant Code Examples
```terraform
resource "aws_lambda_permission" "negative1" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.logging.function_name
  principal     = "logs.eu-west-1.amazonaws.com"
  source_arn    = "${aws_cloudwatch_log_group.default.arn}:*"
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_lambda_permission" "positive1" {
  action        = "lambda:DeleteFunction"
  function_name = aws_lambda_function.logging.function_name
  principal     = "logs.eu-west-1.amazonaws.com"
  source_arn    = "${aws_cloudwatch_log_group.default.arn}:*"
}

```