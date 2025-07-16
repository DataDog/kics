grep -rl "get_module_equivalent_key" ./assets | while read file; do                                                                                                                                             ─╯
    dest="./no_module_support$(echo "$file" | sed 's/^\.\/assets//')"
    mkdir -p "$(dirname "$dest")";  cp "$file" "$dest"
done
