---
title: "Stack Notifications Disabled"
meta:
  name: "aws/stack_notifications_disabled"
  id: "b72d0026-f649-4c91-a9ea-15d8f681ac09"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/stack_notifications_disabled`

**Id:** `b72d0026-f649-4c91-a9ea-15d8f681ac09`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
AWS CloudFormation should have stack notifications enabled to be notified when an event occurs

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudformation_stack" "positive1" {

  name = "networking-stack"

  parameters = {
    VPCCidr = "10.0.0.0/16"
  }


}
```

## Compliant Code Examples
```terraform
resource "aws_cloudformation_stack" "negative1" {

  name = "networking-stack"

  parameters = {
    VPCCidr = "10.0.0.0/16"
  }


  notification_arns = ["a","b"]

}
```