# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.compute import network_pb2
from google3.cloud.graphite.mmv2.services.google.compute import network_pb2_grpc

from typing import List


class Network(object):
    def __init__(
        self,
        description: str = None,
        gateway_ipv4: str = None,
        ipv4_range: str = None,
        name: str = None,
        auto_create_subnetworks: bool = None,
        routing_config: dict = None,
        project: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.name = name
        self.auto_create_subnetworks = auto_create_subnetworks
        self.routing_config = routing_config
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = network_pb2_grpc.ComputeBetaNetworkServiceStub(channel.Channel())
        request = network_pb2.ApplyComputeBetaNetworkRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.auto_create_subnetworks):
            request.resource.auto_create_subnetworks = Primitive.to_proto(
                self.auto_create_subnetworks
            )

        if NetworkRoutingConfig.to_proto(self.routing_config):
            request.resource.routing_config.CopyFrom(
                NetworkRoutingConfig.to_proto(self.routing_config)
            )
        else:
            request.resource.ClearField("routing_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaNetwork(request)
        self.description = Primitive.from_proto(response.description)
        self.gateway_ipv4 = Primitive.from_proto(response.gateway_ipv4)
        self.ipv4_range = Primitive.from_proto(response.ipv4_range)
        self.name = Primitive.from_proto(response.name)
        self.auto_create_subnetworks = Primitive.from_proto(
            response.auto_create_subnetworks
        )
        self.routing_config = NetworkRoutingConfig.from_proto(response.routing_config)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = network_pb2_grpc.ComputeBetaNetworkServiceStub(channel.Channel())
        request = network_pb2.DeleteComputeBetaNetworkRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.auto_create_subnetworks):
            request.resource.auto_create_subnetworks = Primitive.to_proto(
                self.auto_create_subnetworks
            )

        if NetworkRoutingConfig.to_proto(self.routing_config):
            request.resource.routing_config.CopyFrom(
                NetworkRoutingConfig.to_proto(self.routing_config)
            )
        else:
            request.resource.ClearField("routing_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaNetwork(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = network_pb2_grpc.ComputeBetaNetworkServiceStub(channel.Channel())
        request = network_pb2.ListComputeBetaNetworkRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeBetaNetwork(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = network_pb2.ComputeBetaNetwork()
        any_proto.Unpack(res_proto)

        res = Network()
        res.description = Primitive.from_proto(res_proto.description)
        res.gateway_ipv4 = Primitive.from_proto(res_proto.gateway_ipv4)
        res.ipv4_range = Primitive.from_proto(res_proto.ipv4_range)
        res.name = Primitive.from_proto(res_proto.name)
        res.auto_create_subnetworks = Primitive.from_proto(
            res_proto.auto_create_subnetworks
        )
        res.routing_config = NetworkRoutingConfig.from_proto(res_proto.routing_config)
        res.project = Primitive.from_proto(res_proto.project)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        return res

    def to_proto(self):
        resource = network_pb2.ComputeBetaNetwork()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.auto_create_subnetworks):
            resource.auto_create_subnetworks = Primitive.to_proto(
                self.auto_create_subnetworks
            )
        if NetworkRoutingConfig.to_proto(self.routing_config):
            resource.routing_config.CopyFrom(
                NetworkRoutingConfig.to_proto(self.routing_config)
            )
        else:
            resource.ClearField("routing_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NetworkRoutingConfig(object):
    def __init__(self, routing_mode: str = None):
        self.routing_mode = routing_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = network_pb2.ComputeBetaNetworkRoutingConfig()
        if NetworkRoutingConfigRoutingModeEnum.to_proto(resource.routing_mode):
            res.routing_mode = NetworkRoutingConfigRoutingModeEnum.to_proto(
                resource.routing_mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NetworkRoutingConfig(routing_mode=resource.routing_mode,)


class NetworkRoutingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NetworkRoutingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NetworkRoutingConfig.from_proto(i) for i in resources]


class NetworkRoutingConfigRoutingModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return network_pb2.ComputeBetaNetworkRoutingConfigRoutingModeEnum.Value(
            "ComputeBetaNetworkRoutingConfigRoutingModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return network_pb2.ComputeBetaNetworkRoutingConfigRoutingModeEnum.Name(
            resource
        )[len("ComputeBetaNetworkRoutingConfigRoutingModeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
