// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.
 *
 * ## Import
 *
 * WirelessController WtpProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerWtpProfile:WirelessControllerWtpProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerWtpProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerWtpProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerWtpProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerWtpProfile {
        return new WirelessControllerWtpProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerWtpProfile:WirelessControllerWtpProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerWtpProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerWtpProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerWtpProfile.__pulumiType;
    }

    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).
     */
    public readonly apCountry!: pulumi.Output<string>;
    /**
     * Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly apHandoff!: pulumi.Output<string>;
    /**
     * AP local configuration profile name.
     */
    public readonly apcfgProfile!: pulumi.Output<string>;
    /**
     * Bluetooth Low Energy profile name.
     */
    public readonly bleProfile!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable CAPWAP control message data channel offload.
     */
    public readonly controlMessageOffload!: pulumi.Output<string>;
    /**
     * List of MAC addresses that are denied access to this WTP, FortiAP, or AP. The structure of `denyMacList` block is documented below.
     */
    public readonly denyMacLists!: pulumi.Output<outputs.WirelessControllerWtpProfileDenyMacList[] | undefined>;
    /**
     * Enable/disable data channel DTLS in kernel. Valid values: `enable`, `disable`.
     */
    public readonly dtlsInKernel!: pulumi.Output<string>;
    /**
     * WTP data channel DTLS policy (default = clear-text). Valid values: `clear-text`, `dtls-enabled`, `ipsec-vpn`.
     */
    public readonly dtlsPolicy!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable use of energy efficient Ethernet on WTP. Valid values: `enable`, `disable`.
     */
    public readonly energyEfficientEthernet!: pulumi.Output<string>;
    /**
     * Enable/disable station/VAP/radio extension information. Valid values: `enable`, `disable`.
     */
    public readonly extInfoEnable!: pulumi.Output<string>;
    /**
     * Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly frequencyHandoff!: pulumi.Output<string>;
    /**
     * Enable/disable client load balancing during roaming to avoid roaming delay (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly handoffRoaming!: pulumi.Output<string>;
    /**
     * Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
     */
    public readonly handoffRssi!: pulumi.Output<number>;
    /**
     * Threshold value for AP handoff.
     */
    public readonly handoffStaThresh!: pulumi.Output<number>;
    /**
     * Select how to prevent IP fragmentation for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    public readonly ipFragmentPreventing!: pulumi.Output<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    public readonly lan!: pulumi.Output<outputs.WirelessControllerWtpProfileLan | undefined>;
    /**
     * Set various location based service (LBS) options. The structure of `lbs` block is documented below.
     */
    public readonly lbs!: pulumi.Output<outputs.WirelessControllerWtpProfileLbs | undefined>;
    /**
     * Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ledSchedules` block is documented below.
     */
    public readonly ledSchedules!: pulumi.Output<outputs.WirelessControllerWtpProfileLedSchedule[] | undefined>;
    /**
     * Enable/disable use of LEDs on WTP (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly ledState!: pulumi.Output<string>;
    /**
     * Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly lldp!: pulumi.Output<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    public readonly loginPasswd!: pulumi.Output<string | undefined>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    public readonly loginPasswdChange!: pulumi.Output<string>;
    /**
     * Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
     */
    public readonly maxClients!: pulumi.Output<number>;
    /**
     * Virtual Access Point (VAP) name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WTP, FortiAP, or AP platform. The structure of `platform` block is documented below.
     */
    public readonly platform!: pulumi.Output<outputs.WirelessControllerWtpProfilePlatform | undefined>;
    /**
     * Set the WTP, FortiAP, or AP's PoE mode.
     */
    public readonly poeMode!: pulumi.Output<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    public readonly radio1!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio1 | undefined>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    public readonly radio2!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio2 | undefined>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    public readonly radio3!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio3 | undefined>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    public readonly radio4!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio4 | undefined>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly splitTunnelingAclLocalApSubnet!: pulumi.Output<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    public readonly splitTunnelingAclPath!: pulumi.Output<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    public readonly splitTunnelingAcls!: pulumi.Output<outputs.WirelessControllerWtpProfileSplitTunnelingAcl[] | undefined>;
    /**
     * Downlink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    public readonly tunMtuDownlink!: pulumi.Output<number>;
    /**
     * Uplink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    public readonly tunMtuUplink!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable using a WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    public readonly wanPortMode!: pulumi.Output<string>;

    /**
     * Create a WirelessControllerWtpProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerWtpProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerWtpProfileArgs | WirelessControllerWtpProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerWtpProfileState | undefined;
            inputs["allowaccess"] = state ? state.allowaccess : undefined;
            inputs["apCountry"] = state ? state.apCountry : undefined;
            inputs["apHandoff"] = state ? state.apHandoff : undefined;
            inputs["apcfgProfile"] = state ? state.apcfgProfile : undefined;
            inputs["bleProfile"] = state ? state.bleProfile : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["controlMessageOffload"] = state ? state.controlMessageOffload : undefined;
            inputs["denyMacLists"] = state ? state.denyMacLists : undefined;
            inputs["dtlsInKernel"] = state ? state.dtlsInKernel : undefined;
            inputs["dtlsPolicy"] = state ? state.dtlsPolicy : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["energyEfficientEthernet"] = state ? state.energyEfficientEthernet : undefined;
            inputs["extInfoEnable"] = state ? state.extInfoEnable : undefined;
            inputs["frequencyHandoff"] = state ? state.frequencyHandoff : undefined;
            inputs["handoffRoaming"] = state ? state.handoffRoaming : undefined;
            inputs["handoffRssi"] = state ? state.handoffRssi : undefined;
            inputs["handoffStaThresh"] = state ? state.handoffStaThresh : undefined;
            inputs["ipFragmentPreventing"] = state ? state.ipFragmentPreventing : undefined;
            inputs["lan"] = state ? state.lan : undefined;
            inputs["lbs"] = state ? state.lbs : undefined;
            inputs["ledSchedules"] = state ? state.ledSchedules : undefined;
            inputs["ledState"] = state ? state.ledState : undefined;
            inputs["lldp"] = state ? state.lldp : undefined;
            inputs["loginPasswd"] = state ? state.loginPasswd : undefined;
            inputs["loginPasswdChange"] = state ? state.loginPasswdChange : undefined;
            inputs["maxClients"] = state ? state.maxClients : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["poeMode"] = state ? state.poeMode : undefined;
            inputs["radio1"] = state ? state.radio1 : undefined;
            inputs["radio2"] = state ? state.radio2 : undefined;
            inputs["radio3"] = state ? state.radio3 : undefined;
            inputs["radio4"] = state ? state.radio4 : undefined;
            inputs["splitTunnelingAclLocalApSubnet"] = state ? state.splitTunnelingAclLocalApSubnet : undefined;
            inputs["splitTunnelingAclPath"] = state ? state.splitTunnelingAclPath : undefined;
            inputs["splitTunnelingAcls"] = state ? state.splitTunnelingAcls : undefined;
            inputs["tunMtuDownlink"] = state ? state.tunMtuDownlink : undefined;
            inputs["tunMtuUplink"] = state ? state.tunMtuUplink : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["wanPortMode"] = state ? state.wanPortMode : undefined;
        } else {
            const args = argsOrState as WirelessControllerWtpProfileArgs | undefined;
            inputs["allowaccess"] = args ? args.allowaccess : undefined;
            inputs["apCountry"] = args ? args.apCountry : undefined;
            inputs["apHandoff"] = args ? args.apHandoff : undefined;
            inputs["apcfgProfile"] = args ? args.apcfgProfile : undefined;
            inputs["bleProfile"] = args ? args.bleProfile : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["controlMessageOffload"] = args ? args.controlMessageOffload : undefined;
            inputs["denyMacLists"] = args ? args.denyMacLists : undefined;
            inputs["dtlsInKernel"] = args ? args.dtlsInKernel : undefined;
            inputs["dtlsPolicy"] = args ? args.dtlsPolicy : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["energyEfficientEthernet"] = args ? args.energyEfficientEthernet : undefined;
            inputs["extInfoEnable"] = args ? args.extInfoEnable : undefined;
            inputs["frequencyHandoff"] = args ? args.frequencyHandoff : undefined;
            inputs["handoffRoaming"] = args ? args.handoffRoaming : undefined;
            inputs["handoffRssi"] = args ? args.handoffRssi : undefined;
            inputs["handoffStaThresh"] = args ? args.handoffStaThresh : undefined;
            inputs["ipFragmentPreventing"] = args ? args.ipFragmentPreventing : undefined;
            inputs["lan"] = args ? args.lan : undefined;
            inputs["lbs"] = args ? args.lbs : undefined;
            inputs["ledSchedules"] = args ? args.ledSchedules : undefined;
            inputs["ledState"] = args ? args.ledState : undefined;
            inputs["lldp"] = args ? args.lldp : undefined;
            inputs["loginPasswd"] = args ? args.loginPasswd : undefined;
            inputs["loginPasswdChange"] = args ? args.loginPasswdChange : undefined;
            inputs["maxClients"] = args ? args.maxClients : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["poeMode"] = args ? args.poeMode : undefined;
            inputs["radio1"] = args ? args.radio1 : undefined;
            inputs["radio2"] = args ? args.radio2 : undefined;
            inputs["radio3"] = args ? args.radio3 : undefined;
            inputs["radio4"] = args ? args.radio4 : undefined;
            inputs["splitTunnelingAclLocalApSubnet"] = args ? args.splitTunnelingAclLocalApSubnet : undefined;
            inputs["splitTunnelingAclPath"] = args ? args.splitTunnelingAclPath : undefined;
            inputs["splitTunnelingAcls"] = args ? args.splitTunnelingAcls : undefined;
            inputs["tunMtuDownlink"] = args ? args.tunMtuDownlink : undefined;
            inputs["tunMtuUplink"] = args ? args.tunMtuUplink : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["wanPortMode"] = args ? args.wanPortMode : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WirelessControllerWtpProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerWtpProfile resources.
 */
export interface WirelessControllerWtpProfileState {
    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).
     */
    apCountry?: pulumi.Input<string>;
    /**
     * Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable`, `disable`.
     */
    apHandoff?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    apcfgProfile?: pulumi.Input<string>;
    /**
     * Bluetooth Low Energy profile name.
     */
    bleProfile?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable CAPWAP control message data channel offload.
     */
    controlMessageOffload?: pulumi.Input<string>;
    /**
     * List of MAC addresses that are denied access to this WTP, FortiAP, or AP. The structure of `denyMacList` block is documented below.
     */
    denyMacLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileDenyMacList>[]>;
    /**
     * Enable/disable data channel DTLS in kernel. Valid values: `enable`, `disable`.
     */
    dtlsInKernel?: pulumi.Input<string>;
    /**
     * WTP data channel DTLS policy (default = clear-text). Valid values: `clear-text`, `dtls-enabled`, `ipsec-vpn`.
     */
    dtlsPolicy?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of energy efficient Ethernet on WTP. Valid values: `enable`, `disable`.
     */
    energyEfficientEthernet?: pulumi.Input<string>;
    /**
     * Enable/disable station/VAP/radio extension information. Valid values: `enable`, `disable`.
     */
    extInfoEnable?: pulumi.Input<string>;
    /**
     * Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable`, `disable`.
     */
    frequencyHandoff?: pulumi.Input<string>;
    /**
     * Enable/disable client load balancing during roaming to avoid roaming delay (default = disable). Valid values: `enable`, `disable`.
     */
    handoffRoaming?: pulumi.Input<string>;
    /**
     * Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
     */
    handoffRssi?: pulumi.Input<number>;
    /**
     * Threshold value for AP handoff.
     */
    handoffStaThresh?: pulumi.Input<number>;
    /**
     * Select how to prevent IP fragmentation for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    ipFragmentPreventing?: pulumi.Input<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    lan?: pulumi.Input<inputs.WirelessControllerWtpProfileLan>;
    /**
     * Set various location based service (LBS) options. The structure of `lbs` block is documented below.
     */
    lbs?: pulumi.Input<inputs.WirelessControllerWtpProfileLbs>;
    /**
     * Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ledSchedules` block is documented below.
     */
    ledSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileLedSchedule>[]>;
    /**
     * Enable/disable use of LEDs on WTP (default = disable). Valid values: `enable`, `disable`.
     */
    ledState?: pulumi.Input<string>;
    /**
     * Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = disable). Valid values: `enable`, `disable`.
     */
    lldp?: pulumi.Input<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    loginPasswd?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswdChange?: pulumi.Input<string>;
    /**
     * Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
     */
    maxClients?: pulumi.Input<number>;
    /**
     * Virtual Access Point (VAP) name.
     */
    name?: pulumi.Input<string>;
    /**
     * WTP, FortiAP, or AP platform. The structure of `platform` block is documented below.
     */
    platform?: pulumi.Input<inputs.WirelessControllerWtpProfilePlatform>;
    /**
     * Set the WTP, FortiAP, or AP's PoE mode.
     */
    poeMode?: pulumi.Input<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    radio1?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio1>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    radio2?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio2>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    radio3?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio3>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    radio4?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio4>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    splitTunnelingAclPath?: pulumi.Input<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileSplitTunnelingAcl>[]>;
    /**
     * Downlink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    tunMtuDownlink?: pulumi.Input<number>;
    /**
     * Uplink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    tunMtuUplink?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable using a WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    wanPortMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerWtpProfile resource.
 */
export interface WirelessControllerWtpProfileArgs {
    /**
     * Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).
     */
    apCountry?: pulumi.Input<string>;
    /**
     * Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable`, `disable`.
     */
    apHandoff?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    apcfgProfile?: pulumi.Input<string>;
    /**
     * Bluetooth Low Energy profile name.
     */
    bleProfile?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable CAPWAP control message data channel offload.
     */
    controlMessageOffload?: pulumi.Input<string>;
    /**
     * List of MAC addresses that are denied access to this WTP, FortiAP, or AP. The structure of `denyMacList` block is documented below.
     */
    denyMacLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileDenyMacList>[]>;
    /**
     * Enable/disable data channel DTLS in kernel. Valid values: `enable`, `disable`.
     */
    dtlsInKernel?: pulumi.Input<string>;
    /**
     * WTP data channel DTLS policy (default = clear-text). Valid values: `clear-text`, `dtls-enabled`, `ipsec-vpn`.
     */
    dtlsPolicy?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of energy efficient Ethernet on WTP. Valid values: `enable`, `disable`.
     */
    energyEfficientEthernet?: pulumi.Input<string>;
    /**
     * Enable/disable station/VAP/radio extension information. Valid values: `enable`, `disable`.
     */
    extInfoEnable?: pulumi.Input<string>;
    /**
     * Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable`, `disable`.
     */
    frequencyHandoff?: pulumi.Input<string>;
    /**
     * Enable/disable client load balancing during roaming to avoid roaming delay (default = disable). Valid values: `enable`, `disable`.
     */
    handoffRoaming?: pulumi.Input<string>;
    /**
     * Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
     */
    handoffRssi?: pulumi.Input<number>;
    /**
     * Threshold value for AP handoff.
     */
    handoffStaThresh?: pulumi.Input<number>;
    /**
     * Select how to prevent IP fragmentation for CAPWAP tunneled control and data packets (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.
     */
    ipFragmentPreventing?: pulumi.Input<string>;
    /**
     * WTP LAN port mapping. The structure of `lan` block is documented below.
     */
    lan?: pulumi.Input<inputs.WirelessControllerWtpProfileLan>;
    /**
     * Set various location based service (LBS) options. The structure of `lbs` block is documented below.
     */
    lbs?: pulumi.Input<inputs.WirelessControllerWtpProfileLbs>;
    /**
     * Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ledSchedules` block is documented below.
     */
    ledSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileLedSchedule>[]>;
    /**
     * Enable/disable use of LEDs on WTP (default = disable). Valid values: `enable`, `disable`.
     */
    ledState?: pulumi.Input<string>;
    /**
     * Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = disable). Valid values: `enable`, `disable`.
     */
    lldp?: pulumi.Input<string>;
    /**
     * Set the managed WTP, FortiAP, or AP's administrator password.
     */
    loginPasswd?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswdChange?: pulumi.Input<string>;
    /**
     * Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
     */
    maxClients?: pulumi.Input<number>;
    /**
     * Virtual Access Point (VAP) name.
     */
    name?: pulumi.Input<string>;
    /**
     * WTP, FortiAP, or AP platform. The structure of `platform` block is documented below.
     */
    platform?: pulumi.Input<inputs.WirelessControllerWtpProfilePlatform>;
    /**
     * Set the WTP, FortiAP, or AP's PoE mode.
     */
    poeMode?: pulumi.Input<string>;
    /**
     * Configuration options for radio 1. The structure of `radio1` block is documented below.
     */
    radio1?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio1>;
    /**
     * Configuration options for radio 2. The structure of `radio2` block is documented below.
     */
    radio2?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio2>;
    /**
     * Configuration options for radio 3. The structure of `radio3` block is documented below.
     */
    radio3?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio3>;
    /**
     * Configuration options for radio 4. The structure of `radio4` block is documented below.
     */
    radio4?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio4>;
    /**
     * Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable`, `disable`.
     */
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    /**
     * Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.
     */
    splitTunnelingAclPath?: pulumi.Input<string>;
    /**
     * Split tunneling ACL filter list. The structure of `splitTunnelingAcl` block is documented below.
     */
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileSplitTunnelingAcl>[]>;
    /**
     * Downlink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    tunMtuDownlink?: pulumi.Input<number>;
    /**
     * Uplink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
     */
    tunMtuUplink?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable using a WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.
     */
    wanPortMode?: pulumi.Input<string>;
}