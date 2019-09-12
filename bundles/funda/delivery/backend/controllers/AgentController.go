package controllers

import (
	"fmt"
	"github.com/fspace/ecm/bundles/funda/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type AgentController struct {
	interactor *usecases.ContactAgentInteractor
}

func NewAgentController(itr *usecases.ContactAgentInteractor) *AgentController {
	c := &AgentController{}
	c.interactor = itr
	return c
}
func (ctrl *AgentController) Contact(c *gin.Context) {
	// ContactAgentRequestModel
	// used for gin binding !
	requestModel := struct {
		CustomerEmailAddress string `form:"email" json:"email" xml:"email"  binding:"required"`
	}{}

	if err := c.ShouldBind(&requestModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 模型转存：
	req := usecases.ContactAgentRequestMessage{}
	copier.Copy(&req, &requestModel)

	resp := ctrl.interactor.Handle(req)
	fmt.Println("response is : ", resp)
	if resp.ValidationResult != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.ValidationResult})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": resp.ValidationResult})
}
