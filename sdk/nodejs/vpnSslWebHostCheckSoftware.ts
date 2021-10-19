// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * SSL-VPN host check software.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.VpnSslWebHostCheckSoftware("trname", {
 *     osType: "windows",
 *     type: "fw",
 * });
 * ```
 *
 * ## Import
 *
 * VpnSslWeb HostCheckSoftware can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnSslWebHostCheckSoftware:VpnSslWebHostCheckSoftware labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VpnSslWebHostCheckSoftware extends pulumi.CustomResource {
    /**
     * Get an existing VpnSslWebHostCheckSoftware resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnSslWebHostCheckSoftwareState, opts?: pulumi.CustomResourceOptions): VpnSslWebHostCheckSoftware {
        return new VpnSslWebHostCheckSoftware(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnSslWebHostCheckSoftware:VpnSslWebHostCheckSoftware';

    /**
     * Returns true if the given object is an instance of VpnSslWebHostCheckSoftware.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnSslWebHostCheckSoftware {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnSslWebHostCheckSoftware.__pulumiType;
    }

    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    public readonly checkItemLists!: pulumi.Output<outputs.VpnSslWebHostCheckSoftwareCheckItemList[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Globally unique ID.
     */
    public readonly guid!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * Type. Valid values: `file`, `registry`, `process`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Version.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a VpnSslWebHostCheckSoftware resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnSslWebHostCheckSoftwareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnSslWebHostCheckSoftwareArgs | VpnSslWebHostCheckSoftwareState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnSslWebHostCheckSoftwareState | undefined;
            inputs["checkItemLists"] = state ? state.checkItemLists : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["guid"] = state ? state.guid : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as VpnSslWebHostCheckSoftwareArgs | undefined;
            inputs["checkItemLists"] = args ? args.checkItemLists : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["guid"] = args ? args.guid : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpnSslWebHostCheckSoftware.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnSslWebHostCheckSoftware resources.
 */
export interface VpnSslWebHostCheckSoftwareState {
    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    checkItemLists?: pulumi.Input<pulumi.Input<inputs.VpnSslWebHostCheckSoftwareCheckItemList>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Globally unique ID.
     */
    guid?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Type. Valid values: `file`, `registry`, `process`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Version.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnSslWebHostCheckSoftware resource.
 */
export interface VpnSslWebHostCheckSoftwareArgs {
    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    checkItemLists?: pulumi.Input<pulumi.Input<inputs.VpnSslWebHostCheckSoftwareCheckItemList>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Globally unique ID.
     */
    guid?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Type. Valid values: `file`, `registry`, `process`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Version.
     */
    version?: pulumi.Input<string>;
}
