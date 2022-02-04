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

__all__ = ['FirewallShapingProfileArgs', 'FirewallShapingProfile']

@pulumi.input_type
class FirewallShapingProfileArgs:
    def __init__(__self__, *,
                 default_class_id: pulumi.Input[int],
                 profile_name: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 shaping_entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallShapingProfile resource.
        :param pulumi.Input[int] default_class_id: Default class ID to handle unclassified packets (including all local traffic).
        :param pulumi.Input[str] profile_name: Shaping profile name.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]] shaping_entries: Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        :param pulumi.Input[str] type: Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "default_class_id", default_class_id)
        pulumi.set(__self__, "profile_name", profile_name)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if shaping_entries is not None:
            pulumi.set(__self__, "shaping_entries", shaping_entries)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="defaultClassId")
    def default_class_id(self) -> pulumi.Input[int]:
        """
        Default class ID to handle unclassified packets (including all local traffic).
        """
        return pulumi.get(self, "default_class_id")

    @default_class_id.setter
    def default_class_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "default_class_id", value)

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Input[str]:
        """
        Shaping profile name.
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "profile_name", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

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
    @pulumi.getter(name="shapingEntries")
    def shaping_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]]:
        """
        Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        """
        return pulumi.get(self, "shaping_entries")

    @shaping_entries.setter
    def shaping_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]]):
        pulumi.set(self, "shaping_entries", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
class _FirewallShapingProfileState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_class_id: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 shaping_entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallShapingProfile resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] default_class_id: Default class ID to handle unclassified packets (including all local traffic).
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] profile_name: Shaping profile name.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]] shaping_entries: Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        :param pulumi.Input[str] type: Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_class_id is not None:
            pulumi.set(__self__, "default_class_id", default_class_id)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if profile_name is not None:
            pulumi.set(__self__, "profile_name", profile_name)
        if shaping_entries is not None:
            pulumi.set(__self__, "shaping_entries", shaping_entries)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultClassId")
    def default_class_id(self) -> Optional[pulumi.Input[int]]:
        """
        Default class ID to handle unclassified packets (including all local traffic).
        """
        return pulumi.get(self, "default_class_id")

    @default_class_id.setter
    def default_class_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_class_id", value)

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
    @pulumi.getter(name="profileName")
    def profile_name(self) -> Optional[pulumi.Input[str]]:
        """
        Shaping profile name.
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile_name", value)

    @property
    @pulumi.getter(name="shapingEntries")
    def shaping_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]]:
        """
        Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        """
        return pulumi.get(self, "shaping_entries")

    @shaping_entries.setter
    def shaping_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallShapingProfileShapingEntryArgs']]]]):
        pulumi.set(self, "shaping_entries", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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


class FirewallShapingProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_class_id: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 shaping_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallShapingProfileShapingEntryArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure shaping profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallShapingProfile("trname",
            default_class_id=2,
            profile_name="shapingprofiles1",
            shaping_entries=[fortios.FirewallShapingProfileShapingEntryArgs(
                class_id=2,
                guaranteed_bandwidth_percentage=33,
                id=1,
                maximum_bandwidth_percentage=88,
                priority="high",
            )])
        ```

        ## Import

        Firewall ShapingProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallShapingProfile:FirewallShapingProfile labelname {{profile_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] default_class_id: Default class ID to handle unclassified packets (including all local traffic).
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] profile_name: Shaping profile name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallShapingProfileShapingEntryArgs']]]] shaping_entries: Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        :param pulumi.Input[str] type: Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallShapingProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure shaping profiles.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallShapingProfile("trname",
            default_class_id=2,
            profile_name="shapingprofiles1",
            shaping_entries=[fortios.FirewallShapingProfileShapingEntryArgs(
                class_id=2,
                guaranteed_bandwidth_percentage=33,
                id=1,
                maximum_bandwidth_percentage=88,
                priority="high",
            )])
        ```

        ## Import

        Firewall ShapingProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallShapingProfile:FirewallShapingProfile labelname {{profile_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallShapingProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallShapingProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_class_id: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 shaping_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallShapingProfileShapingEntryArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallShapingProfileArgs.__new__(FirewallShapingProfileArgs)

            __props__.__dict__["comment"] = comment
            if default_class_id is None and not opts.urn:
                raise TypeError("Missing required property 'default_class_id'")
            __props__.__dict__["default_class_id"] = default_class_id
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'profile_name'")
            __props__.__dict__["profile_name"] = profile_name
            __props__.__dict__["shaping_entries"] = shaping_entries
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallShapingProfile, __self__).__init__(
            'fortios:index/firewallShapingProfile:FirewallShapingProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            default_class_id: Optional[pulumi.Input[int]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            profile_name: Optional[pulumi.Input[str]] = None,
            shaping_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallShapingProfileShapingEntryArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallShapingProfile':
        """
        Get an existing FirewallShapingProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] default_class_id: Default class ID to handle unclassified packets (including all local traffic).
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] profile_name: Shaping profile name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallShapingProfileShapingEntryArgs']]]] shaping_entries: Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        :param pulumi.Input[str] type: Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallShapingProfileState.__new__(_FirewallShapingProfileState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["default_class_id"] = default_class_id
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["profile_name"] = profile_name
        __props__.__dict__["shaping_entries"] = shaping_entries
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallShapingProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultClassId")
    def default_class_id(self) -> pulumi.Output[int]:
        """
        Default class ID to handle unclassified packets (including all local traffic).
        """
        return pulumi.get(self, "default_class_id")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Output[str]:
        """
        Shaping profile name.
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="shapingEntries")
    def shaping_entries(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallShapingProfileShapingEntry']]]:
        """
        Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
        """
        return pulumi.get(self, "shaping_entries")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

