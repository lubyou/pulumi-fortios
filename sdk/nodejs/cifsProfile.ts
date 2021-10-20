// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure CIFS profile. Applies to FortiOS Version `<= 6.4.0`.
 *
 * ## Import
 *
 * Cifs Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/cifsProfile:CifsProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class CifsProfile extends pulumi.CustomResource {
    /**
     * Get an existing CifsProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CifsProfileState, opts?: pulumi.CustomResourceOptions): CifsProfile {
        return new CifsProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/cifsProfile:CifsProfile';

    /**
     * Returns true if the given object is an instance of CifsProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CifsProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CifsProfile.__pulumiType;
    }

    /**
     * Domain for which to decrypt CIFS traffic.
     */
    public readonly domainController!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    public readonly fileFilter!: pulumi.Output<outputs.CifsProfileFileFilter | undefined>;
    /**
     * File type name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
     */
    public readonly serverCredentialType!: pulumi.Output<string>;
    /**
     * Server keytab. The structure of `serverKeytab` block is documented below.
     */
    public readonly serverKeytabs!: pulumi.Output<outputs.CifsProfileServerKeytab[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a CifsProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CifsProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CifsProfileArgs | CifsProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CifsProfileState | undefined;
            inputs["domainController"] = state ? state.domainController : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["fileFilter"] = state ? state.fileFilter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["serverCredentialType"] = state ? state.serverCredentialType : undefined;
            inputs["serverKeytabs"] = state ? state.serverKeytabs : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CifsProfileArgs | undefined;
            inputs["domainController"] = args ? args.domainController : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["fileFilter"] = args ? args.fileFilter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["serverCredentialType"] = args ? args.serverCredentialType : undefined;
            inputs["serverKeytabs"] = args ? args.serverKeytabs : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CifsProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CifsProfile resources.
 */
export interface CifsProfileState {
    /**
     * Domain for which to decrypt CIFS traffic.
     */
    domainController?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.CifsProfileFileFilter>;
    /**
     * File type name.
     */
    name?: pulumi.Input<string>;
    /**
     * CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
     */
    serverCredentialType?: pulumi.Input<string>;
    /**
     * Server keytab. The structure of `serverKeytab` block is documented below.
     */
    serverKeytabs?: pulumi.Input<pulumi.Input<inputs.CifsProfileServerKeytab>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CifsProfile resource.
 */
export interface CifsProfileArgs {
    /**
     * Domain for which to decrypt CIFS traffic.
     */
    domainController?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.CifsProfileFileFilter>;
    /**
     * File type name.
     */
    name?: pulumi.Input<string>;
    /**
     * CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
     */
    serverCredentialType?: pulumi.Input<string>;
    /**
     * Server keytab. The structure of `serverKeytab` block is documented below.
     */
    serverKeytabs?: pulumi.Input<pulumi.Input<inputs.CifsProfileServerKeytab>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}