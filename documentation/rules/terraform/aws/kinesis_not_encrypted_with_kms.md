---
title: "Kinesis Not Encrypted With KMS"
group-id: "rules/terraform/aws"
meta:
  name: "aws/kinesis_not_encrypted_with_kms"
  id: "862fe4bf-3eec-4767-a517-40f378886b88"
  display_name: "Kinesis Not Encrypted With KMS"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `862fe4bf-3eec-4767-a517-40f378886b88`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kinesis_stream)

### Description

 AWS Kinesis Streams contain potentially sensitive data that should be encrypted at rest using AWS KMS to enhance security. When Kinesis streams are not encrypted with KMS, data stored in them is vulnerable to unauthorized access if the underlying storage is compromised. To properly secure Kinesis streams, both the encryption_type must be set to 'KMS' and a valid kms_key_id must be specified, as shown in the following example:

```
resource "aws_kinesis_stream" "secure_example" {
  name             = "terraform-kinesis-test"
  // ... other configurations ...
  
  encryption_type = "KMS"
  kms_key_id = "alias/aws/kinesis"
}
```


## Compliant Code Examples
```terraform
resource "aws_kinesis_stream" "negative1" {
  name             = "terraform-kinesis-test"
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  tags = {
    Environment = "test"
  }


  encryption_type = "KMS"

  kms_key_id = "alias/aws/kinesis"
}


```
## Non-Compliant Code Examples
```terraform
resource "aws_kinesis_stream" "positive1" {
  name             = "terraform-kinesis-test"
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  tags = {
    Environment = "test"
  }
}




resource "aws_kinesis_stream" "positive2" {
  name             = "terraform-kinesis-test"
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  tags = {
    Environment = "test"
  }


  encryption_type = "NONE"
}





resource "aws_kinesis_stream" "positive3" {
  name             = "terraform-kinesis-test"
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  tags = {
    Environment = "test"
  }


  encryption_type = "KMS"
}



```