---
title: "MQ Broker Is Publicly Accessible"
group-id: "rules/terraform/aws"
meta:
  name: "aws/mq_broker_is_publicly_accessible"
  id: "4eb5f791-c861-4afd-9f94-f2a6a3fe49cb"
  display_name: "MQ Broker Is Publicly Accessible"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `4eb5f791-c861-4afd-9f94-f2a6a3fe49cb`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/mq_broker)

### Description

 This check verifies if AWS MQ Brokers have the 'publicly_accessible' attribute set to true, which makes them accessible from the internet. When an MQ Broker is publicly accessible, it increases the attack surface and exposure to potential unauthorized access, potentially leading to data breaches or service disruption. To secure your MQ Broker, omit the 'publicly_accessible' attribute or explicitly set it to false, as shown below:

```
resource "aws_mq_broker" "secure_broker" {
  broker_name = "example"
  engine_type = "ActiveMQ"
  // Other configurations
  
  // Either omit this line completely or set to false
  publicly_accessible = false
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
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_mq_broker" "positive1" {
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

  publicly_accessible = true
}
```