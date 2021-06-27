package requests

import "go.mongodb.org/mongo-driver/bson/primitive"

var EmptyCreateAccount = AccountCreateRequest{}

type AccountCreateRequest struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name    string             `json:"name" binding:"required,min=2,max=100"`
	Email   string             `json:"email" binding:"required,email"`
	Zipcode int32              `json:"zip_code" binding:"required,min=2,max=10"`
	Dept    Department         `json:"dept" binding:"required"`
}

type Department struct {
	DeptId   primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	DeptName string             `json:"DeptName" binding:"required,min=2,max=100"`
	DeptType string             `json:"DeptType" binding:"required,min=2,max=100"`
}

type AccountUpdateRequest struct {
	Name    string     `json:"name" binding:"required,min=2,max=100"`
	Zipcode int32      `json:"zip_code" binding:"required,min=2,max=10"`
	Dept    Department `json:"dept" binding:"required"`
}
