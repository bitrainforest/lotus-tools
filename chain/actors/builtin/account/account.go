package account

import (
	"github.com/filecoin-project/lotus/chain/actors"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtin8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"

)

func init() {

	builtin.RegisterActorState(builtin0.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version0, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load0(store, root)
		})
	}

	builtin.RegisterActorState(builtin2.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version2, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load2(store, root)
		})
	}

	builtin.RegisterActorState(builtin3.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version3, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load3(store, root)
		})
	}

	builtin.RegisterActorState(builtin4.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load4(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version4, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load4(store, root)
		})
	}

	builtin.RegisterActorState(builtin5.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load5(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version5, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load5(store, root)
		})
	}

	builtin.RegisterActorState(builtin6.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load6(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version6, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load6(store, root)
		})
	}

	builtin.RegisterActorState(builtin7.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load7(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version7, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load7(store, root)
		})
	}

	builtin.RegisterActorState(builtin8.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load8(store, root)
	})

	if c, ok := actors.GetActorCodeID(actors.Version8, "account"); ok {
    	builtin.RegisterActorState(c, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
			return load8(store, root)
		})
	}
}

var Methods = builtin4.MethodsAccount

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := actors.GetActorMetaByCode(act.Code); ok {
       if name != "account" {
          return nil, xerrors.Errorf("actor code is not account: %s", name)
       }

       switch av {
            
			case actors.Version0:
                 return load0(store, act.Head)
            
			case actors.Version2:
                 return load2(store, act.Head)
            
			case actors.Version3:
                 return load3(store, act.Head)
            
			case actors.Version4:
                 return load4(store, act.Head)
            
			case actors.Version5:
                 return load5(store, act.Head)
            
			case actors.Version6:
                 return load6(store, act.Head)
            
			case actors.Version7:
                 return load7(store, act.Head)
            
			case actors.Version8:
                 return load8(store, act.Head)
            
            default:
                return nil, xerrors.Errorf("unknown actor version: %d", av)
       }
	}

	switch act.Code {

	case builtin0.AccountActorCodeID:
		return load0(store, act.Head)

	case builtin2.AccountActorCodeID:
		return load2(store, act.Head)

	case builtin3.AccountActorCodeID:
		return load3(store, act.Head)

	case builtin4.AccountActorCodeID:
		return load4(store, act.Head)

	case builtin5.AccountActorCodeID:
		return load5(store, act.Head)

	case builtin6.AccountActorCodeID:
		return load6(store, act.Head)

	case builtin7.AccountActorCodeID:
		return load7(store, act.Head)

	case builtin8.AccountActorCodeID:
		return load8(store, act.Head)

	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

func MakeState(store adt.Store, av actors.Version, addr address.Address) (State, error) {
	switch av {

	case actors.Version0:
		return make0(store, addr)

	case actors.Version2:
		return make2(store, addr)

	case actors.Version3:
		return make3(store, addr)

	case actors.Version4:
		return make4(store, addr)

	case actors.Version5:
		return make5(store, addr)

	case actors.Version6:
		return make6(store, addr)

	case actors.Version7:
		return make7(store, addr)

	case actors.Version8:
		return make8(store, addr)

}
	return nil, xerrors.Errorf("unknown actor version %d", av)
}

func GetActorCodeID(av actors.Version) (cid.Cid, error) {
    if c, ok := actors.GetActorCodeID(av, "account"); ok {
       return c, nil
    }

	switch av {

	case actors.Version0:
		return builtin0.AccountActorCodeID, nil

	case actors.Version2:
		return builtin2.AccountActorCodeID, nil

	case actors.Version3:
		return builtin3.AccountActorCodeID, nil

	case actors.Version4:
		return builtin4.AccountActorCodeID, nil

	case actors.Version5:
		return builtin5.AccountActorCodeID, nil

	case actors.Version6:
		return builtin6.AccountActorCodeID, nil

	case actors.Version7:
		return builtin7.AccountActorCodeID, nil

	case actors.Version8:
		return builtin8.AccountActorCodeID, nil

	}

	return cid.Undef, xerrors.Errorf("unknown actor version %d", av)
}

type State interface {
	cbor.Marshaler

	PubkeyAddress() (address.Address, error)
	GetState() interface{}
}
