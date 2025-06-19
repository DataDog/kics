#!/usr/bin/env -S uv run --script
# /// script
# dependencies = [
#   "ruamel.yaml",
# ]
# ///

import os
import sys
import json
import glob
import argparse
import shutil
import yaml
from itertools import islice
from pathlib import Path

def parse_args():
    parser = argparse.ArgumentParser(description="Generate documentation from metadata.json and test files")
    parser.add_argument("input_dir", type=Path, help="Base directory containing all the rules")
    parser.add_argument("--resources-json", type=str, required=True, help="JSON file listing resources and providers to document")
    parser.add_argument("--output-dir", type=str, default="provider_docs", help="Directory for generated markdown files")
    parser.add_argument("--list-json", type=str, default="list.json", help="Path to write list.json")
    parser.add_argument("--frontmatter-yaml", type=str, default="frontmatter.yaml", help="Path to write frontmatter.yaml")
    parser.add_argument("--max-examples", type=int, default=3, help="Max number of compliant and non-compliant examples to add to each markdown")
    return parser.parse_args()

def read_file_contents(filepath):
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            return f.read()
    except Exception as e:
        print(f"Warning: Failed to read {filepath}: {e}")
        return ""

def get_code_snippets(test_dir, resource_type, max_examples):
    compliant, non_compliant = [], []
    for tf_file in islice(glob.iglob(str(test_dir / "negative*.tf")), max_examples):
        if (code := read_file_contents(tf_file)):
            compliant.append(f"```{resource_type}\n{code}\n```")
    for tf_file in islice(glob.iglob(str(test_dir / "positive*.tf")), max_examples):
        if (code := read_file_contents(tf_file)):
            non_compliant.append(f"```{resource_type}\n{code}\n```")
    return compliant, non_compliant

def build_markdown(rule_path, metadata, cloud_provider, resource_type, max_examples):
    rule_name = rule_path.name
    title = metadata.get("queryName", "Untitled Rule")
    rule_id = metadata.get("id", "unknown-id")
    severity = metadata.get("severity", "INFO").upper()
    category = metadata.get("category", "Unknown")
    description = (
        metadata.get("override", {})
                .get("detailedDescriptionText", {})
                .get("descriptionText")
        or metadata.get("descriptionText", "No description provided.")
    )
    description_url = metadata.get("descriptionUrl")
    compliant, non_compliant = get_code_snippets(rule_path / "test", resource_type, max_examples)
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
    if compliant:
        markdown += "\n\n## Compliant Code Examples\n" + "\n\n".join(compliant)
    if non_compliant:
        markdown += "\n## Non-Compliant Code Examples\n" + "\n\n".join(non_compliant)
    return markdown

def load_list(path):
    try:
        with open(path, "r") as f:
            return json.load(f)
    except Exception as e:
        sys.exit(f"Error loading providers JSON: {e}")

def process_provider(provider, resource_type, input_dir, output_dir, max_examples, list_json_data, dict_frontmatter):
    provider_path = input_dir / resource_type / provider
    if not provider_path.is_dir():
        print(f"Warning: Missing provider path: {provider_path}")
        return

    output_provider_path = output_dir / resource_type / provider
    output_provider_path.mkdir(parents=True, exist_ok=True)

    provider_entry = {
        "name": provider,
        "short_description": f"{provider.upper()} Rules",
        "rules": []
    }

    dict_frontmatter[provider] = {}

    for rule_dir in provider_path.iterdir():
        if not rule_dir.is_dir():
            continue

        metadata_file = rule_dir / "metadata.json"
        if not metadata_file.exists():
            print(f"Skipping {rule_dir.name} — missing metadata.json")
            continue

        try:
            with open(metadata_file, "r", encoding="utf-8") as f:
                metadata = json.load(f)
        except Exception as e:
            print(f"Failed to parse metadata for {rule_dir}: {e}")
            continue

        rule_name = rule_dir.name
        rule_desc = metadata.get("queryName", "No description provided")

        provider_entry["rules"].append({
            "name": rule_name,
            "short_description": rule_desc
        })

        dict_frontmatter[provider][rule_name] = {
            "title": rule_desc,
            "description": rule_desc
        }

        md_content = build_markdown(rule_dir, metadata, provider, resource_type, max_examples)
        output_file = output_provider_path / f"{rule_name}.md"
        with open(output_file, "w", encoding="utf-8") as f:
            f.write(md_content)
        print(f"Generated: {output_file}")

    provider_entry["rules"].sort(key=lambda r: r["name"])
    list_json_data.append(provider_entry)

def main():
    args = parse_args()
    input_dir = args.input_dir
    output_dir = Path(args.output_dir)
    list_json_path = args.list_json
    dict_yaml_path = args.frontmatter_yaml
    max_examples = args.max_examples

    if output_dir.exists():
        shutil.rmtree(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    resource_type_dict = load_list(args.resources_json)

    list_json_data = []
    dict_yaml_data = {"rules":{}}
    
    for resource_type, providers in resource_type_dict.items():
        resource_path = input_dir / resource_type
        if not resource_path.is_dir():
            if resource_type != "default": print(f"Warning: Missing resource path: {resource_path}")
            continue

        resource_entry = {
            "name": resource_type,
            "providers" : []
        }
        list_json_data.append(resource_entry)
        dict_yaml_data["rules"][resource_type] = {}

        providers = providers if len(providers) > 0 else resource_type_dict["default"]
        for provider in providers:
            process_provider(provider, resource_type, input_dir, output_dir, max_examples, list_json_data[-1]["providers"], dict_yaml_data["rules"][resource_type])
        
        list_json_data[-1]["providers"].sort(key=lambda p: p["name"])

    list_json_data.sort(key=lambda p: p["name"])

    try:
        with open(dict_yaml_path, "w", encoding="utf-8") as f:
            yaml.dump(dict_yaml_data, f)
        print(f"Generated frontmatter yaml: {dict_yaml_path}")
    except:
        sys.exit("Failed to write frontmatter.yaml")

    try:
        with open(list_json_path, "w", encoding="utf-8") as f:
            json.dump(list_json_data, f, indent=2, ensure_ascii=False)
        print(f"Generated list JSON: {list_json_path}")
    except Exception as e:
        print(f"Failed to write list.json: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()