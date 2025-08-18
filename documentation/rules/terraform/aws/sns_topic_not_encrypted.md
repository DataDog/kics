---
title: "SNS topic not encrypted"
group_id: "rules/terraform/aws"
meta:
  name: "aws/sns_topic_not_encrypted"
  id: "28545147-2fc6-42d5-a1f9-cf226658e591"
  display_name: "SNS topic not encrypted"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `28545147-2fc6-42d5-a1f9-cf226658e591`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic#kms_master_key_id)

### Description

 Amazon SNS topics should be encrypted at rest using AWS KMS to protect sensitive message content. Without encryption, data stored in SNS topics is vulnerable to unauthorized access if the service or underlying storage is compromised. Use the `kms_master_key_id` attribute to specify a KMS key for encryption, as shown in this example: `kms_master_key_id = "alias/MyAlias"`. Leaving this attribute empty or unspecified (as in `kms_master_key_id = ""`) results in unencrypted SNS topics, exposing sensitive data to potential security breaches.


## Compliant Code Examples
```terraform
provider "aws2" {
  region = "us-east-1"
}

resource "aws_sns_topic" "test2" {
  name              = "sns_ecnrypted"
  kms_master_key_id = "alias/MyAlias"
}

```

```terraform
module "sns_topic" {
  source = "terraform-aws-modules/sns/aws"
  version = "~> 2.0"

  name  = "my-topic"
  display_name = "My awesome sns topic"

  # We recommend using FIFO topics whenever possible
  fifo_topic = true

  # Use some valid AWS account ID here
  cross_account_iam_role_arns = ["arn:aws:iam::817568760441:role/sns-publish-my-topic"]

  subscriptions = [
    {
      protocol = "lambda"
      endpoint = "arn:aws:lambda:eu-west-1:835367859851:function:example"
    }
  ]

  tags = {
    Name = "my-sns-topic"
  }
  kms_master_key_id = "alias/aws/sns/"
}
```
## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

resource "aws_sns_topic" "test" {
  name = "sns_not_ecnrypted"
}

```

```terraform
resource "aws_sns_topic" "user_updates" {
  name              = "user-updates-topic"
  kms_master_key_id = ""
}

```

```terraform
module "sns_topic" {
  source = "terraform-aws-modules/sns/aws"
  version = "~> 2.0"

  name  = "my-topic"
  display_name = "My awesome sns topic"

  # We recommend using FIFO topics whenever possible
  fifo_topic = true

  # Use some valid AWS account ID here
  cross_account_iam_role_arns = ["arn:aws:iam::817568760441:role/sns-publish-my-topic"]

  subscriptions = [
    {
      protocol = "lambda"
      endpoint = "arn:aws:lambda:eu-west-1:835367859851:function:example"
    }
  ]

  tags = {
    Name = "my-sns-topic"
  }
  kms_master_key_id = ""
}
```