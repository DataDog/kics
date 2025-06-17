---
title: "AmazonMQ Broker Encryption Disabled"
meta:
  name: "terraform/amazon_mq_broker_encryption_disabled"
  id: "3db3f534-e3a3-487f-88c7-0a9fbf64b702"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/amazon_mq_broker_encryption_disabled`
**Id:** `3db3f534-e3a3-487f-88c7-0a9fbf64b702`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
Amazon MQ brokers should have encryption options defined to ensure messages are encrypted at rest. Without proper encryption, sensitive data in message queues could be exposed if storage is compromised. To secure your broker, add an encryption_options block to your aws_mq_broker resource, either with a custom KMS key (recommended) or with the default AWS owned keys. Example of secure configuration: `encryption_options { kms_key_id = "your-kms-key-arn" use_aws_owned_key = false }` or simply `encryption_options {}` to use AWS-owned keys.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/mq_broker)

## Non-Compliant Code Examples
```aws
resource "aws_mq_broker" "positive1" {
  broker_name = "example"

  configuration {
    id       = aws_mq_configuration.test.id
    revision = aws_mq_configuration.test.latest_revision
  }

  engine_type        = "ActiveMQ"
  engine_version     = "5.15.9"
  host_instance_type = "mq.t2.micro"
  security_groups    = [aws_security_group.test.id]

  user {
    username = "ExampleUser"
    password = "MindTheGap"
  }
}

```

## Compliant Code Examples
```aws
resource "aws_mq_broker" "negative1" {
  broker_name = "example"

  configuration {
    id       = aws_mq_configuration.test.id
    revision = aws_mq_configuration.test.latest_revision
  }

  engine_type        = "ActiveMQ"
  engine_version     = "5.15.9"
  host_instance_type = "mq.t2.micro"
  security_groups    = [aws_security_group.test.id]

  user {
    username = "ExampleUser"
    password = "MindTheGap"
  }

  encryption_options {
    kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
    use_aws_owned_key = false
  }
}

resource "aws_mq_broker" "negative2" {
  broker_name = "example"

  configuration {
    id       = aws_mq_configuration.test.id
    revision = aws_mq_configuration.test.latest_revision
  }

  engine_type        = "ActiveMQ"
  engine_version     = "5.15.9"
  host_instance_type = "mq.t2.micro"
  security_groups    = [aws_security_group.test.id]

  user {
    username = "ExampleUser"
    password = "MindTheGap"
  }

  encryption_options {
  }
}

```