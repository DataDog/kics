---
title: "MQ Broker Logging Disabled"
meta:
  name: "aws/mq_broker_logging_disabled"
  id: "31245f98-a6a9-4182-9fc1-45482b9d030a"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/mq_broker_logging_disabled`

**Id:** `31245f98-a6a9-4182-9fc1-45482b9d030a`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Check if MQ Brokers don't have logging enabled in any of the two options possible (audit and general).

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/mq_broker)

## Non-Compliant Code Examples
```terraform
resource "aws_mq_broker" "positive1" {
  broker_name = "no-logging"
}

resource "aws_mq_broker" "positive2" {
  broker_name = "partial-logging"

  logs {
      general = true
  }
}

resource "aws_mq_broker" "positive3" {
  broker_name = "disabled-logging"

  logs {
      general = false
      audit = true
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_mq_broker" "negative1" {
  broker_name = "example"

  configuration {
    id       = aws_mq_configuration.test.id
    revision = aws_mq_configuration.test.latest_revision
  }

  engine_type        = "ActiveMQ"
  engine_version     = "5.15.0"
  host_instance_type = "mq.t2.micro"
  security_groups    = [aws_security_group.test.id]

  user {
    username = "ExampleUser"
    password = "MindTheGap"
  }

  logs {
      general = true
      audit = true
  }
}
```