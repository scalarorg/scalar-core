package btc

import "github.com/scalarorg/scalar-core/utils/log"

func (c *BtcClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}


