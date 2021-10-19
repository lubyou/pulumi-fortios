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

__all__ = ['FirewallAddress6TemplateArgs', 'FirewallAddress6Template']

@pulumi.input_type
class FirewallAddress6TemplateArgs:
    def __init__(__self__, *,
                 ip6: pulumi.Input[str],
                 subnet_segment_count: pulumi.Input[int],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_segments: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallAddress6Template resource.
        :param pulumi.Input[str] ip6: IPv6 address prefix.
        :param pulumi.Input[int] subnet_segment_count: Number of IPv6 subnet segments.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Subnet segment value name.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]] subnet_segments: IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "ip6", ip6)
        pulumi.set(__self__, "subnet_segment_count", subnet_segment_count)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnet_segments is not None:
            pulumi.set(__self__, "subnet_segments", subnet_segments)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ip6(self) -> pulumi.Input[str]:
        """
        IPv6 address prefix.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter(name="subnetSegmentCount")
    def subnet_segment_count(self) -> pulumi.Input[int]:
        """
        Number of IPv6 subnet segments.
        """
        return pulumi.get(self, "subnet_segment_count")

    @subnet_segment_count.setter
    def subnet_segment_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "subnet_segment_count", value)

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
        Subnet segment value name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subnetSegments")
    def subnet_segments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]]:
        """
        IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        """
        return pulumi.get(self, "subnet_segments")

    @subnet_segments.setter
    def subnet_segments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]]):
        pulumi.set(self, "subnet_segments", value)

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
class _FirewallAddress6TemplateState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_segment_count: Optional[pulumi.Input[int]] = None,
                 subnet_segments: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallAddress6Template resources.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] ip6: IPv6 address prefix.
        :param pulumi.Input[str] name: Subnet segment value name.
        :param pulumi.Input[int] subnet_segment_count: Number of IPv6 subnet segments.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]] subnet_segments: IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnet_segment_count is not None:
            pulumi.set(__self__, "subnet_segment_count", subnet_segment_count)
        if subnet_segments is not None:
            pulumi.set(__self__, "subnet_segments", subnet_segments)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address prefix.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet segment value name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subnetSegmentCount")
    def subnet_segment_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of IPv6 subnet segments.
        """
        return pulumi.get(self, "subnet_segment_count")

    @subnet_segment_count.setter
    def subnet_segment_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "subnet_segment_count", value)

    @property
    @pulumi.getter(name="subnetSegments")
    def subnet_segments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]]:
        """
        IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        """
        return pulumi.get(self, "subnet_segments")

    @subnet_segments.setter
    def subnet_segments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddress6TemplateSubnetSegmentArgs']]]]):
        pulumi.set(self, "subnet_segments", value)

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


class FirewallAddress6Template(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_segment_count: Optional[pulumi.Input[int]] = None,
                 subnet_segments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddress6TemplateSubnetSegmentArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 address templates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallAddress6Template("trname",
            ip6="2001:db8:0:b::/64",
            subnet_segments=[
                fortios.FirewallAddress6TemplateSubnetSegmentArgs(
                    bits=4,
                    exclusive="disable",
                    id=1,
                    name="country",
                ),
                fortios.FirewallAddress6TemplateSubnetSegmentArgs(
                    bits=4,
                    exclusive="disable",
                    id=2,
                    name="state",
                ),
            ],
            subnet_segment_count=2)
        ```

        ## Import

        Firewall Address6Template can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallAddress6Template:FirewallAddress6Template labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] ip6: IPv6 address prefix.
        :param pulumi.Input[str] name: Subnet segment value name.
        :param pulumi.Input[int] subnet_segment_count: Number of IPv6 subnet segments.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddress6TemplateSubnetSegmentArgs']]]] subnet_segments: IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallAddress6TemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 address templates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallAddress6Template("trname",
            ip6="2001:db8:0:b::/64",
            subnet_segments=[
                fortios.FirewallAddress6TemplateSubnetSegmentArgs(
                    bits=4,
                    exclusive="disable",
                    id=1,
                    name="country",
                ),
                fortios.FirewallAddress6TemplateSubnetSegmentArgs(
                    bits=4,
                    exclusive="disable",
                    id=2,
                    name="state",
                ),
            ],
            subnet_segment_count=2)
        ```

        ## Import

        Firewall Address6Template can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallAddress6Template:FirewallAddress6Template labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallAddress6TemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallAddress6TemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_segment_count: Optional[pulumi.Input[int]] = None,
                 subnet_segments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddress6TemplateSubnetSegmentArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallAddress6TemplateArgs.__new__(FirewallAddress6TemplateArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if ip6 is None and not opts.urn:
                raise TypeError("Missing required property 'ip6'")
            __props__.__dict__["ip6"] = ip6
            __props__.__dict__["name"] = name
            if subnet_segment_count is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_segment_count'")
            __props__.__dict__["subnet_segment_count"] = subnet_segment_count
            __props__.__dict__["subnet_segments"] = subnet_segments
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallAddress6Template, __self__).__init__(
            'fortios:index/firewallAddress6Template:FirewallAddress6Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            ip6: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            subnet_segment_count: Optional[pulumi.Input[int]] = None,
            subnet_segments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddress6TemplateSubnetSegmentArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallAddress6Template':
        """
        Get an existing FirewallAddress6Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] ip6: IPv6 address prefix.
        :param pulumi.Input[str] name: Subnet segment value name.
        :param pulumi.Input[int] subnet_segment_count: Number of IPv6 subnet segments.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddress6TemplateSubnetSegmentArgs']]]] subnet_segments: IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallAddress6TemplateState.__new__(_FirewallAddress6TemplateState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["ip6"] = ip6
        __props__.__dict__["name"] = name
        __props__.__dict__["subnet_segment_count"] = subnet_segment_count
        __props__.__dict__["subnet_segments"] = subnet_segments
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallAddress6Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def ip6(self) -> pulumi.Output[str]:
        """
        IPv6 address prefix.
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Subnet segment value name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetSegmentCount")
    def subnet_segment_count(self) -> pulumi.Output[int]:
        """
        Number of IPv6 subnet segments.
        """
        return pulumi.get(self, "subnet_segment_count")

    @property
    @pulumi.getter(name="subnetSegments")
    def subnet_segments(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAddress6TemplateSubnetSegment']]]:
        """
        IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        """
        return pulumi.get(self, "subnet_segments")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

