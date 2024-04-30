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
colorize "blue" "Running: dagger --src ./app call echo --string-arg hello stdout"
colorize "green" $(dagger --src ./app call echo --string-arg hello stdout)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app message text"
colorize "green" $(dagger call --src ./app message text)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app message from"
colorize "green" $(dagger call --src ./app message from)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app hello"
colorize "green" $(dagger call --src ./app hello)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app print-greet"
colorize "green" $(dagger call --src ./app print-greet)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app new-greet print-greet"
colorize "green" $(dagger call --src ./app new-greet print-greet)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app deps test output"
colorize "green" $(dagger call --src ./app deps test output)

colorize "yellow" "-------------------------------"

colorize "blue" "Running: dagger call --src ./app deps build publish"
colorize "green" $(dagger call --src ./app deps build publish)

colorize "yellow" "-------------------------------"

colorize "blue" "Verifying: docker run --name app --rm -it ttl.sh/go-dagger-container:latest /app/build"
colorize "green" $(docker run --name app --rm -it ttl.sh/go-dagger-container:latest /app/build)

colorize "yellow" "-------------------------------"
