# Overview

- [Install](#install)
    - [install docker with convenience script](#install-docker-with-convenience-script)
- [Delete all containers](#delete-all-containers)
- [Stop all containers](#stop-all-containers)
- [stop all containers](#stop-all-containers)
- [Delete all images](#delete-all-images)


# Install

```shell
## install docker with convenience script
curl -fsSL https://get.docker.com -o get-docker.sh
DRY_RUN=1 sh ./get-docker.sh
```

# Delete all containers

```shell
docker rm $(docker ps -aq)
```

# Stop all containers

```shell
# stop all containers
docker stop $(docker ps -aq)
```

# Delete all images

```shell
docker rmi $(docker images -q)
```

