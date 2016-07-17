package safari

import (
	"context"
	"fmt"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

type game struct {
	tix kvc.KVC
	pdx pbf.PokedexClient
}

func newGame() *game {
	return &game{
		tix: kvc.NewMem(),
	}
}

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
			ttl := time.Duration(tkt.Expires.Time - time.Now().Unix())
			if ttl < 1 {
				return false
			}
			go func() {
				time.Sleep(ttl * time.Second)
				g.tix.Set(k, nil)
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

func (g *game) spawn(zone *pbf.Zone) (*pbf.Pokemon, error) {
	ctx := context.Background()
	num := int32(util.RandRng(0, 150))
	pc, err := g.pdx.GetPokemon(ctx, &pbf.Pokemon{Number: num})
	if err != nil {
		return nil, err
	}
	poke := pc.Pokemon[0]
	poke.Uid = util.GenUID()
	poke.NickName = util.RandName()
	return poke, nil
}
