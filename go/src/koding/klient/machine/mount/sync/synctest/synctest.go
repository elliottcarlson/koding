package synctest

import (
	"context"
	"fmt"

	"koding/klient/machine/index"
	"koding/klient/machine/index/indextest"
	msync "koding/klient/machine/mount/sync"
)

// SyncLocal sends all index changes between local and remote directories to
// provided syncer. Then it calls produced Execers. Returned contex will be
// closed when all changes are consumed and their events executed. The direction
// of changes is marked by dir argument and it must be ChangeMetaRemote and/or
// ChangeMetaRemote
func SyncLocal(rootA, rootB string, dir index.ChangeMeta, s msync.Syncer) (context.Context, context.CancelFunc, error) {
	if dir&^(index.ChangeMetaLocal|index.ChangeMetaRemote) != 0 {
		return nil, nil, fmt.Errorf("invalid change type: %v", dir)
	}

	cs, err := indextest.Compare(rootA, rootB)
	if err != nil {
		return nil, nil, err
	}

	var evs []*msync.Event
	for _, c := range cs {
		c = index.NewChange(c.Path(), c.Meta()|dir)
		evs = append(evs, msync.NewEvent(context.Background(), nil, c))
	}

	var (
		evC = make(chan *msync.Event, 1)
		exC = s.ExecStream(evC)
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		for i := 0; i < len(evs); {
			select {
			case evC <- evs[i]:
			case ex := <-exC:
				ex.Exec()
				i++
			case <-ctx.Done():
				return
			}
		}
	}()

	return ctx, cancel, nil
}
