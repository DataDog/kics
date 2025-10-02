#!/usr/bin/env python3
import json
import sys
from pathlib import Path

def extract_k8s_metadata_ids(k8s_dir):
    """Extract all metadata IDs from k8s query metadata.json files."""
    k8s_path = Path(k8s_dir)

    if not k8s_path.exists():
        print(f"Error: Directory {k8s_dir} does not exist")
        return []

    metadata_ids = []

    # Find all metadata.json files in k8s subdirectories
    for metadata_file in k8s_path.rglob("metadata.json"):
        try:
            with open(metadata_file, "r", encoding="utf-8") as f:
                metadata = json.load(f)

            rule_id = metadata.get("id", "")
            rule_name = metadata.get("queryName", "")

            if rule_id:
                metadata_ids.append({
                    "id": rule_id,
                    "queryName": rule_name,
                    "path": str(metadata_file.relative_to(k8s_path.parent))
                })
            else:
                print(f"Warning: No ID found in {metadata_file}", file=sys.stderr)

        except Exception as e:
            print(f"Error processing {metadata_file}: {e}", file=sys.stderr)

    return metadata_ids

def main():
    if len(sys.argv) != 2:
        print("Usage: python extract_k8s_ids.py <k8s_directory_path>")
        sys.exit(1)

    k8s_dir = sys.argv[1]
    metadata_ids = extract_k8s_metadata_ids(k8s_dir)

    # Extract just the IDs and format them
    formatted_ids = [f"rule_ids: {item['id']}" for item in metadata_ids]

    # Print OR-separated IDs
    or_separated_ids = ' OR '.join(formatted_ids)
    print(or_separated_ids)

    # Save to file as well
    with open("k8s_ids_or_separated.txt", "w") as f:
        f.write(or_separated_ids)

    print(f"\n# Total k8s metadata IDs found: {len(metadata_ids)}", file=sys.stderr)
    print(f"# Results also saved to k8s_ids_or_separated.txt", file=sys.stderr)

if __name__ == "__main__":
    main()
