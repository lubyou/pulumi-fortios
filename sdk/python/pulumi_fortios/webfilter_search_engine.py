# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterSearchEngineArgs', 'WebfilterSearchEngine']

@pulumi.input_type
class WebfilterSearchEngineArgs:
    def __init__(__self__, *,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterSearchEngine resource.
        """
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query is not None:
            pulumi.set(__self__, "query", query)
        if safesearch is not None:
            pulumi.set(__self__, "safesearch", safesearch)
        if safesearch_str is not None:
            pulumi.set(__self__, "safesearch_str", safesearch_str)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def safesearch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "safesearch")

    @safesearch.setter
    def safesearch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch", value)

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "safesearch_str")

    @safesearch_str.setter
    def safesearch_str(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch_str", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _WebfilterSearchEngineState:
    def __init__(__self__, *,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterSearchEngine resources.
        """
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query is not None:
            pulumi.set(__self__, "query", query)
        if safesearch is not None:
            pulumi.set(__self__, "safesearch", safesearch)
        if safesearch_str is not None:
            pulumi.set(__self__, "safesearch_str", safesearch_str)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter
    def safesearch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "safesearch")

    @safesearch.setter
    def safesearch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch", value)

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "safesearch_str")

    @safesearch_str.setter
    def safesearch_str(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safesearch_str", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class WebfilterSearchEngine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WebfilterSearchEngine resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebfilterSearchEngineArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WebfilterSearchEngine resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WebfilterSearchEngineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterSearchEngineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 charset: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 safesearch: Optional[pulumi.Input[str]] = None,
                 safesearch_str: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebfilterSearchEngineArgs.__new__(WebfilterSearchEngineArgs)

            __props__.__dict__["charset"] = charset
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["name"] = name
            __props__.__dict__["query"] = query
            __props__.__dict__["safesearch"] = safesearch
            __props__.__dict__["safesearch_str"] = safesearch_str
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebfilterSearchEngine, __self__).__init__(
            'fortios:index/webfilterSearchEngine:WebfilterSearchEngine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            charset: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query: Optional[pulumi.Input[str]] = None,
            safesearch: Optional[pulumi.Input[str]] = None,
            safesearch_str: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebfilterSearchEngine':
        """
        Get an existing WebfilterSearchEngine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterSearchEngineState.__new__(_WebfilterSearchEngineState)

        __props__.__dict__["charset"] = charset
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["name"] = name
        __props__.__dict__["query"] = query
        __props__.__dict__["safesearch"] = safesearch
        __props__.__dict__["safesearch_str"] = safesearch_str
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return WebfilterSearchEngine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def charset(self) -> pulumi.Output[str]:
        return pulumi.get(self, "charset")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output[str]:
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def safesearch(self) -> pulumi.Output[str]:
        return pulumi.get(self, "safesearch")

    @property
    @pulumi.getter(name="safesearchStr")
    def safesearch_str(self) -> pulumi.Output[str]:
        return pulumi.get(self, "safesearch_str")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

