package repositories

import (
	"censor.services/app/domain/models/mongo_models"
	"censor.services/pkg/utils"
	"context"
	"github.com/henrion-y/base.services/database/mongo"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func getRepositories(t *testing.T) SensitiveWordRepository {
	v := viper.New()
	v.Set("mongo.Hosts", "")
	v.Set("mongo.DB", "")
	db, err := mongo.NewDbProvider(v)
	if err != nil {
		t.Fatal(err)
	}
	return NewSensitiveWordRepository(db)
}

var words = []string{}

func TestSensitiveWordRepository_Create(t *testing.T) {
	repository := getRepositories(t)
	for i := range words {
		sensitiveWord := &mongo_models.TSensitiveWord{
			WordId:        "",
			ReviewContent: words[i],
			Label:         "自定义敏感词-xx",
			Rate:          1,
			CreatedAt:     time.Now(),
		}
		sensitiveWord.WordId, _ = utils.GenerateEntityID(sensitiveWord.WordId)

		_, err := repository.Create(context.Background(), sensitiveWord)
		if err != nil {
			t.Fatal(err)
		}
	}
}
