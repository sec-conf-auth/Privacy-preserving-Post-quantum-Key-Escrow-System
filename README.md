# Privacy-preserving Post-quantum Key-Escrow System
About The Function of Chaincodes
----------
The main implementation of system consists of four chaincodes, namely DataSource, EscrowAgent, Investigator,and Supervisor.<br>
<br>
The `Sender DataSource` provides API to generate source data (i.e., the encrypted secret data C and escrowed session key) and APIs to retrieve source data as well as related publickey/record name. <br>
<br>
The `Investigator chaincode` only provides API to decrypt the encrypted secret data C.<br>
<br>
The `Supervisorr chaincode` can be invoked to generate the supervisor’s public/private key pair, get the supervisor’s public key, initialize/get the global setup, anonymize the source data record by giving a pseudonym, recover/get the final session key and get the global time/pseudonym.<br>
<br>
The `EscrowAgent` is responsible to generate the escrow agent’s public/private key pair, get the escrow agent’s public key and decapsulate/get one share (xi
, yi) of one half of the final session key.<br>
<br>
The `collections_config` is used to creat private collection.

About The Function of Command Lines
----------
If developers do not want to develop client codes to create post-quantum supervised secret data sharing applications, they can also use the command lines to complete the post-quantum supervised secret data sharing operations.

# About the Use of Docker Image
In the experiments , we utilize the liboqs 0.6.0 library [1] together with its Go wrapper [2] to generate public/private key pair, encapsulate and decapsulate the shared secret for all the related post-quantum KEM algorithms. Since the liboqs library is incompatible with the
native Hyperledger Fabric docker image for chaincode execution,
we build a new docker image integrated with liboqs based on Ubuntu 18.04 and use it as the execution environment of our
chaincodes. The versions of docker and docker-compose we use
are 20.10.8 and 1.25.0. For developers, the configuration
files for the client codes, the docker file to generate the new
docker image and information about how to use the new image
are also available.

1.In ccenv dockerfile, use the `Ubuntu18.04` as the docker image environment, download and install `liboqs library` required for the experiment, install and configure the Go env(`go1.16.7.linux-amd64.tar.gz`).

2.Build the dockerfile.

   ```javascript
   docker build -t ccenv:latest ./ 
   ```

3.Modify core.yaml.
Set the `chaincode.builder` and `chaincode.golang.runtime` as `ccenv:latest`, and set the `chaincode.golang.dynamicLink` as `true`. The part of the code that needs to be modified in core.yaml is set as follows:

```html
###############################################################################
#
#    Chaincode section
#
###############################################################################
chaincode:

    # The id is used by the Chaincode stub to register the executing Chaincode
    # ID with the Peer and is generally supplied through ENV variables
    # the `path` form of ID is provided when installing the chaincode.
    # The `name` is used for all other requests and can be any string.
    id:
        path:
        name:

    # Generic builder environment, suitable for most chaincode types
    builder: ccenv:latest

    # Enables/disables force pulling of the base docker images (listed below)
    # during user chaincode instantiation.
    # Useful when using moving image tags (such as :latest)
    pull: false

    golang:
        # golang will never need more than baseos
        runtime: ccenv:latest
        # whether or not golang chaincode should be linked dynamically
        dynamicLink: true
   ```
   
4.Start the fabric network and deploy chaincodes.
 


# References
[1] Liboqs, https://github.com/open-quantum-safe/liboqs

[2] Liboqs-go, https://github.com/open-quantum-safe/liboqs-go
