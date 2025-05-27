---
title: "SSO Permission With Inadequate User Session Duration"
meta:
  name: "aws/sso_permission_with_inadequate_user_session_duration"
  id: "ce9dfce0-5fc8-433b-944a-3b16153111a8"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Access Control"
---

## Metadata
**Name:** `aws/sso_permission_with_inadequate_user_session_duration`

**Id:** `ce9dfce0-5fc8-433b-944a-3b16153111a8`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Access Control

## Description
SSO permissions should be configured to limit user sessions to no longer than 1 hour. Allowing longer sessions can increase the risk of unauthorized access or session hijacking. This is a best practice for security and should be implemented in SSO permission settings.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ssoadmin_permission_set)

## Non-Compliant Code Examples
```terraform
resource "aws_ssoadmin_permission_set" "example3" {
  name             = "Example"
  description      = "An example"
  instance_arn     = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  relay_state      = "https://s3.console.aws.amazon.com/s3/home?region=us-east-1#"
  session_duration = "PT1H1M"
}

resource "aws_ssoadmin_permission_set" "example4" {
  name             = "Example"
  description      = "An example"
  instance_arn     = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  relay_state      = "https://s3.console.aws.amazon.com/s3/home?region=us-east-1#"
  session_duration = "PT2H"
}

```

## Compliant Code Examples
```terraform
resource "aws_ssoadmin_permission_set" "example" {
  name             = "Example"
  description      = "An example"
  instance_arn     = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  relay_state      = "https://s3.console.aws.amazon.com/s3/home?region=us-east-1#"
  session_duration = "PT1H"
}

resource "aws_ssoadmin_permission_set" "example2" {
  name             = "Example"
  description      = "An example"
  instance_arn     = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  relay_state      = "https://s3.console.aws.amazon.com/s3/home?region=us-east-1#"
}


```