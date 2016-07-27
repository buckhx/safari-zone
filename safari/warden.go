package safari

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

type warden struct {
	tix    kvc.KVC
	pdx    pbf.PokedexClient
	ctx    context.Context
	access string
}

func newWarden(opts Opts) (*warden, error) {
	pdx, err := opts.PokedexClient()
	if err != nil {
		return nil, err
	}
	return &warden{
		tix: kvc.NewMem(),
		pdx: pdx,
		ctx: context.Background(),
	}, nil
}

func (w *warden) issueTicket(trainer *pbf.Trainer, zone *pbf.Zone, expiry *pbf.Ticket_Expiry) (*pbf.Ticket, error) {
	k := trainer.Uid
	tkt := &pbf.Ticket{
		Uid:     util.GenUID(),
		Trainer: trainer,
		Zone:    zone,
		Time:    &pbf.Timestamp{Unix: time.Now().Unix()},
		Expires: expiry,
	}
	if !w.tix.CompareAndSet(k, tkt, func() bool {
		if !w.tix.(*kvc.MemKVC).UnsafeHas(k) {
			ttl := time.Duration(tkt.Expires.Time - time.Now().Unix())
			if ttl < 1 {
				return false
			}
			go func() {
				time.Sleep(ttl * time.Second)
				w.tix.Set(k, nil)
			}()
			return true
		}
		return false
	}) {
		return nil, fmt.Errorf("Error issuing ticket for trainer %s", trainer.Uid)
	}
	return tkt, nil
}

func (w *warden) spawn(ctx context.Context) (*pbf.Pokemon, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Ticket can't be identified")
	}
	k := string(claims.Subject)
	var tkt *pbf.Ticket
	w.tix.GetAndSet(k, func(cur kvc.Value) kvc.Value {
		tkt, ok = cur.(*pbf.Ticket)
		if !ok || tkt.Expires.Encounters <= 1 {
			return nil
		}
		tkt.Expires.Encounters -= 1
		return tkt
	})
	if tkt == nil {
		return nil, fmt.Errorf("Ticket has expired")
	}
	num := util.RandRng(0, 150) //TODO set on tkt zone
	poke, err := w.fetchPokemon(num)
	if err == nil {
		poke.Uid = util.GenUID()
		poke.NickName = util.RandName()
	}
	return poke, err
}

func (w *warden) fetchPokemon(num int) (*pbf.Pokemon, error) {
	pc, err := w.pdx.GetPokemon(w.ctx, &pbf.Pokemon{Number: int32(num)})
	if err != nil {
		return nil, err
	}
	return pc.Pokemon[0], nil

}
