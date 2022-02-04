# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SwitchControllerQosQueuePolicyArgs', 'SwitchControllerQosQueuePolicy']

@pulumi.input_type
class SwitchControllerQosQueuePolicyArgs:
    def __init__(__self__, *,
                 rate_by: pulumi.Input[str],
                 schedule: pulumi.Input[str],
                 cos_queues: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerQosQueuePolicy resource.
        :param pulumi.Input[str] rate_by: COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        :param pulumi.Input[str] schedule: COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]] cos_queues: COS queue configuration. The structure of `cos_queue` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Cos queue ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "rate_by", rate_by)
        pulumi.set(__self__, "schedule", schedule)
        if cos_queues is not None:
            pulumi.set(__self__, "cos_queues", cos_queues)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="rateBy")
    def rate_by(self) -> pulumi.Input[str]:
        """
        COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        """
        return pulumi.get(self, "rate_by")

    @rate_by.setter
    def rate_by(self, value: pulumi.Input[str]):
        pulumi.set(self, "rate_by", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="cosQueues")
    def cos_queues(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]]:
        """
        COS queue configuration. The structure of `cos_queue` block is documented below.
        """
        return pulumi.get(self, "cos_queues")

    @cos_queues.setter
    def cos_queues(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]]):
        pulumi.set(self, "cos_queues", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Cos queue ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _SwitchControllerQosQueuePolicyState:
    def __init__(__self__, *,
                 cos_queues: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_by: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerQosQueuePolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]] cos_queues: COS queue configuration. The structure of `cos_queue` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Cos queue ID.
        :param pulumi.Input[str] rate_by: COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        :param pulumi.Input[str] schedule: COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if cos_queues is not None:
            pulumi.set(__self__, "cos_queues", cos_queues)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate_by is not None:
            pulumi.set(__self__, "rate_by", rate_by)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="cosQueues")
    def cos_queues(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]]:
        """
        COS queue configuration. The structure of `cos_queue` block is documented below.
        """
        return pulumi.get(self, "cos_queues")

    @cos_queues.setter
    def cos_queues(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosQueuePolicyCosQueueArgs']]]]):
        pulumi.set(self, "cos_queues", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Cos queue ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rateBy")
    def rate_by(self) -> Optional[pulumi.Input[str]]:
        """
        COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        """
        return pulumi.get(self, "rate_by")

    @rate_by.setter
    def rate_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_by", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

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


class SwitchControllerQosQueuePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cos_queues: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosQueuePolicyCosQueueArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_by: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch QoS egress queue policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchControllerQosQueuePolicy("trname",
            rate_by="kbps",
            schedule="round-robin")
        ```

        ## Import

        SwitchControllerQos QueuePolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosQueuePolicyCosQueueArgs']]]] cos_queues: COS queue configuration. The structure of `cos_queue` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Cos queue ID.
        :param pulumi.Input[str] rate_by: COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        :param pulumi.Input[str] schedule: COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SwitchControllerQosQueuePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch QoS egress queue policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchControllerQosQueuePolicy("trname",
            rate_by="kbps",
            schedule="round-robin")
        ```

        ## Import

        SwitchControllerQos QueuePolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchControllerQosQueuePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerQosQueuePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cos_queues: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosQueuePolicyCosQueueArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_by: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchControllerQosQueuePolicyArgs.__new__(SwitchControllerQosQueuePolicyArgs)

            __props__.__dict__["cos_queues"] = cos_queues
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            if rate_by is None and not opts.urn:
                raise TypeError("Missing required property 'rate_by'")
            __props__.__dict__["rate_by"] = rate_by
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerQosQueuePolicy, __self__).__init__(
            'fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cos_queues: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosQueuePolicyCosQueueArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rate_by: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerQosQueuePolicy':
        """
        Get an existing SwitchControllerQosQueuePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosQueuePolicyCosQueueArgs']]]] cos_queues: COS queue configuration. The structure of `cos_queue` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Cos queue ID.
        :param pulumi.Input[str] rate_by: COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        :param pulumi.Input[str] schedule: COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerQosQueuePolicyState.__new__(_SwitchControllerQosQueuePolicyState)

        __props__.__dict__["cos_queues"] = cos_queues
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["rate_by"] = rate_by
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerQosQueuePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cosQueues")
    def cos_queues(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerQosQueuePolicyCosQueue']]]:
        """
        COS queue configuration. The structure of `cos_queue` block is documented below.
        """
        return pulumi.get(self, "cos_queues")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Cos queue ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rateBy")
    def rate_by(self) -> pulumi.Output[str]:
        """
        COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
        """
        return pulumi.get(self, "rate_by")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

