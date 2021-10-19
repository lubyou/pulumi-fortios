#!/usr/bin/env python3

import os
import re

from jinja2 import Environment, select_autoescape, FileSystemLoader
import yaml


DATA_SOURCE_SCHEMA_START_PATTERN = r"\s*DataSourcesMap.+schema\.Resource.+"
RESOURCE_SCHEMA_START_PATTERN = r"\s*ResourcesMap.+schema\.Resource.+"
RESOURCE_NAME_PATTERN = (
    r'\s*"(?P<snake_case>\w+)"\s*:\s*(?:dataSource|resource)(?P<camel_case>\w+)\(\s*\)'
)
PROVIDER_SCHEMA_START_PATTERN = r"\s*Schema:\s*map\[string\]\*schema\.Schema\s*{\s*"
PROVIDER_SCHEMA_VAR_PATTERN = r"\s*\"(\w+)\":\s*.*{\s*"
PROVIDER_SCHEMA_ENV_VAR_PATTERN = r'\s*DefaultFunc:\s*schema\.EnvDefaultFunc\("(\w+)".+'


def snake_case_to_camel_case(snake_cased_string):
    return "".join([x.capitalize() for x in snake_cased_string.split("_")])


def extract_resources(provider_file: str = None):
    resources = []
    data_sources = []
    provider_options = []

    current_list = None
    current_provider_option = {}

    for line in open(provider_file):
        if re.match(PROVIDER_SCHEMA_START_PATTERN, line):
            current_list = provider_options

        elif current_list is provider_options and (
            m := re.match(PROVIDER_SCHEMA_VAR_PATTERN, line)
        ):
            current_provider_option = {"name": m.groups()[0], "env_var": None}
            current_list.append(current_provider_option)

        elif current_provider_option and (
            m := re.match(PROVIDER_SCHEMA_ENV_VAR_PATTERN, line)
        ):
            current_provider_option["env_var"] = m.groups()[0]
            current_provider_option = None

        elif re.match(DATA_SOURCE_SCHEMA_START_PATTERN, line):
            current_list = data_sources

        elif re.match(RESOURCE_SCHEMA_START_PATTERN, line):
            current_list = resources

        elif m := re.match(RESOURCE_NAME_PATTERN, line, re.IGNORECASE):
            current_list.append(m.groups())

    return {
        "resources": sorted(resources),
        "data_sources": sorted(data_sources),
        "provider_options": provider_options,
    }


def update_resources_go(
    provider_file: str, provider_config_file: str, resources_file: str = None
):
    templates_path = os.path.join(os.path.dirname(__file__), "templates")
    env = Environment(
        loader=FileSystemLoader(templates_path), autoescape=select_autoescape()
    )
    env.filters["snake_case_to_camel_case"] = snake_case_to_camel_case

    extracted_resources = extract_resources(provider_file=provider_file)

    provider_config = yaml.load(open(provider_config_file), Loader=yaml.Loader)
    provider_config = provider_config["provider"]
    from pprint import pprint

    env_var_mapping = provider_config.get("env_var_mapping", {})

    for provider_option in extracted_resources["provider_options"]:
        name = provider_option["name"]
        if name in env_var_mapping:
            provider_option["env_var"] = env_var_mapping[name]

    resource_override_map = provider_config.get("resource_name_override_mapping", {})

    resources = []
    for resource_name, resource_camel_case_name in extracted_resources["resources"]:
        if resource_name in resource_override_map:
            resource_camel_case_name = resource_override_map[resource_name]
        resources.append((resource_name, resource_camel_case_name))
    extracted_resources["resources"] = resources

    data_sources = []
    for data_source_name, data_source_camel_case_name in extracted_resources[
        "data_sources"
    ]:
        if data_source_name in resource_override_map:
            data_source_camel_case_name = resource_override_map[data_source_name]
        data_sources.append((data_source_name, data_source_camel_case_name))
    extracted_resources["data_sources"] = data_sources

    template = env.get_template("resources.go.j2")
    if resources_file:
        template_stream = template.stream(
            provider_config=provider_config, **extracted_resources
        )
        template_stream.dump(resources_file)
    else:
        rendered_template = template.render(
            provider_name=provider_config, **extracted_resources
        )
        print(rendered_template)


import argparse

parser = argparse.ArgumentParser(
    description="Extracts resource names from a terraform's provider.go"
)
parser.add_argument("provider_file", help="path to provider.go")
parser.add_argument(
    "--resources_file",
    "-o",
    "--out-file",
    required=False,
    help="path to resources.go",
)
parser.add_argument(
    "--provider-config-file",
    "-c",
    required=True,
    help="Path to yaml file with provder config file",
)

args = parser.parse_args()
update_resources_go(
    provider_file=args.provider_file,
    resources_file=args.resources_file,
    provider_config_file=args.provider_config_file,
)
