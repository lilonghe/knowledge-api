package models

import (
	"time"

	"strings"

	"errors"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lilonghe.net/knowledge/config"
)

type Knowledge struct {
	//gorm.Model

	Id          int    `json:"id"`
	Title       string `json:"title" sql:"NOT NULL"`
	Description string `json:"description"`
	Importance  int    `json:"importance" sql:"DEFAULT:1 NOT NULL"`
	Level       int    `json:"level" sql:"DEFAULT:1 NOT NULL "`

	CreatedAt  time.Time `json:"created_at"  sql:"DEFAULT:CURRENT_TIMESTAMP NOT NULL"`
	ModifiedAt time.Time `json:"modified_at" sql:"DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL"`

	KnowledgePremise []int `json:"knowledge_premise" sql:"-"`
}

func (Knowledge) TableName() string {
	return "knowledge"
}

func GetKnowledgeList() (error, []Knowledge) {
	datas := make([]Knowledge, 0)
	err := config.Store.Master().Find(&datas).Error
	return err, datas
}

func AddKnowledge(item Knowledge) (error, int) {
	if strings.TrimSpace(item.Title) == "" {
		return errors.New("缺少参数"), 0
	}
	err := config.Store.Master().Create(&item).Error
	if err == nil {
		for _, v := range item.KnowledgePremise {
			err2, _ := AddKnowledgePremise(KnowledgePremise{KnowledgeId: item.Id, PremiseKnowledgeId: v})
			if err2 != nil {
				return err2, item.Id
			}
		}
	}
	return err, item.Id
}
