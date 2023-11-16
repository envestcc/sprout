package vm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/vm/server"
	"github.com/pkg/errors"
)

type Handler struct {
	endpoints   map[Type]string
	instanceMgr *server.Mgr
}

func (r *Handler) Handle(msg *msg.Msg, vmtype Type, code string, expParam string) ([]byte, error) {
	endpoint, ok := r.endpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(msg, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	slog.Debug(fmt.Sprintf("acquire %s instance success", vmtype))
	defer r.instanceMgr.Release(msg.ProjectID, ins)

	res, err := ins.Execute(context.Background(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(endpoints map[Type]string) *Handler {
	return &Handler{
		endpoints:   endpoints,
		instanceMgr: server.NewMgr(),
	}
}
