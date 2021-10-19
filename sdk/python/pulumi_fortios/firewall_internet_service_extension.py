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

__all__ = ['FirewallInternetServiceExtensionArgs', 'FirewallInternetServiceExtension']

@pulumi.input_type
class FirewallInternetServiceExtensionArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallInternetServiceExtension resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]] disable_entries: Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]] entries: Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        :param pulumi.Input[int] fosid: Internet Service ID in the Internet Service database.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if disable_entries is not None:
            pulumi.set(__self__, "disable_entries", disable_entries)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
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
    @pulumi.getter(name="disableEntries")
    def disable_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]]:
        """
        Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        """
        return pulumi.get(self, "disable_entries")

    @disable_entries.setter
    def disable_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]]):
        pulumi.set(self, "disable_entries", value)

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
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]]:
        """
        Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID in the Internet Service database.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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
class _FirewallInternetServiceExtensionState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallInternetServiceExtension resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]] disable_entries: Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]] entries: Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        :param pulumi.Input[int] fosid: Internet Service ID in the Internet Service database.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if disable_entries is not None:
            pulumi.set(__self__, "disable_entries", disable_entries)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
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
    @pulumi.getter(name="disableEntries")
    def disable_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]]:
        """
        Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        """
        return pulumi.get(self, "disable_entries")

    @disable_entries.setter
    def disable_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionDisableEntryArgs']]]]):
        pulumi.set(self, "disable_entries", value)

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
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]]:
        """
        Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallInternetServiceExtensionEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID in the Internet Service database.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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


class FirewallInternetServiceExtension(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionDisableEntryArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Internet Services Extension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallInternetServiceExtension("trname",
            comment="EIWE",
            fosid=65536)
        ```

        ## Import

        Firewall InternetServiceExtension can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionDisableEntryArgs']]]] disable_entries: Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionEntryArgs']]]] entries: Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        :param pulumi.Input[int] fosid: Internet Service ID in the Internet Service database.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallInternetServiceExtensionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Internet Services Extension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallInternetServiceExtension("trname",
            comment="EIWE",
            fosid=65536)
        ```

        ## Import

        Firewall InternetServiceExtension can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallInternetServiceExtensionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallInternetServiceExtensionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionDisableEntryArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
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
            __props__ = FirewallInternetServiceExtensionArgs.__new__(FirewallInternetServiceExtensionArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["disable_entries"] = disable_entries
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["entries"] = entries
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallInternetServiceExtension, __self__).__init__(
            'fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            disable_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionDisableEntryArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionEntryArgs']]]]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallInternetServiceExtension':
        """
        Get an existing FirewallInternetServiceExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionDisableEntryArgs']]]] disable_entries: Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallInternetServiceExtensionEntryArgs']]]] entries: Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        :param pulumi.Input[int] fosid: Internet Service ID in the Internet Service database.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallInternetServiceExtensionState.__new__(_FirewallInternetServiceExtensionState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["disable_entries"] = disable_entries
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["entries"] = entries
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallInternetServiceExtension(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="disableEntries")
    def disable_entries(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallInternetServiceExtensionDisableEntry']]]:
        """
        Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        """
        return pulumi.get(self, "disable_entries")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def entries(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallInternetServiceExtensionEntry']]]:
        """
        Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Internet Service ID in the Internet Service database.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

