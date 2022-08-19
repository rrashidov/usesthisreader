package scheduler

import (
	"rrashidov/usesthisreader/internal/generic"
)

type Scheduler interface {
	Schedule(exec generic.GenericLogic) error
}
