// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * SwitchController DynamicPortPolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerDynamicPortPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerDynamicPortPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerDynamicPortPolicyState, opts?: pulumi.CustomResourceOptions): SwitchControllerDynamicPortPolicy {
        return new SwitchControllerDynamicPortPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy';

    /**
     * Returns true if the given object is an instance of SwitchControllerDynamicPortPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerDynamicPortPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerDynamicPortPolicy.__pulumiType;
    }

    /**
     * Description for the policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    public readonly fortilink!: pulumi.Output<string>;
    /**
     * Policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    public readonly policies!: pulumi.Output<outputs.SwitchControllerDynamicPortPolicyPolicy[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerDynamicPortPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerDynamicPortPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerDynamicPortPolicyArgs | SwitchControllerDynamicPortPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerDynamicPortPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerDynamicPortPolicyArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerDynamicPortPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerDynamicPortPolicy resources.
 */
export interface SwitchControllerDynamicPortPolicyState {
    /**
     * Description for the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.SwitchControllerDynamicPortPolicyPolicy>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerDynamicPortPolicy resource.
 */
export interface SwitchControllerDynamicPortPolicyArgs {
    /**
     * Description for the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this Dynamic port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port policies with matching criteria and actions. The structure of `policy` block is documented below.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.SwitchControllerDynamicPortPolicyPolicy>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
