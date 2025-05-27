---
title: "Shield Advanced Not In Use"
meta:
  name: "aws/shield_advanced_not_in_use"
  id: "084c6686-2a70-4710-91b1-000393e54c12"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/shield_advanced_not_in_use`

**Id:** `084c6686-2a70-4710-91b1-000393e54c12`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
AWS Shield Advanced should be used for Amazon Route 53 hosted zone, AWS Global Accelerator accelerator, Elastic IP Address, Elastic Load Balancing, and Amazon CloudFront Distribution to protect these resources against robust DDoS attacks

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/shield_protection#resource_arn)

## Non-Compliant Code Examples
```terraform
resource "aws_route53_zone" "positive2" {
  name = "example.com"
}

resource "aws_shield_protection" "positive2" {
  name         = "example"
  resource_arn = aws_route53_zone.positive.arn

  tags = {
    Environment = "Dev"
  }
}

```

```terraform
data "aws_availability_zones" "available" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "aws_eip" "positive1" {
  vpc = true
}

resource "aws_shield_protection" "positive1" {
  name         = "example"
  resource_arn = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:eip-allocation/${aws_eip.positive.id}"

  tags = {
    Environment = "Dev"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_route53_zone" "negative2" {
  name = "example.com"
}

resource "aws_shield_protection" "negative2" {
  name         = "example"
  resource_arn = aws_route53_zone.negative2.arn

  tags = {
    Environment = "Dev"
  }
}

```

```terraform
data "aws_availability_zones" "available" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "aws_eip" "negative1" {
  vpc = true
}

resource "aws_shield_protection" "negative1" {
  name         = "example"
  resource_arn = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:eip-allocation/${aws_eip.negative1.id}"

  tags = {
    Environment = "Dev"
  }
}

```