package g2productserver

import (
	pb "github.com/senzing/g2-sdk-proto/go/g2product"
	"github.com/senzing/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// server is used to implement helloworld.GreeterServer.
type G2ProductServer struct {
	pb.UnimplementedG2ProductServer
	isTrace bool
	logger  logging.LoggingInterface
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the g2productserver package found messages having the format "senzing-6999xxxx".
const ProductId = 6016

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for the g2productserver package.
var IdMessages = map[int]string{
	3:    "Enter Destroy(%+v).",
	4:    "Exit  Destroy(%+v) returned (%v).",
	9:    "Enter Init(%+v).",
	10:   "Exit  Init(%+v) returned (%v).",
	11:   "Enter License(%+v).",
	12:   "Exit  License(%+v) returned (%s, %v).",
	13:   "Enter SetLogLevel(%+v).",
	14:   "Exit  SetLogLevel(%+v) returned (%v).",
	15:   "Enter ValidateLicenseFile(%+v).",
	16:   "Exit  ValidateLicenseFile(%+v) returned (%s, %v).",
	17:   "Enter ValidateLicenseStringBase64(%+v).",
	18:   "Exit  ValidateLicenseStringBase64(%+v) returned (%s, %v).",
	19:   "Enter Version(%+v).",
	20:   "Exit  Version(%+v) returned (%s, %v).",
	4001: "Destroy() not supported in gRPC",
	4002: "Init() not supported in gRPC",
	4003: "InitWithConfigID() not supported in gRPC",
	5901: "During setup, call to messagelogger.NewSenzingApiLogger() failed.",
	5902: "During setup, call to g2eg2engineconfigurationjson.BuildSimpleSystemConfigurationJson() failed.",
	5903: "During setup, call to g2engine.Init() failed.",
	5904: "During setup, call to g2engine.PurgeRepository() failed.",
	5905: "During setup, call to g2engine.Destroy() failed.",
	5906: "During setup, call to g2config.Init() failed.",
	5907: "During setup, call to g2config.Create() failed.",
	5908: "During setup, call to g2config.AddDataSource() failed.",
	5909: "During setup, call to g2config.Save() failed.",
	5910: "During setup, call to g2config.Close() failed.",
	5911: "During setup, call to g2config.Destroy() failed.",
	5912: "During setup, call to g2configmgr.Init() failed.",
	5913: "During setup, call to g2configmgr.AddConfig() failed.",
	5914: "During setup, call to g2configmgr.SetDefaultConfigID() failed.",
	5915: "During setup, call to g2configmgr.Destroy() failed.",
	5916: "During setup, call to g2engine.Init() failed.",
	5917: "During setup, call to g2engine.AddRecord() failed.",
	5918: "During setup, call to g2engine.Destroy() failed.",
	5920: "During setup, call to setupSenzingConfig() failed.",
	5921: "During setup, call to setupPurgeRepository() failed.",
	5922: "During setup, call to setupAddRecords() failed.",
	5931: "During setup, call to g2engine.Init() failed.",
	5932: "During setup, call to g2engine.PurgeRepository() failed.",
	5933: "During setup, call to g2engine.Destroy() failed.",
}

// Status strings for specific g2productserver messages.
var IdStatuses = map[int]string{}
