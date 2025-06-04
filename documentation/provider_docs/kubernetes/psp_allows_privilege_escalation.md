---
title: "PSP Allows Privilege Escalation"
meta:
  name: "kubernetes/psp_allows_privilege_escalation"
  id: "2bff9906-4e9b-4f71-9346-8ebedfdf43ef"
  cloud_provider: "kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `kubernetes/psp_allows_privilege_escalation`

**Id:** `2bff9906-4e9b-4f71-9346-8ebedfdf43ef`

**Cloud Provider:** kubernetes

**Severity:** High

**Category:** Insecure Configurations

## Description
PodSecurityPolicy should not allow privilege escalation

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod_security_policy#allow_privilege_escalation)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = true

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

resource "kubernetes_pod_security_policy" "example2" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false

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
```terraform
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