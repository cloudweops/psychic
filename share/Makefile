# eg. make service=user generate
generate:
	kitex -module github.com/CloudWeOps/psychic ./../idl/rpc/$(service).thrift

user:
	make service=user generate

errno:
	kitex -module github.com/CloudWeOps/psychic ./../idl/base/errno.thrift

all:
	make user
	make errno

