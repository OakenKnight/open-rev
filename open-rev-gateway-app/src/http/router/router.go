package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"open-rev.com/config"
	"open-rev.com/interactor"
)

func NewRouter(handler interactor.AppHandler, contract *client.Contract) *gin.Engine {
	router := gin.Default()
	cfg, err := config.GetConfig("./.env")
	if err != nil {
		panic(err)
	}
	defaultPath := cfg.DefaultPath
	router.GET(defaultPath+"/test-cc-version", handler.TestCCVersion)
	router.POST(defaultPath+"/user/register", handler.Register)
	router.POST(defaultPath+"/user/edit", handler.EditUser)
	router.POST(defaultPath+"/user/confirmAccount", handler.ConfirmAccount)

	router.POST(defaultPath+"/sciWork", handler.CreateScientificWork)

	router.POST(defaultPath+"/review", handler.CreateReview)
	router.POST(defaultPath+"/review/quality", handler.CreateReviewQuality)

	router.POST(defaultPath+"/subarea", handler.AddSubArea)
	router.POST(defaultPath+"/area", handler.AddArea)

	router.GET(defaultPath+"/user", handler.GetAllUsers)
	router.GET(defaultPath+"/user/:id", handler.GetUserById)
	router.GET(defaultPath+"/user/info/:id", handler.GetUserInfo)
	router.GET(defaultPath+"/user/:id/review", handler.GetAllReviewsByUser)
	router.GET(defaultPath+"/review/:id/reviewQuality", handler.GetAllReviewQualityByReview)

	router.GET(defaultPath+"/sciWork", handler.GetAllScientificWorks)
	router.GET(defaultPath+"/download/:bucket/:filename", handler.DownloadScientificWork)
	router.GET(defaultPath+"/sciWork/:id/review", handler.GetAllReviewsByScientificWork)
	router.GET(defaultPath+"/sciWork/:id/details", handler.GetScientificWorkDetails)
	router.GET(defaultPath+"/sciWork/:id", handler.GetScientificWorkById)
	router.GET(defaultPath+"/sciWork/userId/:userId", handler.GetAllScientificWorksByUser)
	router.GET(defaultPath+"/sciWork/subareaId/:subareaId", handler.GetAllScientificWorksBySubareaId)
	router.GET(defaultPath+"/review", handler.GetAllReviews)
	router.GET(defaultPath+"/reviewQuality", handler.GetAllReviewQualities)

	router.GET(defaultPath+"/topReviewers", handler.GetTopReviewers)
	router.GET(defaultPath+"/topAuthors", handler.GetTopReviewers)
	router.GET(defaultPath+"/topSciWorks", handler.GetTopScientificWorks)

	router.GET(defaultPath+"/subarea", handler.GetAllSubAreas)
	router.GET(defaultPath+"/area", handler.GetAllAreas)

	router.GET(defaultPath+"/sidebar", handler.GetAllAreasAndSubAreas)
	router.GET(defaultPath+"/dashboard", handler.GetDashboard)

	router.DELETE(defaultPath+"/user/:id", handler.DeleteOpenRevUser)
	router.DELETE(defaultPath+"/sciWork/:id", handler.DeleteScientificWork)
	router.DELETE(defaultPath+"/review/:id", handler.DeleteReview)
	router.DELETE(defaultPath+"/area/:id", handler.DeleteArea)
	router.DELETE(defaultPath+"/subarea/:id", handler.DeleteSubArea)

	return router
}
