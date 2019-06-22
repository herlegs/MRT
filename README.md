# MRT
Github link: https://github.com/herlegs/MRT

Usage guide:

* assume Go is already installed, and environment variables has been set ($GOPATH)

* run command to fetch repo if any dependencies is missing
>go get -u github.com/herlegs/MRT

* under the code folder, run _go run main.go_ to start server

* open another terminal, use curl to test the server:

> curl http://localhost:8080/route?from=hollandvillage&to=Bugis&time=2019-06-22T15:30

Note:

endpoint is /route  
parameters for the endpoint are:

> from: start location  
> to: destination location  
time: in YYYY-MM-DDTHH:MM format  
(location name is case insensitive, space is removed for simplicity of Get request)  
(if time not specified, will search for shortest stations number only)  

example input and output:
```text
curl http://localhost:8080/route\?from\=jurongeast\&to\=hawparvilla\&time\=2019-06-22T15:30
```
```json
{
  "summary": "Travel from Jurong East to Haw Par Villa during normal hours",
  "station_travelled": "7",
  "route": [
    "EW24",
    "EW23",
    "EW22",
    "EW21",
    "CC22",
    "CC23",
    "CC24",
    "CC25"
  ],
  "travel_time": "70 minutes",
  "instruction": "Start at Jurong East\nTake EW line from Jurong East to Clementi\nTake EW line from Clementi to Dover\nTake EW line from Dover to Buona Vista\nChange from EW line to CC line\nTake CC line from Buona Vista to one-north\nTake CC line from one-north to Kent Ridge\nTake CC line from Kent Ridge to Haw Par Villa\nAnd you have arrived Haw Par Villa\n",
  "query_time": "time used: 12.12µs"
}
```