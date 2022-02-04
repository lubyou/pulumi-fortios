// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Filters for remote system server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.LogSyslogd4Filter("trname", {
 *     anomaly: "enable",
 *     dns: "enable",
 *     filterType: "include",
 *     forwardTraffic: "enable",
 *     gtp: "enable",
 *     localTraffic: "enable",
 *     multicastTraffic: "enable",
 *     severity: "information",
 *     snifferTraffic: "enable",
 *     ssh: "enable",
 *     voip: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * LogSyslogd4 Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/logSyslogd4Filter:LogSyslogd4Filter labelname LogSyslogd4Filter
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogSyslogd4Filter extends pulumi.CustomResource {
    /**
     * Get an existing LogSyslogd4Filter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSyslogd4FilterState, opts?: pulumi.CustomResourceOptions): LogSyslogd4Filter {
        return new LogSyslogd4Filter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logSyslogd4Filter:LogSyslogd4Filter';

    /**
     * Returns true if the given object is an instance of LogSyslogd4Filter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSyslogd4Filter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSyslogd4Filter.__pulumiType;
    }

    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    public readonly anomaly!: pulumi.Output<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    public readonly dns!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Free style filter string.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    public readonly filterType!: pulumi.Output<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly forwardTraffic!: pulumi.Output<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    public readonly freeStyles!: pulumi.Output<outputs.LogSyslogd4FilterFreeStyle[] | undefined>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    public readonly gtp!: pulumi.Output<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly localTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly multicastTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    public readonly netscanDiscovery!: pulumi.Output<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    public readonly netscanVulnerability!: pulumi.Output<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly snifferTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    public readonly ssh!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    public readonly voip!: pulumi.Output<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly ztnaTraffic!: pulumi.Output<string>;

    /**
     * Create a LogSyslogd4Filter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogSyslogd4FilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSyslogd4FilterArgs | LogSyslogd4FilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSyslogd4FilterState | undefined;
            resourceInputs["anomaly"] = state ? state.anomaly : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["filterType"] = state ? state.filterType : undefined;
            resourceInputs["forwardTraffic"] = state ? state.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = state ? state.freeStyles : undefined;
            resourceInputs["gtp"] = state ? state.gtp : undefined;
            resourceInputs["localTraffic"] = state ? state.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = state ? state.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = state ? state.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = state ? state.netscanVulnerability : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["snifferTraffic"] = state ? state.snifferTraffic : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["voip"] = state ? state.voip : undefined;
            resourceInputs["ztnaTraffic"] = state ? state.ztnaTraffic : undefined;
        } else {
            const args = argsOrState as LogSyslogd4FilterArgs | undefined;
            resourceInputs["anomaly"] = args ? args.anomaly : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["filterType"] = args ? args.filterType : undefined;
            resourceInputs["forwardTraffic"] = args ? args.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = args ? args.freeStyles : undefined;
            resourceInputs["gtp"] = args ? args.gtp : undefined;
            resourceInputs["localTraffic"] = args ? args.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = args ? args.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = args ? args.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = args ? args.netscanVulnerability : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["snifferTraffic"] = args ? args.snifferTraffic : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["voip"] = args ? args.voip : undefined;
            resourceInputs["ztnaTraffic"] = args ? args.ztnaTraffic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSyslogd4Filter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSyslogd4Filter resources.
 */
export interface LogSyslogd4FilterState {
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Free style filter string.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.LogSyslogd4FilterFreeStyle>[]>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSyslogd4Filter resource.
 */
export interface LogSyslogd4FilterArgs {
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Free style filter string.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.LogSyslogd4FilterFreeStyle>[]>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}
