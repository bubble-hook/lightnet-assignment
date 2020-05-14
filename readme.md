

## top-level directory layout

    .
    ├── client_simulator        # test call proxy server 
    ├── internal                # contains the real logic
    ├── proxy                   # handle request and forward to internal
    ├── shared                  # shared func 




# How to run

run.sh


```bash
echo  "Starting Internal Server"
cd ./internal && PORT=9000 go run main.go &
echo  "Starting Proxy Server"
cd ./proxy && PORT=9001 SERVICE_ENDPOINT="http://localhost:9000" go run main.go &
```
or 

```bash
sh run.sh
```


# How to run test
```bash
cd ./internal && go clean -testcache && go test ./services
```
or

```bash
sh test-service.sh
```








