#!/usr/bin/env bash

PROJECT_NAME="echt"
ENDPOINT="main.go"
TARGET_DIR="target"

MIRROR_PACKAGES=("sys", "net")

function glide_mirrors {
    for pkg in ${MIRROR_PACKAGES[*]}
    do
        echo "Setting mirror for package ${pkg}."
        bash -c "glide mirror set https://golang.org/x/${pkg} https://github.com/golang/${pkg} --vcs git"
    done
}

function install_requirements {
    echo "Checking requirements..."
    glide_mirrors
    bash -c "glide install"
}


case "$1" in
    "requirements")
        install_requirements ;;

    "build")
        if [[ "--check-req" = $2 ]]; then
            install_requirements
        fi

        if [[ ! -d ${TARGET_DIR} ]]; then
            mkdir ${TARGET_DIR}
        fi

        bash -c "go build -o $(pwd)/${TARGET_DIR}/${PROJECT_NAME}"

        echo ""
        echo "Successfully built $(pwd)/${TARGET_DIR}/${PROJECT_NAME}" ;;

    "server")
        bash -c "go run ${ENDPOINT}" ;;

    *)
esac
