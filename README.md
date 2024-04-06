# go-kafka
go-kafka sample


First execute the docker-compose using the below command

`docker-compose up`

then it will run the zookeeper server for kafka streaming

then create a topic using the command in `topic.txt`

once created topic

run the `producer.go` using command
>`go run producer.go`

run the consumer using command
>`go run consumer.go`

you will get the message what you have published.
