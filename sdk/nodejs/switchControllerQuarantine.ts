// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch quarantine support.
 *
 * ## Import
 *
 * SwitchController Quarantine can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerQuarantine:SwitchControllerQuarantine labelname SwitchControllerQuarantine
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerQuarantine extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerQuarantine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerQuarantineState, opts?: pulumi.CustomResourceOptions): SwitchControllerQuarantine {
        return new SwitchControllerQuarantine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerQuarantine:SwitchControllerQuarantine';

    /**
     * Returns true if the given object is an instance of SwitchControllerQuarantine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerQuarantine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerQuarantine.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    public readonly quarantine!: pulumi.Output<string>;
    /**
     * Quarantine MACs. The structure of `targets` block is documented below.
     */
    public readonly targets!: pulumi.Output<outputs.SwitchControllerQuarantineTarget[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerQuarantine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerQuarantineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerQuarantineArgs | SwitchControllerQuarantineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerQuarantineState | undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["quarantine"] = state ? state.quarantine : undefined;
            inputs["targets"] = state ? state.targets : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerQuarantineArgs | undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["quarantine"] = args ? args.quarantine : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerQuarantine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerQuarantine resources.
 */
export interface SwitchControllerQuarantineState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    quarantine?: pulumi.Input<string>;
    /**
     * Quarantine MACs. The structure of `targets` block is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.SwitchControllerQuarantineTarget>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerQuarantine resource.
 */
export interface SwitchControllerQuarantineArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    quarantine?: pulumi.Input<string>;
    /**
     * Quarantine MACs. The structure of `targets` block is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.SwitchControllerQuarantineTarget>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
