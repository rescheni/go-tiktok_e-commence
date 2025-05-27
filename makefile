# Makefile for GoMall

.PHONY: /dev/null
/dev/null: makefile
	@make help


.PHONY: help
help:
	@echo "可用命令："
	@echo "运行代码:"	
	@echo "  	make run-frontend           		# 启动前端服务"	
	@echo "  	make air-test           			# 热启动前端"
	@echo "  	make run-rpc-user      				# 运行用户rpc服务代码"
	@echo "  	make run-rpc-product       			# 运行商品rpc服务代码"
	@echo "  	make run-rpc-cart       			# 运行购物车rpc服务代码"
	@echo "		make run-rpc-payment 				# 运行支付 rpc 服务代码"
	@echo "		make run-rpc-checkout 				# 运行结账 rpc 服务代码"
	@echo "		make run-rpc-order  				# 运行结账 rpc 服务代码"
	
	@echo "frontend 代码生成"
	@echo "  	make gen-frontend-home  			# 生成前端主页代码"
	@echo "  	make gen-frontend-auth  			# 生成用户认证前端代码"
	@echo "  	make gen-frontend-product			# 生成商品服务前端代码"
	@echo "		make gen-frontend-checkout			# 生成购物结算前端代码"
	@echo "		make gen-frontend-order				# 生成订单服务前端代码"

	@echo "rpc 服务端 代码生成"
	
	@echo "  	make gen-rpc-user-client      		# 生成用户服务代码-客户端"
	@echo "  	make gen-rpc-user-server      		# 生成用户服务代码-服务端"
	
	@echo "  	make gen-rpc-product-client      	# 生成商品服务代码-客户端"
	@echo "  	make gen-rpc-product-server      	# 生成商品服务代码-服务端"

	@echo "  	make gen-rpc-cart-client      		# 生成购物车服务代码-客户端"
	@echo "  	make gen-rpc-cart-server      		# 生成购物车服务代码-服务端"
	
	@echo "  	make gen-rpc-payment-client      	# 生成支付服务代码-客户端"
	@echo "  	make gen-rpc-payment-server      	# 生成支付服务代码-服务端"

	@echo "  	make gen-rpc-checkout-client      	# 生成结账服务代码-客户端"
	@echo "  	make gen-rpc-checkout-server      	# 生成结账服务代码-服务端"

	@echo "  	make gen-rpc-order-client      		# 生成订单服务代码-客户端"
	@echo "  	make gen-rpc-order-server      		# 生成订单服务代码-服务端"

	@echo "  	make gen-rpc-email-client      		# 生成订单服务代码-客户端"
	@echo "  	make gen-rpc-email-server      		# 生成订单服务代码-服务端"


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
# 前端框架 用户认证服务代码生成
.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto --service frontend -module gomall -I ../../idl

# 前端框架 商品服务代码生成
.PHONY: gen-frontend-product
gen-frontend-product:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/product_page.proto --service frontend -module gomall -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/category_page.proto --service frontend -module gomall -I ../../idl
# 前端框架 购物车服务代码生成
.PHONY: gen-frontend-cart 
gen-frontend-cart:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/cart_page.proto --service frontend -module gomall -I ../../idl

# 前端框架 购物结算代码生成
.PHONY: gen-frontend-checkout 
gen-frontend-checkout:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/checkout_page.proto --service frontend -module gomall -I ../../idl


# 前端框架 订单服务代码生成 
.PHONY: gen-frontend-order
gen-frontend-order:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/order_page.proto --service frontend -module gomall -I ../../idl










# 用户服务
.PHONY: run-rpc-user
run-rpc-user:
	@cd app/user && go run .

.PHONY: gen-rpc-user-client
gen-rpc-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module e-commence -I ../idl --idl ../idl/service/user.proto	

.PHONY: gen-rpc-user-server
gen-rpc-user-server:
	@cd app/user && cwgo server --type RPC --service user --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/user.proto


# 商品服务
.PHONY: run-rpc-product
run-rpc-product:
	@cd app/product && go run .

.PHONY: gen-rpc-product-client
gen-rpc-product-client:
	@cd rpc_gen && cwgo client --type RPC --service product --module e-commence -I ../idl --idl ../idl/service/product.proto	

.PHONY: gen-rpc-product-server
gen-rpc-product-server:
	@cd app/product && cwgo server --type RPC --service product --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/product.proto


# 购物车服务

.PHONY: run-rpc-cart 
run-rpc-cart:
	@cd app/cart  && go run .

.PHONY: gen-rpc-cart-client 
gen-rpc-cart-client:
	@cd rpc_gen && cwgo client --type RPC --service cart --module e-commence -I ../idl --idl ../idl/service/cart.proto	

.PHONY: gen-rpc-cart-server
gen-rpc-cart-server:
	@cd app/cart && cwgo server --type RPC --service cart --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/cart.proto


# 支付服务

.PHONY: run-rpc-payment
run-rpc-payment:
	@cd app/payment  && go run .

.PHONY: gen-rpc-payment-client 
gen-rpc-payment-client:
	@cd rpc_gen && cwgo client --type RPC --service payment --module e-commence -I ../idl --idl ../idl/service/payment.proto	

.PHONY: gen-rpc-payment-server
gen-rpc-payment-server:
	@cd app/payment && cwgo server --type RPC --service payment --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl --idl ../../idl/service/payment.proto

# 结账服务
.PHONY: run-rpc-checkout
run-rpc-checkout:
	@cd app/checkout  && go run .

.PHONY: gen-rpc-checkout-client 
gen-rpc-checkout-client:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module e-commence -I ../idl/service --idl ../idl/service/checkout.proto	

.PHONY: gen-rpc-checkout-server
gen-rpc-checkout-server:
	@cd app/checkout && cwgo server --type RPC --service checkout --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl/service --idl ../../idl/service/checkout.proto


# 订单服务
.PHONY: run-rpc-order
run-rpc-order:
	@cd app/order  && go run .

.PHONY: gen-rpc-order-client 
gen-rpc-order-client:
	@cd rpc_gen && cwgo client --type RPC --service order --module e-commence -I ../idl/service --idl ../idl/service/order.proto	

.PHONY: gen-rpc-order-server
gen-rpc-order-server:
	@cd app/order && cwgo server --type RPC --service order --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl/service --idl ../../idl/service/order.proto



# 邮件服务
.PHONY: run-rpc-email
run-rpc-email:
	@cd app/email  && go run .

.PHONY: gen-rpc-email-client 
gen-rpc-email-client:
	@cd rpc_gen && cwgo client --type RPC --service email --module e-commence -I ../idl/service --idl ../idl/service/email.proto	

.PHONY: gen-rpc-email-server
gen-rpc-email-server:
	@cd app/email && cwgo server --type RPC --service email --module e-commence --pass "-use e-commence/rpc_gen/kitex_gen"  -I ../../idl/service --idl ../../idl/service/email.proto

