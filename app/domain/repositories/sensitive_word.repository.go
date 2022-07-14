package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"censor.services/app/domain/models/mongo_models"
	"censor.services/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type SensitiveWordRepository interface {
	Create(c context.Context, sensitiveWord *mongo_models.TSensitiveWord) (string, error)
	FindByRate(c context.Context, page int, pageSize int, rate float64) ([]*mongo_models.TSensitiveWord, error)
}

type sensitiveWordRepository struct {
	mongoDb *mongo.Database
}

func (r sensitiveWordRepository) FindByRate(c context.Context, page int, pageSize int, rate float64) ([]*mongo_models.TSensitiveWord, error) {
	word := mongo_models.TSensitiveWord{}
	list := make([]*mongo_models.TSensitiveWord, 0)

	filter := bson.M{"rate": bson.M{"$gte": rate}}
	limit := int64(pageSize)
	skip := int64((page - 1) * pageSize)
	findOpt := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	cur, err := r.mongoDb.Collection(word.TableName()).Find(c, filter, findOpt)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// todo 这里如果量大考虑换成next遍历， 一个个添加到 ac自动机 中
	err = cur.All(c, &list)
	if err != nil {
		return nil, err
	}
	cur.Close(c)
	return list, nil
}

func (r sensitiveWordRepository) Create(c context.Context, sensitiveWord *mongo_models.TSensitiveWord) (string, error) {
	sensitiveWord.WordId, _ = utils.GenerateEntityID(sensitiveWord.WordId)
	_, err := r.mongoDb.Collection(sensitiveWord.TableName()).InsertOne(c, sensitiveWord)
	return sensitiveWord.WordId, err
}

func NewSensitiveWordRepository(mongoDb *mongo.Database) SensitiveWordRepository {
	return &sensitiveWordRepository{
		mongoDb: mongoDb,
	}
}
