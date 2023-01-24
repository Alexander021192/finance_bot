generate:
	mockgen \
	-source=internal/model/msg/in_msg.go \
	-destination=internal/mocks/msg/msg_mocks.go