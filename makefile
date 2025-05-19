# Makefile for GoMall

.PHONY: /dev/null
/dev/null: makefile
	@make help


.PHONY: help
help:
	@echo "可用命令："
	@echo "  make air-test           # 热启动前端"
	@echo "  make gen-frontend-home  # 生成前端主页代码"
	@echo "  make gen-frontend-auth  # 生成前端认证代码"
	@echo "  make gen-rpc-user-client      # 生成用户服务代码-客户端"
	@echo "  make gen-rpc-user-server      # 生成用户服务代码-客户端"
	@echo "  run_rpc-user      		# 运行用户rpc服务代码"
	@echo "  run-frontend           # 启动前端服务"	

# 热启动测试运行
.PHONY: air-test
air-test:
	@cd app/frontend && air

.PHONY:启动前端
run-frontend:
	@cd app/frontend && go run .

# 前端框架主页生成
.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/home.proto --service frontend -module gomall -I ../../idl
# 前端框架 用户认证代码生成
.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto --service frontend -module gomall -I ../../idl

.PHONY: run_rpc_user
run-rpc-user:
	@cd app/user && go run .

.PHONY: gen-rpc-user-client
gen-rpc-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module e-commence -I ../idl --idl ../idl/service/user.proto	

.PHONY: gen-rpc-user-server
gen-rpc-user-server:
	@cd app/user && cwgo server --type RPC --service user --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/user.proto
