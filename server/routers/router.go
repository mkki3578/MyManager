package routers

import (
	. "github.com/fishedee/web"
	. "mymanager/controllers"

	_ "mymanager/models/register"
)

func init() {
	InitRoute("user", &UserController{})
	InitRoute("login", &LoginController{})
	InitRoute("category", &CategoryController{})
	InitRoute("card", &CardController{})
	InitRoute("account", &AccountController{})
	InitRoute("blog", &BlogController{})
}
