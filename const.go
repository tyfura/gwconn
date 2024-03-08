package gwconn

const (
	GW_RpcPort = "42000"
)

// Bridge targets actions
const (
	BridgeTargetActAdd = iota + 1
	BridgeTargetActDel
	BridgeTargetActDisable
	BridgeTargetActEnable
)

// Bridge targets type
const (
	BridgeTargetTypPlain = iota + 1
	BridgeTargetTypSsl
)
