[private]
default:
    just --list

# run eridian app
[group: 'go']
run:
    go run main.go

# build python app and install as tool
[group: 'python']
update-py:
    uv tool install . --force

# fetch openapi spec and regenerate go client                                                                                                                                                
[group: 'python']                                                                                                                                                                            
openapi port:                                                                                                                                                        
    curl -s http://127.0.0.1:{{port}}/api/openapi.json | python -m json.tool > internal/mlservice/openapi.json
    oapi-codegen -generate types,client -package mlservice -o internal/mlservice/client.gen.go internal/mlservice/openapi.json 