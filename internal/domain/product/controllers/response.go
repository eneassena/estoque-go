package controllers

type Response struct {
	Code       int         `json:"code"`
	CountItems int         `json:"count_items"`
	Data       interface{} `json:"data"`
}

type ProductRequestByID struct {
	ID int `json:"id" binding:"required"`
}

type ProductRequestUpdateCount struct {
	ID    int `json:"id" binding:"required"`
	Count int `json:"count" binding:"required"`
}

type ProductRequestCreate struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Count int     `json:"count" binding:"required"`
}
