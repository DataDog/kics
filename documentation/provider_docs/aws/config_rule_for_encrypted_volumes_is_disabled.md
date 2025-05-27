---
title: "Config Rule For Encrypted Volumes Disabled"
meta:
  name: "aws/config_rule_for_encrypted_volumes_is_disabled"
  id: "abdb29d4-5ca1-4e91-800b-b3569bbd788c"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Encryption"
---

## Metadata
**Name:** `aws/config_rule_for_encrypted_volumes_is_disabled`

**Id:** `abdb29d4-5ca1-4e91-800b-b3569bbd788c`

**Cloud Provider:** aws

**Severity:** High

**Category:** Encryption

## Description
This check verifies if AWS Config rules include the ENCRYPTED_VOLUMES source identifier, which monitors whether EBS volumes that are attached to EC2 instances are encrypted. Without this rule in place, organizations may unknowingly deploy unencrypted EBS volumes, potentially exposing sensitive data to unauthorized access in case of data breaches or improper access controls. 

To address this vulnerability, ensure at least one AWS Config rule uses ENCRYPTED_VOLUMES as the source identifier as shown below:

```terraform
resource "aws_config_config_rule" "encrypted_volumes" {
  name = "encrypted_vols_rule"

  source {
    owner             = "AWS"
    source_identifier = "ENCRYPTED_VOLUMES"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/config_config_rule)

## Non-Compliant Code Examples
```terraform
resource "aws_config_config_rule" "positive1" {
  name = "some_rule"

  source {
    owner             = "AWS"
    source_identifier = "IAM_PASSWORD_POLICY"
  }
}

resource "aws_config_config_rule" "positive2" {
  name = "another_rule"

  source {
    owner             = "AWS"
    source_identifier = "IAM_PASSWORD_POLICY"
  }
}
```

## Compliant Code Examples
```terraform
resource "aws_config_config_rule" "negative1" {
  name = "encrypted_vols_rule"

  source {
    owner             = "AWS"
    source_identifier = "ENCRYPTED_VOLUMES"
  }
}

resource "aws_config_config_rule" "negative2" {
  name = "another_rule"

  source {
    owner             = "AWS"
    source_identifier = "IAM_PASSWORD_POLICY"
  }
}
```