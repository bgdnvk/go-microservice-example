# go-microservice-example  
WIP example for my blog: bognov.tech

# start docker
docker compose up --build

# curl commands  

curl -X GET localhost:8080/comments | jq  
curl -X POST localhost:8080/comments -d "@request.json" | jq  
curl -X PUT localhost:8080/comments/1 -d "@request.json" | jq