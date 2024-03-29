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

__all__ = ['WebProxyProfileArgs', 'WebProxyProfile']

@pulumi.input_type
class WebProxyProfileArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 header_client_ip: Optional[pulumi.Input[str]] = None,
                 header_front_end_https: Optional[pulumi.Input[str]] = None,
                 header_via_request: Optional[pulumi.Input[str]] = None,
                 header_via_response: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_groups: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_user: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_client_cert: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_for: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]] = None,
                 log_header_change: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strip_encoding: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebProxyProfile resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if header_client_ip is not None:
            pulumi.set(__self__, "header_client_ip", header_client_ip)
        if header_front_end_https is not None:
            pulumi.set(__self__, "header_front_end_https", header_front_end_https)
        if header_via_request is not None:
            pulumi.set(__self__, "header_via_request", header_via_request)
        if header_via_response is not None:
            pulumi.set(__self__, "header_via_response", header_via_response)
        if header_x_authenticated_groups is not None:
            pulumi.set(__self__, "header_x_authenticated_groups", header_x_authenticated_groups)
        if header_x_authenticated_user is not None:
            pulumi.set(__self__, "header_x_authenticated_user", header_x_authenticated_user)
        if header_x_forwarded_client_cert is not None:
            pulumi.set(__self__, "header_x_forwarded_client_cert", header_x_forwarded_client_cert)
        if header_x_forwarded_for is not None:
            pulumi.set(__self__, "header_x_forwarded_for", header_x_forwarded_for)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if log_header_change is not None:
            pulumi.set(__self__, "log_header_change", log_header_change)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if strip_encoding is not None:
            pulumi.set(__self__, "strip_encoding", strip_encoding)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="headerClientIp")
    def header_client_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_client_ip")

    @header_client_ip.setter
    def header_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_client_ip", value)

    @property
    @pulumi.getter(name="headerFrontEndHttps")
    def header_front_end_https(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_front_end_https")

    @header_front_end_https.setter
    def header_front_end_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_front_end_https", value)

    @property
    @pulumi.getter(name="headerViaRequest")
    def header_via_request(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_via_request")

    @header_via_request.setter
    def header_via_request(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_via_request", value)

    @property
    @pulumi.getter(name="headerViaResponse")
    def header_via_response(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_via_response")

    @header_via_response.setter
    def header_via_response(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_via_response", value)

    @property
    @pulumi.getter(name="headerXAuthenticatedGroups")
    def header_x_authenticated_groups(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_authenticated_groups")

    @header_x_authenticated_groups.setter
    def header_x_authenticated_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_authenticated_groups", value)

    @property
    @pulumi.getter(name="headerXAuthenticatedUser")
    def header_x_authenticated_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_authenticated_user")

    @header_x_authenticated_user.setter
    def header_x_authenticated_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_authenticated_user", value)

    @property
    @pulumi.getter(name="headerXForwardedClientCert")
    def header_x_forwarded_client_cert(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_forwarded_client_cert")

    @header_x_forwarded_client_cert.setter
    def header_x_forwarded_client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_forwarded_client_cert", value)

    @property
    @pulumi.getter(name="headerXForwardedFor")
    def header_x_forwarded_for(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_forwarded_for")

    @header_x_forwarded_for.setter
    def header_x_forwarded_for(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_forwarded_for", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="logHeaderChange")
    def log_header_change(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_header_change")

    @log_header_change.setter
    def log_header_change(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_header_change", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="stripEncoding")
    def strip_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "strip_encoding")

    @strip_encoding.setter
    def strip_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strip_encoding", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _WebProxyProfileState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 header_client_ip: Optional[pulumi.Input[str]] = None,
                 header_front_end_https: Optional[pulumi.Input[str]] = None,
                 header_via_request: Optional[pulumi.Input[str]] = None,
                 header_via_response: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_groups: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_user: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_client_cert: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_for: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]] = None,
                 log_header_change: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strip_encoding: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebProxyProfile resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if header_client_ip is not None:
            pulumi.set(__self__, "header_client_ip", header_client_ip)
        if header_front_end_https is not None:
            pulumi.set(__self__, "header_front_end_https", header_front_end_https)
        if header_via_request is not None:
            pulumi.set(__self__, "header_via_request", header_via_request)
        if header_via_response is not None:
            pulumi.set(__self__, "header_via_response", header_via_response)
        if header_x_authenticated_groups is not None:
            pulumi.set(__self__, "header_x_authenticated_groups", header_x_authenticated_groups)
        if header_x_authenticated_user is not None:
            pulumi.set(__self__, "header_x_authenticated_user", header_x_authenticated_user)
        if header_x_forwarded_client_cert is not None:
            pulumi.set(__self__, "header_x_forwarded_client_cert", header_x_forwarded_client_cert)
        if header_x_forwarded_for is not None:
            pulumi.set(__self__, "header_x_forwarded_for", header_x_forwarded_for)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if log_header_change is not None:
            pulumi.set(__self__, "log_header_change", log_header_change)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if strip_encoding is not None:
            pulumi.set(__self__, "strip_encoding", strip_encoding)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="headerClientIp")
    def header_client_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_client_ip")

    @header_client_ip.setter
    def header_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_client_ip", value)

    @property
    @pulumi.getter(name="headerFrontEndHttps")
    def header_front_end_https(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_front_end_https")

    @header_front_end_https.setter
    def header_front_end_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_front_end_https", value)

    @property
    @pulumi.getter(name="headerViaRequest")
    def header_via_request(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_via_request")

    @header_via_request.setter
    def header_via_request(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_via_request", value)

    @property
    @pulumi.getter(name="headerViaResponse")
    def header_via_response(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_via_response")

    @header_via_response.setter
    def header_via_response(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_via_response", value)

    @property
    @pulumi.getter(name="headerXAuthenticatedGroups")
    def header_x_authenticated_groups(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_authenticated_groups")

    @header_x_authenticated_groups.setter
    def header_x_authenticated_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_authenticated_groups", value)

    @property
    @pulumi.getter(name="headerXAuthenticatedUser")
    def header_x_authenticated_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_authenticated_user")

    @header_x_authenticated_user.setter
    def header_x_authenticated_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_authenticated_user", value)

    @property
    @pulumi.getter(name="headerXForwardedClientCert")
    def header_x_forwarded_client_cert(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_forwarded_client_cert")

    @header_x_forwarded_client_cert.setter
    def header_x_forwarded_client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_forwarded_client_cert", value)

    @property
    @pulumi.getter(name="headerXForwardedFor")
    def header_x_forwarded_for(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "header_x_forwarded_for")

    @header_x_forwarded_for.setter
    def header_x_forwarded_for(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_x_forwarded_for", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebProxyProfileHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="logHeaderChange")
    def log_header_change(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_header_change")

    @log_header_change.setter
    def log_header_change(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_header_change", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="stripEncoding")
    def strip_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "strip_encoding")

    @strip_encoding.setter
    def strip_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strip_encoding", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class WebProxyProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 header_client_ip: Optional[pulumi.Input[str]] = None,
                 header_front_end_https: Optional[pulumi.Input[str]] = None,
                 header_via_request: Optional[pulumi.Input[str]] = None,
                 header_via_response: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_groups: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_user: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_client_cert: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_for: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebProxyProfileHeaderArgs']]]]] = None,
                 log_header_change: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strip_encoding: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WebProxyProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebProxyProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WebProxyProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WebProxyProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebProxyProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 header_client_ip: Optional[pulumi.Input[str]] = None,
                 header_front_end_https: Optional[pulumi.Input[str]] = None,
                 header_via_request: Optional[pulumi.Input[str]] = None,
                 header_via_response: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_groups: Optional[pulumi.Input[str]] = None,
                 header_x_authenticated_user: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_client_cert: Optional[pulumi.Input[str]] = None,
                 header_x_forwarded_for: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebProxyProfileHeaderArgs']]]]] = None,
                 log_header_change: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strip_encoding: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebProxyProfileArgs.__new__(WebProxyProfileArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["header_client_ip"] = header_client_ip
            __props__.__dict__["header_front_end_https"] = header_front_end_https
            __props__.__dict__["header_via_request"] = header_via_request
            __props__.__dict__["header_via_response"] = header_via_response
            __props__.__dict__["header_x_authenticated_groups"] = header_x_authenticated_groups
            __props__.__dict__["header_x_authenticated_user"] = header_x_authenticated_user
            __props__.__dict__["header_x_forwarded_client_cert"] = header_x_forwarded_client_cert
            __props__.__dict__["header_x_forwarded_for"] = header_x_forwarded_for
            __props__.__dict__["headers"] = headers
            __props__.__dict__["log_header_change"] = log_header_change
            __props__.__dict__["name"] = name
            __props__.__dict__["strip_encoding"] = strip_encoding
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebProxyProfile, __self__).__init__(
            'fortios:index/webProxyProfile:WebProxyProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            header_client_ip: Optional[pulumi.Input[str]] = None,
            header_front_end_https: Optional[pulumi.Input[str]] = None,
            header_via_request: Optional[pulumi.Input[str]] = None,
            header_via_response: Optional[pulumi.Input[str]] = None,
            header_x_authenticated_groups: Optional[pulumi.Input[str]] = None,
            header_x_authenticated_user: Optional[pulumi.Input[str]] = None,
            header_x_forwarded_client_cert: Optional[pulumi.Input[str]] = None,
            header_x_forwarded_for: Optional[pulumi.Input[str]] = None,
            headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebProxyProfileHeaderArgs']]]]] = None,
            log_header_change: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            strip_encoding: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebProxyProfile':
        """
        Get an existing WebProxyProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebProxyProfileState.__new__(_WebProxyProfileState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["header_client_ip"] = header_client_ip
        __props__.__dict__["header_front_end_https"] = header_front_end_https
        __props__.__dict__["header_via_request"] = header_via_request
        __props__.__dict__["header_via_response"] = header_via_response
        __props__.__dict__["header_x_authenticated_groups"] = header_x_authenticated_groups
        __props__.__dict__["header_x_authenticated_user"] = header_x_authenticated_user
        __props__.__dict__["header_x_forwarded_client_cert"] = header_x_forwarded_client_cert
        __props__.__dict__["header_x_forwarded_for"] = header_x_forwarded_for
        __props__.__dict__["headers"] = headers
        __props__.__dict__["log_header_change"] = log_header_change
        __props__.__dict__["name"] = name
        __props__.__dict__["strip_encoding"] = strip_encoding
        __props__.__dict__["vdomparam"] = vdomparam
        return WebProxyProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="headerClientIp")
    def header_client_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_client_ip")

    @property
    @pulumi.getter(name="headerFrontEndHttps")
    def header_front_end_https(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_front_end_https")

    @property
    @pulumi.getter(name="headerViaRequest")
    def header_via_request(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_via_request")

    @property
    @pulumi.getter(name="headerViaResponse")
    def header_via_response(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_via_response")

    @property
    @pulumi.getter(name="headerXAuthenticatedGroups")
    def header_x_authenticated_groups(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_x_authenticated_groups")

    @property
    @pulumi.getter(name="headerXAuthenticatedUser")
    def header_x_authenticated_user(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_x_authenticated_user")

    @property
    @pulumi.getter(name="headerXForwardedClientCert")
    def header_x_forwarded_client_cert(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_x_forwarded_client_cert")

    @property
    @pulumi.getter(name="headerXForwardedFor")
    def header_x_forwarded_for(self) -> pulumi.Output[str]:
        return pulumi.get(self, "header_x_forwarded_for")

    @property
    @pulumi.getter
    def headers(self) -> pulumi.Output[Optional[Sequence['outputs.WebProxyProfileHeader']]]:
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="logHeaderChange")
    def log_header_change(self) -> pulumi.Output[str]:
        return pulumi.get(self, "log_header_change")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="stripEncoding")
    def strip_encoding(self) -> pulumi.Output[str]:
        return pulumi.get(self, "strip_encoding")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

