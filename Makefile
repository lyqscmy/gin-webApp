person.pb.go: models/person.proto
	protoc --go_out=paths=source_relative:./ $?
