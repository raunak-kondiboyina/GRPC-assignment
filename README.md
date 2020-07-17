# GRPC-assignment

# GRPC client server model

## make sure you have following installed on your mac before you start running the program

grpc, redis, redigo,  protoc-gen-go

else to install, run following commands

$ make installprotoc

$ make installgrpc

$ make installredis

$ make installredigo


## Steps to run this program on your machine

1) clone the repository in the .github.com folder on your mac

2) move to GRPC-assignment folder 

   ($ cd github.com/GRPC-assignment)

3) to run server

   $ make runserver 

4) open another terminal in the GRPC-assignment directory to run client

   $ make runclient

   enter any string you want to store in redis database and press ENTER

## To see data stored in database

$ make printDBlist

## To run unit testing on isvalid function

$ make UT_isvalid


