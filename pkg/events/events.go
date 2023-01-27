package events

import (
	"fmt"

	"testing/internal/entity"
	"testing/pkg/config"

	"github.com/rs/zerolog"
)

type EvtMonitor struct {
	Db          entity.DbService
	Logger      *zerolog.Logger
	Recovery    entity.RecoverInterface
	Config      config.Config
	subscribers MapSubscriber
	History     entity.History
}

type ArraySubs []entity.Subcriber
type MapSubscriber map[entity.EventInt]ArraySubs

var monitor EvtMonitor

func NewMonitor(a entity.App) entity.EventMonitor {
	monitor = EvtMonitor{
		Db:       a.GetDbService(),
		Logger:   a.GetLogger().Logger,
		Recovery: a.GetRecovery(),
		Config:   *a.GetConfig(),
		History:  a.GetHistory(),
	}
	monitor.subscribers = make(MapSubscriber)
	return &monitor
}

func (s *EvtMonitor) Attach(o entity.Subcriber, evt entity.EventInt) (bool, error) {
	defer s.Recovery.RecoverLog("events:attache")
	if _, ok := s.subscribers[evt]; !ok {
		s.subscribers[evt] = make(ArraySubs, 0)
	}
	subs, ok := s.subscribers[evt]
	if !ok {
		return false, fmt.Errorf("EvtMonitor Attach() evt = %v not register subscriber=%v", evt, o.NameSubcriber())
	}
	for _, observer := range subs {
		if observer == o {
			return false, fmt.Errorf("EvtMonitor Attach() subscriber=%v exists", o.NameSubcriber())
		}
	}
	subs = append(subs, o)
	s.subscribers[evt] = subs
	return true, nil
}

func (s *EvtMonitor) Detach(o entity.Subcriber, evt entity.EventInt) (bool, error) {
	defer s.Recovery.RecoverLog("events:detach")
	subs, ok := s.subscribers[evt]
	if !ok {
		return false, fmt.Errorf("EvtMonitor Detach() evt = %v not register subscriber=%v", evt, o.NameSubcriber())
	}
	for i, observer := range subs {
		if observer == o {
			subs = append(subs[:i], subs[i+1:]...)
			s.subscribers[evt] = subs
			return true, nil
		}
	}
	return false, fmt.Errorf("EvtMonitor Detach() evt = %v subscriber=%v not found", evt, o.NameSubcriber())
}

func (s *EvtMonitor) Notify(evt entity.EventInt, task string) (bool, error) {
	defer s.Recovery.RecoverLogWithStack("events:Notify")
	switch evt {
	case entity.NEED_UPDATE_STATUS_BAR:
	case entity.NEED_TICKS_EVERY_SECOND:
	default:
		s.History.Push(fmt.Sprintf("%v [%v]", task, evt.String()))
	}

	// fmt.Println("EvtMonitor new Event : " + evt.String())
	alls := s.subscribers[entity.ALL]
	subs := s.subscribers[evt]
	// чтобы дважды не оповещать объединяем списки
	for _, observer := range alls {
		if !isPresent(observer, subs) {
			subs = append(subs, observer)
		}
	}
	for _, observer := range subs {
		// fmt.Println("EvtMonitor Notify : " + observer.NameSubcriber() + " evt:" + evt.String())
		observer.UpdateSubcriber(task, evt)
	}
	return true, nil
}

func isPresent(subscriber entity.Subcriber, arr ArraySubs) bool {
	for _, sub := range arr {
		if sub == subscriber {
			return true
		}
	}
	return false
}
