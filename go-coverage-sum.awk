#!/usr/bin/awk -f

BEGIN {
    count = 0;
    totalPercentage = 0;
}

{
    # Echo all lines
    print $0
   
}

# Match the release lines from skaffold
/^coverage:\s(\w+.\w+)% of statements/ {
    totalPercentage+=$2
    count++
}

END {
    if (count > 0) {
        printf "coverage percentage of all packages: %.1f%\n", totalPercentage / count    
    } else {
        print "coverage percentage of all packages: 0.0%"
    }
   
}

