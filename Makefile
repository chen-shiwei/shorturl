server_name=shorturl
server_path=.
.PHONY: test
test:##@test 测试.
	go test -v ./... -cover



.PHONY:run
run:
	go run . api --addr=:7999

.PHONY:build-linux
build-linux: ##@build 构建二进制文件.
	@echo "构建二进制文件"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(server_name) $(server_path)
	upx $(server_name)
	@echo "可执行文件大小为:`du -sh $(server_name)`"

.PHONY:help
help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)


HELP_FUN = \
	%help; \
	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	print "usage: make [target]\n\n"; \
	for (sort keys %help) { \
		print "${WHITE}$$_:${RESET}\n"; \
		for (@{$$help{$$_}}) { \
			$$sep = " " x (32 - length $$_->[0]); \
			print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; }