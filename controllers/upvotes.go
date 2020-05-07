package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/repositories"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UpVotesController interface {
	GetUpVotesCount(ctx *gin.Context)
}

type upVoteController struct {
	upVotesRepo repositories.UpVotesRepo
}

func NewUpVotesController(upVotesRepo repositories.UpVotesRepo) UpVotesController {
	return &upVoteController{
		upVotesRepo: upVotesRepo,
	}
}

func (ctrl upVoteController) GetUpVotesCount(ctx *gin.Context) {
	var voteRequest models.GetUpVoteRequest
	err := ctx.ShouldBindBodyWith(&voteRequest, binding.JSON)
	if err != nil {
		log.Error("unable to get vote due to ", err)
		ctx.Abort()
		return
	}

	votes, err := ctrl.upVotesRepo.GetUpVoteCounts(voteRequest.SiteId, voteRequest.PageId)
	if err != nil {
		log.Printf("unable to get vote : %+v", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, votes)
}
