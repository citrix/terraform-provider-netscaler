/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package lb

/**
* Binding class showing the vserver that can be bound to lbvserver.
*/
type Lbvservervserverbinding struct {
	/**
	* Cache virtual server.
	*/
	Cachevserver string `json:"cachevserver,omitempty"`
	/**
	* Name of the policy bound to the LB vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Cache type.
	*/
	Cachetype string `json:"cachetype,omitempty"`
	/**
	* Priority.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* Number of hits.
	*/
	Hits uint32 `json:"hits,omitempty"`
	/**
	* Number of hits.
	*/
	Pipolicyhits uint32 `json:"pipolicyhits,omitempty"`
	Policysubtype uint32 `json:"policysubtype,omitempty"`
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
	*/
	Labelname string `json:"labelname,omitempty"`


}