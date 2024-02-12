package counsellorhandler

import (
	"hr/app/service"
	"hr/app/utils"
	"hr/configs/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
)

const savePath = utils.Information
const maxSizeLimit = 20

type information struct {
	ItemName     string `json:"itemName"`
	AcademicYear string `json:"academicYear"`
	CorrectGrade string `json:"correctGrade"`
}

// 改正成绩的
func CorrectGrade(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var information information
	userId := c.Param("userID")
	err := c.ShouldBindJSON(&information)
	if err != nil {
		c.Error(utils.GetError(utils.PARAM_ERROR, err.Error()))
		c.Abort()
		return
	}
	filter := bson.M{
		"_id":          userId,
		"academicYear": information.AcademicYear,
		"itemName":     information.ItemName,
	}
	modified := bson.M{
		"$set": bson.M{
			"grade": information.CorrectGrade,
		},
	}
	_ = service.UpdateOne(c, utils.MongodbName, utils.Score, filter, modified)
	utils.ResponseSuccess(c, nil)
}

func ImportCounsellor(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	err := c.Request.ParseMultipartForm(maxSizeLimit << 20) // 最大文件限制
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}

	// 获取上传的文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	defer file.Close()
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	// 遍历每个工作表和行
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	} // 表头
	for _, row := range rows[1:] { // 切片，去除表头
		userId := row[0]
		userName := row[1]
		grade := row[2]
		profession := row[3]
		user := models.Counsellor{
			UserId:     userId,
			UserName:   userName,
			Grade:      grade,
			Profession: profession,
		}
		_ = service.InsertOne(c, utils.MongodbName, utils.Counsellor, user)
	}
}

func ImportStudent(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	err := c.Request.ParseMultipartForm(maxSizeLimit << 20) // 最大文件限制
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}

	// 获取上传的文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	defer file.Close()
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	// 遍历每个工作表和行
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	for _, row := range rows[1:] { // 切片，去除表头
		userId := row[0]
		userName := row[1]
		grade := row[2]
		profession := row[3]
		class := row[4]
		user := models.Student{
			UserId:     userId,
			UserName:   userName,
			Class:      class,
			Profession: profession,
			Grade:      grade,
		}
		_ = service.InsertOne(c, utils.MongodbName, utils.Student, user)
	}
	utils.ResponseSuccess(c, nil)
}

func ImportMark(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	err := c.Request.ParseMultipartForm(maxSizeLimit << 20) // 最大文件限制
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}

	// 获取上传的文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	defer file.Close()
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	// 遍历每个工作表和行
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		c.Error(utils.GetError(utils.FILE_ERROR, err.Error()))
		c.Abort()
		return
	}
	item := rows[0]
	for _, row := range rows[1:] { // 切片，去除表头
		userId := row[0]
		academicYear := row[2]

		for colIndex, colValue := range row[6:] { //第六列后面就是成绩了
			itemName := item[colIndex]
			colvalue := colValue
			value, err := strconv.Atoi(colvalue)
			if err != nil {
				c.Error(utils.GetError(utils.PARAM_ERROR, err.Error()))
				c.Abort()
				return
			}
			sorce := models.Score{
				UserId:       userId,
				AcademicYear: academicYear,
				ItemName:     itemName,
				Mark:         int64(value),
			}
			// 新增成绩
			_ = service.InsertOne(c, utils.MongodbName, utils.Score, sorce)
		}
	}
}
