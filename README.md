# GoWebTest

Requirements: docker
Go to https://www.docker.com/ and download and install docker

open a terminal and execute the following to verify installation
$ docker version


Download repository

Open terminal to Webtest directory and run the following 
$ docker build -t webtest . 
This will look inside the current directory "." and execute the Dockerfile located there
Don't forget to add the "." at the end or you won't be providing a directory to load from 

You should see step 1/2 executing 
    This will build the docker container and it's associated image
After the second step compleates run the following to check the image
$ docker images
    you should now see webtest listed as a repository in the images catalog 

To run the corresponding image execute the following command in your terminal
$ docker run --publish 80:8080 --name test --rm webtest

open a web browsre and navigate to http://localhost/


