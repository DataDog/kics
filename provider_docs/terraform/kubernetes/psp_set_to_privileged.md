---
title: "PSP Set To Privileged"
meta:
  name: "terraform/psp_set_to_privileged"
  id: "a6a4d4fc-4e8f-47d1-969f-e9d4a084f3b9"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/psp_set_to_privileged`
**Id:** `a6a4d4fc-4e8f-47d1-969f-e9d4a084f3b9`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Insecure Configurations
## Description
Do not allow pod to request execution as privileged.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#privileged)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = true
    allow_privilege_escalation = false

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
resource "kubernetes_pod_security_policy" "example2" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = false

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