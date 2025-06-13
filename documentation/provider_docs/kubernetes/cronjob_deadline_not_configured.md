---
title: "CronJob Deadline Not Configured"
meta:
  name: "kubernetes/cronjob_deadline_not_configured"
  id: "58876b44-a690-4e9f-9214-7735fa0dd15d"
  cloud_provider: "kubernetes"
  severity: "LOW"
  category: "Resource Management"
---

## Metadata
**Name:** `kubernetes/cronjob_deadline_not_configured`

**Id:** `58876b44-a690-4e9f-9214-7735fa0dd15d`

**Cloud Provider:** kubernetes

**Severity:** Low

**Category:** Resource Management

## Description
Cronjobs must have a configured deadline, which means the attribute 'starting_deadline_seconds' must be defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cron_job#starting_deadline_seconds)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_cron_job" "demo" {
  metadata {
    name = "demo"
  }
  spec {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "1 0 * * *"
    successful_jobs_history_limit = 10
    job_template {
      metadata {}
      spec {
        backoff_limit              = 2
        ttl_seconds_after_finished = 10
        template {
          metadata {}
          spec {
            container {
              name    = "hello"
              image   = "busybox"
              command = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]
            }
          }
        }
      }
    }
  }
}

```

## Compliant Code Examples
```terraform
resource "kubernetes_cron_job" "demo2" {
  metadata {
    name = "demo"
  }
  spec {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "1 0 * * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 10
    job_template {
      metadata {}
      spec {
        backoff_limit              = 2
        ttl_seconds_after_finished = 10
        template {
          metadata {}
          spec {
            container {
              name    = "hello"
              image   = "busybox"
              command = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]
            }
          }
        }
      }
    }
  }
}

```