#!/usr/bin/awk -f

BEGIN {
    count = 0;
}

{
    # Echo all lines
    print $0
}

# Match the release lines from skaffold
/^nlxio\/[a-z-]* -> nlxio\/[a-z-]*:[0-9a-z.-]*@[0-9a-f]*$/ {
    cmdTag=sprintf("docker tag %s %s:latest", $3, $1 )
    print cmdTag
    c = system(cmdTag)
    if (c!=0) {
        print "Executing `" cmdTag "` failed."
        exit 1
    }

    cmdPush=sprintf("docker push %s:latest", $1)
    print cmdPush
    c = system(cmdPush)
    if (c!=0) {
        print "Executing `" cmdPush "` failed."
        exit 1
    }

    count++
}

END {
    print "Tagged and pushed " count " images with the `:latest` tag."
    if (count==0) {
        print "Error: no :latest tag images were pushed."
        exit 1
    }
}
