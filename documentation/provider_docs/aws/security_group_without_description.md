---
title: "Security Group Rule Without Description"
meta:
  name: "aws/security_group_without_description"
  id: "cb3f5ed6-0d18-40de-a93d-b3538db31e8c"
  cloud_provider: "aws"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/security_group_without_description`

**Id:** `cb3f5ed6-0d18-40de-a93d-b3538db31e8c`

**Cloud Provider:** aws

**Severity:** Info

**Category:** Best Practices

## Description
It's considered a best practice for AWS Security Group to have a description

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group#description)

## Non-Compliant Code Examples
```terraform
resource "aws_security_group" "allow_tls" {
  name        = "allow_tls"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "TLS from VPC"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  tags = {
    Name = "allow_tls"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_security_group" "allow_tls" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "TLS from VPC"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  tags = {
    Name = "allow_tls"
  }
}

```