---
title: "Run block injection"
group_id: "rules/cicd/github"
meta:
  name: "github/run_block_injection"
  id: "20f14e1a-a899-4e79-9f09-b6a84cd4649b"
  display_name: "Run block injection"
  cloud_provider: "github"
  framework: "CICD"
  platform: "CICD"
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


## Compliant Code Examples
```yaml
name: check-go-coverage

on:
  pull_request_target:
    branches: [master]

jobs:
  coverage:
    name: Check Go coverage
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Source
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Set up Go 1.22.x
        uses: actions/setup-go@v5
        with:
          go-version: 1.22.x
      - name: Run test metrics script
        id: testcov
        run: |
          make test-coverage-report | tee test-results
          echo "coverage=$(cat test-results | grep "Total coverage: " test-results | cut -d ":" -f 2 | bc)" >> $GITHUB_ENV
      - name: Checks if Go coverage is at least 80%
        if: env.coverage < 80
        run: |
          echo "Go coverage is lower than 80%: ${{ env.coverage }}%"
          exit 1

```

```yaml
name: Author Workflow

on:
  author:
    types:
      - created

jobs:
  process_author:
    runs-on: ubuntu-latest
    steps:
      - name: Greet the New Author
        run: |
          echo "Hello, a new author has been created!"

```

```yaml
name: Workflow Run Workflow

on:
  workflow_run:
    workflows:
      - "Your Workflow Name" # Replace with the name of your specific workflow

jobs:
  process_workflow_run:
    runs-on: ubuntu-latest
    steps:
      - name: Greet the New Workflow Run
        run: |
          echo "Hello, a new workflow run has started for 'Your Workflow Name'!"

```
## Non-Compliant Code Examples
```yaml
name: Web Page To Markdown
on:
  issues:
    types: [opened]
jobs:
  WebPageToMarkdown:
    runs-on: ubuntu-latest
    steps:
      - name: Does the issue need to be converted to markdown
        run: |
          if [ "${{ github.event.issue.body }}" ]; then
            if [[ "${{ github.event.issue.title }}" =~ ^\[Auto\]* ]]; then
              :
            else
              echo "This issue does not need to generate a markdown file." 1>&2
              exit 1;
            fi;
          else
            echo "The description of the issue is empty." 1>&2
            exit 1;
          fi;
        shell: bash
      - name: Checkout
        uses: actions/checkout@v4
        with:
          ref: ${{ github.head_ref }}
      - name: Crawl pages and generate Markdown files
        uses: freeCodeCamp-China/article-webpage-to-markdown-action@v0.1.8
        with:
          newsLink: '${{ github.event.issue.body }}'
          markDownFilePath: './chinese/articles/'
          githubToken: ${{ github.token }}
      - name: Git Auto Commit
        uses: stefanzweifel/git-auto-commit-action@v4.9.2
        with:
          commit_message: '${{ github.event.issue.title }}'
          file_pattern: chinese/articles/*.md
          commit_user_name: PageToMarkdown Bot
          commit_user_email: PageToMarkdown-bot@freeCodeCamp.org

```

```yaml
name: Author Workflow

on:
  author:
    types:
      - created

jobs:
  process_author:
    runs-on: ubuntu-latest
    steps:
      - name: Echo Author's Username
        run: |
          echo "Author's Name: ${{ github.event.authors.name }}"

```

```yaml
name: Issue Comment Workflow

on:
  issue_comment:
    types:
      - created

jobs:
  process_issue_comment:
    runs-on: ubuntu-latest
    steps:
      - name: Echo Issue Comment Body
        run: |
          echo "Issue Comment Body: ${{ github.event.comment.body }}"

```