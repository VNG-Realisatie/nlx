#!/usr/bin/awk -f

BEGIN {
    statementCount = 0;
    coverageCount = 0;
}

{
    # Echo all lines
    print $0
   
}

# Match the release lines from skaffold
/,\w+.\w+\s(\w+)\s(\w+)/ {
    print("le match")
    print($2)
    print($3)
    statementCount+=$2
    coverageCount+=$3
   
}

END {
    coverage = (coverageCount / statementCount) * 100
    print(coverage)
    if (coverage > 0) {
        printf "coverage percentage of all packages: %.1f%\n", coverage
    } else {
        print "coverage percentage of all packages: 0.0%"
    }
   
}

