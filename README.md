# go-microservice-example  
Example for my blog: https://bognov.tech/  
You can find the article here:  https://bognov.tech/modern-api-design-with-golang-postgresql-and-docker

# start docker
docker compose up --build

# curl commands  

curl -X GET localhost:8080/comments | jq  
curl -X POST localhost:8080/comments -d "@request.json" | jq  
curl -X PUT localhost:8080/comments/1 -d "@request.json" | jq  
curl -X DELETE localhost:8080/comments/2 -d "@request.json" | jq