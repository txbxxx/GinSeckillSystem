/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:31
 * @File:  Cors
 * @Software: GoLand
 **/

package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Cors() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	// 允许跨域请的头
	corsConf.AddAllowHeaders("Content-Type", "Authorization", "Token")
	// 允许跨域请求的响应头
	corsConf.AddExposeHeaders("Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type")
	// 允许跨域请求的域名
	corsConf.AllowOrigins = []string{"*"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		corsConf.AllowOrigins = []string{"http://www.example.com"}
	} else {
		// 测试环境下模糊匹配本地开头的请求，如果是本地则直接允许跨域
		corsConf.AllowOriginFunc = func(origin string) bool {
			//如果是本地开头的请求，则允许跨域
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	// 用户允许跨域携带凭证比如:cookie,token,SSL等
	corsConf.AllowCredentials = true
	return cors.New(corsConf)
}
