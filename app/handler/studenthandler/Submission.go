package studenthandler

import (
	"context"
	"fmt"
	"hr/app/service"
	"hr/app/utils"
	"hr/configs/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const savePath = utils.Evidence

func Submission(c *gin.Context) {
	// 上传申报
	c.Header("Content-Type", "application/json")
	userID := c.Param("UserID")
	itemName := c.PostForm("ItemName")
	academicYear := c.PostForm("AcademicYear")
	msg := c.PostForm("Msg")
	data, err := c.MultipartForm()
	if err != nil {
		c.Error(utils.GetError(utils.PARAM_ERROR, err.Error()))
		c.Abort()
		return
	}
	files := data.File["Evidence"]
	fmt.Println(files[0].Filename)
	destPaths := make([]string, len(files))

	for i, file := range files {
		dst := savePath + "/" + file.Filename
		fmt.Println(dst)
		destPaths[i] = dst
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			c.Error(utils.GetError(utils.INNER_ERROR, err.Error()))
			c.Abort()
			return
		}
	}

	// 从上下文中获取用户信息
	currentUser := service.GetCurrentUser(c)
	if currentUser.UserID != userID {
		c.Error(utils.GetError(utils.UNAUTHORIZED, nil))
		c.Abort()
		return
	}
	newSubmission := models.SubmitInformation{
		CurrentUser:  currentUser,
		ItemName:     itemName,
		AcademicYear: academicYear,
		Msg:          msg,
		Evidence:     destPaths,
		Status:       false,
		CreateAt:     time.Now().Unix(),
	}
	insertResult := service.InsertOne(c, utils.MongodbName, utils.Submission, newSubmission)
	utils.ResponseSuccess(c, insertResult.InsertedID)
}

func GetSubmissionStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	userID := c.Query("UserID")

	filter := bson.M{
		"CurrentUser.UserID": userID,
	}

	result := service.Find(c, utils.MongodbName, utils.Submission, filter)
	var forms []models.SubmitInformation
	if err := result.All(context.Background(), &forms); err != nil {
		c.Error(utils.GetError(utils.DECODE_ERROR, err.Error()))
		c.Abort()
		return
	}
	utils.ResponseSuccess(c, forms)
}

// 这个东西要改
