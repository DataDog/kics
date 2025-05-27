---
title: "Lambda Permission Principal Is Wildcard"
meta:
  name: "aws/lambda_permission_principal_is_wildcard"
  id: "e08ed7eb-f3ef-494d-9d22-2e3db756a347"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/lambda_permission_principal_is_wildcard`

**Id:** `e08ed7eb-f3ef-494d-9d22-2e3db756a347`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Lambda Permission Principal should not contain a wildcard.

#### Learn More

 - [Provider Reference](https://docs.ansible.com/ansible/latest/collections/community/aws/lambda_policy_module.html)

## Non-Compliant Code Examples
```terraform
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
```terraform
resource "aws_lambda_permission" "negative1" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily"
  qualifier     = aws_lambda_alias.test_alias.name
}

```