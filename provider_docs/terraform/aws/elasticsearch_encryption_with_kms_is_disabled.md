---
title: "ElasticSearch Encryption With KMS Disabled"
meta:
  name: "terraform/elasticsearch_encryption_with_kms_is_disabled"
  id: "7af2f4a3-00d9-47f3-8d15-ca0888f4e5b2"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/elasticsearch_encryption_with_kms_is_disabled`
**Id:** `7af2f4a3-00d9-47f3-8d15-ca0888f4e5b2`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
ElasticSearch domains should use AWS Key Management Service (KMS) for encryption at rest to provide enhanced security. While enabling basic encryption at rest is important, not specifying a KMS key ID means ElasticSearch will use default AWS-managed keys rather than customer-managed keys, reducing your control over the encryption process. Without KMS encryption, sensitive data stored in ElasticSearch could be at risk if unauthorized access to the storage media occurs.

To properly implement KMS encryption, ensure the 'encrypt_at_rest' block includes both 'enabled = true' and a specific 'kms_key_id' as shown below:

```terraform
encrypt_at_rest {
    enabled = true
    kms_key_id = "your-kms-key-id"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain)

## Non-Compliant Code Examples
```aws
resource "aws_elasticsearch_domain" "positive1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
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
      kms_key_id = "some-key-id"
  }
}
```