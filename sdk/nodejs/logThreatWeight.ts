// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure threat weight settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.LogThreatWeight("trname", {
 *     applications: [
 *         {
 *             category: 2,
 *             id: 1,
 *             level: "low",
 *         },
 *         {
 *             category: 6,
 *             id: 2,
 *             level: "medium",
 *         },
 *     ],
 *     blockedConnection: "high",
 *     failedConnection: "low",
 *     ips: {
 *         criticalSeverity: "critical",
 *         highSeverity: "high",
 *         infoSeverity: "disable",
 *         lowSeverity: "low",
 *         mediumSeverity: "medium",
 *     },
 *     level: {
 *         critical: 50,
 *         high: 30,
 *         low: 5,
 *         medium: 10,
 *     },
 *     malware: {
 *         botnetConnection: "critical",
 *         commandBlocked: "disable",
 *         contentDisarm: "medium",
 *         fileBlocked: "low",
 *         mimefragmented: "disable",
 *         oversized: "disable",
 *         switchProto: "disable",
 *         virusFileTypeExecutable: "medium",
 *         virusInfected: "critical",
 *         virusOutbreakPrevention: "critical",
 *         virusScanError: "high",
 *     },
 *     status: "enable",
 *     urlBlockDetected: "high",
 *     webs: [
 *         {
 *             category: 26,
 *             id: 1,
 *             level: "high",
 *         },
 *         {
 *             category: 61,
 *             id: 2,
 *             level: "high",
 *         },
 *         {
 *             category: 86,
 *             id: 3,
 *             level: "high",
 *         },
 *         {
 *             category: 1,
 *             id: 4,
 *             level: "medium",
 *         },
 *         {
 *             category: 3,
 *             id: 5,
 *             level: "medium",
 *         },
 *         {
 *             category: 4,
 *             id: 6,
 *             level: "medium",
 *         },
 *         {
 *             category: 5,
 *             id: 7,
 *             level: "medium",
 *         },
 *         {
 *             category: 6,
 *             id: 8,
 *             level: "medium",
 *         },
 *         {
 *             category: 12,
 *             id: 9,
 *             level: "medium",
 *         },
 *         {
 *             category: 59,
 *             id: 10,
 *             level: "medium",
 *         },
 *         {
 *             category: 62,
 *             id: 11,
 *             level: "medium",
 *         },
 *         {
 *             category: 83,
 *             id: 12,
 *             level: "medium",
 *         },
 *         {
 *             category: 72,
 *             id: 13,
 *             level: "low",
 *         },
 *         {
 *             category: 14,
 *             id: 14,
 *             level: "low",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Log ThreatWeight can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/logThreatWeight:LogThreatWeight labelname LogThreatWeight
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogThreatWeight extends pulumi.CustomResource {
    /**
     * Get an existing LogThreatWeight resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogThreatWeightState, opts?: pulumi.CustomResourceOptions): LogThreatWeight {
        return new LogThreatWeight(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logThreatWeight:LogThreatWeight';

    /**
     * Returns true if the given object is an instance of LogThreatWeight.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogThreatWeight {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogThreatWeight.__pulumiType;
    }

    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    public readonly applications!: pulumi.Output<outputs.LogThreatWeightApplication[] | undefined>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly blockedConnection!: pulumi.Output<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly botnetConnectionDetected!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly failedConnection!: pulumi.Output<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    public readonly geolocations!: pulumi.Output<outputs.LogThreatWeightGeolocation[] | undefined>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    public readonly ips!: pulumi.Output<outputs.LogThreatWeightIps | undefined>;
    /**
     * Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly level!: pulumi.Output<outputs.LogThreatWeightLevel | undefined>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    public readonly malware!: pulumi.Output<outputs.LogThreatWeightMalware | undefined>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    public readonly urlBlockDetected!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    public readonly webs!: pulumi.Output<outputs.LogThreatWeightWeb[] | undefined>;

    /**
     * Create a LogThreatWeight resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogThreatWeightArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogThreatWeightArgs | LogThreatWeightState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogThreatWeightState | undefined;
            inputs["applications"] = state ? state.applications : undefined;
            inputs["blockedConnection"] = state ? state.blockedConnection : undefined;
            inputs["botnetConnectionDetected"] = state ? state.botnetConnectionDetected : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["failedConnection"] = state ? state.failedConnection : undefined;
            inputs["geolocations"] = state ? state.geolocations : undefined;
            inputs["ips"] = state ? state.ips : undefined;
            inputs["level"] = state ? state.level : undefined;
            inputs["malware"] = state ? state.malware : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["urlBlockDetected"] = state ? state.urlBlockDetected : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["webs"] = state ? state.webs : undefined;
        } else {
            const args = argsOrState as LogThreatWeightArgs | undefined;
            inputs["applications"] = args ? args.applications : undefined;
            inputs["blockedConnection"] = args ? args.blockedConnection : undefined;
            inputs["botnetConnectionDetected"] = args ? args.botnetConnectionDetected : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["failedConnection"] = args ? args.failedConnection : undefined;
            inputs["geolocations"] = args ? args.geolocations : undefined;
            inputs["ips"] = args ? args.ips : undefined;
            inputs["level"] = args ? args.level : undefined;
            inputs["malware"] = args ? args.malware : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["urlBlockDetected"] = args ? args.urlBlockDetected : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["webs"] = args ? args.webs : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LogThreatWeight.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogThreatWeight resources.
 */
export interface LogThreatWeightState {
    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightApplication>[]>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    blockedConnection?: pulumi.Input<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    botnetConnectionDetected?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    failedConnection?: pulumi.Input<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    geolocations?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightGeolocation>[]>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    ips?: pulumi.Input<inputs.LogThreatWeightIps>;
    /**
     * Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    level?: pulumi.Input<inputs.LogThreatWeightLevel>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    malware?: pulumi.Input<inputs.LogThreatWeightMalware>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    urlBlockDetected?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    webs?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightWeb>[]>;
}

/**
 * The set of arguments for constructing a LogThreatWeight resource.
 */
export interface LogThreatWeightArgs {
    /**
     * Application-control threat weight settings. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightApplication>[]>;
    /**
     * Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    blockedConnection?: pulumi.Input<string>;
    /**
     * Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    botnetConnectionDetected?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    failedConnection?: pulumi.Input<string>;
    /**
     * Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
     */
    geolocations?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightGeolocation>[]>;
    /**
     * IPS threat weight settings. The structure of `ips` block is documented below.
     */
    ips?: pulumi.Input<inputs.LogThreatWeightIps>;
    /**
     * Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    level?: pulumi.Input<inputs.LogThreatWeightLevel>;
    /**
     * Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
     */
    malware?: pulumi.Input<inputs.LogThreatWeightMalware>;
    /**
     * Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
     */
    urlBlockDetected?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web filtering threat weight settings. The structure of `web` block is documented below.
     */
    webs?: pulumi.Input<pulumi.Input<inputs.LogThreatWeightWeb>[]>;
}
