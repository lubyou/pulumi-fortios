// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Designate logical storage for DLP fingerprint database.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.DlpSettings("trname", {
 *     cacheMemPercent: 2,
 *     chunkSize: 2800,
 *     dbMode: "stop-adding",
 *     size: 16,
 * });
 * ```
 *
 * ## Import
 *
 * Dlp Settings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/dlpSettings:DlpSettings labelname DlpSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DlpSettings extends pulumi.CustomResource {
    /**
     * Get an existing DlpSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpSettingsState, opts?: pulumi.CustomResourceOptions): DlpSettings {
        return new DlpSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpSettings:DlpSettings';

    /**
     * Returns true if the given object is an instance of DlpSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpSettings.__pulumiType;
    }

    /**
     * Maximum percentage of available memory allocated to caching (1 - 15%).
     */
    public readonly cacheMemPercent!: pulumi.Output<number>;
    /**
     * Maximum fingerprint chunk size.  **Changing will flush the entire database**.
     */
    public readonly chunkSize!: pulumi.Output<number>;
    /**
     * Behaviour when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.
     */
    public readonly dbMode!: pulumi.Output<string>;
    /**
     * Maximum total size of files within the storage (MB).
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Storage device name.
     */
    public readonly storageDevice!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a DlpSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpSettingsArgs | DlpSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpSettingsState | undefined;
            resourceInputs["cacheMemPercent"] = state ? state.cacheMemPercent : undefined;
            resourceInputs["chunkSize"] = state ? state.chunkSize : undefined;
            resourceInputs["dbMode"] = state ? state.dbMode : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["storageDevice"] = state ? state.storageDevice : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DlpSettingsArgs | undefined;
            resourceInputs["cacheMemPercent"] = args ? args.cacheMemPercent : undefined;
            resourceInputs["chunkSize"] = args ? args.chunkSize : undefined;
            resourceInputs["dbMode"] = args ? args.dbMode : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["storageDevice"] = args ? args.storageDevice : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DlpSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpSettings resources.
 */
export interface DlpSettingsState {
    /**
     * Maximum percentage of available memory allocated to caching (1 - 15%).
     */
    cacheMemPercent?: pulumi.Input<number>;
    /**
     * Maximum fingerprint chunk size.  **Changing will flush the entire database**.
     */
    chunkSize?: pulumi.Input<number>;
    /**
     * Behaviour when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.
     */
    dbMode?: pulumi.Input<string>;
    /**
     * Maximum total size of files within the storage (MB).
     */
    size?: pulumi.Input<number>;
    /**
     * Storage device name.
     */
    storageDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DlpSettings resource.
 */
export interface DlpSettingsArgs {
    /**
     * Maximum percentage of available memory allocated to caching (1 - 15%).
     */
    cacheMemPercent?: pulumi.Input<number>;
    /**
     * Maximum fingerprint chunk size.  **Changing will flush the entire database**.
     */
    chunkSize?: pulumi.Input<number>;
    /**
     * Behaviour when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.
     */
    dbMode?: pulumi.Input<string>;
    /**
     * Maximum total size of files within the storage (MB).
     */
    size?: pulumi.Input<number>;
    /**
     * Storage device name.
     */
    storageDevice?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
