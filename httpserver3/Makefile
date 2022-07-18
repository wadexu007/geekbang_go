export tag=1.0.5

hello:
	echo "Welcome to pizza order system"

local:
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t asia.gcr.io/devops-apac-mgmt/exercise-pizza:${tag} .

push: build
	echo "pushing to my gcr"
	docker push asia.gcr.io/devops-apac-mgmt/exercise-pizza:${tag}