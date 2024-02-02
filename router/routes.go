package router

import (
	"net/http"

	"github.com/sangeeth518/learning-routes/constant"
	"github.com/sangeeth518/learning-routes/controller"
)

//this routes represent user

var user = Routes{
	Route{"Hello User", http.MethodPost, constant.HelloUserPath, controller.HelloUser},
}
