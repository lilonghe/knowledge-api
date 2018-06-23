package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lilonghe.net/knowledge/config"
)

type KnowledgePremise struct {
	//gorm.Model

	Id                 int `json:"id"`
	PremiseKnowledgeId int `json:"premise_knowledge_id" sql:"NOT NULL"`
	KnowledgeId        int `json:"knowledge_id" sql:"NOT NULL"`

	CreatedAt  time.Time `json:"created_at"  sql:"DEFAULT:now()"`
	ModifiedAt time.Time `json:"modified_at"  sql:"DEFAULT:now() ON UPDATE CURRENT_TIMESTAMP"`
}

func (KnowledgePremise) TableName() string {
	return "knowledge_premise"
}

func AddKnowledgePremise(data KnowledgePremise) (error, int) {
	err := config.Store.Master().Create(&data).Error
	return err, data.Id
}

func GetAllKnowledgePremise() (error, []KnowledgePremise) {
	data := make([]KnowledgePremise, 0)
	err := config.Store.Master().Find(&data).Error
	return err, data
}
