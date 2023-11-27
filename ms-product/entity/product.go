package entity

type Products struct {
	Id          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Category    string `json:"category,omitempty" bson:"category,omitempty"`
	Price       int64  `json:"price,omitempty" bson:"price,omitempty"`
	Stock       int32  `json:"stock,omitempty" bson:"stock,omitempty"`
	CreatedAt   string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
