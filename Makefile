run: setup-local build-local run-local
	
setup-local:
	docker network create -d bridge form-task-network
	docker pull mysql:8.0.25
	docker run -d --name form_database --network=form-task-network -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw mysql:8.0.25

build-local:
	docker build -t forms-api-image .
	docker tag forms-api-image forms-api-image:latest

run-local:
	docker stop forms_api
	docker run --rm --name=forms_api --network=form-task-network -d -p 5000:5000 forms-api-image:latest