package utils

import (
	"fmt"
	"odin/ent"
)

func TxRollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}

	return err
}
