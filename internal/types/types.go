// Code generated by goctl. DO NOT EDIT.
package types

type SimpleRequest struct {}

type Request struct {
	User_id string `path:"userId,options=you|me"` // parameters are auto validated
}

type Response struct {
	Message string `json:"message"`
}
