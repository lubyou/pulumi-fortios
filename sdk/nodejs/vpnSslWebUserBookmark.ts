// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure SSL VPN user bookmark.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.VpnSslWebUserBookmark("trname", {
 *     customLang: "big5",
 * });
 * ```
 *
 * ## Import
 *
 * VpnSslWeb UserBookmark can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnSslWebUserBookmark:VpnSslWebUserBookmark labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnSslWebUserBookmark:VpnSslWebUserBookmark labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VpnSslWebUserBookmark extends pulumi.CustomResource {
    /**
     * Get an existing VpnSslWebUserBookmark resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnSslWebUserBookmarkState, opts?: pulumi.CustomResourceOptions): VpnSslWebUserBookmark {
        return new VpnSslWebUserBookmark(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnSslWebUserBookmark:VpnSslWebUserBookmark';

    /**
     * Returns true if the given object is an instance of VpnSslWebUserBookmark.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnSslWebUserBookmark {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnSslWebUserBookmark.__pulumiType;
    }

    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    public readonly bookmarks!: pulumi.Output<outputs.VpnSslWebUserBookmarkBookmark[] | undefined>;
    /**
     * Personal language.
     */
    public readonly customLang!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnSslWebUserBookmark resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnSslWebUserBookmarkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnSslWebUserBookmarkArgs | VpnSslWebUserBookmarkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnSslWebUserBookmarkState | undefined;
            resourceInputs["bookmarks"] = state ? state.bookmarks : undefined;
            resourceInputs["customLang"] = state ? state.customLang : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnSslWebUserBookmarkArgs | undefined;
            resourceInputs["bookmarks"] = args ? args.bookmarks : undefined;
            resourceInputs["customLang"] = args ? args.customLang : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnSslWebUserBookmark.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnSslWebUserBookmark resources.
 */
export interface VpnSslWebUserBookmarkState {
    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    bookmarks?: pulumi.Input<pulumi.Input<inputs.VpnSslWebUserBookmarkBookmark>[]>;
    /**
     * Personal language.
     */
    customLang?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnSslWebUserBookmark resource.
 */
export interface VpnSslWebUserBookmarkArgs {
    /**
     * Bookmark table. The structure of `bookmarks` block is documented below.
     */
    bookmarks?: pulumi.Input<pulumi.Input<inputs.VpnSslWebUserBookmarkBookmark>[]>;
    /**
     * Personal language.
     */
    customLang?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
