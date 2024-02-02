package router

import (
	"net/http"

	"github.com/sangeeth518/learning-routes/constant"
	"github.com/sangeeth518/learning-routes/controller"
)

var user = Routes{
	Route{"Hello User", http.MethodGet, constant.HelloUserPath, controller.HelloUser},
}
