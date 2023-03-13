// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class UserExchange extends pulumi.CustomResource {
    /**
     * Get an existing UserExchange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserExchangeState, opts?: pulumi.CustomResourceOptions): UserExchange {
        return new UserExchange(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userExchange:UserExchange';

    /**
     * Returns true if the given object is an instance of UserExchange.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserExchange {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserExchange.__pulumiType;
    }

    public readonly authLevel!: pulumi.Output<string>;
    public readonly authType!: pulumi.Output<string>;
    public readonly autoDiscoverKdc!: pulumi.Output<string>;
    public readonly connectProtocol!: pulumi.Output<string>;
    public readonly domainName!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly httpAuthType!: pulumi.Output<string>;
    public readonly ip!: pulumi.Output<string>;
    public readonly kdcIps!: pulumi.Output<outputs.UserExchangeKdcIp[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly serverName!: pulumi.Output<string>;
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserExchange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserExchangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserExchangeArgs | UserExchangeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserExchangeState | undefined;
            resourceInputs["authLevel"] = state ? state.authLevel : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["autoDiscoverKdc"] = state ? state.autoDiscoverKdc : undefined;
            resourceInputs["connectProtocol"] = state ? state.connectProtocol : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["httpAuthType"] = state ? state.httpAuthType : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["kdcIps"] = state ? state.kdcIps : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserExchangeArgs | undefined;
            resourceInputs["authLevel"] = args ? args.authLevel : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["autoDiscoverKdc"] = args ? args.autoDiscoverKdc : undefined;
            resourceInputs["connectProtocol"] = args ? args.connectProtocol : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["httpAuthType"] = args ? args.httpAuthType : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["kdcIps"] = args ? args.kdcIps : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UserExchange.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserExchange resources.
 */
export interface UserExchangeState {
    authLevel?: pulumi.Input<string>;
    authType?: pulumi.Input<string>;
    autoDiscoverKdc?: pulumi.Input<string>;
    connectProtocol?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    httpAuthType?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    kdcIps?: pulumi.Input<pulumi.Input<inputs.UserExchangeKdcIp>[]>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    serverName?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserExchange resource.
 */
export interface UserExchangeArgs {
    authLevel?: pulumi.Input<string>;
    authType?: pulumi.Input<string>;
    autoDiscoverKdc?: pulumi.Input<string>;
    connectProtocol?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    httpAuthType?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    kdcIps?: pulumi.Input<pulumi.Input<inputs.UserExchangeKdcIp>[]>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    serverName?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
