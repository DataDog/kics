---
title: "VPC Without Network Firewall"
meta:
  name: "aws/vpc_without_network_firewall"
  id: "fd632aaf-b8a1-424d-a4d1-0de22fd3247a"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/vpc_without_network_firewall`

**Id:** `fd632aaf-b8a1-424d-a4d1-0de22fd3247a`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
VPC should have a Network Firewall associated

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/networkfirewall_firewall#vpc_id)

## Non-Compliant Code Examples
```terraform
resource "aws_vpc" "positive" {
  cidr_block = "10.0.0.0/16"
}

```

## Compliant Code Examples
```terraform
resource "aws_vpc" "negative" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_networkfirewall_firewall" "example" {
  name                = "example"
  firewall_policy_arn = aws_networkfirewall_firewall_policy.example.arn
  vpc_id              = aws_vpc.negative.id
  subnet_mapping {
    subnet_id = aws_subnet.example.id
  }

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}

```