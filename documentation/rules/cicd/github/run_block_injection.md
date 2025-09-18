---
title: "Run block injection"
group_id: "rules/cicd/github"
meta:
  name: "github/run_block_injection"
  id: "20f14e1a-a899-4e79-9f09-b6a84cd4649b"
  display_name: "Run block injection"
  cloud_provider: "github"
  framework: "CICD"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `20f14e1a-a899-4e79-9f09-b6a84cd4649b`

**Cloud Provider:** github

**Framework:** CICD

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://securitylab.github.com/research/github-actions-untrusted-input/)

### Description

 GitHub Actions workflows can be triggered by a variety of events. Each trigger includes a GitHub context with details about the event, such as the user who triggered it, the branch name, and other relevant information. Some of this data, like the base repository name, changeset hash, or pull request number, is typically not controlled by the user and is unlikely to be used for injection.
