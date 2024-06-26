# service.place

The database schema

![image](https://github.com/turistikrota/service.place/assets/76786120/d82ce0ce-0856-42e7-8d43-64ef01f0b3b1)

## Run Project

How to run project

### 1. add swarm network

```bash
docker network create --driver overlay --attachable turistikrota

```

### 2. add secrets

```bash
docker secret create jwt_private_key ./jwtRS256.key
docker secret create jwt_public_key ./jwtRS256.key.pub

```

### 3. build image

```bash
docker build --build-arg GITHUB_USER=<GITHUB_USER> --build-arg GITHUB_TOKEN=<GITHUB_TOKEN> -t github.com/turistikrota/service.place .  
```

### 4. run container

```bash
docker service create --name place-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6019:6019 github.com/turistikrota/service.place:latest
```
