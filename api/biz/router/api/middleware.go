// Code generated by hertz generator.

package api

import (
	"bibi/api/biz/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bibiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _infoMw() []app.HandlerFunc {
	return nil
}

func _switch2faMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avatarMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avatar0Mw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login0Mw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.LoginHandler,
	}
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _register0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _access_tokenMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getaccesstokenMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtRefreshMiddleware.MiddlewareFunc(),
		//jwt.JwtMiddleware.RefreshToken,
	}
}
