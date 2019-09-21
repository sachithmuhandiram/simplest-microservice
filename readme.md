## The simplest Microservice

Here I am developing a simple microservice based system using `golang`.

This system has an API gateway (mainmodule) and four microservices,
   * Add (addmodule)
   * Substraction 
   * Divide
   * Multiply


First API gateway will develop to handle just single operation in one request. Later it will be modified to handle complex mathematical operations.

Eg:
When API gateway gets `1+10(8-6)/2`
1. Sends (8-6) to substration service gets the answer.
2. Sends 2/2 to divide service and gets the answer.
3. Sends 1*10 to multiply service and gets the answer.
4. Sends 1+10 to add module and gets the final answer.

API gateway and other services will be developed using minimal 3rd party packages.

### Important

* All containers must be in the same network. `network_mode: "host"`, Otherwise it will give 

    > Get http://container1:port/: dial tcp 0.0.0.0:port: connect: connection refused

* Microservice containers should be running on `0.0.0.0`. [The host binding to 0.0.0.0 instead of localhost](https://chat.stackoverflow.com/rooms/198447/discussion-between-maartendev-and-sachith)

    Eg : `http.ListenAndServe("0.0.0.0:7070", nil)`

* Add microservice name and IP to `/etc/hosts`. [Connecting to docker with curl](https://stackoverflow.com/questions/41887775/connecting-to-docker-with-curl/41895590#41895590)

    Eg:
    `0.0.0.0        addmodule`

## To Do

- [ ] Add Unit tests for each module
- [ ] 2 phase builds of docker images. As I use four alpine:golang instances.
- [ ] API gateway functionalities one by one.
    - [ ] - Implementing circuite-braker pattern
    - [ ] - Service Descovery
    - [ ] - Security (http requests etc)
    - [ ] - Logical expression parsing

### How to run

`git clone https://github.com/sachithmuhandiram/simplest-microservice.git`

`sudo docker-compose up --build`

https://www.reddit.com/r/golang/comments/cweswg/golang_part_time_developer_beginners_level/ 