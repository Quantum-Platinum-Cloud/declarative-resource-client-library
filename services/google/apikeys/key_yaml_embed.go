// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package apikeys -var YAML_key blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apikeys/key.yaml

package apikeys

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apikeys/key.yaml
var YAML_key = []byte("info:\n  title: Apikeys/Key\n  description: The Apikeys Key resource\n  x-dcl-struct-name: Key\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Key\n    parameters:\n    - name: Key\n      required: true\n      description: A full instance of a Key\n  apply:\n    description: The function used to apply information about a Key\n    parameters:\n    - name: Key\n      required: true\n      description: A full instance of a Key\n  delete:\n    description: The function used to delete a Key\n    parameters:\n    - name: Key\n      required: true\n      description: A full instance of a Key\n  deleteAll:\n    description: The function used to delete all Key\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Key\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Key:\n      title: Key\n      x-dcl-id: projects/{{project}}/locations/global/keys/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. A timestamp identifying the time this API key\n            was originally created.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. A timestamp when this key was deleted. If the\n            resource is not deleted, this must be empty.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Human-readable display name of this API key. Modifiable by\n            user.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Output only. A checksum computed by the server based on the\n            current value of the Key resource. This may be sent on update and delete\n            requests to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        keyString:\n          type: string\n          x-dcl-go-name: KeyString\n          readOnly: true\n          description: Output only. An encrypted and signed value held by this key.\n            This field can be accessed only through the `GetKeyString` method.\n          x-kubernetes-immutable: true\n          x-dcl-sensitive: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of the key. The name must be unique within\n            the project, must conform with RFC-1034, is restricted to lower-cased\n            letters, and has a maximum length of 63 characters. In another word, the\n            name must match the regular expression: [a-z]([a-z0-9-]{0,61}[a-z0-9])?.'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        restrictions:\n          type: object\n          x-dcl-go-name: Restrictions\n          x-dcl-go-type: KeyRestrictions\n          description: Key restrictions.\n          properties:\n            androidKeyRestrictions:\n              type: object\n              x-dcl-go-name: AndroidKeyRestrictions\n              x-dcl-go-type: KeyRestrictionsAndroidKeyRestrictions\n              description: The Android apps that are allowed to use the key.\n              x-dcl-conflicts:\n              - browserKeyRestrictions\n              - serverKeyRestrictions\n              - iosKeyRestrictions\n              properties:\n                allowedApplications:\n                  type: array\n                  x-dcl-go-name: AllowedApplications\n                  description: A list of Android applications that are allowed to\n                    make API calls with this key.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: KeyRestrictionsAndroidKeyRestrictionsAllowedApplications\n                    properties:\n                      packageName:\n                        type: string\n                        x-dcl-go-name: PackageName\n                        description: The package name of the application.\n                      sha1Fingerprint:\n                        type: string\n                        x-dcl-go-name: Sha1Fingerprint\n                        description: 'The SHA1 fingerprint of the application. For\n                          example, both sha1 formats are acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09\n                          or DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format\n                          is the latter.'\n            apiTargets:\n              type: array\n              x-dcl-go-name: ApiTargets\n              description: A restriction for a specific service and optionally one\n                or more specific methods. Requests are allowed if they match any of\n                these restrictions. If no restrictions are specified, all targets\n                are allowed.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: KeyRestrictionsApiTargets\n                properties:\n                  methods:\n                    type: array\n                    x-dcl-go-name: Methods\n                    description: 'Optional. List of one or more methods that can be\n                      called. If empty, all methods for the service are allowed. A\n                      wildcard (*) can be used as the last symbol. Valid examples:\n                      `google.cloud.translate.v2.TranslateService.GetSupportedLanguage`\n                      `TranslateText` `Get*` `translate.googleapis.com.Get*`'\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  service:\n                    type: string\n                    x-dcl-go-name: Service\n                    description: 'The service for this restriction. It should be the\n                      canonical service name, for example: `translate.googleapis.com`.\n                      You can use `gcloud services list` to get a list of services\n                      that are enabled in the project.'\n            browserKeyRestrictions:\n              type: object\n              x-dcl-go-name: BrowserKeyRestrictions\n              x-dcl-go-type: KeyRestrictionsBrowserKeyRestrictions\n              description: The HTTP referrers (websites) that are allowed to use the\n                key.\n              x-dcl-conflicts:\n              - serverKeyRestrictions\n              - androidKeyRestrictions\n              - iosKeyRestrictions\n              properties:\n                allowedReferrers:\n                  type: array\n                  x-dcl-go-name: AllowedReferrers\n                  description: A list of regular expressions for the referrer URLs\n                    that are allowed to make API calls with this key.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n            iosKeyRestrictions:\n              type: object\n              x-dcl-go-name: IosKeyRestrictions\n              x-dcl-go-type: KeyRestrictionsIosKeyRestrictions\n              description: The iOS apps that are allowed to use the key.\n              x-dcl-conflicts:\n              - browserKeyRestrictions\n              - serverKeyRestrictions\n              - androidKeyRestrictions\n              properties:\n                allowedBundleIds:\n                  type: array\n                  x-dcl-go-name: AllowedBundleIds\n                  description: A list of bundle IDs that are allowed when making API\n                    calls with this key.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n            serverKeyRestrictions:\n              type: object\n              x-dcl-go-name: ServerKeyRestrictions\n              x-dcl-go-type: KeyRestrictionsServerKeyRestrictions\n              description: The IP addresses of callers that are allowed to use the\n                key.\n              x-dcl-conflicts:\n              - browserKeyRestrictions\n              - androidKeyRestrictions\n              - iosKeyRestrictions\n              properties:\n                allowedIps:\n                  type: array\n                  x-dcl-go-name: AllowedIps\n                  description: A list of the caller IP addresses that are allowed\n                    to make API calls with this key.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Unique id in UUID4 format.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. A timestamp identifying the time this key was\n            last updated.\n          x-kubernetes-immutable: true\n")

// 10071 bytes
// MD5: 25f70be93f783e6cc2e1b1cc19452b1f
