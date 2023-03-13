// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemSnmpMibView extends pulumi.CustomResource {
    /**
     * Get an existing SystemSnmpMibView resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSnmpMibViewState, opts?: pulumi.CustomResourceOptions): SystemSnmpMibView {
        return new SystemSnmpMibView(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSnmpMibView:SystemSnmpMibView';

    /**
     * Returns true if the given object is an instance of SystemSnmpMibView.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSnmpMibView {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSnmpMibView.__pulumiType;
    }

    public readonly exclude!: pulumi.Output<string>;
    public readonly include!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSnmpMibView resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSnmpMibViewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSnmpMibViewArgs | SystemSnmpMibViewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSnmpMibViewState | undefined;
            resourceInputs["exclude"] = state ? state.exclude : undefined;
            resourceInputs["include"] = state ? state.include : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSnmpMibViewArgs | undefined;
            resourceInputs["exclude"] = args ? args.exclude : undefined;
            resourceInputs["include"] = args ? args.include : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSnmpMibView.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSnmpMibView resources.
 */
export interface SystemSnmpMibViewState {
    exclude?: pulumi.Input<string>;
    include?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSnmpMibView resource.
 */
export interface SystemSnmpMibViewArgs {
    exclude?: pulumi.Input<string>;
    include?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
