package application

import "testing/internal/entity"

func (a *applicationType) SetMonitor(mon entity.EventMonitor) error {
	a.monitor = mon
	return nil
}

func (a *applicationType) GetMonitor() entity.EventMonitor {
	return a.monitor
}
