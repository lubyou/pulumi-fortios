// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 access proxy. Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * Firewall AccessProxy6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAccessProxy6:FirewallAccessProxy6 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAccessProxy6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAccessProxy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAccessProxy6State, opts?: pulumi.CustomResourceOptions): FirewallAccessProxy6 {
        return new FirewallAccessProxy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAccessProxy6:FirewallAccessProxy6';

    /**
     * Returns true if the given object is an instance of FirewallAccessProxy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAccessProxy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAccessProxy6.__pulumiType;
    }

    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    public readonly apiGateway6s!: pulumi.Output<outputs.FirewallAccessProxy6ApiGateway6[] | undefined>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    public readonly apiGateways!: pulumi.Output<outputs.FirewallAccessProxy6ApiGateway[] | undefined>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    public readonly authPortal!: pulumi.Output<string>;
    /**
     * Virtual host for authentication portal.
     */
    public readonly authVirtualHost!: pulumi.Output<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    public readonly clientCert!: pulumi.Output<string>;
    /**
     * Decrypted traffic mirror.
     */
    public readonly decryptedTrafficMirror!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Action of an empty client certificate. Valid values: `accept`, `block`.
     */
    public readonly emptyCertAction!: pulumi.Output<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    public readonly logBlockedTraffic!: pulumi.Output<string>;
    /**
     * Server host key name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Virtual IP name.
     */
    public readonly vip!: pulumi.Output<string>;

    /**
     * Create a FirewallAccessProxy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAccessProxy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAccessProxy6Args | FirewallAccessProxy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAccessProxy6State | undefined;
            resourceInputs["apiGateway6s"] = state ? state.apiGateway6s : undefined;
            resourceInputs["apiGateways"] = state ? state.apiGateways : undefined;
            resourceInputs["authPortal"] = state ? state.authPortal : undefined;
            resourceInputs["authVirtualHost"] = state ? state.authVirtualHost : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["decryptedTrafficMirror"] = state ? state.decryptedTrafficMirror : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["emptyCertAction"] = state ? state.emptyCertAction : undefined;
            resourceInputs["logBlockedTraffic"] = state ? state.logBlockedTraffic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
        } else {
            const args = argsOrState as FirewallAccessProxy6Args | undefined;
            resourceInputs["apiGateway6s"] = args ? args.apiGateway6s : undefined;
            resourceInputs["apiGateways"] = args ? args.apiGateways : undefined;
            resourceInputs["authPortal"] = args ? args.authPortal : undefined;
            resourceInputs["authVirtualHost"] = args ? args.authVirtualHost : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["decryptedTrafficMirror"] = args ? args.decryptedTrafficMirror : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["emptyCertAction"] = args ? args.emptyCertAction : undefined;
            resourceInputs["logBlockedTraffic"] = args ? args.logBlockedTraffic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vip"] = args ? args.vip : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAccessProxy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAccessProxy6 resources.
 */
export interface FirewallAccessProxy6State {
    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    apiGateway6s?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxy6ApiGateway6>[]>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    apiGateways?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxy6ApiGateway>[]>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    authPortal?: pulumi.Input<string>;
    /**
     * Virtual host for authentication portal.
     */
    authVirtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror.
     */
    decryptedTrafficMirror?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action of an empty client certificate. Valid values: `accept`, `block`.
     */
    emptyCertAction?: pulumi.Input<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    logBlockedTraffic?: pulumi.Input<string>;
    /**
     * Server host key name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    vip?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAccessProxy6 resource.
 */
export interface FirewallAccessProxy6Args {
    /**
     * Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
     */
    apiGateway6s?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxy6ApiGateway6>[]>;
    /**
     * Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
     */
    apiGateways?: pulumi.Input<pulumi.Input<inputs.FirewallAccessProxy6ApiGateway>[]>;
    /**
     * Enable/disable authentication portal. Valid values: `disable`, `enable`.
     */
    authPortal?: pulumi.Input<string>;
    /**
     * Virtual host for authentication portal.
     */
    authVirtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable to request client certificate. Valid values: `disable`, `enable`.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror.
     */
    decryptedTrafficMirror?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action of an empty client certificate. Valid values: `accept`, `block`.
     */
    emptyCertAction?: pulumi.Input<string>;
    /**
     * Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
     */
    logBlockedTraffic?: pulumi.Input<string>;
    /**
     * Server host key name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual IP name.
     */
    vip?: pulumi.Input<string>;
}
