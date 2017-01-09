#!/usr/bin/env bash

function main {
    case "$1" in

    "install")
    install;;

    esac
}

function terraform-not-installed {
    echo "checking if terraform is installed"
    if ! type "$terraform" > /dev/null; then
        echo "Terraform not installed"
        return 0
    fi
    return 1
}

function download-terraform {
    curl -o terraform.zip https://releases.hashicorp.com/terraform/0.8.2/terraform_0.8.2_darwin_amd64.zip
    unzip terraform.zip
}

function install-terraform {
    echo "Installing terraform"
    download-terraform
}

function cleanup {
    echo "deleting installation file"
    rm -rf terraform.zip
}

function install {
    if terraform-not-installed; then
        install-terraform
    fi
    echo "running terraform"
    
    cleanup
}

function fail {
    echo "Please use one of the expected commands 'install'"
    exit 1
}

main $@
exit 0
