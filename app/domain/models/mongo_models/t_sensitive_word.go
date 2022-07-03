package mongo_models

import "time"

type TSensitiveWord struct {
	WordId        string    `json:"word_id" bson:"word_id"`               // 敏感词id
	ReviewContent string    `json:"review_content" bson:"review_content"` // 敏感内容
	Label         string    `json:"label" bson:"label"`                   // 检测分类
	Rate          float64   `json:"rate" bson:"rate"`                     // 相似度
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`         // 创建时间
}

func (t *TSensitiveWord) TableName() string {
	return "t_sensitive_word"
}
