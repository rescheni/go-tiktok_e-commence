# Makefile for GoMall

.PHONY:
/dev/null: makefile
	@make help


.PHONY: help
help:
	@echo "可用命令："
	@echo "  make air-test           # 热启动前端"
	@echo "  make gen-frontend-home  # 生成前端主页代码"
	@echo "  make gen-frontend-auth  # 生成前端认证代码"
	@echo "  make gen-rpc-user-client      # 生成用户服务代码-客户端"

# 热启动测试运行
.PHONY: air-test
air-test:
	@cd app/frontend && air


# 前端框架主页生成
.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/home.proto --service frontend -module gomall -I ../../idl
# 前端框架 用户认证代码生成
.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto --service frontend -module gomall -I ../../idl

.PHONY: gen-rpc-user-client
gen-rpc-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module e-commence -I ../idl --idl ../idl/service/user.proto	

.PHONY: gen-rpc-user-server
gen-rpc-user-server:
	@cd app/user && cwgo server --type RPC --service user --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/user.proto
