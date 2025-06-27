---
title: "Stack Retention Disabled"
meta:
  name: "aws/stack_retention_disabled"
  id: "6e0e2f68-3fd9-4cd8-a5e4-e2213ef0df97"
  display_name: "Stack Retention Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Backup"
---
## Metadata

**Name:** `aws/stack_retention_disabled`

**Query Name** `Stack Retention Disabled`

**Id:** `6e0e2f68-3fd9-4cd8-a5e4-e2213ef0df97`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Backup

## Description
When defining an `aws_cloudformation_stack_set_instance` resource in Terraform, it is important to set the `retain_stack` attribute to `true`. If `retain_stack` is set to `false` or omitted (the default value is `false`), the underlying CloudFormation stack and all associated resources will be deleted when the stack set instance is destroyed or removed from the configuration. This creates a risk of accidental and irreversible data loss, as resources could be unintentionally deleted during operations such as stack set updates, deletions, or when Terraform destroy is executed. Ensuring that `retain_stack` is enabled (`retain_stack = true`) helps protect critical infrastructure by leaving the stack and its resources intact even after the stack set instance is removed, allowing for manual intervention or recovery if needed.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack_set_instance#stack_set_name)


## Compliant Code Examples
```terraform
resource "aws_cloudformation_stack_set_instance" "negative1" {
  account_id     = "123456789012"
  region         = "us-east-1"
  stack_set_name = aws_cloudformation_stack_set.example.name
  retain_stack     = true
}
```
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