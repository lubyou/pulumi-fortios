// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.
 *
 * ## Import
 *
 * SwitchController StpInstance can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerStpInstance:SwitchControllerStpInstance labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerStpInstance extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerStpInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerStpInstanceState, opts?: pulumi.CustomResourceOptions): SwitchControllerStpInstance {
        return new SwitchControllerStpInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerStpInstance:SwitchControllerStpInstance';

    /**
     * Returns true if the given object is an instance of SwitchControllerStpInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerStpInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerStpInstance.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Instance ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    public readonly vlanRanges!: pulumi.Output<outputs.SwitchControllerStpInstanceVlanRange[] | undefined>;

    /**
     * Create a SwitchControllerStpInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerStpInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerStpInstanceArgs | SwitchControllerStpInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerStpInstanceState | undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlanRanges"] = state ? state.vlanRanges : undefined;
        } else {
            const args = argsOrState as SwitchControllerStpInstanceArgs | undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlanRanges"] = args ? args.vlanRanges : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerStpInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerStpInstance resources.
 */
export interface SwitchControllerStpInstanceState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    vlanRanges?: pulumi.Input<pulumi.Input<inputs.SwitchControllerStpInstanceVlanRange>[]>;
}

/**
 * The set of arguments for constructing a SwitchControllerStpInstance resource.
 */
export interface SwitchControllerStpInstanceArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    vlanRanges?: pulumi.Input<pulumi.Input<inputs.SwitchControllerStpInstanceVlanRange>[]>;
}
