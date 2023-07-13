// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemIke extends pulumi.CustomResource {
    /**
     * Get an existing SystemIke resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemIkeState, opts?: pulumi.CustomResourceOptions): SystemIke {
        return new SystemIke(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemIke:SystemIke';

    /**
     * Returns true if the given object is an instance of SystemIke.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemIke {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemIke.__pulumiType;
    }

    public readonly dhGroup1!: pulumi.Output<outputs.SystemIkeDhGroup1>;
    public readonly dhGroup14!: pulumi.Output<outputs.SystemIkeDhGroup14>;
    public readonly dhGroup15!: pulumi.Output<outputs.SystemIkeDhGroup15>;
    public readonly dhGroup16!: pulumi.Output<outputs.SystemIkeDhGroup16>;
    public readonly dhGroup17!: pulumi.Output<outputs.SystemIkeDhGroup17>;
    public readonly dhGroup18!: pulumi.Output<outputs.SystemIkeDhGroup18>;
    public readonly dhGroup19!: pulumi.Output<outputs.SystemIkeDhGroup19>;
    public readonly dhGroup2!: pulumi.Output<outputs.SystemIkeDhGroup2>;
    public readonly dhGroup20!: pulumi.Output<outputs.SystemIkeDhGroup20>;
    public readonly dhGroup21!: pulumi.Output<outputs.SystemIkeDhGroup21>;
    public readonly dhGroup27!: pulumi.Output<outputs.SystemIkeDhGroup27>;
    public readonly dhGroup28!: pulumi.Output<outputs.SystemIkeDhGroup28>;
    public readonly dhGroup29!: pulumi.Output<outputs.SystemIkeDhGroup29>;
    public readonly dhGroup30!: pulumi.Output<outputs.SystemIkeDhGroup30>;
    public readonly dhGroup31!: pulumi.Output<outputs.SystemIkeDhGroup31>;
    public readonly dhGroup32!: pulumi.Output<outputs.SystemIkeDhGroup32>;
    public readonly dhGroup5!: pulumi.Output<outputs.SystemIkeDhGroup5>;
    public readonly dhKeypairCache!: pulumi.Output<string>;
    public readonly dhKeypairCount!: pulumi.Output<number>;
    public readonly dhKeypairThrottle!: pulumi.Output<string>;
    public readonly dhMode!: pulumi.Output<string>;
    public readonly dhMultiprocess!: pulumi.Output<string>;
    public readonly dhWorkerCount!: pulumi.Output<number>;
    public readonly embryonicLimit!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemIke resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemIkeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemIkeArgs | SystemIkeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemIkeState | undefined;
            resourceInputs["dhGroup1"] = state ? state.dhGroup1 : undefined;
            resourceInputs["dhGroup14"] = state ? state.dhGroup14 : undefined;
            resourceInputs["dhGroup15"] = state ? state.dhGroup15 : undefined;
            resourceInputs["dhGroup16"] = state ? state.dhGroup16 : undefined;
            resourceInputs["dhGroup17"] = state ? state.dhGroup17 : undefined;
            resourceInputs["dhGroup18"] = state ? state.dhGroup18 : undefined;
            resourceInputs["dhGroup19"] = state ? state.dhGroup19 : undefined;
            resourceInputs["dhGroup2"] = state ? state.dhGroup2 : undefined;
            resourceInputs["dhGroup20"] = state ? state.dhGroup20 : undefined;
            resourceInputs["dhGroup21"] = state ? state.dhGroup21 : undefined;
            resourceInputs["dhGroup27"] = state ? state.dhGroup27 : undefined;
            resourceInputs["dhGroup28"] = state ? state.dhGroup28 : undefined;
            resourceInputs["dhGroup29"] = state ? state.dhGroup29 : undefined;
            resourceInputs["dhGroup30"] = state ? state.dhGroup30 : undefined;
            resourceInputs["dhGroup31"] = state ? state.dhGroup31 : undefined;
            resourceInputs["dhGroup32"] = state ? state.dhGroup32 : undefined;
            resourceInputs["dhGroup5"] = state ? state.dhGroup5 : undefined;
            resourceInputs["dhKeypairCache"] = state ? state.dhKeypairCache : undefined;
            resourceInputs["dhKeypairCount"] = state ? state.dhKeypairCount : undefined;
            resourceInputs["dhKeypairThrottle"] = state ? state.dhKeypairThrottle : undefined;
            resourceInputs["dhMode"] = state ? state.dhMode : undefined;
            resourceInputs["dhMultiprocess"] = state ? state.dhMultiprocess : undefined;
            resourceInputs["dhWorkerCount"] = state ? state.dhWorkerCount : undefined;
            resourceInputs["embryonicLimit"] = state ? state.embryonicLimit : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemIkeArgs | undefined;
            resourceInputs["dhGroup1"] = args ? args.dhGroup1 : undefined;
            resourceInputs["dhGroup14"] = args ? args.dhGroup14 : undefined;
            resourceInputs["dhGroup15"] = args ? args.dhGroup15 : undefined;
            resourceInputs["dhGroup16"] = args ? args.dhGroup16 : undefined;
            resourceInputs["dhGroup17"] = args ? args.dhGroup17 : undefined;
            resourceInputs["dhGroup18"] = args ? args.dhGroup18 : undefined;
            resourceInputs["dhGroup19"] = args ? args.dhGroup19 : undefined;
            resourceInputs["dhGroup2"] = args ? args.dhGroup2 : undefined;
            resourceInputs["dhGroup20"] = args ? args.dhGroup20 : undefined;
            resourceInputs["dhGroup21"] = args ? args.dhGroup21 : undefined;
            resourceInputs["dhGroup27"] = args ? args.dhGroup27 : undefined;
            resourceInputs["dhGroup28"] = args ? args.dhGroup28 : undefined;
            resourceInputs["dhGroup29"] = args ? args.dhGroup29 : undefined;
            resourceInputs["dhGroup30"] = args ? args.dhGroup30 : undefined;
            resourceInputs["dhGroup31"] = args ? args.dhGroup31 : undefined;
            resourceInputs["dhGroup32"] = args ? args.dhGroup32 : undefined;
            resourceInputs["dhGroup5"] = args ? args.dhGroup5 : undefined;
            resourceInputs["dhKeypairCache"] = args ? args.dhKeypairCache : undefined;
            resourceInputs["dhKeypairCount"] = args ? args.dhKeypairCount : undefined;
            resourceInputs["dhKeypairThrottle"] = args ? args.dhKeypairThrottle : undefined;
            resourceInputs["dhMode"] = args ? args.dhMode : undefined;
            resourceInputs["dhMultiprocess"] = args ? args.dhMultiprocess : undefined;
            resourceInputs["dhWorkerCount"] = args ? args.dhWorkerCount : undefined;
            resourceInputs["embryonicLimit"] = args ? args.embryonicLimit : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemIke.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemIke resources.
 */
export interface SystemIkeState {
    dhGroup1?: pulumi.Input<inputs.SystemIkeDhGroup1>;
    dhGroup14?: pulumi.Input<inputs.SystemIkeDhGroup14>;
    dhGroup15?: pulumi.Input<inputs.SystemIkeDhGroup15>;
    dhGroup16?: pulumi.Input<inputs.SystemIkeDhGroup16>;
    dhGroup17?: pulumi.Input<inputs.SystemIkeDhGroup17>;
    dhGroup18?: pulumi.Input<inputs.SystemIkeDhGroup18>;
    dhGroup19?: pulumi.Input<inputs.SystemIkeDhGroup19>;
    dhGroup2?: pulumi.Input<inputs.SystemIkeDhGroup2>;
    dhGroup20?: pulumi.Input<inputs.SystemIkeDhGroup20>;
    dhGroup21?: pulumi.Input<inputs.SystemIkeDhGroup21>;
    dhGroup27?: pulumi.Input<inputs.SystemIkeDhGroup27>;
    dhGroup28?: pulumi.Input<inputs.SystemIkeDhGroup28>;
    dhGroup29?: pulumi.Input<inputs.SystemIkeDhGroup29>;
    dhGroup30?: pulumi.Input<inputs.SystemIkeDhGroup30>;
    dhGroup31?: pulumi.Input<inputs.SystemIkeDhGroup31>;
    dhGroup32?: pulumi.Input<inputs.SystemIkeDhGroup32>;
    dhGroup5?: pulumi.Input<inputs.SystemIkeDhGroup5>;
    dhKeypairCache?: pulumi.Input<string>;
    dhKeypairCount?: pulumi.Input<number>;
    dhKeypairThrottle?: pulumi.Input<string>;
    dhMode?: pulumi.Input<string>;
    dhMultiprocess?: pulumi.Input<string>;
    dhWorkerCount?: pulumi.Input<number>;
    embryonicLimit?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemIke resource.
 */
export interface SystemIkeArgs {
    dhGroup1?: pulumi.Input<inputs.SystemIkeDhGroup1>;
    dhGroup14?: pulumi.Input<inputs.SystemIkeDhGroup14>;
    dhGroup15?: pulumi.Input<inputs.SystemIkeDhGroup15>;
    dhGroup16?: pulumi.Input<inputs.SystemIkeDhGroup16>;
    dhGroup17?: pulumi.Input<inputs.SystemIkeDhGroup17>;
    dhGroup18?: pulumi.Input<inputs.SystemIkeDhGroup18>;
    dhGroup19?: pulumi.Input<inputs.SystemIkeDhGroup19>;
    dhGroup2?: pulumi.Input<inputs.SystemIkeDhGroup2>;
    dhGroup20?: pulumi.Input<inputs.SystemIkeDhGroup20>;
    dhGroup21?: pulumi.Input<inputs.SystemIkeDhGroup21>;
    dhGroup27?: pulumi.Input<inputs.SystemIkeDhGroup27>;
    dhGroup28?: pulumi.Input<inputs.SystemIkeDhGroup28>;
    dhGroup29?: pulumi.Input<inputs.SystemIkeDhGroup29>;
    dhGroup30?: pulumi.Input<inputs.SystemIkeDhGroup30>;
    dhGroup31?: pulumi.Input<inputs.SystemIkeDhGroup31>;
    dhGroup32?: pulumi.Input<inputs.SystemIkeDhGroup32>;
    dhGroup5?: pulumi.Input<inputs.SystemIkeDhGroup5>;
    dhKeypairCache?: pulumi.Input<string>;
    dhKeypairCount?: pulumi.Input<number>;
    dhKeypairThrottle?: pulumi.Input<string>;
    dhMode?: pulumi.Input<string>;
    dhMultiprocess?: pulumi.Input<string>;
    dhWorkerCount?: pulumi.Input<number>;
    embryonicLimit?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
