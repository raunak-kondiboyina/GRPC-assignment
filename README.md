# GRPC-assignment

## GRPC client server model

make sure you have following installed on your mac before you start running the program
grpc,redis,redigo,protoc-gen-go

else
$ make installprotoc
$ make installgrpc
$ make installredis
$ make installredigo

Steps to run this program on your machine

1) clone the repository in the .github.com folder on your mac

2) move to GRPC-assignment folder ($ cd .github.com/GRPC-assignment)

3) $ make runserver (to run server)

4) open another terminal in the GRPC-assignment directory
   $ make runclient (to run client)
   enter any string you want to store in redis database and press ENTER

To see data stored in database
$ make printDBlist

To run unit testing on isvalid function
$ make UT_isvalid


