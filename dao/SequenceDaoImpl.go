package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-pdf-poc/db"
	"go-pdf-poc/model"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

type SequenceDaoImpl struct {

}

func (i SequenceDaoImpl) GetSequence() (*int64,error) {
	db := db.GetDb()
	var seq model.Sequence
	update := bson.D{
		{"$inc", bson.D{
			{"seq", 1},
		}},
	}
	//id,err := primitive.ObjectIDFromHex("5f8ec7b0b8c20c274a085927")

	err := db.Collection(os.Getenv("SEQUENCE_MONGODB_COLLECTION")).FindOneAndUpdate(context.TODO(), bson.M{"_id": "medicineSequence"}, update).Decode(&seq)
	if err != nil {
		log.Printf("error in saving Medicine")
		log.Println(err)
		return nil, err
	}
	return seq.Seq,nil
}