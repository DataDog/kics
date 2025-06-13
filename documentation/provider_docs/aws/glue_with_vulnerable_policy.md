---
title: "Glue With Vulnerable Policy"
meta:
  name: "aws/glue_with_vulnerable_policy"
  id: "d25edb51-07fb-4a73-97d4-41cecdc53a22"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/glue_with_vulnerable_policy`

**Id:** `d25edb51-07fb-4a73-97d4-41cecdc53a22`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Glue policy should avoid wildcard in 'principals' and 'actions'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/glue_resource_policy#policy)

## Non-Compliant Code Examples
```terraform
data "aws_iam_policy_document" "glue-example-policy" {
  statement {
    actions = [
      "glue:*",
    ]
    resources = ["arn:data.aws_partition.current.partition:glue:data.aws_region.current.name:data.aws_caller_identity.current.account_id:*"]
    principals {
      identifiers = ["*"]
      type        = "AWS"
    }
  }
}

resource "aws_glue_resource_policy" "example" {
  policy = data.aws_iam_policy_document.glue-example-policy.json
}

```

## Compliant Code Examples
```terraform
data "aws_iam_policy_document" "glue-example-policy2" {
  statement {
    actions = [
      "glue:CreateTable",
    ]
    resources = ["arn:data.aws_partition.current.partition:glue:data.aws_region.current.name:data.aws_caller_identity.current.account_id:*"]
    principals {
      identifiers = ["arn:aws:iam::var.account_id:saml-provider/var.provider_name"]
      type        = "AWS"
    }
  }
}

resource "aws_glue_resource_policy" "example2" {
  policy = data.aws_iam_policy_document.glue-example-policy2.json
}

```