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

func (g *game) issueTicket(trainer *pbf.Trainer, zone *pbf.Zone, expiry *pbf.Ticket_Expiry) (*pbf.Ticket, error) {
	k := trainer.Uid
	tkt := &pbf.Ticket{
		Uid:     util.GenUID(),
		Trainer: trainer,
		Zone:    zone,
		Time:    &pbf.Timestamp{Unix: time.Now().Unix()},
		Expires: expiry,
	}
	ok := g.tix.CompareAndSet(k, tkt, func() bool {
		if !g.tix.(*kvc.MemKVC).UnsafeHas(k) {
			ttl := (tkt.Expires.Time - time.Now().Unix())
			if ttl < 1 {
				return false
			}
			go func() {
				time.Sleep(ttl * time.Second)
				c.Set(k, nil)
			}()
			return true
		}
		return false
	})
	if !ok {
		return nil, fmt.Errorf("Error issuing ticket for trainer %s", trainer.Uid)
	}
	return tkt, nil
}
