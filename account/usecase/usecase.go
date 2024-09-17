package usecase

import (
	"sync"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/account/domain/permission"
)

type useCase[T account.Model] struct {
	logger log.Logger
	config config

	repo a.Repository[T]

	startMtx  *sync.RWMutex
	closeChan chan bool
	started   bool

	accountRole permission.AccountRole

	manager a.Manager[T]
}

func New[T account.Model](logger log.Logger, registry registry.Registry,
	repo a.Repository[T]) domain.UseCase[T] {
	uc := &useCase[T]{
		logger:   logger,
		repo:     repo,
		startMtx: new(sync.RWMutex),
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}

func (uc *useCase[T]) Initialize(accountRole permission.AccountRole) {
	uc.accountRole = accountRole
}

func (uc *useCase[T]) SetManager(manager a.Manager[T]) {
	uc.manager = manager
}

func (uc *useCase[T]) Start() {
	if uc.started {
		return
	}
	uc.closeChan = make(chan bool)
	uc.started = true
	go uc.startCron()
}

func (uc *useCase[T]) Stop() {
	if !uc.started {
		return
	}
	uc.started = false
	uc.startMtx.Lock()
	defer uc.startMtx.Unlock()
}
