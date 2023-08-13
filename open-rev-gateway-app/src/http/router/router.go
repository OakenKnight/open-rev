package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"open-rev.com/interactor"
)

func NewRouter(handler interactor.AppHandler, contract *client.Contract) *gin.Engine {
	router := gin.Default()
	router.GET("/test-cc-version", handler.TestCCVersion)
	router.POST("/user/register", handler.Register)
	router.POST("/user/edit", handler.EditUser)
	router.POST("/user/confirmAccount", handler.ConfirmAccount)

	router.POST("/sciWork", handler.CreateScientificWork)

	router.POST("/review", handler.CreateReview)
	router.POST("/review/quality", handler.CreateReviewQuality)

	router.POST("/subarea", handler.AddSubArea)
	router.POST("/area", handler.AddArea)
	router.POST("/review/fix/:id", handler.FixReviewId)
	router.POST("/reviewQuality/fix/:id", handler.FixReviewQualityId)

	router.GET("/user", handler.GetAllUsers)
	router.GET("/user/:id", handler.GetUserById)
	router.GET("/user/info/:id", handler.GetUserInfo)
	router.GET("/user/:id/review", handler.GetAllReviewsByUser)
	router.GET("/review/:id/reviewQuality", handler.GetAllReviewQualityByReview)

	router.GET("/sciWork", handler.GetAllScientificWorks)
	router.GET("/download/:bucket/:filename", handler.DownloadScientificWork)
	router.GET("/sciWork/:id/review", handler.GetAllReviewsByScientificWork)
	router.GET("/sciWork/:id/details", handler.GetScientificWorkDetails)
	router.GET("/sciWork/:id", handler.GetScientificWorkById)
	router.GET("/sciWork/userId/:userId", handler.GetAllScientificWorksByUser)
	router.GET("/sciWork/subareaId/:subareaId", handler.GetAllScientificWorksBySubareaId)
	router.GET("/review", handler.GetAllReviews)
	router.GET("/reviewQuality", handler.GetAllReviewQualities)

	router.GET("/topReviewers", handler.GetTopReviewers)
	router.GET("/topAuthors", handler.GetTopReviewers)
	router.GET("/topSciWorks", handler.GetTopScientificWorks)

	router.GET("/subarea", handler.GetAllSubAreas)
	router.GET("/area", handler.GetAllAreas)

	router.GET("/sidebar", handler.GetAllAreasAndSubAreas)
	router.GET("/dashboard", handler.GetDashboard)

	router.DELETE("/user/:id", handler.DeleteOpenRevUser)
	router.DELETE("/sciWork/:id", handler.DeleteScientificWork)
	router.DELETE("/review/:id", handler.DeleteReview)
	router.DELETE("/area/:id", handler.DeleteArea)
	router.DELETE("/subarea/:id", handler.DeleteSubArea)

	return router
}
