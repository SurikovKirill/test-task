.PHONY:

build-image:
	docker build -t test-task:v1.0 .

start-container:
	docker run -d --restart always --name test-task -p 80:80 test-task:v1.0

logs-container:
	docker logs test-task