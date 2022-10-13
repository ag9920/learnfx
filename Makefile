gen_kitex:
	kitex -service learn.fx.item learn_fx.thrift 

install_kitex:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/thriftgo@latest