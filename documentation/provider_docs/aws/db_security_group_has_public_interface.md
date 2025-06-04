---
title: "DB Security Group Has Public Interface"
meta:
  name: "aws/db_security_group_has_public_interface"
  id: "f0d8781f-99bf-4958-9917-d39283b168a0"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/db_security_group_has_public_interface`

**Id:** `f0d8781f-99bf-4958-9917-d39283b168a0`

**Cloud Provider:** aws

**Severity:** High

**Category:** Insecure Configurations

## Description
AWS DB Security Groups control access to RDS database instances by defining which IP addresses or Amazon EC2 instances can connect to them. Configuring a DB Security Group with a public interface (0.0.0.0/0) allows unrestricted access from any IP address, potentially exposing your database to unauthorized access and attacks from the internet.

This vulnerability significantly increases the risk of data breaches, unauthorized data access, and potential compromise of sensitive information stored in your database. Instead of using public interfaces, you should restrict access to specific IP ranges or VPC CIDR blocks.

Secure example:
```
resource "aws_db_security_group" "example" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/8"
  }
}
```

Insecure example:
```
resource "aws_db_security_group" "example" {
  name = "rds_sg"

  ingress {
    cidr = "0.0.0.0/0"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_security_group)

## Non-Compliant Code Examples
```terraform
resource "aws_db_security_group" "positive1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/8"
  }

  ingress {
    cidr = "0.0.0.0/0"
  }
}

```

```terraform
resource "aws_db_security_group" "positive1" {
  name = "rds_sg"

  ingress {
    cidr = "0.0.0.0/0"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_db_security_group" "negative1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/8"
  }
}

```