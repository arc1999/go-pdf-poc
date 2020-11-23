package model

type Sequence struct {
	ID  string `json:"id,omitempty" bson:"_id,omitempty"`
	Seq *int64  `json:"seq" bson:"seq"`
}
