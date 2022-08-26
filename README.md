# flonkerton
Welcome to flockerton, a web application used for testing data protection controls. The application is designed to be portable so that it can be deployed behind where ever it is needed to test the controls. Specifically it can be used to test a ZTNAs ability to block sensitive data movement when users are accessing a private application.


## Requirements
Before installing you must have the following requirements installed.  

- Ubuntu or Mac (These have been tested, but others may work as well.)
- [hugo] (https://gohugo.io/getting-started/installing/)
- [Docker] (https://docs.docker.com/engine/install/)
- [Docker Compose*] (https://docs.docker.com/compose/install/)
* Docker Compose is included in many docker instalations, so you may not need to add this separately.


## Install
1.) Clone this repo
```sh
git clone https://github.com/stjbrown/flonkerton.git
```
1.) Navigate to flonkerton directory
```sh
cd flonkerton
```
1.) Modify the `config.yml` file so that the base url contains the domain you will be using for flonkerton.

`baseURL: "flonkerton.example.com"`
1.) hugo -D
1.) From the flonkerton directory start the docker immages using docker compose.
```sh
docker compose ud -d
```
