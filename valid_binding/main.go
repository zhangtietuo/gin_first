package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	Birthday time.Time `form:"birth_day, customValid" binding:"required" time_format:"1993-01-30"`
	Currentday time.Time `form:"current_day" binding:"required,gtfield=Birthday" time_format:"2019-09-03"`
}

/**
自定义
 */
func customValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok:=field.Interface().(time.Time); ok {
		today:=time.Now()
		if date.Unix()>today.Unix() {
			return true
		}
	}

	return false
}


func main() {
	r:=gin.Default()
	if v, ok:=binding.Validator.Engine().(*validator.Validate);ok {
		v.RegisterValidation("customValid", customValid)
	}
	r.GET("/person", func(ctx *gin.Context) {
		var person Person
		if err:=ctx.ShouldBind(&person); err!=nil {
			ctx.String(500, "%v", err)
			ctx.Abort()
			return
		}
		ctx.String(200, "%v", person)
	})
	r.Run()
}