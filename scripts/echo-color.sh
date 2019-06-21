#!/bin/bash

function echoSuccess {
    # 32:green
    echo -e "\033[32m$1\033[m"
}

function echoError {
    # 31:red
    echo -e "\033[31m$1\033[m"
}

function echoWarning {
    # 33:green
    echo -e "\033[33m$1\033[m"
}

function echoInfo {
    # 34:blue
    echo -e "\033[34m$1\033[m"
}
