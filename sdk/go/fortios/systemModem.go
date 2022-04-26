// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure MODEM. Applies to FortiOS Version `>= 7.0.4`.
//
// ## Import
//
// System Modem can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemModem:SystemModem labelname SystemModem
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemModem:SystemModem labelname SystemModem
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemModem struct {
	pulumi.CustomResourceState

	// Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
	Altmode pulumi.StringOutput `pulumi:"altmode"`
	// Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype1 pulumi.StringOutput `pulumi:"authtype1"`
	// Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype2 pulumi.StringOutput `pulumi:"authtype2"`
	// Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype3 pulumi.StringOutput `pulumi:"authtype3"`
	// Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
	AutoDial pulumi.StringOutput `pulumi:"autoDial"`
	// Connection completion timeout (30 - 255 sec, default = 90).
	ConnectTimeout pulumi.IntOutput `pulumi:"connectTimeout"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd1 pulumi.StringOutput `pulumi:"dialCmd1"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd2 pulumi.StringOutput `pulumi:"dialCmd2"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd3 pulumi.StringOutput `pulumi:"dialCmd3"`
	// Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringOutput `pulumi:"dialOnDemand"`
	// Distance of learned routes (1 - 255, default = 1).
	Distance pulumi.IntOutput `pulumi:"distance"`
	// Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
	DontSendCr1 pulumi.StringOutput `pulumi:"dontSendCr1"`
	// Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
	DontSendCr2 pulumi.StringOutput `pulumi:"dontSendCr2"`
	// Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
	DontSendCr3 pulumi.StringOutput `pulumi:"dontSendCr3"`
	// Extra initialization string to ISP 1.
	ExtraInit1 pulumi.StringOutput `pulumi:"extraInit1"`
	// Extra initialization string to ISP 2.
	ExtraInit2 pulumi.StringOutput `pulumi:"extraInit2"`
	// Extra initialization string to ISP 3.
	ExtraInit3 pulumi.StringOutput `pulumi:"extraInit3"`
	// Hold down timer in seconds (1 - 60 sec).
	HolddownTimer pulumi.IntOutput `pulumi:"holddownTimer"`
	// MODEM connection idle time (1 - 9999 min, default = 5).
	IdleTimer pulumi.IntOutput `pulumi:"idleTimer"`
	// Name of redundant interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Allow connection only to the specified Location Area Code (LAC).
	LockdownLac pulumi.StringOutput `pulumi:"lockdownLac"`
	// Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
	NetworkInit pulumi.StringOutput `pulumi:"networkInit"`
	// Password to access the specified dialup account.
	Passwd1 pulumi.StringPtrOutput `pulumi:"passwd1"`
	// Password to access the specified dialup account.
	Passwd2 pulumi.StringPtrOutput `pulumi:"passwd2"`
	// Password to access the specified dialup account.
	Passwd3 pulumi.StringPtrOutput `pulumi:"passwd3"`
	// Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem1 pulumi.StringOutput `pulumi:"peerModem1"`
	// Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem2 pulumi.StringOutput `pulumi:"peerModem2"`
	// Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem3 pulumi.StringOutput `pulumi:"peerModem3"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone1 pulumi.StringOutput `pulumi:"phone1"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone2 pulumi.StringOutput `pulumi:"phone2"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone3 pulumi.StringOutput `pulumi:"phone3"`
	// AT command to set the PIN (AT+PIN=<pin>).
	PinInit pulumi.StringOutput `pulumi:"pinInit"`
	// Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
	PppEchoRequest1 pulumi.StringOutput `pulumi:"pppEchoRequest1"`
	// Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
	PppEchoRequest2 pulumi.StringOutput `pulumi:"pppEchoRequest2"`
	// Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
	PppEchoRequest3 pulumi.StringOutput `pulumi:"pppEchoRequest3"`
	// Priority of learned routes (0 - 4294967295, default = 0).
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
	Redial pulumi.StringOutput `pulumi:"redial"`
	// Number of dial attempts before resetting modem (0 = never reset).
	Reset pulumi.IntOutput `pulumi:"reset"`
	// Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable traffic-check. Valid values: `enable`, `disable`.
	TrafficCheck pulumi.StringOutput `pulumi:"trafficCheck"`
	// User name to access the specified dialup account.
	Username1 pulumi.StringOutput `pulumi:"username1"`
	// User name to access the specified dialup account.
	Username2 pulumi.StringOutput `pulumi:"username2"`
	// User name to access the specified dialup account.
	Username3 pulumi.StringOutput `pulumi:"username3"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
	WirelessPort pulumi.IntOutput `pulumi:"wirelessPort"`
}

// NewSystemModem registers a new resource with the given unique name, arguments, and options.
func NewSystemModem(ctx *pulumi.Context,
	name string, args *SystemModemArgs, opts ...pulumi.ResourceOption) (*SystemModem, error) {
	if args == nil {
		args = &SystemModemArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemModem
	err := ctx.RegisterResource("fortios:index/systemModem:SystemModem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemModem gets an existing SystemModem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemModem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemModemState, opts ...pulumi.ResourceOption) (*SystemModem, error) {
	var resource SystemModem
	err := ctx.ReadResource("fortios:index/systemModem:SystemModem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemModem resources.
type systemModemState struct {
	// Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
	Action *string `pulumi:"action"`
	// Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
	Altmode *string `pulumi:"altmode"`
	// Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype1 *string `pulumi:"authtype1"`
	// Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype2 *string `pulumi:"authtype2"`
	// Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype3 *string `pulumi:"authtype3"`
	// Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
	AutoDial *string `pulumi:"autoDial"`
	// Connection completion timeout (30 - 255 sec, default = 90).
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd1 *string `pulumi:"dialCmd1"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd2 *string `pulumi:"dialCmd2"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd3 *string `pulumi:"dialCmd3"`
	// Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
	DialOnDemand *string `pulumi:"dialOnDemand"`
	// Distance of learned routes (1 - 255, default = 1).
	Distance *int `pulumi:"distance"`
	// Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
	DontSendCr1 *string `pulumi:"dontSendCr1"`
	// Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
	DontSendCr2 *string `pulumi:"dontSendCr2"`
	// Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
	DontSendCr3 *string `pulumi:"dontSendCr3"`
	// Extra initialization string to ISP 1.
	ExtraInit1 *string `pulumi:"extraInit1"`
	// Extra initialization string to ISP 2.
	ExtraInit2 *string `pulumi:"extraInit2"`
	// Extra initialization string to ISP 3.
	ExtraInit3 *string `pulumi:"extraInit3"`
	// Hold down timer in seconds (1 - 60 sec).
	HolddownTimer *int `pulumi:"holddownTimer"`
	// MODEM connection idle time (1 - 9999 min, default = 5).
	IdleTimer *int `pulumi:"idleTimer"`
	// Name of redundant interface.
	Interface *string `pulumi:"interface"`
	// Allow connection only to the specified Location Area Code (LAC).
	LockdownLac *string `pulumi:"lockdownLac"`
	// Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
	Mode *string `pulumi:"mode"`
	// AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
	NetworkInit *string `pulumi:"networkInit"`
	// Password to access the specified dialup account.
	Passwd1 *string `pulumi:"passwd1"`
	// Password to access the specified dialup account.
	Passwd2 *string `pulumi:"passwd2"`
	// Password to access the specified dialup account.
	Passwd3 *string `pulumi:"passwd3"`
	// Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem1 *string `pulumi:"peerModem1"`
	// Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem2 *string `pulumi:"peerModem2"`
	// Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem3 *string `pulumi:"peerModem3"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone1 *string `pulumi:"phone1"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone2 *string `pulumi:"phone2"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone3 *string `pulumi:"phone3"`
	// AT command to set the PIN (AT+PIN=<pin>).
	PinInit *string `pulumi:"pinInit"`
	// Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
	PppEchoRequest1 *string `pulumi:"pppEchoRequest1"`
	// Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
	PppEchoRequest2 *string `pulumi:"pppEchoRequest2"`
	// Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
	PppEchoRequest3 *string `pulumi:"pppEchoRequest3"`
	// Priority of learned routes (0 - 4294967295, default = 0).
	Priority *int `pulumi:"priority"`
	// Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
	Redial *string `pulumi:"redial"`
	// Number of dial attempts before resetting modem (0 = never reset).
	Reset *int `pulumi:"reset"`
	// Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable traffic-check. Valid values: `enable`, `disable`.
	TrafficCheck *string `pulumi:"trafficCheck"`
	// User name to access the specified dialup account.
	Username1 *string `pulumi:"username1"`
	// User name to access the specified dialup account.
	Username2 *string `pulumi:"username2"`
	// User name to access the specified dialup account.
	Username3 *string `pulumi:"username3"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
	WirelessPort *int `pulumi:"wirelessPort"`
}

type SystemModemState struct {
	// Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
	Action pulumi.StringPtrInput
	// Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
	Altmode pulumi.StringPtrInput
	// Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype1 pulumi.StringPtrInput
	// Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype2 pulumi.StringPtrInput
	// Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype3 pulumi.StringPtrInput
	// Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
	AutoDial pulumi.StringPtrInput
	// Connection completion timeout (30 - 255 sec, default = 90).
	ConnectTimeout pulumi.IntPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd1 pulumi.StringPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd2 pulumi.StringPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd3 pulumi.StringPtrInput
	// Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringPtrInput
	// Distance of learned routes (1 - 255, default = 1).
	Distance pulumi.IntPtrInput
	// Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
	DontSendCr1 pulumi.StringPtrInput
	// Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
	DontSendCr2 pulumi.StringPtrInput
	// Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
	DontSendCr3 pulumi.StringPtrInput
	// Extra initialization string to ISP 1.
	ExtraInit1 pulumi.StringPtrInput
	// Extra initialization string to ISP 2.
	ExtraInit2 pulumi.StringPtrInput
	// Extra initialization string to ISP 3.
	ExtraInit3 pulumi.StringPtrInput
	// Hold down timer in seconds (1 - 60 sec).
	HolddownTimer pulumi.IntPtrInput
	// MODEM connection idle time (1 - 9999 min, default = 5).
	IdleTimer pulumi.IntPtrInput
	// Name of redundant interface.
	Interface pulumi.StringPtrInput
	// Allow connection only to the specified Location Area Code (LAC).
	LockdownLac pulumi.StringPtrInput
	// Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
	Mode pulumi.StringPtrInput
	// AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
	NetworkInit pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd1 pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd2 pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd3 pulumi.StringPtrInput
	// Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem1 pulumi.StringPtrInput
	// Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem2 pulumi.StringPtrInput
	// Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem3 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone1 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone2 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone3 pulumi.StringPtrInput
	// AT command to set the PIN (AT+PIN=<pin>).
	PinInit pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
	PppEchoRequest1 pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
	PppEchoRequest2 pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
	PppEchoRequest3 pulumi.StringPtrInput
	// Priority of learned routes (0 - 4294967295, default = 0).
	Priority pulumi.IntPtrInput
	// Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
	Redial pulumi.StringPtrInput
	// Number of dial attempts before resetting modem (0 = never reset).
	Reset pulumi.IntPtrInput
	// Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable traffic-check. Valid values: `enable`, `disable`.
	TrafficCheck pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username1 pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username2 pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username3 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
	WirelessPort pulumi.IntPtrInput
}

func (SystemModemState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemModemState)(nil)).Elem()
}

type systemModemArgs struct {
	// Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
	Action *string `pulumi:"action"`
	// Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
	Altmode *string `pulumi:"altmode"`
	// Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype1 *string `pulumi:"authtype1"`
	// Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype2 *string `pulumi:"authtype2"`
	// Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype3 *string `pulumi:"authtype3"`
	// Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
	AutoDial *string `pulumi:"autoDial"`
	// Connection completion timeout (30 - 255 sec, default = 90).
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd1 *string `pulumi:"dialCmd1"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd2 *string `pulumi:"dialCmd2"`
	// Dial command (this is often an ATD or ATDT command).
	DialCmd3 *string `pulumi:"dialCmd3"`
	// Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
	DialOnDemand *string `pulumi:"dialOnDemand"`
	// Distance of learned routes (1 - 255, default = 1).
	Distance *int `pulumi:"distance"`
	// Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
	DontSendCr1 *string `pulumi:"dontSendCr1"`
	// Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
	DontSendCr2 *string `pulumi:"dontSendCr2"`
	// Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
	DontSendCr3 *string `pulumi:"dontSendCr3"`
	// Extra initialization string to ISP 1.
	ExtraInit1 *string `pulumi:"extraInit1"`
	// Extra initialization string to ISP 2.
	ExtraInit2 *string `pulumi:"extraInit2"`
	// Extra initialization string to ISP 3.
	ExtraInit3 *string `pulumi:"extraInit3"`
	// Hold down timer in seconds (1 - 60 sec).
	HolddownTimer *int `pulumi:"holddownTimer"`
	// MODEM connection idle time (1 - 9999 min, default = 5).
	IdleTimer *int `pulumi:"idleTimer"`
	// Name of redundant interface.
	Interface *string `pulumi:"interface"`
	// Allow connection only to the specified Location Area Code (LAC).
	LockdownLac *string `pulumi:"lockdownLac"`
	// Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
	Mode *string `pulumi:"mode"`
	// AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
	NetworkInit *string `pulumi:"networkInit"`
	// Password to access the specified dialup account.
	Passwd1 *string `pulumi:"passwd1"`
	// Password to access the specified dialup account.
	Passwd2 *string `pulumi:"passwd2"`
	// Password to access the specified dialup account.
	Passwd3 *string `pulumi:"passwd3"`
	// Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem1 *string `pulumi:"peerModem1"`
	// Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem2 *string `pulumi:"peerModem2"`
	// Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem3 *string `pulumi:"peerModem3"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone1 *string `pulumi:"phone1"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone2 *string `pulumi:"phone2"`
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone3 *string `pulumi:"phone3"`
	// AT command to set the PIN (AT+PIN=<pin>).
	PinInit *string `pulumi:"pinInit"`
	// Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
	PppEchoRequest1 *string `pulumi:"pppEchoRequest1"`
	// Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
	PppEchoRequest2 *string `pulumi:"pppEchoRequest2"`
	// Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
	PppEchoRequest3 *string `pulumi:"pppEchoRequest3"`
	// Priority of learned routes (0 - 4294967295, default = 0).
	Priority *int `pulumi:"priority"`
	// Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
	Redial *string `pulumi:"redial"`
	// Number of dial attempts before resetting modem (0 = never reset).
	Reset *int `pulumi:"reset"`
	// Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable traffic-check. Valid values: `enable`, `disable`.
	TrafficCheck *string `pulumi:"trafficCheck"`
	// User name to access the specified dialup account.
	Username1 *string `pulumi:"username1"`
	// User name to access the specified dialup account.
	Username2 *string `pulumi:"username2"`
	// User name to access the specified dialup account.
	Username3 *string `pulumi:"username3"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
	WirelessPort *int `pulumi:"wirelessPort"`
}

// The set of arguments for constructing a SystemModem resource.
type SystemModemArgs struct {
	// Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
	Action pulumi.StringPtrInput
	// Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
	Altmode pulumi.StringPtrInput
	// Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype1 pulumi.StringPtrInput
	// Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype2 pulumi.StringPtrInput
	// Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
	Authtype3 pulumi.StringPtrInput
	// Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
	AutoDial pulumi.StringPtrInput
	// Connection completion timeout (30 - 255 sec, default = 90).
	ConnectTimeout pulumi.IntPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd1 pulumi.StringPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd2 pulumi.StringPtrInput
	// Dial command (this is often an ATD or ATDT command).
	DialCmd3 pulumi.StringPtrInput
	// Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
	DialOnDemand pulumi.StringPtrInput
	// Distance of learned routes (1 - 255, default = 1).
	Distance pulumi.IntPtrInput
	// Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
	DontSendCr1 pulumi.StringPtrInput
	// Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
	DontSendCr2 pulumi.StringPtrInput
	// Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
	DontSendCr3 pulumi.StringPtrInput
	// Extra initialization string to ISP 1.
	ExtraInit1 pulumi.StringPtrInput
	// Extra initialization string to ISP 2.
	ExtraInit2 pulumi.StringPtrInput
	// Extra initialization string to ISP 3.
	ExtraInit3 pulumi.StringPtrInput
	// Hold down timer in seconds (1 - 60 sec).
	HolddownTimer pulumi.IntPtrInput
	// MODEM connection idle time (1 - 9999 min, default = 5).
	IdleTimer pulumi.IntPtrInput
	// Name of redundant interface.
	Interface pulumi.StringPtrInput
	// Allow connection only to the specified Location Area Code (LAC).
	LockdownLac pulumi.StringPtrInput
	// Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
	Mode pulumi.StringPtrInput
	// AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
	NetworkInit pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd1 pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd2 pulumi.StringPtrInput
	// Password to access the specified dialup account.
	Passwd3 pulumi.StringPtrInput
	// Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem1 pulumi.StringPtrInput
	// Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem2 pulumi.StringPtrInput
	// Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
	PeerModem3 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone1 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone2 pulumi.StringPtrInput
	// Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
	Phone3 pulumi.StringPtrInput
	// AT command to set the PIN (AT+PIN=<pin>).
	PinInit pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
	PppEchoRequest1 pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
	PppEchoRequest2 pulumi.StringPtrInput
	// Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
	PppEchoRequest3 pulumi.StringPtrInput
	// Priority of learned routes (0 - 4294967295, default = 0).
	Priority pulumi.IntPtrInput
	// Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
	Redial pulumi.StringPtrInput
	// Number of dial attempts before resetting modem (0 = never reset).
	Reset pulumi.IntPtrInput
	// Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable traffic-check. Valid values: `enable`, `disable`.
	TrafficCheck pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username1 pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username2 pulumi.StringPtrInput
	// User name to access the specified dialup account.
	Username3 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
	WirelessPort pulumi.IntPtrInput
}

func (SystemModemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemModemArgs)(nil)).Elem()
}

type SystemModemInput interface {
	pulumi.Input

	ToSystemModemOutput() SystemModemOutput
	ToSystemModemOutputWithContext(ctx context.Context) SystemModemOutput
}

func (*SystemModem) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemModem)(nil)).Elem()
}

func (i *SystemModem) ToSystemModemOutput() SystemModemOutput {
	return i.ToSystemModemOutputWithContext(context.Background())
}

func (i *SystemModem) ToSystemModemOutputWithContext(ctx context.Context) SystemModemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemModemOutput)
}

// SystemModemArrayInput is an input type that accepts SystemModemArray and SystemModemArrayOutput values.
// You can construct a concrete instance of `SystemModemArrayInput` via:
//
//          SystemModemArray{ SystemModemArgs{...} }
type SystemModemArrayInput interface {
	pulumi.Input

	ToSystemModemArrayOutput() SystemModemArrayOutput
	ToSystemModemArrayOutputWithContext(context.Context) SystemModemArrayOutput
}

type SystemModemArray []SystemModemInput

func (SystemModemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemModem)(nil)).Elem()
}

func (i SystemModemArray) ToSystemModemArrayOutput() SystemModemArrayOutput {
	return i.ToSystemModemArrayOutputWithContext(context.Background())
}

func (i SystemModemArray) ToSystemModemArrayOutputWithContext(ctx context.Context) SystemModemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemModemArrayOutput)
}

// SystemModemMapInput is an input type that accepts SystemModemMap and SystemModemMapOutput values.
// You can construct a concrete instance of `SystemModemMapInput` via:
//
//          SystemModemMap{ "key": SystemModemArgs{...} }
type SystemModemMapInput interface {
	pulumi.Input

	ToSystemModemMapOutput() SystemModemMapOutput
	ToSystemModemMapOutputWithContext(context.Context) SystemModemMapOutput
}

type SystemModemMap map[string]SystemModemInput

func (SystemModemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemModem)(nil)).Elem()
}

func (i SystemModemMap) ToSystemModemMapOutput() SystemModemMapOutput {
	return i.ToSystemModemMapOutputWithContext(context.Background())
}

func (i SystemModemMap) ToSystemModemMapOutputWithContext(ctx context.Context) SystemModemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemModemMapOutput)
}

type SystemModemOutput struct{ *pulumi.OutputState }

func (SystemModemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemModem)(nil)).Elem()
}

func (o SystemModemOutput) ToSystemModemOutput() SystemModemOutput {
	return o
}

func (o SystemModemOutput) ToSystemModemOutputWithContext(ctx context.Context) SystemModemOutput {
	return o
}

type SystemModemArrayOutput struct{ *pulumi.OutputState }

func (SystemModemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemModem)(nil)).Elem()
}

func (o SystemModemArrayOutput) ToSystemModemArrayOutput() SystemModemArrayOutput {
	return o
}

func (o SystemModemArrayOutput) ToSystemModemArrayOutputWithContext(ctx context.Context) SystemModemArrayOutput {
	return o
}

func (o SystemModemArrayOutput) Index(i pulumi.IntInput) SystemModemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemModem {
		return vs[0].([]*SystemModem)[vs[1].(int)]
	}).(SystemModemOutput)
}

type SystemModemMapOutput struct{ *pulumi.OutputState }

func (SystemModemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemModem)(nil)).Elem()
}

func (o SystemModemMapOutput) ToSystemModemMapOutput() SystemModemMapOutput {
	return o
}

func (o SystemModemMapOutput) ToSystemModemMapOutputWithContext(ctx context.Context) SystemModemMapOutput {
	return o
}

func (o SystemModemMapOutput) MapIndex(k pulumi.StringInput) SystemModemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemModem {
		return vs[0].(map[string]*SystemModem)[vs[1].(string)]
	}).(SystemModemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemModemInput)(nil)).Elem(), &SystemModem{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemModemArrayInput)(nil)).Elem(), SystemModemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemModemMapInput)(nil)).Elem(), SystemModemMap{})
	pulumi.RegisterOutputType(SystemModemOutput{})
	pulumi.RegisterOutputType(SystemModemArrayOutput{})
	pulumi.RegisterOutputType(SystemModemMapOutput{})
}
