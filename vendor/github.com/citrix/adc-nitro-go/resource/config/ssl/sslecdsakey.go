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
* Configuration for ecdsa key resource.
*/
type Sslecdsakey struct {
	/**
	* Name for and, optionally, path to the ECDSA key file. /nsconfig/ssl/ is the default path.
	*/
	Keyfile string `json:"keyfile,omitempty"`
	/**
	* Curve id to generate ECDSA key. Only P_256 and P_384 are supported
	*/
	Curve string `json:"curve,omitempty"`
	/**
	* Format in which the ECDSA key file is stored on the appliance.
	*/
	Keyform string `json:"keyform,omitempty"`
	/**
	* Encrypt the generated ECDSA key by using the DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.
	*/
	Des bool `json:"des,omitempty"`
	/**
	* Encrypt the generated ECDSA key by using the Triple-DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.
	*/
	Des3 bool `json:"des3,omitempty"`
	/**
	* Encrypt the generated ECDSA key by using the AES algorithm.
	*/
	Aes256 bool `json:"aes256,omitempty"`
	/**
	* Pass phrase to use for encryption if DES or DES3 option is selected.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Create the private key in PKCS#8 format.
	*/
	Pkcs8 bool `json:"pkcs8,omitempty"`

}
