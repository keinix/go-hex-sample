package handler

import (
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/domain/login"
)

type LoginHandler interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type loginHandler struct {
	service login.Service
}

func NewLoginHandler(s login.Service) LoginHandler {
	return &loginHandler{service: s}
}

func (h *loginHandler) Login(c *gin.Context) {

}

func (h *loginHandler) SignUp(c *gin.Context) {

}
