package dto

type HeightWeightInput struct {
	Height *float32 `form:"height" binding:"required,numeric"`
	Weight *float32 `form:"weight" binding:"required,numeric"`
}
