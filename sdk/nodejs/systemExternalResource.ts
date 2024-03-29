// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemExternalResource extends pulumi.CustomResource {
    /**
     * Get an existing SystemExternalResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemExternalResourceState, opts?: pulumi.CustomResourceOptions): SystemExternalResource {
        return new SystemExternalResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemExternalResource:SystemExternalResource';

    /**
     * Returns true if the given object is an instance of SystemExternalResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemExternalResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemExternalResource.__pulumiType;
    }

    public readonly category!: pulumi.Output<number>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly refreshRate!: pulumi.Output<number>;
    public readonly resource!: pulumi.Output<string>;
    public readonly serverIdentityCheck!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly updateMethod!: pulumi.Output<string>;
    public readonly userAgent!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemExternalResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemExternalResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemExternalResourceArgs | SystemExternalResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemExternalResourceState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["refreshRate"] = state ? state.refreshRate : undefined;
            resourceInputs["resource"] = state ? state.resource : undefined;
            resourceInputs["serverIdentityCheck"] = state ? state.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updateMethod"] = state ? state.updateMethod : undefined;
            resourceInputs["userAgent"] = state ? state.userAgent : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemExternalResourceArgs | undefined;
            if ((!args || args.refreshRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'refreshRate'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["refreshRate"] = args ? args.refreshRate : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["serverIdentityCheck"] = args ? args.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["updateMethod"] = args ? args.updateMethod : undefined;
            resourceInputs["userAgent"] = args ? args.userAgent : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemExternalResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemExternalResource resources.
 */
export interface SystemExternalResourceState {
    category?: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    refreshRate?: pulumi.Input<number>;
    resource?: pulumi.Input<string>;
    serverIdentityCheck?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    updateMethod?: pulumi.Input<string>;
    userAgent?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemExternalResource resource.
 */
export interface SystemExternalResourceArgs {
    category?: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    refreshRate: pulumi.Input<number>;
    resource: pulumi.Input<string>;
    serverIdentityCheck?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    updateMethod?: pulumi.Input<string>;
    userAgent?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
