// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IpsCustom extends pulumi.CustomResource {
    /**
     * Get an existing IpsCustom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsCustomState, opts?: pulumi.CustomResourceOptions): IpsCustom {
        return new IpsCustom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ipsCustom:IpsCustom';

    /**
     * Returns true if the given object is an instance of IpsCustom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsCustom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsCustom.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly application!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    public readonly log!: pulumi.Output<string>;
    public readonly logPacket!: pulumi.Output<string>;
    public readonly os!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly ruleId!: pulumi.Output<number>;
    public readonly severity!: pulumi.Output<string>;
    public readonly sigName!: pulumi.Output<string>;
    public readonly signature!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly tag!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IpsCustom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpsCustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsCustomArgs | IpsCustomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsCustomState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["log"] = state ? state.log : undefined;
            resourceInputs["logPacket"] = state ? state.logPacket : undefined;
            resourceInputs["os"] = state ? state.os : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["sigName"] = state ? state.sigName : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IpsCustomArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["log"] = args ? args.log : undefined;
            resourceInputs["logPacket"] = args ? args.logPacket : undefined;
            resourceInputs["os"] = args ? args.os : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["sigName"] = args ? args.sigName : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpsCustom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsCustom resources.
 */
export interface IpsCustomState {
    action?: pulumi.Input<string>;
    application?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    log?: pulumi.Input<string>;
    logPacket?: pulumi.Input<string>;
    os?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    ruleId?: pulumi.Input<number>;
    severity?: pulumi.Input<string>;
    sigName?: pulumi.Input<string>;
    signature?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsCustom resource.
 */
export interface IpsCustomArgs {
    action?: pulumi.Input<string>;
    application?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    log?: pulumi.Input<string>;
    logPacket?: pulumi.Input<string>;
    os?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    ruleId?: pulumi.Input<number>;
    severity?: pulumi.Input<string>;
    sigName?: pulumi.Input<string>;
    signature?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
