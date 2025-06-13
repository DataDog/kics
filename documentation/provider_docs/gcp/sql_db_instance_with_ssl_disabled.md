---
title: "SQL DB Instance With SSL Disabled"
meta:
  name: "gcp/sql_db_instance_with_ssl_disabled"
  id: "02474449-71aa-40a1-87ae-e14497747b00"
  cloud_provider: "gcp"
  severity: "HIGH"
  category: "Encryption"
---

## Metadata
**Name:** `gcp/sql_db_instance_with_ssl_disabled`

**Id:** `02474449-71aa-40a1-87ae-e14497747b00`

**Cloud Provider:** gcp

**Severity:** High

**Category:** Encryption

## Description
Google Cloud SQL instances without SSL enabled allow unencrypted connections, which can lead to data exposure through network eavesdropping and man-in-the-middle attacks. SSL encryption provides an essential layer of security for database connections by encrypting data in transit between the client and server. To secure your Google Cloud SQL Database, you should explicitly set 'require_ssl = true' in the ip_configuration block as shown below:

```
settings {
  ip_configuration {
    require_ssl = true
  }
}
```

Without this configuration, sensitive data such as credentials, personal information, and proprietary data could be intercepted when transmitted over the network.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance#require_ssl)

## Non-Compliant Code Examples
```terraform
resource "google_sql_database_instance" "positive1" {
  provider = google-beta

  name   = "private-instance-${random_id.db_name_suffix.hex}"
  region = "us-central1"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database_instance" "positive2" {
  provider = google-beta

  name   = "private-instance-${random_id.db_name_suffix.hex}"
  region = "us-central1"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.private_network.id
    }
  }
}

resource "google_sql_database_instance" "positive3" {
  provider = google-beta

  name   = "private-instance-${random_id.db_name_suffix.hex}"
  region = "us-central1"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.private_network.id
	    require_ssl 	  = false
    }
  }
}
```

## Compliant Code Examples
```terraform
resource "google_sql_database_instance" "negative1" {
  provider = google-beta

  name   = "private-instance-${random_id.db_name_suffix.hex}"
  region = "us-central1"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.private_network.id
	  require_ssl 	  = true
    }
  }
}
```