---
title: "Security Group Rule Without Description"
meta:
  name: "aws/security_group_without_description"
  id: "cb3f5ed6-0d18-40de-a93d-b3538db31e8c"
  display_name: "Security Group Rule Without Description"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "INFO"
  category: "Best Practices"
---
## Metadata
**Name:** `aws/security_group_without_description`
**Query Name** `Security Group Rule Without Description`
**Id:** `cb3f5ed6-0d18-40de-a93d-b3538db31e8c`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Info
**Category:** Best Practices
## Description
It is a best practice for AWS Security Groups to include a meaningful `description` attribute in their Terraform configuration, as in:

```
description = "Allow TLS inbound traffic"
```

Omitting the description field, as shown below,

```
resource "aws_security_group" "allow_tls" {
  name   = "allow_tls"
  vpc_id = aws_vpc.main.id
  // missing description
  ...
}
```

can lead to confusion and hinder effective management or auditing of security groups, especially in environments with many resources. Without clear descriptions, security teams may struggle to quickly identify the purpose of a group, increasing the risk of misconfigurations and delayed incident response.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group#description)


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