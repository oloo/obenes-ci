#!/usr/bin/env bash

function main {
    case "$1" in

    "execute")
        assign-inputs $@
        execute
        ;;
    *)
        fail
        ;;

    esac
}

function assign-inputs {
     echo "Assign inputs"
     ACCESS_KEY="$2"
     SECRET_KEY="$3"
     if [[ -z $ACCESS_KEY || -z $SECRET_KEY ]]; then
        echo "Please enter ACCESS_KEY and SECRET_KEY"
        exit 1
     fi
}

function terraform-not-installed {
    echo "checking if terraform file exists"
    if [ ! -f terraform ]; then
        echo "Terraform file not found"
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

function run-terraform {
    echo "Applying terraform for concourse CI"
    ./terraform apply \
        -var "access_key=${ACCESS_KEY}" \
        -var "secret_key=${SECRET_KEY}"
}

function install-terraform-if-needed {
     if terraform-not-installed; then
        install-terraform
    fi
}

function cleanup {
    echo "deleting installation file"
    rm -rf terraform.zip
}

function fail {
    echo "Please use one of the expected commands 'install'"
    exit 1
}

function execute {
   install-terraform-if-needed
   run-terraform
   cleanup
}

main $@
exit 0
