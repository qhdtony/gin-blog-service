package errcode
var {
	Success							:= new NewError(0, "成功")
	ServerError						:= new NewError(400001, "服务内部错误")
	InvalidParams					:= new NewError(400002, "入参错误")
	NotFound						:= new NewError(400002, "找不到")
	UnauthorizedAuthNotExist		:= new NewError(400003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError			:= new NewError(400004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout		:= new NewError(400005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate		:= new NewError(400006, "鉴权失败，Token 生成失败")
	TooManyRequests					:= new NewError(400007, "请求过多")
}