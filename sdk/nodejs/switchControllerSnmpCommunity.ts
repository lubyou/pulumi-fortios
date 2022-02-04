// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch SNMP v1/v2c communities globally. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController SnmpCommunity can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerSnmpCommunity:SwitchControllerSnmpCommunity labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerSnmpCommunity extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerSnmpCommunity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerSnmpCommunityState, opts?: pulumi.CustomResourceOptions): SwitchControllerSnmpCommunity {
        return new SwitchControllerSnmpCommunity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerSnmpCommunity:SwitchControllerSnmpCommunity';

    /**
     * Returns true if the given object is an instance of SwitchControllerSnmpCommunity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerSnmpCommunity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerSnmpCommunity.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
     */
    public readonly events!: pulumi.Output<string>;
    /**
     * SNMP community ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    public readonly hosts!: pulumi.Output<outputs.SwitchControllerSnmpCommunityHost[] | undefined>;
    /**
     * SNMP community name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    public readonly queryV1Port!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
     */
    public readonly queryV1Status!: pulumi.Output<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    public readonly queryV2cPort!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
     */
    public readonly queryV2cStatus!: pulumi.Output<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    public readonly trapV1Lport!: pulumi.Output<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    public readonly trapV1Rport!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
     */
    public readonly trapV1Status!: pulumi.Output<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    public readonly trapV2cLport!: pulumi.Output<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    public readonly trapV2cRport!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
     */
    public readonly trapV2cStatus!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerSnmpCommunity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerSnmpCommunityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerSnmpCommunityArgs | SwitchControllerSnmpCommunityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerSnmpCommunityState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["events"] = state ? state.events : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queryV1Port"] = state ? state.queryV1Port : undefined;
            resourceInputs["queryV1Status"] = state ? state.queryV1Status : undefined;
            resourceInputs["queryV2cPort"] = state ? state.queryV2cPort : undefined;
            resourceInputs["queryV2cStatus"] = state ? state.queryV2cStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trapV1Lport"] = state ? state.trapV1Lport : undefined;
            resourceInputs["trapV1Rport"] = state ? state.trapV1Rport : undefined;
            resourceInputs["trapV1Status"] = state ? state.trapV1Status : undefined;
            resourceInputs["trapV2cLport"] = state ? state.trapV2cLport : undefined;
            resourceInputs["trapV2cRport"] = state ? state.trapV2cRport : undefined;
            resourceInputs["trapV2cStatus"] = state ? state.trapV2cStatus : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerSnmpCommunityArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["events"] = args ? args.events : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryV1Port"] = args ? args.queryV1Port : undefined;
            resourceInputs["queryV1Status"] = args ? args.queryV1Status : undefined;
            resourceInputs["queryV2cPort"] = args ? args.queryV2cPort : undefined;
            resourceInputs["queryV2cStatus"] = args ? args.queryV2cStatus : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["trapV1Lport"] = args ? args.trapV1Lport : undefined;
            resourceInputs["trapV1Rport"] = args ? args.trapV1Rport : undefined;
            resourceInputs["trapV1Status"] = args ? args.trapV1Status : undefined;
            resourceInputs["trapV2cLport"] = args ? args.trapV2cLport : undefined;
            resourceInputs["trapV2cRport"] = args ? args.trapV2cRport : undefined;
            resourceInputs["trapV2cStatus"] = args ? args.trapV2cStatus : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerSnmpCommunity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerSnmpCommunity resources.
 */
export interface SwitchControllerSnmpCommunityState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
     */
    events?: pulumi.Input<string>;
    /**
     * SNMP community ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.SwitchControllerSnmpCommunityHost>[]>;
    /**
     * SNMP community name.
     */
    name?: pulumi.Input<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    queryV1Port?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
     */
    queryV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    queryV2cPort?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
     */
    queryV2cStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV1Lport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV1Rport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
     */
    trapV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV2cLport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV2cRport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
     */
    trapV2cStatus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerSnmpCommunity resource.
 */
export interface SwitchControllerSnmpCommunityArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
     */
    events?: pulumi.Input<string>;
    /**
     * SNMP community ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.SwitchControllerSnmpCommunityHost>[]>;
    /**
     * SNMP community name.
     */
    name?: pulumi.Input<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    queryV1Port?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
     */
    queryV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    queryV2cPort?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
     */
    queryV2cStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV1Lport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV1Rport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
     */
    trapV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV2cLport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV2cRport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
     */
    trapV2cStatus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
