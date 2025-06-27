---
title: "ElasticSearch Not Encrypted At Rest"
meta:
  name: "aws/elasticsearch_not_encrypted_at_rest"
  id: "24e16922-4330-4e9d-be8a-caa90299466a"
  display_name: "ElasticSearch Not Encrypted At Rest"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `aws/elasticsearch_not_encrypted_at_rest`
**Query Name** `ElasticSearch Not Encrypted At Rest`
**Id:** `24e16922-4330-4e9d-be8a-caa90299466a`
**Cloud Provider:** aws
**Platform** Terraform
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


## Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "negative1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
  }
}
```
## Non-Compliant Code Examples
```terraform
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