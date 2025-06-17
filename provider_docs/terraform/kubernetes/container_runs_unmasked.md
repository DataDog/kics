---
title: "Container Runs Unmasked"
meta:
  name: "terraform/container_runs_unmasked"
  id: "0ad60203-c050-4115-83b6-b94bde92541d"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/container_runs_unmasked`
**Id:** `0ad60203-c050-4115-83b6-b94bde92541d`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Insecure Configurations
## Description
Check if a container has full access (unmasked) to the host's /proc command, which would allow to retrieve sensitive information and possibly change the kernel parameters in runtime.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod_security_policy#allowed_proc_mount_types)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = false
    allowed_proc_mount_types   = ["Unmasked"]

    volumes = [
      "configMap",
      "emptyDir",
      "projected",
      "secret",
      "downwardAPI",
      "persistentVolumeClaim",
    ]

    run_as_user {
      rule = "MustRunAsNonRoot"
    }

    se_linux {
      rule = "RunAsAny"
    }

    supplemental_groups {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    fs_group {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    read_only_root_filesystem = true
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = false
    allowed_proc_mount_types   = ["Default"]

    volumes = [
      "configMap",
      "emptyDir",
      "projected",
      "secret",
      "downwardAPI",
      "persistentVolumeClaim",
    ]

    run_as_user {
      rule = "MustRunAsNonRoot"
    }

    se_linux {
      rule = "RunAsAny"
    }

    supplemental_groups {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    fs_group {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    read_only_root_filesystem = true
  }
}

```