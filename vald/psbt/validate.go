package psbt

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/scalar-core/utils/clog"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (mgr Mgr) ValidatePsbt(client *rpcclient.Client, p covenantTypes.Psbt) error {
	psbt, err := psbt.NewFromRawBytes(bytes.NewReader(p), false)
	if err != nil {
		return fmt.Errorf("failed to parse PSBT: %w", err)
	}

	if len(psbt.Inputs) == 0 {
		return fmt.Errorf("PSBT has no inputs")
	}

	// TODO: use go routines to validate each input

	// For each PSBT input, verify it exists as a UTXO
	for i, input := range psbt.Inputs {
		txIn := psbt.UnsignedTx.TxIn[i]

		clog.Redf("input.WitnessUtxo: %+v", input.WitnessUtxo)
		clog.Redf("txIn: %+v", txIn)

		// Create outpoint hash string
		// txHash := txIn.PreviousOutPoint.Hash.String()

		// // Get the transaction from the Bitcoin node
		// tx, err := client.GetRawTransaction(&txIn.PreviousOutPoint.Hash)
		// if err != nil {
		// 	return fmt.Errorf("input %d: failed to fetch transaction %s: %w", i, txHash, err)
		// }

		// // Check if the referenced output index exists
		// if int(txIn.PreviousOutPoint.Index) >= len(tx.MsgTx().TxOut) {
		// 	return fmt.Errorf("input %d: invalid output index %d for transaction %s",
		// 		i, txIn.PreviousOutPoint.Index, txHash)
		// }

		// // Verify the UTXO is unspent using GetTxOut
		// txOut, err := client.GetTxOut(&txIn.PreviousOutPoint.Hash, txIn.PreviousOutPoint.Index, true)
		// if err != nil {
		// 	return fmt.Errorf("input %d: failed to check UTXO status: %w", i, err)
		// }
		// if txOut == nil {
		// 	return fmt.Errorf("input %d: UTXO is already spent", i)
		// }

		// // Verify witness utxo if present
		// if input.WitnessUtxo != nil {
		// 	outputIndex := txIn.PreviousOutPoint.Index
		// 	if input.WitnessUtxo.Value != tx.MsgTx().TxOut[outputIndex].Value {
		// 		return fmt.Errorf("input %d: witness UTXO amount mismatch", i)
		// 	}
		// 	if !bytes.Equal(input.WitnessUtxo.PkScript, tx.MsgTx().TxOut[outputIndex].PkScript) {
		// 		return fmt.Errorf("input %d: witness UTXO script mismatch", i)
		// 	}
		// }
	}

	return nil
}
