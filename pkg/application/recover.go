package application

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"testing/internal/entity"

	"github.com/rs/zerolog"
)

type recoverInterface struct {
	logger zerolog.Logger
}

var _ entity.RecoverInterface = &recoverInterface{}

func NewRecoveryInterface(l zerolog.Logger) entity.RecoverInterface {
	return &recoverInterface{
		logger: l,
	}
}

func (ri *recoverInterface) RecoverLog(str string) {
	if r := recover(); r != nil {
		// stack := make([]byte, 8096)
		// stack = stack[:runtime.Stack(stack, false)]
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		ri.logger.Error().AnErr("RecoverLog", err).Send()
		// ri.logger.Error().Msgf("%s", string(stack))
	}
}

func (ri *recoverInterface) RecoverLogWithStack(str string) {
	if r := recover(); r != nil {
		stack := make([]byte, 8096)
		stack = stack[:runtime.Stack(stack, false)]
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		ri.logger.Error().AnErr("RecoverLogWithStack", err).Send()
		ri.logger.Error().Msgf("%s", string(stack))
	}
}

func (ri *recoverInterface) RecoverFunc(str string) {
	// applicationInstance.Logger.Logger.Debug().Msg("RecoverFunc")
	if r := recover(); r != nil {
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		_ = os.WriteFile("error.txt", []byte(err.Error()), 0644)
		ri.logger.Error().AnErr("RecoverFunc", err).Send()
	}
}

func (ri *recoverInterface) RecoverExit(str string) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%v %v %v", str, time.Now(), r)
		_ = os.WriteFile("error.txt", []byte(err.Error()), 0644)
		ri.logger.Error().AnErr("RecoverExit", err).Send()
		os.Exit(1)
	}
}

// func (ri *recoverInterface) RecoverFuncGo(str string) {
// 	if r := recover(); r != nil {
// 		err := fmt.Sprintf("%v %v %v", str, time.Now(), r)
// 		d := []byte(err)
// 		_ = ioutil.WriteFile("error.txt", d, 0644)
// 		applicationInstance.ErrorLog().Str("Recovery", err).Send()
// 		time.Sleep(5 * time.Second)
// 	}
// }
