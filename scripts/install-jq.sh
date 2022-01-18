#!/bin/bash
wget -O jq https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && chmod +x jq
PATH=$(pwd):$PATH
export PATH
