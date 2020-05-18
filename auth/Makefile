docker-clear:
	sudo docker image prune -f

build-python-requirements:
	pipenv run pip freeze > requirements.txt

build-container:
	sudo docker build --no-cache -t auth-proto-service -f _infrastructure/Dockerfile .

build:
	make build-python-requirements && make build-container && make docker-clear
