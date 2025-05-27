#!/usr/bin/env -S uv run --script
# /// script
# dependencies = [
#   "ruamel.yaml",
# ]
# ///

import os, sys, json, glob, argparse, shutil
from pathlib import Path

def parse_args():
    parser = argparse.ArgumentParser(description="Generate documentation from metadata.json and test files")
    parser.add_argument("input_dir", type=Path, help="Base directory containing provider subfolders")
    parser.add_argument("--providers-json", type=str, required=True, help="JSON file listing providers to document")
    parser.add_argument("--output-dir", type=str, default="provider_docs", help="Directory for generated markdown files")
    return parser.parse_args()

def read_file_contents(filepath):
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            return f.read()
    except Exception as e:
        print(f"Warning: Failed to read {filepath}: {e}")
        return ""

def get_code_snippets(test_dir):
    compliant, non_compliant = [], []
    for tf_file in glob.glob(str(test_dir / "negative*.tf")):
        code = read_file_contents(tf_file)
        if code:
            compliant.append(f"```terraform\n{code}\n```")
    for tf_file in glob.glob(str(test_dir / "positive*.tf")):
        code = read_file_contents(tf_file)
        if code:
            non_compliant.append(f"```terraform\n{code}\n```")
    return compliant, non_compliant

def build_markdown(rule_path, metadata, cloud_provider):
    rule_name = rule_path.name
    title = metadata.get("queryName", "Untitled Rule")
    rule_id = metadata.get("id", "unknown-id")
    severity = metadata.get("severity", "INFO").upper()
    category = metadata.get("category", "Unknown")

    # Prefer override.detailedDescriptionText.descriptionText if available
    description = (
        metadata.get("override", {})
                .get("detailedDescriptionText", {})
                .get("descriptionText")
        or metadata.get("descriptionText", "No description provided.")
    )

    description_url = metadata.get("descriptionUrl")

    compliant, non_compliant = get_code_snippets(rule_path / "test")

    meta_name = f"{cloud_provider}/{rule_name}"

    markdown = f"""---
title: {json.dumps(title)}
meta:
  name: "{meta_name}"
  id: "{rule_id}"
  cloud_provider: "{cloud_provider}"
  severity: "{severity}"
  category: "{category}"
---

## Metadata
**Name:** `{meta_name}`

**Id:** `{rule_id}`

**Cloud Provider:** {cloud_provider}

**Severity:** {severity.capitalize()}

**Category:** {category}

## Description
{description}
"""
    if description_url:
        markdown += f"\n#### Learn More\n\n - [Provider Reference]({description_url})\n"

    if non_compliant:
        markdown += "\n## Non-Compliant Code Examples\n" + "\n\n".join(non_compliant)
    if compliant:
        markdown += "\n\n## Compliant Code Examples\n" + "\n\n".join(compliant)

    return markdown

def main():
    args = parse_args()
    input_dir = args.input_dir
    output_dir = Path(args.output_dir)

    # Remove and recreate the output directory
    if output_dir.exists():
        shutil.rmtree(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    # Load list of providers
    try:
        with open(args.providers_json, 'r') as f:
            provider_list = json.load(f)
    except Exception as e:
        print(f"Error loading providers JSON: {e}")
        sys.exit(1)

    for provider in provider_list:
        provider_path = input_dir / provider
        if not provider_path.is_dir():
            print(f"Warning: Missing provider path: {provider_path}")
            continue

        output_provider_path = output_dir / provider
        output_provider_path.mkdir(parents=True, exist_ok=True)

        for rule_dir in provider_path.iterdir():
            if not rule_dir.is_dir():
                continue

            metadata_file = rule_dir / "metadata.json"
            if not metadata_file.exists():
                print(f"Skipping {rule_dir.name} — missing metadata.json")
                continue

            try:
                with open(metadata_file, 'r', encoding='utf-8') as f:
                    metadata = json.load(f)
            except Exception as e:
                print(f"Failed to parse metadata for {rule_dir}: {e}")
                continue

            md_content = build_markdown(rule_dir, metadata, provider)
            output_file = output_provider_path / f"{rule_dir.name}.md"
            with open(output_file, "w", encoding='utf-8') as f:
                f.write(md_content)
            print(f"Generated: {output_file}")

if __name__ == "__main__":
    main()
