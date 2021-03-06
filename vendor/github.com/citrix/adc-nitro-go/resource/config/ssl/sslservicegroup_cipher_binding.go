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

package ssl

/**
* Binding class showing the cipher that can be bound to sslservicegroup.
*/
type Sslservicegroupcipherbinding struct {
	/**
	* The name of the cipher group/alias/name configured for the SSL service group.
	*/
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	/**
	* The description of the cipher.
	*/
	Description string `json:"description,omitempty"`
	/**
	* The name of the SSL service to which the SSL policy needs to be bound.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.
	*/
	Ciphername string `json:"ciphername,omitempty"`


}