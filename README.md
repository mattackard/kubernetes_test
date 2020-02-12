# kubernetes_test
This project uses Go to create an HTTP server that echoes the IP address of the server sending the request.

Running the build script will build the binary for the project including dependencies needed to run inside a docker image from scratch.

The deploy.yaml file is used to define the kubernetes deployment.

The kubernetes deployment relies on pulling the ip-echo image from dockerhub, so if you want to change how the go program works, you will need to create your own image on dockerhub and reference that image from within the deploy.yaml file.
