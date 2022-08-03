package api

import "github.com/go-playground/validator/v10"

//这是一个函数变量，接受一个FieldLevel 作为输入，返回bool变量
// FieldLevel是一个包含了所有用来帮助该函数验证字段的信息的类型
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	// Field函数返回字段的值，interface将值转换为Interface{}(可以理解为java中的Object类），
	// 最后使用go自带的类型转换，类型转换会返回转换的值和是否转换成功
	currency, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return IsSupportCurrency(currency)

}

// 后期开发可以扩展
const (
	USD = "USD"
	CAD = "CAD"
)

func IsSupportCurrency(currency string) bool {
	//balabala

	return true
}
