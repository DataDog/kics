---
title: "SQL DB Instance Backup Disabled"
meta:
  name: "gcp/sql_db_instance_backup_disabled"
  id: "cf3c7631-cd1e-42f3-8801-a561214a6e79"
  display_name: "SQL DB Instance Backup Disabled"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Backup"
---
## Metadata
**Name:** `gcp/sql_db_instance_backup_disabled`
**Query Name** `SQL DB Instance Backup Disabled`
**Id:** `cf3c7631-cd1e-42f3-8801-a561214a6e79`
**Cloud Provider:** gcp
**Platform** Terraform
**Severity:** Medium
**Category:** Backup
## Description
This check verifies whether automated backups are enabled for all Google Cloud SQL database instances by ensuring the `backup_configuration` block has the `enabled = true` attribute. Without automated backups, as shown below:

```
settings {
    backup_configuration {
        enabled = false
    }
}
```

or if the `backup_configuration` block is missing, databases are at risk of unrecoverable data loss in case of accidental deletion, corruption, or other failures. Enabling backups like this:

```
settings {
    backup_configuration {
        enabled = true
    }
}
```

ensures that point-in-time recovery is possible and critical business data can be restored when needed.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance)


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