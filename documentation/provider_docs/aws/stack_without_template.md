---
title: "Stack Without Template"
meta:
  name: "aws/stack_without_template"
  id: "91bea7b8-0c31-4863-adc9-93f6177266c4"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Build Process"
---

## Metadata
**Name:** `aws/stack_without_template`

**Id:** `91bea7b8-0c31-4863-adc9-93f6177266c4`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Build Process

## Description
AWS CloudFormation should have a template defined through the attribute template_url or attribute template_body

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

     template_url = "sometemplateurl"
}



resource "aws_cloudformation_stack" "negative2" {

     name = "networking-stack"

     parameters = {
     VPCCidr = "10.0.0.0/16"
     }

     template_body = "sometemplatebody"
}

```