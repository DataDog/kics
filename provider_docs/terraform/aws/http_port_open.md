---
title: "HTTP Port Open To Internet"
meta:
  name: "terraform/http_port_open"
  id: "ffac8a12-322e-42c1-b9b9-81ff85c39ef7"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/http_port_open`
**Id:** `ffac8a12-322e-42c1-b9b9-81ff85c39ef7`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Opening HTTP ports (for example, using `from_port = 80` and `to_port = 80`) to the internet in a Security Group by setting `cidr_blocks = ["0.0.0.0/0"]` allows unrestricted access from any IP address. This misconfiguration exposes resources to potential unauthorized access and attacks, such as brute force or exploitation of unpatched web services. To mitigate this risk, restrict `cidr_blocks` to trusted IP addresses or use security controls that limit ingress to necessary sources only, as in the secure example below:

```
resource "aws_security_group" "secure_http" {
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["203.0.113.0/24"] // Restrict to trusted IP ranges only
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```aws
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
```aws
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