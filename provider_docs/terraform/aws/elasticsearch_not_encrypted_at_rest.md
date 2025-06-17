---
title: "ElasticSearch Not Encrypted At Rest"
meta:
  name: "terraform/elasticsearch_not_encrypted_at_rest"
  id: "24e16922-4330-4e9d-be8a-caa90299466a"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/elasticsearch_not_encrypted_at_rest`
**Id:** `24e16922-4330-4e9d-be8a-caa90299466a`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
This check ensures that AWS ElasticSearch domains have encryption at rest enabled, which protects sensitive data stored in ElasticSearch indices from unauthorized access if physical storage is compromised. Without encryption at rest, data stored in ElasticSearch is vulnerable to exposure if someone gains access to the underlying storage media, potentially leading to data breaches and compliance violations.

To properly secure ElasticSearch, you must explicitly configure the encrypt_at_rest block with enabled set to true:

```
resource "aws_elasticsearch_domain" "example" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
  }
}
```

Insecure configurations either omit the encrypt_at_rest block entirely or explicitly set enabled to false.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain)

## Non-Compliant Code Examples
```aws
resource "aws_elasticsearch_domain" "positive1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"
}

resource "aws_elasticsearch_domain" "positive2" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = false
  }
}
```

## Compliant Code Examples
```aws
resource "aws_elasticsearch_domain" "negative1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
  }
}
```