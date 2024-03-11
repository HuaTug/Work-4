package router

import (
	"4-TikTok/IM"
	"4-TikTok/configs"
	"4-TikTok/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	configs.Init()
	N1 := r.Group("user")
	{
		N1.POST("register", handler.UserRegister)
		N1.POST("login", handler.UserLogin)
		N1.GET("info", handler.UserInfo)
		N1.PUT("avatar/upload", handler.UpLoad)
	}
	N2 := r.Group("video")
	{
		N2.GET("feed/", handler.GetVideoFeed)
		N2.POST("publish", handler.Publish)
		N2.GET("list", handler.VideoFeedList)
		N2.GET("popular")
		N2.POST("search", handler.VideoSearch)
	}
	N3 := r.Group("like")
	{
		N3.POST("action", handler.Favorite)
		N3.GET("list", handler.FavorList)

	}
	N4 := r.Group("comment")
	{
		N4.POST("publish", handler.CommentCreate)
		N4.GET("list", handler.CommentList)
		N4.DELETE("delete", handler.CommentDelete)
	}
	N5 := r.Group("follow")
	{
		N5.POST("relation/action", handler.Following)
		N5.GET("following/list", handler.List)
		N5.GET("follower/list", handler.List2)
	}
	r.GET("ws", IM.WsHandler)
	return r
}
