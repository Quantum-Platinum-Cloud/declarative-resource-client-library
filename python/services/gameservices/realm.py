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
from google3.cloud.graphite.mmv2.services.google.gameservices import realm_pb2
from google3.cloud.graphite.mmv2.services.google.gameservices import realm_pb2_grpc

from typing import List


class Realm(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        time_zone: str = None,
        description: str = None,
        location: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.time_zone = time_zone
        self.description = description
        self.location = location
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = realm_pb2_grpc.GameservicesRealmServiceStub(channel.Channel())
        request = realm_pb2.ApplyGameservicesRealmRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGameservicesRealm(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.time_zone = Primitive.from_proto(response.time_zone)
        self.description = Primitive.from_proto(response.description)
        self.location = Primitive.from_proto(response.location)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = realm_pb2_grpc.GameservicesRealmServiceStub(channel.Channel())
        request = realm_pb2.DeleteGameservicesRealmRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.time_zone):
            request.resource.time_zone = Primitive.to_proto(self.time_zone)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteGameservicesRealm(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = realm_pb2_grpc.GameservicesRealmServiceStub(channel.Channel())
        request = realm_pb2.ListGameservicesRealmRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListGameservicesRealm(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = realm_pb2.GameservicesRealm()
        any_proto.Unpack(res_proto)

        res = Realm()
        res.name = Primitive.from_proto(res_proto.name)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.time_zone = Primitive.from_proto(res_proto.time_zone)
        res.description = Primitive.from_proto(res_proto.description)
        res.location = Primitive.from_proto(res_proto.location)
        res.project = Primitive.from_proto(res_proto.project)
        return res

    def to_proto(self):
        resource = realm_pb2.GameservicesRealm()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.time_zone):
            resource.time_zone = Primitive.to_proto(self.time_zone)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
