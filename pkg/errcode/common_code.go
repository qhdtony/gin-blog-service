package errcode
var (
	Success							= NewError(0, "成功")
	ServerError						= NewError(400001, "服务内部错误")
	InvalidParams					= NewError(400002, "入参错误")
	NotFound						= NewError(400003, "找不到")
	UnauthorizedAuthNotExist		= NewError(400004, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError			= NewError(400005, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout		= NewError(400006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate		= NewError(400007, "鉴权失败，Token 生成失败")
	TooManyRequests					= NewError(400008, "请求过多")
)