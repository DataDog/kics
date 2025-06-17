---
title: "SQS With SSE Disabled"
meta:
  name: "terraform/sqs_with_sse_disabled"
  id: "6e8849c1-3aa7-40e3-9063-b85ee300f29f"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/sqs_with_sse_disabled`
**Id:** `6e8849c1-3aa7-40e3-9063-b85ee300f29f`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Encryption
## Description
Amazon Simple Queue Service (SQS) queues should use Server-Side Encryption (SSE) to protect the contents of their messages while at rest. Without specifying the `kms_master_key_id` attribute in the Terraform configuration (e.g., `resource "aws_sqs_queue" "positive1"`), messages sent to the queue are stored without encryption, exposing sensitive data to unauthorized access if AWS infrastructure is compromised. Enabling SSE by setting `kms_master_key_id` (as shown below) ensures that all messages are encrypted using a customer-managed key, significantly reducing the risk of data leakage.

```
resource "aws_sqs_queue" "example" {
  name                  = "terraform-example-queue"
  kms_master_key_id     = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue)

## Non-Compliant Code Examples
```aws
resource "aws_sqs_queue" "positive7" {
  name                    = "terraform-example-queue"
  sqs_managed_sse_enabled = false
}

```

```aws
resource "aws_sqs_queue" "positive1" {
  name                              = "terraform-example-queue"
  kms_data_key_reuse_period_seconds = 300
}

```

```aws
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }
}

```

## Compliant Code Examples
```aws
resource "aws_sqs_queue" "negative1" {
  name                              = "terraform-example-queue"
  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
}

```

```aws
resource "aws_sqs_queue" "negative3" {
  name                    = "terraform-example-queue"
  sqs_managed_sse_enabled = true
}

```

```aws
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }

  kms_master_key_id = "alias/aws/sqs"

}

```