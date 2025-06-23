---
title: "MSK Cluster Encryption Disabled"
meta:
  name: "aws/msk_cluster_encryption_disabled"
  id: "6db52fa6-d4da-4608-908a-89f0c59e743e"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `aws/msk_cluster_encryption_disabled`
**Id:** `6db52fa6-d4da-4608-908a-89f0c59e743e`
**Cloud Provider:** aws
**Severity:** High
**Category:** Encryption
## Description
Amazon MSK Clusters store sensitive data that should be protected both at rest and in transit through encryption. When encryption is disabled, data may be exposed to unauthorized users during storage or transmission, creating significant security vulnerabilities. To properly secure an MSK cluster, you should specify both encryption_at_rest_kms_key_arn (for data at rest) and encryption_in_transit with client_broker set to 'TLS' and in_cluster set to true (for data in transit). Here's a secure example:

```terraform
resource "aws_msk_cluster" "example" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3
  
  encryption_info {
    encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
    encryption_in_transit {
      client_broker = "TLS"
      in_cluster = true
    }
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/msk_cluster#encryption_info)


## Compliant Code Examples
```terraform
resource "aws_msk_cluster" "negative1" {  
  encryption_info {
    encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
  }
}

resource "aws_msk_cluster" "negative2" {  
  encryption_info {
    encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
    encryption_in_transit {
      client_broker = "TLS"
      in_cluster = true
    }
  }
}

resource "aws_msk_cluster" "negative3" {  
  encryption_info {
    encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
    encryption_in_transit {
      client_broker = "TLS"
    }
  }
}

resource "aws_msk_cluster" "negative4" {  
  encryption_info {
    encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
    encryption_in_transit {
      in_cluster = true
    }
  }
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_msk_cluster" "positive1" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3
}

resource "aws_msk_cluster" "positive2" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3
  
  encryption_info {
    encryption_in_transit {
      client_broker = "PLAINTEXT"
    }
  }
}

resource "aws_msk_cluster" "positive3" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3
  
  encryption_info {
    encryption_in_transit {
      in_cluster = false
    }
  }
}

resource "aws_msk_cluster" "positive4" {
  cluster_name           = "example"
  kafka_version          = "2.4.1"
  number_of_broker_nodes = 3
  
  encryption_info {
    encryption_in_transit {
      client_broker = "PLAINTEXT"
      in_cluster = false
    }
  }
}
```