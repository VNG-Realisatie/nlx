#!/bin/ash
ARCH=$(uname -m)
case $ARCH in
    armv5*) ARCH="ARM";;
    armv6*) ARCH="ARM";;
    armv7*) ARCH="ARM";;
    aarch64) ARCH="ARM64";;
    x86) ARCH="32bit";;
    x86_64) ARCH="64bit";;
    i686) ARCH="386";;
    i386) ARCH="386";;
    *)
        echo "Sorry, could not find the Hugo binary for your architecture: $ARCH"
        exit 1
        ;;
esac

wget -O - "https://github.com/gohugoio/hugo/releases/download/v0.42.1/hugo_0.42.1_Linux-${ARCH}.tar.gz" | tar --no-same-owner -C /usr/local/bin/ -xz hugo
