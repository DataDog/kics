---
title: "Sagemaker Endpoint Configuration Encryption Disabled"
meta:
  name: "aws/sagemaker_endpoint_configuration_encryption_disabled"
  id: "58b35504-0287-4154-bf69-02c0573deab8"
  display_name: "Sagemaker Endpoint Configuration Encryption Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Name:** `aws/sagemaker_endpoint_configuration_encryption_disabled`

**Query Name** `Sagemaker Endpoint Configuration Encryption Disabled`

**Id:** `58b35504-0287-4154-bf69-02c0573deab8`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** High

**Category:** Encryption

## Description
Amazon SageMaker endpoint configurations should have encryption enabled using a KMS key to protect sensitive data at rest. Without proper encryption, data stored within SageMaker endpoints may be vulnerable to unauthorized access if the underlying storage is compromised. This represents a significant security risk as machine learning endpoints often process and store sensitive information.

To address this vulnerability, specify the `kms_key_arn` attribute in your SageMaker endpoint configuration. For example, the secure implementation uses:
```
resource "aws_sagemaker_endpoint_configuration" "example" {
  // other configuration
  kms_key_arn = "aws_kms_key.example.arn"
}
```
while the insecure version omits this critical encryption setting.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sagemaker_endpoint_configuration#kms_key_arn)


## Compliant Code Examples
```terraform
resource "aws_sagemaker_endpoint_configuration" "negative" {
  name = "my-endpoint-config"

  production_variants {
    variant_name           = "variant-1"
    model_name             = aws_sagemaker_model.m.name
    initial_instance_count = 1
    instance_type          = "ml.t2.medium"
  }

  tags = {
    Name = "foo"
  }

  kms_key_arn = "aws_kms_key.example.arn"
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_sagemaker_endpoint_configuration" "positive" {
  name = "my-endpoint-config"

  production_variants {
    variant_name           = "variant-1"
    model_name             = aws_sagemaker_model.m.name
    initial_instance_count = 1
    instance_type          = "ml.t2.medium"
  }

  tags = {
    Name = "foo"
  }
}

```