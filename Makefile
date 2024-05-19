build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.place . 

run:
	docker service create --name place-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6019:6019 github.com/turistikrota/service.place:latest

remove:
	docker service rm place-api-turistikrota-com

stop:
	docker service scale place-api-turistikrota-com=0

start:
	docker service scale place-api-turistikrota-com=1

restart: remove build run
	