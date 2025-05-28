---
title: "HTTP Port Open To Internet"
meta:
  name: "aws/http_port_open"
  id: "ffac8a12-322e-42c1-b9b9-81ff85c39ef7"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/http_port_open`

**Id:** `ffac8a12-322e-42c1-b9b9-81ff85c39ef7`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
The HTTP port is open to the internet in a Security Group

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```terraform
resource "aws_security_group" "positive1" {
  name        = "http_positive_tcp_1"
  description = "Gets the HTTP port open with the tcp protocol"

  ingress {
    description = "HTTP port open"
    from_port   = 78
    to_port     = 91
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "positive2" {
  name        = "http_positive_tcp_2"
  description = "Gets the HTTP port open with the tcp protocol"

  ingress {
    description = "HTTP port open"
    from_port   = 60
    to_port     = 85
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.2/0"]
  }

  ingress {
    description = "HTTP port open"
    from_port   = 65
    to_port     = 81
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_security_group" "negative1" {
  name        = "negative_http"
  description = "Doesn't get the HTTP port open"
}

resource "aws_security_group" "negative2" {

  ingress {
    from_port   = 70
    to_port     = 81
    protocol    = "tcp"
  }
}

resource "aws_security_group" "negative3" {

  ingress {
    from_port   = 79
    to_port     = 100
    protocol    = "tcp"
    cidr_blocks = ["0.1.0.0/0"]
  }
}

```