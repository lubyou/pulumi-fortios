// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Access Proxy SSH client certificate. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Firewall AccessProxySshClientCert can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessProxySshClientCert:FirewallAccessProxySshClientCert labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessProxySshClientCert:FirewallAccessProxySshClientCert labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAccessProxySshClientCert extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAccessProxySshClientCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAccessProxySshClientCertState, opts?: pulumi.CustomResourceOptions): FirewallAccessProxySshClientCert {
        return new FirewallAccessProxySshClientCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAccessProxySshClientCert:FirewallAccessProxySshClientCert';

    /**
     * Returns true if the given object is an instance of FirewallAccessProxySshClientCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAccessProxySshClientCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAccessProxySshClientCert.__pulumiType;
    }

    /**
     * Name of the SSH server public key authentication CA.
     */
    public readonly authCa!: pulumi.Output<string>;
    /**
     * Configure certificate extension for user certificate. The structure of `certExtension` block is documented below.
     */
    public readonly certExtensions!: pulumi.Output<outputs.FirewallAccessProxySshClientCertCertExtension[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name of certificate extension.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    public readonly permitAgentForwarding!: pulumi.Output<string>;
    /**
     * Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    public readonly permitPortForwarding!: pulumi.Output<string>;
    /**
     * Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
     */
    public readonly permitPty!: pulumi.Output<string>;
    /**
     * Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
     */
    public readonly permitUserRc!: pulumi.Output<string>;
    /**
     * Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    public readonly permitX11Forwarding!: pulumi.Output<string>;
    /**
     * Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
     */
    public readonly sourceAddress!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallAccessProxySshClientCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAccessProxySshClientCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAccessProxySshClientCertArgs | FirewallAccessProxySshClientCertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAccessProxySshClientCertState | undefined;
            resourceInputs["authCa"] = state ? state.authCa : undefined;
            resourceInputs["certExtensions"] = state ? state.certExtensions : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permitAgentForwarding"] = state ? state.permitAgentForwarding : undefined;
            resourceInputs["permitPortForwarding"] = state ? state.permitPortForwarding : undefined;
            resourceInputs["permitPty"] = state ? state.permitPty : undefined;
            resourceInputs["permitUserRc"] = state ? state.permitUserRc : undefined;
            resourceInputs["permitX11Forwarding"] = state ? state.permitX11Forwarding : undefined;
            resourceInputs["sourceAddress"] = state ? state.sourceAddress : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallAccessProxySshClientCertArgs | undefined;
            resourceInputs["authCa"] = args ? args.authCa : undefined;
            resourceInputs["certExtensions"] = args ? args.certExtensions : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permitAgentForwarding"] = args ? args.permitAgentForwarding : undefined;
            resourceInputs["permitPortForwarding"] = args ? args.permitPortForwarding : undefined;
            resourceInputs["permitPty"] = args ? args.permitPty : undefined;
            resourceInputs["permitUserRc"] = args ? args.permitUserRc : undefined;
            resourceInputs["permitX11Forwarding"] = args ? args.permitX11Forwarding : undefined;
            resourceInputs["sourceAddress"] = args ? args.sourceAddress : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAccessProxySshClientCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAccessProxySshClientCert resources.
 */
export interface FirewallAccessProxySshClientCertState {
    /**
     * Name of the SSH server public key authentication CA.
     */
    authCa?: pulumi.Input<string>;
    /**
     * Configure certificate extension for user certificate. The structure of `certExtension` block is documented below.
     */
    certExtensions?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxySshClientCertCertExtension>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of certificate extension.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitAgentForwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitPortForwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
     */
    permitPty?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
     */
    permitUserRc?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitX11Forwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
     */
    sourceAddress?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAccessProxySshClientCert resource.
 */
export interface FirewallAccessProxySshClientCertArgs {
    /**
     * Name of the SSH server public key authentication CA.
     */
    authCa?: pulumi.Input<string>;
    /**
     * Configure certificate extension for user certificate. The structure of `certExtension` block is documented below.
     */
    certExtensions?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxySshClientCertCertExtension>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of certificate extension.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitAgentForwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitPortForwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
     */
    permitPty?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
     */
    permitUserRc?: pulumi.Input<string>;
    /**
     * Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
     */
    permitX11Forwarding?: pulumi.Input<string>;
    /**
     * Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
     */
    sourceAddress?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
