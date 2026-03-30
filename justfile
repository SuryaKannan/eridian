[private]
default:
    just --list

# run eridian app
[group: 'go']
run:
    go run main.go