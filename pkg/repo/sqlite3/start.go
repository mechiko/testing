package sqlite3

import (
	"fmt"

	"testing/pkg/application"
)

func (r *Repository) Start() error {
	defer r.Recovery.RecoverLog("Start()")
	dbService := application.Get().GetDbService()
	// создаем структуру БД если файл БД был только что создан
	if !dbService.IsCreated() {
		if err := r.create(); err != nil {
			return fmt.Errorf("r.create() %w", err)
		}
		dbService.SetCreated()
	}

	// rules := rules.New()
	// if err := r.InsertRules(); err != nil {
	// 	return fmt.Errorf("InsertRules() %w", err)
	// }
	// rl, err := rules.Get()
	// if err != nil {
	// 	// application.Get().ErrorLog().AnErr("repo.Start()", err).Send()
	// 	return fmt.Errorf("rules.Get() %w", err)
	// }
	// if len(rl.Items) == 0 {
	// 	if err := r.InsertRules(); err != nil {
	// 		return fmt.Errorf("InsertRules() %w", err)
	// 	}
	// }
	// if err := r.InsertStates(); err != nil {
	// 	return fmt.Errorf("InsertStates() %w", err)
	// }
	return nil
}
