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

__all__ = ['EmailfilterBwlArgs', 'EmailfilterBwl']

@pulumi.input_type
class EmailfilterBwlArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EmailfilterBwl resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]] entries: Anti-spam black/white list entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] name: Name of table.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Optional comments.
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
    @pulumi.getter
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]]:
        """
        Anti-spam black/white list entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of table.
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
class _EmailfilterBwlState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EmailfilterBwl resources.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]] entries: Anti-spam black/white list entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] name: Name of table.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Optional comments.
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
    @pulumi.getter
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]]:
        """
        Anti-spam black/white list entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EmailfilterBwlEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of table.
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


class EmailfilterBwl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EmailfilterBwlEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure anti-spam black/white list. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.2`.

        ## Import

        Emailfilter Bwl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/emailfilterBwl:EmailfilterBwl labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EmailfilterBwlEntryArgs']]]] entries: Anti-spam black/white list entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] name: Name of table.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EmailfilterBwlArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure anti-spam black/white list. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.2`.

        ## Import

        Emailfilter Bwl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/emailfilterBwl:EmailfilterBwl labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param EmailfilterBwlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmailfilterBwlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EmailfilterBwlEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = EmailfilterBwlArgs.__new__(EmailfilterBwlArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["entries"] = entries
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(EmailfilterBwl, __self__).__init__(
            'fortios:index/emailfilterBwl:EmailfilterBwl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EmailfilterBwlEntryArgs']]]]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'EmailfilterBwl':
        """
        Get an existing EmailfilterBwl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EmailfilterBwlEntryArgs']]]] entries: Anti-spam black/white list entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] name: Name of table.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EmailfilterBwlState.__new__(_EmailfilterBwlState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["entries"] = entries
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return EmailfilterBwl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def entries(self) -> pulumi.Output[Optional[Sequence['outputs.EmailfilterBwlEntry']]]:
        """
        Anti-spam black/white list entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of table.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

