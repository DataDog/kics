---
title: "Sensitive Port Is Exposed To Entire Network"
meta:
  name: "terraform/sensitive_port_is_exposed_to_entire_network"
  id: "381c3f2a-ef6f-4eff-99f7-b169cda3422c"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/sensitive_port_is_exposed_to_entire_network`
**Id:** `381c3f2a-ef6f-4eff-99f7-b169cda3422c`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Networking and Firewall
## Description
This vulnerability occurs when security groups allow inbound traffic to sensitive ports (such as SSH, RDP, database, or administrative service ports) from the entire internet (0.0.0.0/0 or /0). Exposing sensitive ports to the internet creates a significant security risk as it allows potential attackers from anywhere to attempt connections, potentially leading to unauthorized access, data breaches, or service exploitation.

To remediate this issue, restrict access to sensitive ports by specifying narrower CIDR ranges or specific IP addresses in your security group rules. For example, instead of using `cidr_blocks = ["0.0.0.0/0"]`, use specific IP ranges like `cidr_blocks = [aws_vpc.main.cidr_block]` or `cidr_blocks = ["10.0.0.0/16"]` that only allow traffic from trusted networks.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```aws
resource "aws_security_group" "positive1" {
  name        = "allow_tls1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2200
    to_port     = 2500
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_security_group" "positive2" {
  name        = "allow_tls2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 20
    to_port     = 60
    protocol    = "tcp"
    cidr_blocks = ["/0"]
  }
}


resource "aws_security_group" "positive3" {
  name        = "allow_tls3"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 5000
    to_port     = 6000
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_security_group" "positive4" {
  name        = "allow_tls4"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 20
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["/0"]
  }
}


resource "aws_security_group" "positive5" {
  name        = "allow_tls5"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 445
    to_port     = 500
    protocol    = "udp"
    cidr_blocks = ["1.1.1.1/1","0.0.0.0/0",  "2.2.3.4/12"]
  }
}


resource "aws_security_group" "positive6" {
  name        = "allow_tls6"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 135
    to_port     = 170
    protocol    = "udp"
    cidr_blocks = ["10.68.0.0", "0.0.0.0/0"]
  }
}


resource "aws_security_group" "positive7" {
  name        = "allow_tls7"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "udp"
    cidr_blocks = ["/0", "1.2.3.4/12"]
  }
}


resource "aws_security_group" "positive8" {
  name        = "allow_tls8"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["/0"]
  }
}

```

## Compliant Code Examples
```aws
resource "aws_security_group" "negative1" {
  name        = "allow_tls1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.main.cidr_block]
  }
}


resource "aws_security_group" "negative2" {
  name        = "allow_tls2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2384
    to_port     = 2386
    protocol    = "tcp"
    cidr_blocks = ["/0"]
  }
}


resource "aws_security_group" "negative3" {
  name        = "allow_tls3"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "tcp"
    cidr_blocks = ["1.2.3.4/5"]
  }
}


resource "aws_security_group" "negative4" {
  name        = "allow_tls4"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "tcp"
    cidr_blocks = ["1.2.3.4/5"]
  }
}


resource "aws_security_group" "negative5" {
  name        = "allow_tls5"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "udp"
    cidr_blocks = ["1.2.3.4/5","0.0.0.0/12"]
  }
}


resource "aws_security_group" "negative6" {
  name        = "allow_tls6"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["1.2.3.4/5","0.0.0.0/12"]
  }
}
```