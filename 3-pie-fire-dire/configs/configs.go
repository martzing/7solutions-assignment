package configs

var BeefUrl *string
var GRPCHost *string
var GRPCPort *string
var GRPCGateWayHost *string
var GRPCGateWayPort *string

func BootConfig() {
	env := NewEnvironment()

	BeefUrl = env.GetString("BEEF_URL")
	GRPCHost = env.GetString("GRPC_HOST")
	GRPCPort = env.GetString("GRPC_PORT")
	GRPCGateWayHost = env.GetString("GRPC_GATEWAY_HOST")
	GRPCGateWayPort = env.GetString("GRPC_GATEWAY_PORT")
}
