# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemAffinityInterruptArgs', 'SystemAffinityInterrupt']

@pulumi.input_type
class SystemAffinityInterruptArgs:
    def __init__(__self__, *,
                 affinity_cpumask: pulumi.Input[str],
                 fosid: pulumi.Input[int],
                 interrupt: pulumi.Input[str],
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAffinityInterrupt resource.
        :param pulumi.Input[str] affinity_cpumask: Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        :param pulumi.Input[int] fosid: ID of the interrupt affinity setting.
        :param pulumi.Input[str] interrupt: Interrupt name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "affinity_cpumask", affinity_cpumask)
        pulumi.set(__self__, "fosid", fosid)
        pulumi.set(__self__, "interrupt", interrupt)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="affinityCpumask")
    def affinity_cpumask(self) -> pulumi.Input[str]:
        """
        Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        """
        return pulumi.get(self, "affinity_cpumask")

    @affinity_cpumask.setter
    def affinity_cpumask(self, value: pulumi.Input[str]):
        pulumi.set(self, "affinity_cpumask", value)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Input[int]:
        """
        ID of the interrupt affinity setting.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: pulumi.Input[int]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interrupt(self) -> pulumi.Input[str]:
        """
        Interrupt name.
        """
        return pulumi.get(self, "interrupt")

    @interrupt.setter
    def interrupt(self, value: pulumi.Input[str]):
        pulumi.set(self, "interrupt", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemAffinityInterruptState:
    def __init__(__self__, *,
                 affinity_cpumask: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interrupt: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemAffinityInterrupt resources.
        :param pulumi.Input[str] affinity_cpumask: Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        :param pulumi.Input[int] fosid: ID of the interrupt affinity setting.
        :param pulumi.Input[str] interrupt: Interrupt name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if affinity_cpumask is not None:
            pulumi.set(__self__, "affinity_cpumask", affinity_cpumask)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if interrupt is not None:
            pulumi.set(__self__, "interrupt", interrupt)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="affinityCpumask")
    def affinity_cpumask(self) -> Optional[pulumi.Input[str]]:
        """
        Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        """
        return pulumi.get(self, "affinity_cpumask")

    @affinity_cpumask.setter
    def affinity_cpumask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "affinity_cpumask", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID of the interrupt affinity setting.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interrupt(self) -> Optional[pulumi.Input[str]]:
        """
        Interrupt name.
        """
        return pulumi.get(self, "interrupt")

    @interrupt.setter
    def interrupt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interrupt", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemAffinityInterrupt(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 affinity_cpumask: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interrupt: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure interrupt affinity.

        ## Import

        System AffinityInterrupt can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemAffinityInterrupt:SystemAffinityInterrupt labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemAffinityInterrupt:SystemAffinityInterrupt labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] affinity_cpumask: Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        :param pulumi.Input[int] fosid: ID of the interrupt affinity setting.
        :param pulumi.Input[str] interrupt: Interrupt name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemAffinityInterruptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure interrupt affinity.

        ## Import

        System AffinityInterrupt can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemAffinityInterrupt:SystemAffinityInterrupt labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemAffinityInterrupt:SystemAffinityInterrupt labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemAffinityInterruptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAffinityInterruptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 affinity_cpumask: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interrupt: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemAffinityInterruptArgs.__new__(SystemAffinityInterruptArgs)

            if affinity_cpumask is None and not opts.urn:
                raise TypeError("Missing required property 'affinity_cpumask'")
            __props__.__dict__["affinity_cpumask"] = affinity_cpumask
            if fosid is None and not opts.urn:
                raise TypeError("Missing required property 'fosid'")
            __props__.__dict__["fosid"] = fosid
            if interrupt is None and not opts.urn:
                raise TypeError("Missing required property 'interrupt'")
            __props__.__dict__["interrupt"] = interrupt
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemAffinityInterrupt, __self__).__init__(
            'fortios:index/systemAffinityInterrupt:SystemAffinityInterrupt',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            affinity_cpumask: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            interrupt: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemAffinityInterrupt':
        """
        Get an existing SystemAffinityInterrupt resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] affinity_cpumask: Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        :param pulumi.Input[int] fosid: ID of the interrupt affinity setting.
        :param pulumi.Input[str] interrupt: Interrupt name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAffinityInterruptState.__new__(_SystemAffinityInterruptState)

        __props__.__dict__["affinity_cpumask"] = affinity_cpumask
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["interrupt"] = interrupt
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemAffinityInterrupt(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="affinityCpumask")
    def affinity_cpumask(self) -> pulumi.Output[str]:
        """
        Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        """
        return pulumi.get(self, "affinity_cpumask")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID of the interrupt affinity setting.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def interrupt(self) -> pulumi.Output[str]:
        """
        Interrupt name.
        """
        return pulumi.get(self, "interrupt")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

