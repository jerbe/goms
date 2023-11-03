package scripting

import (
	"time"

	"github.com/dop251/goja"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/21 17:46
  @describe :
*/

type EventManager struct {
	chId          int
	name          string
	scriptRuntime *goja.Runtime
}

func NewEventManager(chId int, scriptRuntime *goja.Runtime, name string) *EventManager {
	return &EventManager{chId: chId, scriptRuntime: scriptRuntime, name: name}
}

func (m *EventManager) Cancel() {
}

func (m *EventManager) Schedule(method string, delay time.Duration) {
	go func() {
		//	m.scriptRuntime.
	}()
}

func (m *EventManager) ScheduleWithInstance(method string, delay time.Duration, instance *EventInstanceManager) {

}

func (m *EventManager) ScheduleAt(method string, at time.Time) {

}
