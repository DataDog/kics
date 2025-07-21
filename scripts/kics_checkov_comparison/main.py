import argparse
import csv
import json
import sys
import pandas as pd

from pathlib import Path


def parse_args():
    parser = argparse.ArgumentParser(
        description="Automatically compare Kics and Checkov analysis"
    )
    parser.add_argument("checkov_csv", type=Path, help="Path to checkov's results")
    parser.add_argument(
        "--sarif", default="kics-result.sarif", type=str, help="Path to kics results"
    )
    parser.add_argument(
        "--compare", default="script/kics_checkov_comparison/Checkov_vs_KICS_Diff.json"
    )


def load_kics_results(path: Path):
    try:
        with open(path, "r") as f:
            raw_kics_results = json.load(f)
    except Exception as e:
        sys.exit(f"Error loading results JSON: {e}")

    if len(raw_kics_results["runs"]) == 0:
        return
    kics_results = {}
    for result in raw_kics_results["runs"][0]["results"]:
        lower = result["ruleId"].lower()
        if not (lower in kics_results):
            kics_results[lower] = 0
        kics_results[lower] = kics_results[lower] + 1
    return kics_results


def clean(path: Path):
    try:
        with open(path, "r") as f:
            rows = list(csv.DictReader(f))
    except Exception as e:
        sys.exit(f"Error loading results CSV: {e}")

    translation_dict = {}
    for row in rows:
        if not (row["Checkov_RuleID"] in translation_dict):
            translation_dict[row["Checkov_RuleID"]] = ""
        translation_dict[row["Checkov_RuleID"]] = (
            row["KICS_Description"].lower()
            if row["KICS_Description"] != ""
            else translation_dict[row["Checkov_RuleID"]]
        )
    try:
        with open(
            "scripts/kics_checkov_comparison/Checkov_vs_KICS_Diff.json", "w"
        ) as f:
            json.dump(translation_dict, f)
    except:
        sys.exit(f"Error dumping results JSON: {e}")
    return translation_dict


def load_checkov_results(path: Path):
    try:
        with open(path, "r") as f:
            checkov_results = {}
            for row in list(csv.DictReader(f)):
                checkov_results[row["group"][11::]] = row["value"]
            return checkov_results
    except Exception as e:
        sys.exit(f"Error loading results CSV: {e}")


def load_translation(path: Path):
    try:
        with open(path, "r") as f:
            return json.load(f)
    except Exception as e:
        sys.exit(f"Error loading translation JSON: {e}")


def load_severity(path: Path):
    try:
        with open(path, "r") as f:
            kics_severity = {}
            for row in list(csv.DictReader(f)):
                kics_severity[row["Query Name"].lower()] = row["Severity"]
            return kics_severity
    except Exception as e:
        sys.exit(f"Error loading severity CSV: {e}")


def compare(
    checkov_path: Path, kics_path: Path, translation_path: Path, severity_path: Path
):
    kics_results = load_kics_results(kics_path)
    checkov_results = load_checkov_results(checkov_path)
    translation_dict = load_translation(translation_path)
    severity_dict = load_severity(severity_path)
    comparisons = []
    for result in checkov_results:
        comparison = {}
        comparison["Checkov_ID"] = result
        comparison["Checkov_number"] = checkov_results[result]
        comparison["Kics_ID"] = translation_dict[result]
        comparison["Kics_number"] = (
            0
            if not (comparison["Kics_ID"] in kics_results)
            else kics_results[comparison["Kics_ID"]]
        )
        comparison["Ratio"] = float(comparison["Kics_number"]) / float(
            comparison["Checkov_number"]
        )
        comparison["Severity"] = (
            "?"
            if not (comparison["Kics_ID"] in severity_dict)
            else severity_dict[comparison["Kics_ID"]]
        )
        comparisons.append(comparison)
    pd.DataFrame(comparisons).to_csv("out.csv", index=False)


if __name__ == "__main__":
    # print(load_checkov_results(Path("findings-by-check-id.csv")))
    # print(load_kics_results(Path("kics-result.sarif")))
    # print(clean("scripts/kics_checkov_comparison/Checkov_vs_KICS_Diff.csv"))
    compare(
        Path("findings-by-check-id.csv"),
        Path("kics-result.sarif"),
        Path("scripts/kics_checkov_comparison/Checkov_vs_KICS_Diff.json"),
        Path("scripts/kics_checkov_comparison/KICS_Rules.csv"),
    )
