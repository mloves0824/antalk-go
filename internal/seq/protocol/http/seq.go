package http

import (
	"antalk-go/internal/common"
	"antalk-go/internal/seq/controller"
	"context"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine     *gin.Engine
	controller *controller.SeqController
}

type resp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

const (
	OK = 0
	RequestErr = -400
	ServerErr = -500
)


func New(c *common.Config) *Server {
	engine := gin.New()

	go func() {
		if err := engine.Run(c.GetString("httpServer.addr")); err != nil {
			panic(err)
		}
	}()

	s := &Server{
		engine:     engine,
		controller: nil,
	}
	s.initRouter()
	return s
}

func (s *Server) initRouter() {
	group := s.engine.Group("/antalk")
	group.POST("/seq/get", s.seqGet)
}

func (s *Server) seqGet(c *gin.Context) {
	var arg struct {
		Uid string
	}

	if err := c.BindQuery(&arg); err != nil {
		c.Set("context/err/code", RequestErr)
		c.JSON(200, resp{
			Code:    RequestErr,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	getSeqReq := &controller.GetSeqReq{}
	getSeqReq.Uid = arg.Uid
	getSeqResp, err := s.controller.GetSeq(context.Background(), getSeqReq)
	if err != nil {
		c.Set("context/err/code", ServerErr)
		c.JSON(200, resp{
			Code:    ServerErr,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, resp {
		Code: OK,
		Message: "ok",
		Data: getSeqResp.SeqID,
	})
	return
}

func (s *Server) Close() {

}
