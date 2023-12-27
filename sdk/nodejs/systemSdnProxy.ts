// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemSdnProxy extends pulumi.CustomResource {
    /**
     * Get an existing SystemSdnProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSdnProxyState, opts?: pulumi.CustomResourceOptions): SystemSdnProxy {
        return new SystemSdnProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSdnProxy:SystemSdnProxy';

    /**
     * Returns true if the given object is an instance of SystemSdnProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSdnProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSdnProxy.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly server!: pulumi.Output<string>;
    public readonly serverPort!: pulumi.Output<number>;
    public readonly type!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSdnProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSdnProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSdnProxyArgs | SystemSdnProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSdnProxyState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["serverPort"] = state ? state.serverPort : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSdnProxyArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["serverPort"] = args ? args.serverPort : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSdnProxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSdnProxy resources.
 */
export interface SystemSdnProxyState {
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    serverPort?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSdnProxy resource.
 */
export interface SystemSdnProxyArgs {
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    serverPort?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}