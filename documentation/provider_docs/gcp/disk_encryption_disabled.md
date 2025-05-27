---
title: "Disk Encryption Disabled"
meta:
  name: "gcp/disk_encryption_disabled"
  id: "b1d51728-7270-4991-ac2f-fc26e2695b38"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `gcp/disk_encryption_disabled`

**Id:** `b1d51728-7270-4991-ac2f-fc26e2695b38`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Encryption

## Description
VM disks for critical VMs must be encrypted with Customer Supplied Encryption Keys (CSEK) or with Customer-managed encryption keys (CMEK), which means the attribute 'disk_encryption_key' must be defined and its sub attributes 'raw_key' or 'kms_key_self_link' must also be defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_disk)

## Non-Compliant Code Examples
```terraform
resource "google_compute_disk" "positive3" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096

  disk_encryption_key {
      raw_key = ""
      sha256 = "A"
  }
}

```

```terraform
resource "google_compute_disk" "positive4" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096

  disk_encryption_key {
    kms_key_self_link = ""
    sha256 = "A"
  }
}

```

```terraform
resource "google_compute_disk" "positive1" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

resource "google_compute_disk" "positive2" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096

  disk_encryption_key {
    sha256 = "A"
  }
}

```

## Compliant Code Examples
```terraform
resource "google_compute_disk" "negative1" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096

  disk_encryption_key {
      kms_key_self_link = "disk-crypto-key"
      sha256 = "A"
  }
}

```

```terraform
resource "google_compute_disk" "negative1" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096

  disk_encryption_key {
      raw_key = "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0="
      sha256 = "A"
  }
}

```