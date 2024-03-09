package gwconn

const (
	GW_RpcPort = "42000"
)

// General actions
const (
	GeneralActAdd = iota + 1
	GeneralActDel
	GeneralActDisable
	GeneralActEnable
)

// Bridge targets type
const (
	BridgeTargetTypPlain = iota + 1
	BridgeTargetTypSsl
)
