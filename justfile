alias b := build
alias c := check
alias f := fmt
alias i := install
alias l := lint
alias r := run
alias up := upgrade
alias upt := upgradet

@default:
    just --choose

build:
    go build -ldflags='-s -w' .

check:
    golangci-lint run

fmt:
    gofumpt -w .

install:
    go install -ldflags='-s -w' .

lint:
    golangci-lint run

run:
    go run .

upgrade:
    go get -u ./...

upgradet:
    go get -t -u ./...
