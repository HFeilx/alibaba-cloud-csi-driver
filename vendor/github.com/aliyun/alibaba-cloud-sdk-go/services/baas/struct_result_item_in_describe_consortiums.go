package baas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ResultItemInDescribeConsortiums is a nested struct in baas response
type ResultItemInDescribeConsortiums struct {
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ConsortiumId      string `json:"ConsortiumId" xml:"ConsortiumId"`
	Name              string `json:"Name" xml:"Name"`
	CodeName          string `json:"CodeName" xml:"CodeName"`
	OwnerBid          string `json:"OwnerBid" xml:"OwnerBid"`
	OwnerUid          int    `json:"OwnerUid" xml:"OwnerUid"`
	OwnerName         string `json:"OwnerName" xml:"OwnerName"`
	ChannelPolicy     string `json:"ChannelPolicy" xml:"ChannelPolicy"`
	OrganizationCount int    `json:"OrganizationCount" xml:"OrganizationCount"`
	ChannelCount      int    `json:"ChannelCount" xml:"ChannelCount"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	State             string `json:"State" xml:"State"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
	Domain            string `json:"Domain" xml:"Domain"`
}