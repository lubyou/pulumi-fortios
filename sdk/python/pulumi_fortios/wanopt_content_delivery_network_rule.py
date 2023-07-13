# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['WanoptContentDeliveryNetworkRuleArgs', 'WanoptContentDeliveryNetworkRule']

@pulumi.input_type
class WanoptContentDeliveryNetworkRuleArgs:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 host_domain_name_suffixes: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_cache_control: Optional[pulumi.Input[str]] = None,
                 response_cache_control: Optional[pulumi.Input[str]] = None,
                 response_expires: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 text_response_vcache: Optional[pulumi.Input[str]] = None,
                 updateserver: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WanoptContentDeliveryNetworkRule resource.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if host_domain_name_suffixes is not None:
            pulumi.set(__self__, "host_domain_name_suffixes", host_domain_name_suffixes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_cache_control is not None:
            pulumi.set(__self__, "request_cache_control", request_cache_control)
        if response_cache_control is not None:
            pulumi.set(__self__, "response_cache_control", response_cache_control)
        if response_expires is not None:
            pulumi.set(__self__, "response_expires", response_expires)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if text_response_vcache is not None:
            pulumi.set(__self__, "text_response_vcache", text_response_vcache)
        if updateserver is not None:
            pulumi.set(__self__, "updateserver", updateserver)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="hostDomainNameSuffixes")
    def host_domain_name_suffixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]:
        return pulumi.get(self, "host_domain_name_suffixes")

    @host_domain_name_suffixes.setter
    def host_domain_name_suffixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]):
        pulumi.set(self, "host_domain_name_suffixes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestCacheControl")
    def request_cache_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_cache_control")

    @request_cache_control.setter
    def request_cache_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_cache_control", value)

    @property
    @pulumi.getter(name="responseCacheControl")
    def response_cache_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_cache_control")

    @response_cache_control.setter
    def response_cache_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_cache_control", value)

    @property
    @pulumi.getter(name="responseExpires")
    def response_expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_expires")

    @response_expires.setter
    def response_expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_expires", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="textResponseVcache")
    def text_response_vcache(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "text_response_vcache")

    @text_response_vcache.setter
    def text_response_vcache(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "text_response_vcache", value)

    @property
    @pulumi.getter
    def updateserver(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updateserver")

    @updateserver.setter
    def updateserver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updateserver", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _WanoptContentDeliveryNetworkRuleState:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 host_domain_name_suffixes: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_cache_control: Optional[pulumi.Input[str]] = None,
                 response_cache_control: Optional[pulumi.Input[str]] = None,
                 response_expires: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 text_response_vcache: Optional[pulumi.Input[str]] = None,
                 updateserver: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WanoptContentDeliveryNetworkRule resources.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if host_domain_name_suffixes is not None:
            pulumi.set(__self__, "host_domain_name_suffixes", host_domain_name_suffixes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_cache_control is not None:
            pulumi.set(__self__, "request_cache_control", request_cache_control)
        if response_cache_control is not None:
            pulumi.set(__self__, "response_cache_control", response_cache_control)
        if response_expires is not None:
            pulumi.set(__self__, "response_expires", response_expires)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if text_response_vcache is not None:
            pulumi.set(__self__, "text_response_vcache", text_response_vcache)
        if updateserver is not None:
            pulumi.set(__self__, "updateserver", updateserver)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="hostDomainNameSuffixes")
    def host_domain_name_suffixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]:
        return pulumi.get(self, "host_domain_name_suffixes")

    @host_domain_name_suffixes.setter
    def host_domain_name_suffixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]):
        pulumi.set(self, "host_domain_name_suffixes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestCacheControl")
    def request_cache_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_cache_control")

    @request_cache_control.setter
    def request_cache_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_cache_control", value)

    @property
    @pulumi.getter(name="responseCacheControl")
    def response_cache_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_cache_control")

    @response_cache_control.setter
    def response_cache_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_cache_control", value)

    @property
    @pulumi.getter(name="responseExpires")
    def response_expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_expires")

    @response_expires.setter
    def response_expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_expires", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WanoptContentDeliveryNetworkRuleRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="textResponseVcache")
    def text_response_vcache(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "text_response_vcache")

    @text_response_vcache.setter
    def text_response_vcache(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "text_response_vcache", value)

    @property
    @pulumi.getter
    def updateserver(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updateserver")

    @updateserver.setter
    def updateserver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updateserver", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class WanoptContentDeliveryNetworkRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 host_domain_name_suffixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_cache_control: Optional[pulumi.Input[str]] = None,
                 response_cache_control: Optional[pulumi.Input[str]] = None,
                 response_expires: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleRuleArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 text_response_vcache: Optional[pulumi.Input[str]] = None,
                 updateserver: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WanoptContentDeliveryNetworkRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WanoptContentDeliveryNetworkRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WanoptContentDeliveryNetworkRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WanoptContentDeliveryNetworkRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WanoptContentDeliveryNetworkRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 host_domain_name_suffixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_cache_control: Optional[pulumi.Input[str]] = None,
                 response_cache_control: Optional[pulumi.Input[str]] = None,
                 response_expires: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleRuleArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 text_response_vcache: Optional[pulumi.Input[str]] = None,
                 updateserver: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WanoptContentDeliveryNetworkRuleArgs.__new__(WanoptContentDeliveryNetworkRuleArgs)

            __props__.__dict__["category"] = category
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["host_domain_name_suffixes"] = host_domain_name_suffixes
            __props__.__dict__["name"] = name
            __props__.__dict__["request_cache_control"] = request_cache_control
            __props__.__dict__["response_cache_control"] = response_cache_control
            __props__.__dict__["response_expires"] = response_expires
            __props__.__dict__["rules"] = rules
            __props__.__dict__["status"] = status
            __props__.__dict__["text_response_vcache"] = text_response_vcache
            __props__.__dict__["updateserver"] = updateserver
            __props__.__dict__["vdomparam"] = vdomparam
        super(WanoptContentDeliveryNetworkRule, __self__).__init__(
            'fortios:index/wanoptContentDeliveryNetworkRule:WanoptContentDeliveryNetworkRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            category: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            host_domain_name_suffixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleHostDomainNameSuffixArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            request_cache_control: Optional[pulumi.Input[str]] = None,
            response_cache_control: Optional[pulumi.Input[str]] = None,
            response_expires: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WanoptContentDeliveryNetworkRuleRuleArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            text_response_vcache: Optional[pulumi.Input[str]] = None,
            updateserver: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WanoptContentDeliveryNetworkRule':
        """
        Get an existing WanoptContentDeliveryNetworkRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WanoptContentDeliveryNetworkRuleState.__new__(_WanoptContentDeliveryNetworkRuleState)

        __props__.__dict__["category"] = category
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["host_domain_name_suffixes"] = host_domain_name_suffixes
        __props__.__dict__["name"] = name
        __props__.__dict__["request_cache_control"] = request_cache_control
        __props__.__dict__["response_cache_control"] = response_cache_control
        __props__.__dict__["response_expires"] = response_expires
        __props__.__dict__["rules"] = rules
        __props__.__dict__["status"] = status
        __props__.__dict__["text_response_vcache"] = text_response_vcache
        __props__.__dict__["updateserver"] = updateserver
        __props__.__dict__["vdomparam"] = vdomparam
        return WanoptContentDeliveryNetworkRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="hostDomainNameSuffixes")
    def host_domain_name_suffixes(self) -> pulumi.Output[Optional[Sequence['outputs.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix']]]:
        return pulumi.get(self, "host_domain_name_suffixes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="requestCacheControl")
    def request_cache_control(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_cache_control")

    @property
    @pulumi.getter(name="responseCacheControl")
    def response_cache_control(self) -> pulumi.Output[str]:
        return pulumi.get(self, "response_cache_control")

    @property
    @pulumi.getter(name="responseExpires")
    def response_expires(self) -> pulumi.Output[str]:
        return pulumi.get(self, "response_expires")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.WanoptContentDeliveryNetworkRuleRule']]]:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="textResponseVcache")
    def text_response_vcache(self) -> pulumi.Output[str]:
        return pulumi.get(self, "text_response_vcache")

    @property
    @pulumi.getter
    def updateserver(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updateserver")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

