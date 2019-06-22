# MRT
Github link: https://github.com/herlegs/MRT

Usage guide:

* assume Go is already installed, and environment variables has been set ($GOPATH)

* run command to fetch repo if any dependencies is missing
>>go get -u github.com/herlegs/MRT

* under the code folder, run ~go run main.go~ to start server

* open another terminal, use curl to test the server:

>> curl http://localhost:8080/route?from=hollandvillage&to=Bugis&time=2019-06-22T15:30

Note:

endpoint is /route  
parameters for the endpoint are:

> from: start location  
> to: destination location  
time: in YYYY-MM-DDTHH:MM format  
(if time not specified, will search for shortest stations number only)