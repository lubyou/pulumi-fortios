// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure VDOM links.
 *
 * ## Import
 *
 * System VdomLink can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemVdomLink:SystemVdomLink labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemVdomLink extends pulumi.CustomResource {
    /**
     * Get an existing SystemVdomLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVdomLinkState, opts?: pulumi.CustomResourceOptions): SystemVdomLink {
        return new SystemVdomLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVdomLink:SystemVdomLink';

    /**
     * Returns true if the given object is an instance of SystemVdomLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVdomLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVdomLink.__pulumiType;
    }

    /**
     * VDOM link name (maximum = 8 characters).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VDOM link type: PPP or Ethernet. Valid values: `ppp`, `ethernet`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
     */
    public readonly vcluster!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemVdomLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemVdomLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVdomLinkArgs | SystemVdomLinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVdomLinkState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vcluster"] = state ? state.vcluster : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemVdomLinkArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vcluster"] = args ? args.vcluster : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVdomLink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVdomLink resources.
 */
export interface SystemVdomLinkState {
    /**
     * VDOM link name (maximum = 8 characters).
     */
    name?: pulumi.Input<string>;
    /**
     * VDOM link type: PPP or Ethernet. Valid values: `ppp`, `ethernet`.
     */
    type?: pulumi.Input<string>;
    /**
     * Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
     */
    vcluster?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemVdomLink resource.
 */
export interface SystemVdomLinkArgs {
    /**
     * VDOM link name (maximum = 8 characters).
     */
    name?: pulumi.Input<string>;
    /**
     * VDOM link type: PPP or Ethernet. Valid values: `ppp`, `ethernet`.
     */
    type?: pulumi.Input<string>;
    /**
     * Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
     */
    vcluster?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
