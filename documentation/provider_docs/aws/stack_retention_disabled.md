---
title: "Stack Retention Disabled"
meta:
  name: "aws/stack_retention_disabled"
  id: "6e0e2f68-3fd9-4cd8-a5e4-e2213ef0df97"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Backup"
---

## Metadata
**Name:** `aws/stack_retention_disabled`

**Id:** `6e0e2f68-3fd9-4cd8-a5e4-e2213ef0df97`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Backup

## Description
Make sure that retain_stack is enabled to keep the Stack and it's associated resources during resource destruction

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack_set_instance#stack_set_name)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudformation_stack_set_instance" "positive1" {
  account_id     = "123456789012"
  region         = "us-east-1"
  stack_set_name = aws_cloudformation_stack_set.example.name
  retain_stack   = false
}

resource "aws_cloudformation_stack_set_instance" "positive2" {
  account_id     = "123456789012"
  region         = "us-east-1"
  stack_set_name = aws_cloudformation_stack_set.example.name
}
```

## Compliant Code Examples
```terraform
resource "aws_cloudformation_stack_set_instance" "negative1" {
  account_id     = "123456789012"
  region         = "us-east-1"
  stack_set_name = aws_cloudformation_stack_set.example.name
  retain_stack     = true
}
```