package slice

import (
	"context"
	"errors"

	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/infra/bus"
)

type CommandExecutionResult struct {
	CtxCancFunc     context.CancelFunc
	EventBus        *bus.EventBus
	EventResultChan chan bus.EventResult
}

func (cEr *CommandExecutionResult) Execute(domainEvent []domain.Event) (err error) {

	evPublisher := bus.NewEventPublisher(cEr.EventBus)
	evPublisher.Publish(domainEvent)

	eventResult, resultChanOk := <-cEr.EventResultChan

	if !resultChanOk {
		err = errors.New("result channel closed")
		return
	}

	if eventResult.Err != nil {
		err = eventResult.Err
		cEr.CtxCancFunc()
		return
	}

	return

}
