---
title: "Unrestricted Security Group Ingress"
meta:
  name: "terraform/unrestricted_security_group_ingress"
  id: "4728cd65-a20c-49da-8b31-9c08b423e4db"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/unrestricted_security_group_ingress`
**Id:** `4728cd65-a20c-49da-8b31-9c08b423e4db`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Networking and Firewall
## Description
This check identifies AWS security group rules that allow unrestricted inbound traffic from any IP address (0.0.0.0/0 for IPv4 or ::/0 for IPv6). Such configurations create a significant security vulnerability by exposing your resources to potential unauthorized access from anywhere on the internet, increasing the risk of data breaches and attacks.

Restricting inbound traffic to specific, trusted IP addresses or CIDR ranges is a security best practice that follows the principle of least privilege. Instead of using wide-open rules like `cidr_blocks = ["0.0.0.0/0"]`, configure your security groups with specific CIDR blocks as shown in the secure example: `cidr_blocks = ["0.0.2.0/0"]` or more targeted ranges like corporate IP addresses.

#### Learn More

 - [Provider Reference](https://www.terraform.io/docs/providers/aws/r/security_group.html)

## Non-Compliant Code Examples
```aws
module "web_server_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.3.0"

  name        = "web-server"
  description = "Security group for web-server with HTTP ports open within VPC"
  vpc_id      = "vpc-12345678"

  ingress_ipv6_cidr_blocks  = ["fc00::/8", "::/0"]
}

```

```aws
resource "aws_security_group" "positive7" {
  ingress {
    from_port         = 3306
    to_port           = 3306
    protocol          = "tcp"
    ipv6_cidr_blocks  = ["::/0"]
    security_group_id = aws_security_group.default.id
  }
}

```

```aws
resource "aws_security_group_rule" "positive1" {
  type              = "ingress"
  from_port         = 3306
  to_port           = 3306
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.default.id
}

```

## Compliant Code Examples
```aws
resource "aws_security_group_rule" "negative1" {
  type              = "ingress"
  from_port         = 3306
  to_port           = 3306
  protocol          = "tcp"
  cidr_blocks       = ["0.0.2.0/0"]
  security_group_id = aws_security_group.default.id
}


```

```aws
resource "aws_security_group" "negative3" {
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["1.0.0.0/0"]
  }

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.1.0/0"]
  }
}

```

```aws
resource "aws_security_group" "negative7" {
  ingress {
    from_port         = 3306
    to_port           = 3306
    protocol          = "tcp"
    ipv6_cidr_blocks  = ["fc00::/9"]
  }

  ingress {
    from_port         = 3306
    to_port           = 3306
    protocol          = "tcp"
    ipv6_cidr_blocks  = ["fc00::/8"]
  }
}

```