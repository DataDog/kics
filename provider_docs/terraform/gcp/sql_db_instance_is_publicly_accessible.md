---
title: "SQL DB Instance Publicly Accessible"
meta:
  name: "terraform/sql_db_instance_is_publicly_accessible"
  id: "b187edca-b81e-4fdc-aff4-aab57db45edb"
  cloud_provider: "terraform"
  severity: "CRITICAL"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/sql_db_instance_is_publicly_accessible`
**Id:** `b187edca-b81e-4fdc-aff4-aab57db45edb`
**Cloud Provider:** terraform
**Severity:** Critical
**Category:** Insecure Configurations
## Description
A Google Cloud SQL instance becomes publicly accessible when it has public IP addressing enabled without proper network restrictions, creating a potential attack vector for unauthorized access. This can occur through ipv4_enabled being set to true (default) or by configuring authorized_networks with overly permissive CIDR ranges like '0.0.0.0/0' which allows connections from any IP address. To secure Cloud SQL instances, either disable public IP by setting ipv4_enabled to false and specifying a private_network (e.g., `ipv4_enabled = false` and `private_network = "your-network-id"`) or restrict authorized_networks to specific trusted IP ranges (e.g., `authorized_networks { name = "trusted-network", value = "10.0.0.0/24" }`) rather than using '0.0.0.0/0'.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance)

## Non-Compliant Code Examples
```gcp
resource "google_sql_database_instance" "positive1" {
  name             = "master-instance"
  database_version = "POSTGRES_11"
  region           = "us-central1"

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"
  }
}

resource "google_sql_database_instance" "positive2" {
  name             = "postgres-instance-2"
  database_version = "POSTGRES_11"

  settings {
    tier = "db-f1-micro"

    ip_configuration {

      authorized_networks {
        name  = "pub-network"
        value = "0.0.0.0/0"
      }
    }
  }
}

resource "google_sql_database_instance" "positive3" {
  name             = "master-instance"
  database_version = "POSTGRES_11"
  region           = "us-central1"

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"

    ip_configuration {
        ipv4_enabled = true
    }
  }
}

resource "google_sql_database_instance" "positive4" {
  name             = "master-instance"
  database_version = "POSTGRES_11"
  region           = "us-central1"

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"

    ip_configuration {}
  }
}

```

## Compliant Code Examples
```gcp
resource "google_sql_database_instance" "negative1" {

  name   = "private-instance-1"
  database_version = "POSTGRES_11"
  settings {
    ip_configuration {
      ipv4_enabled = false
      private_network = "some_private_network"
    }
  }
}

resource "google_sql_database_instance" "negative2" {
  name             = "postgres-instance-2"
  database_version = "POSTGRES_11"

  settings {
    tier = "db-f1-micro"

    ip_configuration {

      authorized_networks {

        content {
          name  = "some_trusted_network"
          value = "some_trusted_network_address"
        }
      }

      authorized_networks {

        content {
          name  = "another_trusted_network"
          value = "another_trusted_network_address"
        }
      }
    }
  }
}

```