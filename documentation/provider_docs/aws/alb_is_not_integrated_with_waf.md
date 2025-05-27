---
title: "ALB Is Not Integrated With WAF"
meta:
  name: "aws/alb_is_not_integrated_with_waf"
  id: "0afa6ab8-a047-48cf-be07-93a2f8c34cf7"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/alb_is_not_integrated_with_waf`

**Id:** `0afa6ab8-a047-48cf-be07-93a2f8c34cf7`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
All Application Load Balancers (ALB) must be protected with Web Application Firewall (WAF) service

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafregional_web_acl_association)

## Non-Compliant Code Examples
```terraform
resource "aws_lb" "alb" {
  name               = "test-lb-tf"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.lb_sg.id]
  subnets            = [for subnet in aws_subnet.public : subnet.id]
}

resource "aws_wafv2_web_acl_association" "alb_waf_association" {
  resource_arn = aws_lb.alba.arn
  web_acl_arn  = aws_wafv2_web_acl.example.arn
}
```

```terraform
resource "aws_alb" "foo" {
  internal = false
  subnets  = [aws_subnet.foo.id, aws_subnet.bar.id]
}

resource "aws_wafregional_web_acl_association" "foo_waf" {
  resource_arn = aws_alb.fooooo.arn
  web_acl_id   = aws_wafregional_web_acl.foo.id
}


```

## Compliant Code Examples
```terraform
resource "aws_lb" "alb" {
  name               = "test-lb-tf"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.lb_sg.id]
  subnets            = [for subnet in aws_subnet.public : subnet.id]
}

resource "aws_wafv2_web_acl_association" "alb_waf_association" {
  resource_arn = aws_lb.alb.arn
  web_acl_arn  = aws_wafv2_web_acl.example.arn
}
```

```terraform
resource "aws_alb" "foo33" {
  internal = false
  subnets  = [aws_subnet.foo.id, aws_subnet.bar.id]
}

resource "aws_wafregional_web_acl_association" "foo_waf33" {
  resource_arn = aws_alb.foo33.arn
  web_acl_id   = aws_wafregional_web_acl.foo.id
}
# trigger validation

```