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

__all__ = ['RouterAccessListArgs', 'RouterAccessList']

@pulumi.input_type
class RouterAccessListArgs:
    def __init__(__self__, *,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterAccessList resource.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]] rules: Rule. The structure of `rule` block is documented below.
        """
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]]:
        """
        Rule. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _RouterAccessListState:
    def __init__(__self__, *,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouterAccessList resources.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]] rules: Rule. The structure of `rule` block is documented below.
        """
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]]:
        """
        Rule. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterAccessListRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class RouterAccessList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterAccessListRuleArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure access lists.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.RouterAccessList("trname", comments="test accesslist")
        ```
        ## Note

        The feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.

        ### Example
        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.SystemAutoScript("trname1",
            interval=1,
            output_size=10,
            repeat=1,
            script=\"\"\"config router access-list
        edit "static-redistribution"
        config rule
        edit 10
        set prefix 10.0.0.0 255.255.255.0
        set action permit
        set exact-match enable
        end
        end

        \"\"\",
            start="auto")
        ```

        ## Import

        Router AccessList can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/routerAccessList:RouterAccessList labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerAccessList:RouterAccessList labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterAccessListRuleArgs']]]] rules: Rule. The structure of `rule` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterAccessListArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure access lists.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.RouterAccessList("trname", comments="test accesslist")
        ```
        ## Note

        The feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.

        ### Example
        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.SystemAutoScript("trname1",
            interval=1,
            output_size=10,
            repeat=1,
            script=\"\"\"config router access-list
        edit "static-redistribution"
        config rule
        edit 10
        set prefix 10.0.0.0 255.255.255.0
        set action permit
        set exact-match enable
        end
        end

        \"\"\",
            start="auto")
        ```

        ## Import

        Router AccessList can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/routerAccessList:RouterAccessList labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerAccessList:RouterAccessList labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param RouterAccessListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterAccessListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterAccessListRuleArgs']]]]] = None,
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
            __props__ = RouterAccessListArgs.__new__(RouterAccessListArgs)

            __props__.__dict__["comments"] = comments
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["rules"] = rules
            __props__.__dict__["vdomparam"] = vdomparam
        super(RouterAccessList, __self__).__init__(
            'fortios:index/routerAccessList:RouterAccessList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterAccessListRuleArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'RouterAccessList':
        """
        Get an existing RouterAccessList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterAccessListRuleArgs']]]] rules: Rule. The structure of `rule` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterAccessListState.__new__(_RouterAccessListState)

        __props__.__dict__["comments"] = comments
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["rules"] = rules
        __props__.__dict__["vdomparam"] = vdomparam
        return RouterAccessList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.RouterAccessListRule']]]:
        """
        Rule. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

