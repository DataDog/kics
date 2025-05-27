---
title: "SQL DB Instance Backup Disabled"
meta:
  name: "gcp/sql_db_instance_backup_disabled"
  id: "cf3c7631-cd1e-42f3-8801-a561214a6e79"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Backup"
---

## Metadata
**Name:** `gcp/sql_db_instance_backup_disabled`

**Id:** `cf3c7631-cd1e-42f3-8801-a561214a6e79`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Backup

## Description
Checks if backup configuration is enabled for all Cloud SQL Database instances

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance)

## Non-Compliant Code Examples
```terraform
resource "google_sql_database_instance" "positive1" {
    name             = "master-instance"
    database_version = "POSTGRES_11"
    region           = "us-central1"

    settings {
        tier = "db-f1-micro"
    }
}

resource "google_sql_database_instance" "positive2" {
    name             = "master-instance"
    database_version = "POSTGRES_11"
    region           = "us-central1"

    settings {
        tier = "db-f1-micro"
        backup_configuration{
            binary_log_enabled = true
        }
    }
}

resource "google_sql_database_instance" "positive3" {
    name             = "master-instance"
    database_version = "POSTGRES_11"
    region           = "us-central1"

    settings {
        backup_configuration{
            enabled = false
        }
    }
}


```

## Compliant Code Examples
```terraform
resource "google_sql_database_instance" "negative1" {
    name             = "master-instance"
    database_version = "POSTGRES_11"
    region           = "us-central1"
 
    settings {
        backup_configuration{
            enabled = true
        }
    }
}

```