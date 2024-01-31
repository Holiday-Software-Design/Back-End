package models

type Student struct {
	UserId     int64          `bson:"userId"`
	UserName   string         `bson:"userName"`
	passWord   string         `bson:"-"`
	Class      string         `bson:"class"`
	Profession string         `bson:"profession"`
	Grade      int            `bson:"grade"` // 这里是年级
	Mark       map[string]int `bson:"mark"`
	// FeedbackId   []string       `bson:"feedbackId"`  //申诉
	// AdviceId     []string       `bson:"recommemdId"` // 建议
	// SubmissionId []string       `bson:"formId"`      // 申报表
}

type CurrentUser struct {
	UserId     string
	UserName   string
	Grade      int //这里是年级不是成绩
	Role       string
	Profession string
}

// 有一个值得思考的问题，既然能在submission库中直接通过id找到该用户，那么为什么要增加这几个没有的字段
