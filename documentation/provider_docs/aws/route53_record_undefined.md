---
title: "Route53 Record Undefined"
meta:
  name: "aws/route53_record_undefined"
  id: "25db74bf-fa3b-44da-934e-8c3e005c0453"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/route53_record_undefined`

**Id:** `25db74bf-fa3b-44da-934e-8c3e005c0453`

**Cloud Provider:** aws

**Severity:** High

**Category:** Networking and Firewall

## Description
This check verifies that Route53 record resources have defined values in the 'records' array. Empty record arrays in Route53 configurations may result in DNS resolution failures, causing service disruptions and potentially breaking application functionality that depends on proper name resolution.

A properly configured Route53 record should include appropriate values in the records array as shown in the secure example below:

```terraform
records = [
  aws_route53_zone.example.name_servers[0],
  aws_route53_zone.example.name_servers[1],
  aws_route53_zone.example.name_servers[2],
  aws_route53_zone.example.name_servers[3],
]
```

Insecure configurations leave the records array empty:

```terraform
records = [
  
]
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record)

## Non-Compliant Code Examples
```terraform
resource "aws_route53_record" "example" {
  allow_overwrite = true
  name            = "test.example.com"
  ttl             = 30
  type            = "NS"
  zone_id         = aws_route53_zone.example.zone_id

  records = [
    
  ]
}
```

## Compliant Code Examples
```terraform
resource "aws_route53_record" "example" {
  allow_overwrite = true
  name            = "test.example.com"
  ttl             = 30
  type            = "NS"
  zone_id         = aws_route53_zone.example.zone_id

  records = [
    aws_route53_zone.example.name_servers[0],
    aws_route53_zone.example.name_servers[1],
    aws_route53_zone.example.name_servers[2],
    aws_route53_zone.example.name_servers[3],
  ]
}
```