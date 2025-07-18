---
title: "Lambda With Vulnerable Policy"
group-id: "rules/terraform/aws"
meta:
  name: "aws/lambda_with_vulnerable_policy"
  id: "ad9dabc7-7839-4bae-a957-aa9120013f39"
  display_name: "Lambda With Vulnerable Policy"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `ad9dabc7-7839-4bae-a957-aa9120013f39`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission#action)

### Description

 AWS Lambda permissions with wildcard actions ('lambda:*') grant excessive privileges that violate the principle of least privilege, potentially allowing unauthorized operations on your Lambda functions. When wildcards are used, principals may execute unintended actions against your functions, leading to potential service disruption or data leakage. Instead of using wildcards like `action = "lambda:*"`, specify only the precise permissions needed, such as `action = "lambda:InvokeFunction"` to ensure proper access controls and reduce the attack surface of your Lambda resources.


## Compliant Code Examples
```terraform
resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily"
  qualifier     = aws_lambda_alias.test_alias.name
}

resource "aws_lambda_alias" "test_alias" {
  name             = "testalias"
  description      = "a sample description"
  function_name    = aws_lambda_function.test_lambda.function_name
  function_version = "$LATEST"
}

resource "aws_lambda_function" "test_lambda" {
  filename      = "lambdatest.zip"
  function_name = "lambda_function_name"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "exports.handler"
  runtime       = "nodejs12.x"
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

```
## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "my-lambda" {
  filename = "~/Downloads/lambda.json.zip"
  function_name = "my-lambda"
  role          = aws_iam_role.lambda-role.arn
  handler       = "lambda_function.lambda_handler"
  runtime = "python3.8"
}

resource "aws_iam_role" "lambda-role" {
  name = "lambda-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "all" {
  statement_id  = "AllowAllResources"
  action        = "lambda:*"
  function_name = aws_lambda_function.my-lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = "arn:aws:s3:::delete-me-us-east-1-permissions-tests"
  source_account = "111111111111"
  qualifier     = aws_lambda_alias.my-lambda-alias.name
}


resource "aws_lambda_alias" "my-lambda-alias" {
  name             = "v1"
  description      = "a sample description"
  function_name    = aws_lambda_function.my-lambda.function_name
  function_version = "$LATEST"
}

```