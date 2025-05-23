#!/bin/bash

# This script searches for metadata.json files and replaces hashtags (#)
# in override.detailedDescriptionText.descriptionText with //

# Requires: jq

find . -type f -name "metadata.json" | while read -r file; do
    # Extract the descriptionText
    description=$(jq -r '
        try .override.detailedDescriptionText.descriptionText // empty
    ' "$file" 2>/dev/null)

    # Check if descriptionText contains a hashtag
    if [[ "$description" == *"#"* ]]; then
        # Replace # with //
        new_description=$(echo "$description" | sed 's/#/\/\//g')

        # Use jq to update the field and overwrite the file
        jq --arg new_desc "$new_description" '
            .override.detailedDescriptionText.descriptionText = $new_desc
        ' "$file" > tmp.$$ && mv tmp.$$ "$file"

        echo "Updated: $file"
    fi
done

