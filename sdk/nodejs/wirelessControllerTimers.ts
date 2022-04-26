// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure CAPWAP timers.
 *
 * ## Import
 *
 * WirelessController Timers can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerTimers:WirelessControllerTimers labelname WirelessControllerTimers
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerTimers:WirelessControllerTimers labelname WirelessControllerTimers
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerTimers extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerTimers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerTimersState, opts?: pulumi.CustomResourceOptions): WirelessControllerTimers {
        return new WirelessControllerTimers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerTimers:WirelessControllerTimers';

    /**
     * Returns true if the given object is an instance of WirelessControllerTimers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerTimers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerTimers.__pulumiType;
    }

    /**
     * Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
     */
    public readonly bleScanReportIntv!: pulumi.Output<number>;
    /**
     * Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
     */
    public readonly clientIdleTimeout!: pulumi.Output<number>;
    /**
     * Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    public readonly darrpDay!: pulumi.Output<string>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
     */
    public readonly darrpOptimize!: pulumi.Output<number>;
    /**
     * Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
     */
    public readonly darrpTimes!: pulumi.Output<outputs.WirelessControllerTimersDarrpTime[] | undefined>;
    /**
     * Time between discovery requests (2 - 180 sec, default = 5).
     */
    public readonly discoveryInterval!: pulumi.Output<number>;
    /**
     * Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
     */
    public readonly drmaInterval!: pulumi.Output<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
     */
    public readonly echoInterval!: pulumi.Output<number>;
    /**
     * Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
     */
    public readonly fakeApLog!: pulumi.Output<number>;
    /**
     * Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
     */
    public readonly ipsecIntfCleanup!: pulumi.Output<number>;
    /**
     * Time between running radio reports (1 - 255 sec, default = 15).
     */
    public readonly radioStatsInterval!: pulumi.Output<number>;
    /**
     * Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
     */
    public readonly rogueApLog!: pulumi.Output<number>;
    /**
     * Time between running station capability reports (1 - 255 sec, default = 30).
     */
    public readonly staCapabilityInterval!: pulumi.Output<number>;
    /**
     * Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
     */
    public readonly staLocateTimer!: pulumi.Output<number>;
    /**
     * Time between running client (station) reports (1 - 255 sec, default = 1).
     */
    public readonly staStatsInterval!: pulumi.Output<number>;
    /**
     * Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
     */
    public readonly vapStatsInterval!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerTimers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerTimersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerTimersArgs | WirelessControllerTimersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerTimersState | undefined;
            resourceInputs["bleScanReportIntv"] = state ? state.bleScanReportIntv : undefined;
            resourceInputs["clientIdleTimeout"] = state ? state.clientIdleTimeout : undefined;
            resourceInputs["darrpDay"] = state ? state.darrpDay : undefined;
            resourceInputs["darrpOptimize"] = state ? state.darrpOptimize : undefined;
            resourceInputs["darrpTimes"] = state ? state.darrpTimes : undefined;
            resourceInputs["discoveryInterval"] = state ? state.discoveryInterval : undefined;
            resourceInputs["drmaInterval"] = state ? state.drmaInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["echoInterval"] = state ? state.echoInterval : undefined;
            resourceInputs["fakeApLog"] = state ? state.fakeApLog : undefined;
            resourceInputs["ipsecIntfCleanup"] = state ? state.ipsecIntfCleanup : undefined;
            resourceInputs["radioStatsInterval"] = state ? state.radioStatsInterval : undefined;
            resourceInputs["rogueApLog"] = state ? state.rogueApLog : undefined;
            resourceInputs["staCapabilityInterval"] = state ? state.staCapabilityInterval : undefined;
            resourceInputs["staLocateTimer"] = state ? state.staLocateTimer : undefined;
            resourceInputs["staStatsInterval"] = state ? state.staStatsInterval : undefined;
            resourceInputs["vapStatsInterval"] = state ? state.vapStatsInterval : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerTimersArgs | undefined;
            resourceInputs["bleScanReportIntv"] = args ? args.bleScanReportIntv : undefined;
            resourceInputs["clientIdleTimeout"] = args ? args.clientIdleTimeout : undefined;
            resourceInputs["darrpDay"] = args ? args.darrpDay : undefined;
            resourceInputs["darrpOptimize"] = args ? args.darrpOptimize : undefined;
            resourceInputs["darrpTimes"] = args ? args.darrpTimes : undefined;
            resourceInputs["discoveryInterval"] = args ? args.discoveryInterval : undefined;
            resourceInputs["drmaInterval"] = args ? args.drmaInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["echoInterval"] = args ? args.echoInterval : undefined;
            resourceInputs["fakeApLog"] = args ? args.fakeApLog : undefined;
            resourceInputs["ipsecIntfCleanup"] = args ? args.ipsecIntfCleanup : undefined;
            resourceInputs["radioStatsInterval"] = args ? args.radioStatsInterval : undefined;
            resourceInputs["rogueApLog"] = args ? args.rogueApLog : undefined;
            resourceInputs["staCapabilityInterval"] = args ? args.staCapabilityInterval : undefined;
            resourceInputs["staLocateTimer"] = args ? args.staLocateTimer : undefined;
            resourceInputs["staStatsInterval"] = args ? args.staStatsInterval : undefined;
            resourceInputs["vapStatsInterval"] = args ? args.vapStatsInterval : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerTimers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerTimers resources.
 */
export interface WirelessControllerTimersState {
    /**
     * Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
     */
    bleScanReportIntv?: pulumi.Input<number>;
    /**
     * Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
     */
    clientIdleTimeout?: pulumi.Input<number>;
    /**
     * Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    darrpDay?: pulumi.Input<string>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
     */
    darrpOptimize?: pulumi.Input<number>;
    /**
     * Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
     */
    darrpTimes?: pulumi.Input<pulumi.Input<inputs.WirelessControllerTimersDarrpTime>[]>;
    /**
     * Time between discovery requests (2 - 180 sec, default = 5).
     */
    discoveryInterval?: pulumi.Input<number>;
    /**
     * Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
     */
    drmaInterval?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
     */
    echoInterval?: pulumi.Input<number>;
    /**
     * Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
     */
    fakeApLog?: pulumi.Input<number>;
    /**
     * Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
     */
    ipsecIntfCleanup?: pulumi.Input<number>;
    /**
     * Time between running radio reports (1 - 255 sec, default = 15).
     */
    radioStatsInterval?: pulumi.Input<number>;
    /**
     * Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
     */
    rogueApLog?: pulumi.Input<number>;
    /**
     * Time between running station capability reports (1 - 255 sec, default = 30).
     */
    staCapabilityInterval?: pulumi.Input<number>;
    /**
     * Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
     */
    staLocateTimer?: pulumi.Input<number>;
    /**
     * Time between running client (station) reports (1 - 255 sec, default = 1).
     */
    staStatsInterval?: pulumi.Input<number>;
    /**
     * Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
     */
    vapStatsInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerTimers resource.
 */
export interface WirelessControllerTimersArgs {
    /**
     * Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
     */
    bleScanReportIntv?: pulumi.Input<number>;
    /**
     * Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
     */
    clientIdleTimeout?: pulumi.Input<number>;
    /**
     * Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    darrpDay?: pulumi.Input<string>;
    /**
     * Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
     */
    darrpOptimize?: pulumi.Input<number>;
    /**
     * Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
     */
    darrpTimes?: pulumi.Input<pulumi.Input<inputs.WirelessControllerTimersDarrpTime>[]>;
    /**
     * Time between discovery requests (2 - 180 sec, default = 5).
     */
    discoveryInterval?: pulumi.Input<number>;
    /**
     * Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
     */
    drmaInterval?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
     */
    echoInterval?: pulumi.Input<number>;
    /**
     * Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
     */
    fakeApLog?: pulumi.Input<number>;
    /**
     * Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
     */
    ipsecIntfCleanup?: pulumi.Input<number>;
    /**
     * Time between running radio reports (1 - 255 sec, default = 15).
     */
    radioStatsInterval?: pulumi.Input<number>;
    /**
     * Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
     */
    rogueApLog?: pulumi.Input<number>;
    /**
     * Time between running station capability reports (1 - 255 sec, default = 30).
     */
    staCapabilityInterval?: pulumi.Input<number>;
    /**
     * Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
     */
    staLocateTimer?: pulumi.Input<number>;
    /**
     * Time between running client (station) reports (1 - 255 sec, default = 1).
     */
    staStatsInterval?: pulumi.Input<number>;
    /**
     * Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
     */
    vapStatsInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
