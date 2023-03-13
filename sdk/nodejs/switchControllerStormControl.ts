// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SwitchControllerStormControl extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerStormControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerStormControlState, opts?: pulumi.CustomResourceOptions): SwitchControllerStormControl {
        return new SwitchControllerStormControl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerStormControl:SwitchControllerStormControl';

    /**
     * Returns true if the given object is an instance of SwitchControllerStormControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerStormControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerStormControl.__pulumiType;
    }

    public readonly broadcast!: pulumi.Output<string>;
    public readonly rate!: pulumi.Output<number>;
    public readonly unknownMulticast!: pulumi.Output<string>;
    public readonly unknownUnicast!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerStormControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerStormControlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerStormControlArgs | SwitchControllerStormControlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerStormControlState | undefined;
            resourceInputs["broadcast"] = state ? state.broadcast : undefined;
            resourceInputs["rate"] = state ? state.rate : undefined;
            resourceInputs["unknownMulticast"] = state ? state.unknownMulticast : undefined;
            resourceInputs["unknownUnicast"] = state ? state.unknownUnicast : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerStormControlArgs | undefined;
            resourceInputs["broadcast"] = args ? args.broadcast : undefined;
            resourceInputs["rate"] = args ? args.rate : undefined;
            resourceInputs["unknownMulticast"] = args ? args.unknownMulticast : undefined;
            resourceInputs["unknownUnicast"] = args ? args.unknownUnicast : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerStormControl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerStormControl resources.
 */
export interface SwitchControllerStormControlState {
    broadcast?: pulumi.Input<string>;
    rate?: pulumi.Input<number>;
    unknownMulticast?: pulumi.Input<string>;
    unknownUnicast?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerStormControl resource.
 */
export interface SwitchControllerStormControlArgs {
    broadcast?: pulumi.Input<string>;
    rate?: pulumi.Input<number>;
    unknownMulticast?: pulumi.Input<string>;
    unknownUnicast?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
