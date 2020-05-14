kill $(lsof -ti:9000,9001)
echo  "==================="
echo  "Starting Internal Server"
cd ./internal && PORT=9000 go run main.go &
echo  "Starting Proxy Server"
cd ./proxy && PORT=9001 SERVICE_ENDPOINT="http://localhost:9000" go run main.go &