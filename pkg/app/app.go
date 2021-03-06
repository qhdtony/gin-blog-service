package app
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-blog-service/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}
type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, TotalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager {
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: TotalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg": err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["detail"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}