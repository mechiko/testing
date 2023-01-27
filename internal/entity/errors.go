package entity

import (
	"errors"
	"fmt"
	"time"
)

// для таких ошибок работает метод errors.Is(err, ErrRuleIsTimeOut)
var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given param is not valid")
	// мои придуманные ошибки
	ErrNotPingUtm    = errors.New("utm host:port not ping")
	ErrRuleIsTimeOut = errors.New("правило ждет таймаут")
)

//	для таких ошибок надо if perr, ok := err.(*RuleIsTimeOutError); ok {
//		TimeOut  int     для правила установлен таймаут в секундах
//		TimeLeft int     для правила осталось секунд ожидания
//		Msg      string  сообщение ошибки
type RuleIsTimeOutError struct {
	TimeOut  int
	TimeWait time.Duration
}

func (e *RuleIsTimeOutError) Error() string {
	return fmt.Sprintf("правило ждет тайм аут %v осталось %s", e.TimeOut, e.TimeWait)
}

type UTMOutError struct {
	Err error
	Doc string
}

func (e *UTMOutError) Error() string {
	return fmt.Sprintf("ошибка УТМ  %v получено: '%v'", e.Err.Error(), e.Doc)
}
