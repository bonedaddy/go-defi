// Copyright (c) 2021 Alexandre Trottier (bonedaddy). All rights reserved.
// Use of this source code is governed by the Apache 2.0 License that can be found in
// the LICENSE file.

/* Package config provides a flexible configuration management type, allowing for the usage of YAML files or environment variables.
   For details on how the environment variable based configuration works please see https://github.com/vrischmann/envconfig#how-does-it-work however a basic example is included below
	```
		var conf struct {
			Name string
			Shard struct {
				Host string
				Port int
			}
		}
	```
	The above struct will check for the following keys:
		* NAME or name
    	* SHARD_HOST, or shard_host
    	* SHARD_PORT, or shard_port

*/
package config
