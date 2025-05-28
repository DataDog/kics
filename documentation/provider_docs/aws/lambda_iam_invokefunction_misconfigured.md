---
title: "Lambda IAM InvokeFunction Misconfigured"
meta:
  name: "aws/lambda_iam_invokefunction_misconfigured"
  id: "0ca1017d-3b80-423e-bb9c-6cd5898d34bd"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/lambda_iam_invokefunction_misconfigured`

**Id:** `0ca1017d-3b80-423e-bb9c-6cd5898d34bd`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
Lambda permission may be misconfigured if the action field is not filled in by 'lambda:InvokeFunction'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

## Non-Compliant Code Examples
```terraform
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

```terraform
resource "aws_iam_policy" "positive2policy" {
  name        = "positive2policy"
  path        = "/"
  description = "Positive2 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2022-20-27"
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction",
        ]
        Effect   = "Allow"
        Resource = [
            "arn:aws:lambda:*:*:function:positive2*:*"
        ]
      },
    ]
  })
}

```

```terraform
resource "aws_iam_policy" "positive3policy" {
  name        = "positive3policy"
  path        = "/"
  description = "positive3 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2022-20-27"
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction",
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

```terraform
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

```terraform
resource "aws_iam_policy" "positive5policy" {
  name        = "positive5policy"
  path        = "/"
  description = "positive5 Policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2022-20-27"
    Statement = [
      {
        Action = [
          "*",
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

```terraform
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

## Compliant Code Examples
```terraform
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

```terraform
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

```terraform
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