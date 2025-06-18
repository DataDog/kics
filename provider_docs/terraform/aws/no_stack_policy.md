---
title: "No Stack Policy"
meta:
  name: "terraform/no_stack_policy"
  id: "2f01fb2d-828a-499d-b98e-b83747305052"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata
**Name:** `terraform/no_stack_policy`
**Id:** `2f01fb2d-828a-499d-b98e-b83747305052`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Resource Management
## Description
AWS CloudFormation stacks should implement a stack policy, set using the `policy_body` or `policy_url` attributes in Terraform, to restrict which actions can modify or delete critical resources within the stack. Without a stack policy, any update operation could unintentionally overwrite, disrupt, or remove essential resources, increasing the risk of accidental outages or security issues. Properly configuring a stack policy helps enforce change management controls and reduces the attack surface for unauthorized or accidental actions.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack)

## Non-Compliant Code Examples
```aws
resource "aws_cloudformation_stack" "positive1" {

  name = "networking-stack"

  parameters = {
    VPCCidr = "10.0.0.0/16"
  }

}

```

## Compliant Code Examples
```aws
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