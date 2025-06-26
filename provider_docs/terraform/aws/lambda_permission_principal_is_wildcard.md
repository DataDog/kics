---
title: "Lambda Permission Principal Is Wildcard"
meta:
  name: "terraform/lambda_permission_principal_is_wildcard"
  id: "e08ed7eb-f3ef-494d-9d22-2e3db756a347"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/lambda_permission_principal_is_wildcard`
**Id:** `e08ed7eb-f3ef-494d-9d22-2e3db756a347`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Lambda function permissions should not define the `principal` attribute with a wildcard (`*`) value. Using a wildcard as the principal allows any AWS account or service to invoke the Lambda function, significantly increasing the risk of unauthorized access or unintentional exposure. Instead, the `principal` should be set to the specific AWS service or account that requires access, such as `events.amazonaws.com`, to enforce strict access controls and limit potential abuse.

#### Learn More

 - [Provider Reference](https://docs.ansible.com/ansible/latest/collections/community/aws/lambda_policy_module.html)

## Non-Compliant Code Examples
```aws
resource "aws_lambda_permission" "positive1" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "*"
  source_arn    = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily"
  qualifier     = aws_lambda_alias.test_alias.name
}

```

## Compliant Code Examples
```aws
resource "aws_lambda_permission" "negative1" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily"
  qualifier     = aws_lambda_alias.test_alias.name
}

```