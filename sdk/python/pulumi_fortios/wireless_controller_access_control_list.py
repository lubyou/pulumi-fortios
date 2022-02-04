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

__all__ = ['WirelessControllerAccessControlListArgs', 'WirelessControllerAccessControlList']

@pulumi.input_type
class WirelessControllerAccessControlListArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 layer3_ipv4_rules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]] = None,
                 layer3_ipv6_rules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerAccessControlList resource.
        :param pulumi.Input[str] comment: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]] layer3_ipv4_rules: AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]] layer3_ipv6_rules: AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        :param pulumi.Input[str] name: AP access control list name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if layer3_ipv4_rules is not None:
            pulumi.set(__self__, "layer3_ipv4_rules", layer3_ipv4_rules)
        if layer3_ipv6_rules is not None:
            pulumi.set(__self__, "layer3_ipv6_rules", layer3_ipv6_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
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
    @pulumi.getter(name="layer3Ipv4Rules")
    def layer3_ipv4_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]:
        """
        AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv4_rules")

    @layer3_ipv4_rules.setter
    def layer3_ipv4_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]):
        pulumi.set(self, "layer3_ipv4_rules", value)

    @property
    @pulumi.getter(name="layer3Ipv6Rules")
    def layer3_ipv6_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]:
        """
        AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv6_rules")

    @layer3_ipv6_rules.setter
    def layer3_ipv6_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]):
        pulumi.set(self, "layer3_ipv6_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        AP access control list name.
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
class _WirelessControllerAccessControlListState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 layer3_ipv4_rules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]] = None,
                 layer3_ipv6_rules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerAccessControlList resources.
        :param pulumi.Input[str] comment: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]] layer3_ipv4_rules: AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]] layer3_ipv6_rules: AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        :param pulumi.Input[str] name: AP access control list name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if layer3_ipv4_rules is not None:
            pulumi.set(__self__, "layer3_ipv4_rules", layer3_ipv4_rules)
        if layer3_ipv6_rules is not None:
            pulumi.set(__self__, "layer3_ipv6_rules", layer3_ipv6_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
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
    @pulumi.getter(name="layer3Ipv4Rules")
    def layer3_ipv4_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]:
        """
        AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv4_rules")

    @layer3_ipv4_rules.setter
    def layer3_ipv4_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]):
        pulumi.set(self, "layer3_ipv4_rules", value)

    @property
    @pulumi.getter(name="layer3Ipv6Rules")
    def layer3_ipv6_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]:
        """
        AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv6_rules")

    @layer3_ipv6_rules.setter
    def layer3_ipv6_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]):
        pulumi.set(self, "layer3_ipv6_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        AP access control list name.
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


class WirelessControllerAccessControlList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 layer3_ipv4_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]] = None,
                 layer3_ipv6_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WiFi bridge access control list. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        WirelessController AccessControlList can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerAccessControlList:WirelessControllerAccessControlList labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]] layer3_ipv4_rules: AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]] layer3_ipv6_rules: AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        :param pulumi.Input[str] name: AP access control list name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerAccessControlListArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WiFi bridge access control list. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        WirelessController AccessControlList can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerAccessControlList:WirelessControllerAccessControlList labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerAccessControlListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerAccessControlListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 layer3_ipv4_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]] = None,
                 layer3_ipv6_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]] = None,
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
            __props__ = WirelessControllerAccessControlListArgs.__new__(WirelessControllerAccessControlListArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["layer3_ipv4_rules"] = layer3_ipv4_rules
            __props__.__dict__["layer3_ipv6_rules"] = layer3_ipv6_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerAccessControlList, __self__).__init__(
            'fortios:index/wirelessControllerAccessControlList:WirelessControllerAccessControlList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            layer3_ipv4_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]]] = None,
            layer3_ipv6_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerAccessControlList':
        """
        Get an existing WirelessControllerAccessControlList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv4RuleArgs']]]] layer3_ipv4_rules: AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAccessControlListLayer3Ipv6RuleArgs']]]] layer3_ipv6_rules: AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        :param pulumi.Input[str] name: AP access control list name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerAccessControlListState.__new__(_WirelessControllerAccessControlListState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["layer3_ipv4_rules"] = layer3_ipv4_rules
        __props__.__dict__["layer3_ipv6_rules"] = layer3_ipv6_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerAccessControlList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        Description.
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
    @pulumi.getter(name="layer3Ipv4Rules")
    def layer3_ipv4_rules(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerAccessControlListLayer3Ipv4Rule']]]:
        """
        AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv4_rules")

    @property
    @pulumi.getter(name="layer3Ipv6Rules")
    def layer3_ipv6_rules(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerAccessControlListLayer3Ipv6Rule']]]:
        """
        AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
        """
        return pulumi.get(self, "layer3_ipv6_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        AP access control list name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

