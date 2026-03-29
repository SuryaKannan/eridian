[private]
default:
    just --list

# build cli and install as tool
[group: 'python']
update-py:
    uv tool install . --force
