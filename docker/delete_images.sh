# delete all docker images
docker rmi $(docker images -q)
