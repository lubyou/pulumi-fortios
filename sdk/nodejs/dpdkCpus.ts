// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure CPUs enabled to run engines in each DPDK stage.
 *
 * ## Import
 *
 * Dpdk Cpus can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/dpdkCpus:DpdkCpus labelname DpdkCpus
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DpdkCpus extends pulumi.CustomResource {
    /**
     * Get an existing DpdkCpus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DpdkCpusState, opts?: pulumi.CustomResourceOptions): DpdkCpus {
        return new DpdkCpus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dpdkCpus:DpdkCpus';

    /**
     * Returns true if the given object is an instance of DpdkCpus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DpdkCpus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DpdkCpus.__pulumiType;
    }

    /**
     * CPUs enabled to run DPDK IPS engines.
     */
    public readonly ipsCpus!: pulumi.Output<string>;
    /**
     * CPUs enabled to run DPDK RX engines.
     */
    public readonly rxCpus!: pulumi.Output<string>;
    /**
     * CPUs enabled to run DPDK TX engines.
     */
    public readonly txCpus!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * CPUs enabled to run DPDK VNP engines.
     */
    public readonly vnpCpus!: pulumi.Output<string>;

    /**
     * Create a DpdkCpus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DpdkCpusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DpdkCpusArgs | DpdkCpusState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DpdkCpusState | undefined;
            inputs["ipsCpus"] = state ? state.ipsCpus : undefined;
            inputs["rxCpus"] = state ? state.rxCpus : undefined;
            inputs["txCpus"] = state ? state.txCpus : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vnpCpus"] = state ? state.vnpCpus : undefined;
        } else {
            const args = argsOrState as DpdkCpusArgs | undefined;
            inputs["ipsCpus"] = args ? args.ipsCpus : undefined;
            inputs["rxCpus"] = args ? args.rxCpus : undefined;
            inputs["txCpus"] = args ? args.txCpus : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vnpCpus"] = args ? args.vnpCpus : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DpdkCpus.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DpdkCpus resources.
 */
export interface DpdkCpusState {
    /**
     * CPUs enabled to run DPDK IPS engines.
     */
    ipsCpus?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK RX engines.
     */
    rxCpus?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK TX engines.
     */
    txCpus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK VNP engines.
     */
    vnpCpus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DpdkCpus resource.
 */
export interface DpdkCpusArgs {
    /**
     * CPUs enabled to run DPDK IPS engines.
     */
    ipsCpus?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK RX engines.
     */
    rxCpus?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK TX engines.
     */
    txCpus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * CPUs enabled to run DPDK VNP engines.
     */
    vnpCpus?: pulumi.Input<string>;
}