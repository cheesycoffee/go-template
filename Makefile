proto :
	$(foreach proto_file, $(shell find api/proto -name '*.proto'), \
	protoc --proto_path=api/proto --go_out=api/proto \
	--go_opt=paths=source_relative $(proto_file);)