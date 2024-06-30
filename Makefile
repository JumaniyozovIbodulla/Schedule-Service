CURRENT_DIR := $(shell pwd)

gen-proto:
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/schedules
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/schedules
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/lessons
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/lessons
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/tasks
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/tasks
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/attendances
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/attendances
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/schedules --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/schedules --go-grpc_opt=paths=source_relative protos/schedule_protos/schedules.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/lessons --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/lessons --go-grpc_opt=paths=source_relative protos/schedule_protos/lessons.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/tasks --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/tasks --go-grpc_opt=paths=source_relative protos/schedule_protos/tasks.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/attendances --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/attendances --go-grpc_opt=paths=source_relative protos/schedule_protos/attendances.proto
	
run:
	go run cmd/main.go

git-push:
	git push origin main