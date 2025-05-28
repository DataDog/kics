---
title: "SQL Analysis Services Port 2383 (TCP) Is Publicly Accessible"
meta:
  name: "aws/sql_analysis_services_port_2383_is_publicly_accessible"
  id: "54c417bf-c762-48b9-9d31-b3d87047e3f0"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/sql_analysis_services_port_2383_is_publicly_accessible`

**Id:** `54c417bf-c762-48b9-9d31-b3d87047e3f0`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Check if port 2383 on TCP is publicly accessible by checking the CIDR block range that can access it.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```terraform
resource "aws_security_group" "positive1" {
  name        = "allow_tls_1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2300
    to_port     = 2400
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
resource "aws_security_group" "positive2" {
  name        = "allow_tls_2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2380
    to_port     = 2390
    protocol    = "tcp"
    cidr_blocks = ["0.1.0.0/0"]
  }

   ingress {
    description = "TLS from VPC"
    from_port   = 2350
    to_port     = 2384
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_security_group" "negative1" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "tcp"
    cidr_blocks = ["0.1.0.0/0"]
  }
}

```