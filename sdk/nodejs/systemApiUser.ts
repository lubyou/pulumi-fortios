// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * System ApiUser can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemApiUser:SystemApiUser labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemApiUser:SystemApiUser labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemApiUser extends pulumi.CustomResource {
    /**
     * Get an existing SystemApiUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemApiUserState, opts?: pulumi.CustomResourceOptions): SystemApiUser {
        return new SystemApiUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemApiUser:SystemApiUser';

    /**
     * Returns true if the given object is an instance of SystemApiUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemApiUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemApiUser.__pulumiType;
    }

    /**
     * Admin user access profile.
     */
    public readonly accprofile!: pulumi.Output<string>;
    /**
     * Admin user password.
     */
    public readonly apiKey!: pulumi.Output<string | undefined>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
     */
    public readonly corsAllowOrigin!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Virtual domain name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable peer authentication. Valid values: `enable`, `disable`.
     */
    public readonly peerAuth!: pulumi.Output<string>;
    /**
     * Peer group name.
     */
    public readonly peerGroup!: pulumi.Output<string>;
    /**
     * Schedule name.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * Trusthost. The structure of `trusthost` block is documented below.
     */
    public readonly trusthosts!: pulumi.Output<outputs.SystemApiUserTrusthost[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Virtual domains. The structure of `vdom` block is documented below.
     */
    public readonly vdoms!: pulumi.Output<outputs.SystemApiUserVdom[] | undefined>;

    /**
     * Create a SystemApiUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemApiUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemApiUserArgs | SystemApiUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemApiUserState | undefined;
            resourceInputs["accprofile"] = state ? state.accprofile : undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["corsAllowOrigin"] = state ? state.corsAllowOrigin : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peerAuth"] = state ? state.peerAuth : undefined;
            resourceInputs["peerGroup"] = state ? state.peerGroup : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["trusthosts"] = state ? state.trusthosts : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vdoms"] = state ? state.vdoms : undefined;
        } else {
            const args = argsOrState as SystemApiUserArgs | undefined;
            if ((!args || args.accprofile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accprofile'");
            }
            resourceInputs["accprofile"] = args ? args.accprofile : undefined;
            resourceInputs["apiKey"] = args ? args.apiKey : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["corsAllowOrigin"] = args ? args.corsAllowOrigin : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peerAuth"] = args ? args.peerAuth : undefined;
            resourceInputs["peerGroup"] = args ? args.peerGroup : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["trusthosts"] = args ? args.trusthosts : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vdoms"] = args ? args.vdoms : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemApiUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemApiUser resources.
 */
export interface SystemApiUserState {
    /**
     * Admin user access profile.
     */
    accprofile?: pulumi.Input<string>;
    /**
     * Admin user password.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
     */
    corsAllowOrigin?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Virtual domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable peer authentication. Valid values: `enable`, `disable`.
     */
    peerAuth?: pulumi.Input<string>;
    /**
     * Peer group name.
     */
    peerGroup?: pulumi.Input<string>;
    /**
     * Schedule name.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Trusthost. The structure of `trusthost` block is documented below.
     */
    trusthosts?: pulumi.Input<pulumi.Input<inputs.SystemApiUserTrusthost>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual domains. The structure of `vdom` block is documented below.
     */
    vdoms?: pulumi.Input<pulumi.Input<inputs.SystemApiUserVdom>[]>;
}

/**
 * The set of arguments for constructing a SystemApiUser resource.
 */
export interface SystemApiUserArgs {
    /**
     * Admin user access profile.
     */
    accprofile: pulumi.Input<string>;
    /**
     * Admin user password.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
     */
    corsAllowOrigin?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Virtual domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable peer authentication. Valid values: `enable`, `disable`.
     */
    peerAuth?: pulumi.Input<string>;
    /**
     * Peer group name.
     */
    peerGroup?: pulumi.Input<string>;
    /**
     * Schedule name.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Trusthost. The structure of `trusthost` block is documented below.
     */
    trusthosts?: pulumi.Input<pulumi.Input<inputs.SystemApiUserTrusthost>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual domains. The structure of `vdom` block is documented below.
     */
    vdoms?: pulumi.Input<pulumi.Input<inputs.SystemApiUserVdom>[]>;
}
