// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch global settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerGlobal("trname", {
 *     allowMultipleInterfaces: "disable",
 *     httpsImagePush: "disable",
 *     logMacLimitViolations: "disable",
 *     macAgingInterval: 332,
 *     macRetentionPeriod: 24,
 *     macViolationTimer: 0,
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController Global can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerGlobal:SwitchControllerGlobal labelname SwitchControllerGlobal
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerGlobal extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerGlobal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerGlobalState, opts?: pulumi.CustomResourceOptions): SwitchControllerGlobal {
        return new SwitchControllerGlobal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerGlobal:SwitchControllerGlobal';

    /**
     * Returns true if the given object is an instance of SwitchControllerGlobal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerGlobal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerGlobal.__pulumiType;
    }

    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly allowMultipleInterfaces!: pulumi.Output<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    public readonly bounceQuarantinedLink!: pulumi.Output<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    public readonly customCommands!: pulumi.Output<outputs.SwitchControllerGlobalCustomCommand[] | undefined>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    public readonly defaultVirtualSwitchVlan!: pulumi.Output<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    public readonly disableDiscoveries!: pulumi.Output<outputs.SwitchControllerGlobalDisableDiscovery[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly httpsImagePush!: pulumi.Output<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    public readonly logMacLimitViolations!: pulumi.Output<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    public readonly macAgingInterval!: pulumi.Output<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    public readonly macEventLogging!: pulumi.Output<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    public readonly macRetentionPeriod!: pulumi.Output<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    public readonly macViolationTimer!: pulumi.Output<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    public readonly quarantineMode!: pulumi.Output<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    public readonly snDnsResolution!: pulumi.Output<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    public readonly updateUserDevice!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    public readonly vlanAllMode!: pulumi.Output<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    public readonly vlanOptimization!: pulumi.Output<string>;

    /**
     * Create a SwitchControllerGlobal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerGlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerGlobalArgs | SwitchControllerGlobalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerGlobalState | undefined;
            inputs["allowMultipleInterfaces"] = state ? state.allowMultipleInterfaces : undefined;
            inputs["bounceQuarantinedLink"] = state ? state.bounceQuarantinedLink : undefined;
            inputs["customCommands"] = state ? state.customCommands : undefined;
            inputs["defaultVirtualSwitchVlan"] = state ? state.defaultVirtualSwitchVlan : undefined;
            inputs["disableDiscoveries"] = state ? state.disableDiscoveries : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["httpsImagePush"] = state ? state.httpsImagePush : undefined;
            inputs["logMacLimitViolations"] = state ? state.logMacLimitViolations : undefined;
            inputs["macAgingInterval"] = state ? state.macAgingInterval : undefined;
            inputs["macEventLogging"] = state ? state.macEventLogging : undefined;
            inputs["macRetentionPeriod"] = state ? state.macRetentionPeriod : undefined;
            inputs["macViolationTimer"] = state ? state.macViolationTimer : undefined;
            inputs["quarantineMode"] = state ? state.quarantineMode : undefined;
            inputs["snDnsResolution"] = state ? state.snDnsResolution : undefined;
            inputs["updateUserDevice"] = state ? state.updateUserDevice : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlanAllMode"] = state ? state.vlanAllMode : undefined;
            inputs["vlanOptimization"] = state ? state.vlanOptimization : undefined;
        } else {
            const args = argsOrState as SwitchControllerGlobalArgs | undefined;
            inputs["allowMultipleInterfaces"] = args ? args.allowMultipleInterfaces : undefined;
            inputs["bounceQuarantinedLink"] = args ? args.bounceQuarantinedLink : undefined;
            inputs["customCommands"] = args ? args.customCommands : undefined;
            inputs["defaultVirtualSwitchVlan"] = args ? args.defaultVirtualSwitchVlan : undefined;
            inputs["disableDiscoveries"] = args ? args.disableDiscoveries : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["httpsImagePush"] = args ? args.httpsImagePush : undefined;
            inputs["logMacLimitViolations"] = args ? args.logMacLimitViolations : undefined;
            inputs["macAgingInterval"] = args ? args.macAgingInterval : undefined;
            inputs["macEventLogging"] = args ? args.macEventLogging : undefined;
            inputs["macRetentionPeriod"] = args ? args.macRetentionPeriod : undefined;
            inputs["macViolationTimer"] = args ? args.macViolationTimer : undefined;
            inputs["quarantineMode"] = args ? args.quarantineMode : undefined;
            inputs["snDnsResolution"] = args ? args.snDnsResolution : undefined;
            inputs["updateUserDevice"] = args ? args.updateUserDevice : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlanAllMode"] = args ? args.vlanAllMode : undefined;
            inputs["vlanOptimization"] = args ? args.vlanOptimization : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerGlobal.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerGlobal resources.
 */
export interface SwitchControllerGlobalState {
    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    allowMultipleInterfaces?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceQuarantinedLink?: pulumi.Input<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchControllerGlobalCustomCommand>[]>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    defaultVirtualSwitchVlan?: pulumi.Input<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    disableDiscoveries?: pulumi.Input<pulumi.Input<inputs.SwitchControllerGlobalDisableDiscovery>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    httpsImagePush?: pulumi.Input<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    logMacLimitViolations?: pulumi.Input<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    macAgingInterval?: pulumi.Input<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    macEventLogging?: pulumi.Input<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    macRetentionPeriod?: pulumi.Input<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    macViolationTimer?: pulumi.Input<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    quarantineMode?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    snDnsResolution?: pulumi.Input<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    updateUserDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    vlanAllMode?: pulumi.Input<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    vlanOptimization?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerGlobal resource.
 */
export interface SwitchControllerGlobalArgs {
    /**
     * Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
     */
    allowMultipleInterfaces?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceQuarantinedLink?: pulumi.Input<string>;
    /**
     * List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `customCommand` block is documented below.
     */
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchControllerGlobalCustomCommand>[]>;
    /**
     * Default VLAN for ports when added to the virtual-switch.
     */
    defaultVirtualSwitchVlan?: pulumi.Input<string>;
    /**
     * Prevent this FortiSwitch from discovering. The structure of `disableDiscovery` block is documented below.
     */
    disableDiscoveries?: pulumi.Input<pulumi.Input<inputs.SwitchControllerGlobalDisableDiscovery>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
     */
    httpsImagePush?: pulumi.Input<string>;
    /**
     * Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
     */
    logMacLimitViolations?: pulumi.Input<string>;
    /**
     * Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
     */
    macAgingInterval?: pulumi.Input<number>;
    /**
     * Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
     */
    macEventLogging?: pulumi.Input<string>;
    /**
     * Time in hours after which an inactive MAC is removed from client DB.
     */
    macRetentionPeriod?: pulumi.Input<number>;
    /**
     * Set timeout for Learning Limit Violations (0 = disabled).
     */
    macViolationTimer?: pulumi.Input<number>;
    /**
     * Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
     */
    quarantineMode?: pulumi.Input<string>;
    /**
     * Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
     */
    snDnsResolution?: pulumi.Input<string>;
    /**
     * Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
     */
    updateUserDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
     */
    vlanAllMode?: pulumi.Input<string>;
    /**
     * FortiLink VLAN optimization. Valid values: `enable`, `disable`.
     */
    vlanOptimization?: pulumi.Input<string>;
}