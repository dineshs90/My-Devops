# To Run go app:

go run main.go

go build .

# created executable file name based on go.mod file 
./simple-app

# Build Image based on dockerfile

docker build -t my-app .

# Push Docker Image to docker hub

docker login
docker tag golang-dockercompose:latest dinesh790/golang-dockercompose:latest
docker push dinesh790/golang-dockercompose:latest


# To run the docker image using kubernetes

kubectl apply -f deployment.yaml 

# Jenkins

sudo systemctl stop jenkins
sudo systemctl start jenkins
sudo systemctl status jenkins
sudo systemctl restart jenkins

# Credentials 

# Remote Host:
ssh-keygen -t rsa -b 4096 -C "jenkins@192.168.1.5"
ssh-copy-id osboxes@<IP>
ssh osboxes@<IP>

# Docker:
username:
password

# Github:
ssh-keygen -t rsa -b 4096 -C "jenkins@192.168.1.5"
~/.ssh/id_rsa


# minikube 
minikube service golang-service --url