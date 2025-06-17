---
title: "Lambda IAM InvokeFunction Misconfigured"
meta:
  name: "terraform/lambda_iam_invokefunction_misconfigured"
  id: "0ca1017d-3b80-423e-bb9c-6cd5898d34bd"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/lambda_iam_invokefunction_misconfigured`
**Id:** `0ca1017d-3b80-423e-bb9c-6cd5898d34bd`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Best Practices
## Description
AWS Lambda permissions must be carefully defined so that the `Action` field in the IAM policy explicitly specifies allowed actions, such as `"lambda:InvokeFunction"`. If the `Action` field is omitted or set too broadly, it could inadvertently grant unnecessary permissions, allowing unintended users or services to perform privileged operations on the Lambda function. This misconfiguration increases the risk of unauthorized invocation or modification of Lambda functions, potentially leading to security breaches or the execution of malicious code.

A secure Terraform example ensures the `Action` is correctly specified:

```
policy = jsonencode({
  Version = "2012-10-17"
  Statement = [
    {
      Action = [
        "lambda:InvokeFunction",
      ]
      Effect   = "Allow"
      Resource = [
        "arn:aws:lambda:*:*:function:example-function",
        "arn:aws:lambda:*:*:function:example-function:*"
      ]
    },
  ]
})
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

## Non-Compliant Code Examples
```aws
resource "aws_iam_policy" "positive1policy" {
  name        = "positive1policy"
  path        = "/"
  description = "Positive1 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction",
        ]
        Effect   = "Allow"
        Resource = [
            "arn:aws:lambda:*:*:function:positive1"
        ]
      },
    ]
  })
}

```

```aws
resource "aws_iam_policy" "positive4policy" {
  name        = "positive4policy"
  path        = "/"
  description = "positive4 Policy"
  policy      = data.aws_iam_policy_document.datapositive4policy.json
}
# Terraform's "jsonencode" function converts a
# Terraform expression result to valid JSON syntax.
data "aws_iam_policy_document" "datapositive4policy" {
  statement {
    effect = "Allow"
    actions = [
      "lambda:InvokeFunction"
    ]

    resources = [
      "arn:aws:lambda:*:*:function:*:*"
    ]
  }
}

```

```aws
resource "aws_iam_policy" "positive6policy" {
  name        = "positive6policy"
  path        = "/"
  description = "positive6 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2022-20-27"
    Statement = [
      {
        Action = [
          "lambda:*",
        ]
        Effect   = "Allow"
        Resource = [
            "arn:aws:lambda:*:*:function:*:*"
        ]
      },
    ]
  })
}

```

## Compliant Code Examples
```aws
resource "aws_iam_policy" "negative1policy" {
  name        = "negative1policy"
  path        = "/"
  description = "negative1 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction",
        ]
        Effect   = "Allow"
        Resource = [
            "arn:aws:lambda:*:*:function:negative1",
            "arn:aws:lambda:*:*:function:negative1:*"
        ]
      },
    ]
  })
}

```

```aws
resource "aws_lambda_function" "negative3" {
  function_name = "negative3"
  role          = "negative3_role"
}

resource "aws_iam_policy" "negative3policy" {
  name        = "negative3policy"
  path        = "/"
  description = "negative3 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:*",
        ]
        Effect   = "Allow"
        Resource = [
            aws_lambda_function.negative3.arn,
            "${aws_lambda_function.negative3.arn}:*"
        ]
      },
    ]
  })
}

```

```aws
resource "aws_iam_policy" "negative2policy" {
  name        = "negative2policy"
  path        = "/"
  description = "negative2 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:*",
        ]
        Effect   = "Allow"
        Resource = ["*"]
      },
    ]
  })
}

```