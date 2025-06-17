---
title: "CMK Is Unusable"
meta:
  name: "terraform/cmk_is_unusable"
  id: "7350fa23-dcf7-4938-916d-6a60b0c73b50"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata
**Name:** `terraform/cmk_is_unusable`
**Id:** `7350fa23-dcf7-4938-916d-6a60b0c73b50`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Availability
## Description
This check ensures that AWS Key Management Service (KMS) Customer Master Keys (CMKs) are configured with the 'is_enabled' attribute set to TRUE, making them usable for cryptographic operations such as encryption and decryption. If this attribute is set to FALSE, the CMK becomes unusable, preventing applications and services from accessing encrypted data or generating new data keys. Leaving CMKs disabled can disrupt access to critical resources and services, potentially resulting in application outages, data unavailability, or failed backup and restore operations. Properly enabling CMKs is crucial to maintaining secure and continuous access to resources that depend on AWS KMS.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key#is_enabled)

## Non-Compliant Code Examples
```aws
resource "aws_kms_key" "a" {
  description             = "KMS key 1"
  is_enabled = false
}

```

## Compliant Code Examples
```aws
resource "aws_kms_key" "a3" {
  description             = "KMS key 1"
  is_enabled = true
}

```