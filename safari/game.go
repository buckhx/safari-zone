package safari

import (
	"fmt"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

type game struct {
	tix kvc.KVC
}

func newGame() *game {
	return &game{
		tix: kvc.NewMem(),
	}
}

/*
message Ticket {
	string uid = 1;
	registry.Trainer trainer = 2;
	Zone zone = 3;
	msg.Timestamp time = 4;
	int64 expiration = 5;
}
*/

func (g *game) issueTicket(trainer *pbf.Trainer, zone *pbf.Zone, expiry int64) (*pbf.Ticket, error) {
	k := trainer.Uid
	tkt := &pbf.Ticket{
		Uid:     util.GenUID(),
		Trainer: trainer,
		Zone:    zone,
		Time:    &pbf.Timestamp{Unix: time.Now().Unix()},
		Expiry:  expiry,
	}
	ok := g.tix.CompareAndSet(k, tkt, func() bool {
		return !g.tix.(*kvc.MemKVC).UnsafeHas(k)
	})
	if !ok {
		return nil, fmt.Errorf("Ticket already issued for trainer %s", trainer.Uid)
	}
	return tkt, nil
}
