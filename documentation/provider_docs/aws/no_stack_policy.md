---
title: "No Stack Policy"
meta:
  name: "aws/no_stack_policy"
  id: "2f01fb2d-828a-499d-b98e-b83747305052"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Resource Management"
---

## Metadata
**Name:** `aws/no_stack_policy`

**Id:** `2f01fb2d-828a-499d-b98e-b83747305052`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Resource Management

## Description
AWS CloudFormation Stack should have a stack policy in order to protect stack resources from update actions

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

     policy_url = "somepolicyurl"
}



resource "aws_cloudformation_stack" "negative2" {

     name = "networking-stack"

     parameters = {
     VPCCidr = "10.0.0.0/16"
     }

     policy_body = "somepolicy"
}

```