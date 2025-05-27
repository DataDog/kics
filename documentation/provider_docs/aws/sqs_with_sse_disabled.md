---
title: "SQS With SSE Disabled"
meta:
  name: "aws/sqs_with_sse_disabled"
  id: "6e8849c1-3aa7-40e3-9063-b85ee300f29f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/sqs_with_sse_disabled`

**Id:** `6e8849c1-3aa7-40e3-9063-b85ee300f29f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
Amazon Simple Queue Service (SQS) queue should protect the contents of their messages using Server-Side Encryption (SSE)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue)

## Non-Compliant Code Examples
```terraform
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }

  kms_master_key_id = null

}

```

```terraform
resource "aws_sqs_queue" "positive2" {
  name                              = "terraform-example-queue"
  kms_master_key_id                 = ""
  kms_data_key_reuse_period_seconds = 300
}


```

```terraform
resource "aws_sqs_queue" "positive3" {
  name                              = "terraform-example-queue"
  kms_master_key_id                 = null
  kms_data_key_reuse_period_seconds = 300
}

```

```terraform
resource "aws_sqs_queue" "positive7" {
  name                    = "terraform-example-queue"
  sqs_managed_sse_enabled = false
}

```

```terraform
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

```terraform
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }

  kms_master_key_id = ""
}

```

```terraform
resource "aws_sqs_queue" "positive1" {
  name                              = "terraform-example-queue"
  kms_data_key_reuse_period_seconds = 300
}

```

## Compliant Code Examples
```terraform
resource "aws_sqs_queue" "negative3" {
  name                    = "terraform-example-queue"
  sqs_managed_sse_enabled = true
}

```

```terraform
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

```terraform
resource "aws_sqs_queue" "negative1" {
  name                              = "terraform-example-queue"
  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
}

```