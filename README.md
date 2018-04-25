# CanDIG

Database Access API for mongodb schmea


## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Installing this code:
`$ go get github.com/CanDIG/candig_mds`


Running it then should be as simple as:

```console
$ make build
$ ./bin/candig
```
To generate fake data, run `$ ./bin/candig -g=100`, where 100 is the amount of dummy data requested
To run on a spesific port run `$ ./bin/candig -p=8000`, where 8000 is the desired port



Endpoints:

```sh
# QUERIES

# This is an example query that will return all data from every table in the database
$ curl http://127.0.0.1:8000/Candig/query -X POST -H "content-type:application/json" -d 
'{"selected_fields":["*"],“selected_tables":[],"selected_conditions":[]}'


# INSERTS
# <object_name>
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.<object_name> for object structure`}' 127.0.0.1:8000/Candig/<object_name>
```

# Server Dependencies
- Go
- Go Libraries
- Mongodb

*Note*: The object structure for the database will create itself as long as mongo is running on the default port

