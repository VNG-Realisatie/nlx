#!/bin/bash

rc=0

# flags
f_flag=false

while getopts 'f' flag; do
  case "${flag}" in
    f) f_flag=true ;;
  esac
done

while read -r file; do
    # Look for too many newlines at end of file
    if [ "$(sed -n '$p' "$file")" = "" ]; then
        if [[ "$f_flag" = true ]]; then 
            printf '%s\n' "`cat "$file"`" > "$file"
            echo "fixed too many newlines at end of file $file"; 
        else 
            echo "$file: too many newlines at end of file" >&2;
            rc=1;
        fi
    fi

    # Look for no newlines at end of file
    if [ "$(tail -c 1 "$file")" != "" ]; then
        if [[ "$f_flag" = true ]]; then 
            echo >> "$file";
            echo "fixed no newline at end of file $file"; 
        else 
            echo "$file: no newline at end of file" >&2
            rc=1
        fi   
    fi
done < <(git ls-files | grep -v -E ".*\.(jpg|png|svg|drawio|ico|woff|jpg|jpg|gitkeep)$|management-ui/src/api/.*$")

exit $rc
