#!/bin/bash

# Function to colorize output
colorize() {
    case $1 in
        "green")
            echo -e "\033[0;32m${@:2}\033[0m"
            ;;
        "blue")
            echo -e "\033[0;34m${@:2}\033[0m"
            ;;
        "yellow")
            echo -e "\033[0;33m${@:2}\033[0m"
            ;;
        *)
            echo "${@:2}"
            ;;
    esac
}

# Running commands with colorized output
colorize "blue" "Running: dagger -m dagger call hello"
colorize "green" $(dagger -m dagger call hello)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger -m dagger call test --dir=app"
colorize "green" $(dagger -m dagger call test --dir=app)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger -m dagger call build-and-publish --dir=app --args=\"-C\",\"./app\""
colorize "green" $(dagger -m dagger call build-and-publish --dir=app --args="-C","./")

colorize "yellow" "-------------------------------"

colorize "blue" "Verifying: docker run --name app --rm -it ttl.sh/my-hello-container:latest /usr/local/bin/app"
colorize "green" $(docker run --name app --rm -it ttl.sh/my-hello-container:latest /usr/local/bin/app)

colorize "yellow" "-------------------------------"
