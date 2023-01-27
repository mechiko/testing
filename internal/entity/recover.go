package entity

type RecoverInterface interface {
	RecoverLog(str string)
	RecoverLogWithStack(str string)
	RecoverFunc(str string)
	RecoverExit(str string)
}
