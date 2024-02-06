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

// 有一个值得思考的问题，既然能在submission库中直接通过id找到该用户，那么为什么要增加这几个没有的字段
// 这个表平时改动会比较小，需不需要维护一个索引？
