---
title: "MQ Broker Logging Disabled"
group-id: "rules/terraform/aws"
meta:
  name: "aws/mq_broker_logging_disabled"
  id: "31245f98-a6a9-4182-9fc1-45482b9d030a"
  display_name: "MQ Broker Logging Disabled"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Id:** `31245f98-a6a9-4182-9fc1-45482b9d030a`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/mq_broker)

### Description

 This check ensures that AWS MQ Brokers have both audit and general logging enabled in their configuration. Logging is essential for capturing critical security events and operational information, which aids in monitoring, troubleshooting, and forensic analysis. If logging is not enabled, malicious activity or configuration issues may go undetected, significantly increasing the risk of security breaches and data loss. Unaddressed, the lack of logging impedes compliance efforts and can hinder incident response due to an absence of necessary audit trails.


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