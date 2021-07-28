# PqAuthKEM4Chain
# About the Function of Chaincodes
The main implementation of the PQ user authentication and key
exchange system consists of two chaincodes, namely `PqCa` and
`PqUser`. The `PqCa` chaincode provides
APIs to generate the PQ key pairs together with the related
certificates (for the CA, the user identity key and the user KEM
key), query one certificate (based on a given serial number),
verify one certificate (and its parent certificates in the certificate
chain), revoke one certificate, encapsulate and decapsulate the
shared secret between users. While the `PqUser` chaincode
invokes the corresponding APIs in `PqCa` to query/verify one
certificate and provides APIs to perform the encapsulation and
decapsulation operations for the users.

The data structure in the off-chain private databases is
designed relatively simply for further development. In the off-
chain private databases (e.g., MySQL), one key pair, the
corresponding certificate and related information (e.g., password)
are all stored in one line, and the certificate serial number is set as
the key of the line.

# About the Function of Client Codes
Besides the chaincodes, we also develop client codes to help
developers create their applications based on the PQ user
authentication and key exchange system. 
There are two client codes, namely `PqCa_Cli` and `PqUser_Cli`,
which are written in Go and used to invoke the APIs in the
corresponding chaincodes. The configuration file (e.g.,`connection-profile.yaml`)for the client
codes to invoke the chaincodes is also provided.

# About the Use of Docker Image
In the experiments , we utilize the liboqs 0.6.0 library [1] together with its Go wrapper [2] to generate public/private key pair, encapsulate and decapsulate the shared secret for all the related post-quantum KEM algorithms. Since the liboqs library is incompatible with the
native Hyperledger Fabric docker image for chaincode execution,
we build a new docker image integrated with liboqs based on Ubuntu 18.04 and use it as the execution environment of our
chaincodes. The versions of docker and docker-compose we use
are 19.03.13 and 1.25.0-rc1. For developers, the configuration
files for the client codes, the docker file to generate the new
docker image and information about how to use the new image
are also available.

1.In ccenv dockerfile, use the `Ubuntu18.04` as the docker image environment, download and install `liboqs library` required for the experiment, install and configure the Go env(`go1.15.7.linux-amd64.tar.gz`).

2.Build the dockerfile.

   ```javascript
   docker built -t ccenv:latest ./ 
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
 
# About the Experiment Data
Our experiments involve three VMs on the local server, each one of which is based on Ubuntu 18.04 configured with eight CPU cores of Intel Xeon Gold 6131 throttled to 2.6GHz and 16GB memory, and we deploy one CA peer (for CA administrator) and two user peers (for users Alice and Bob) in one channel separately running on the three VMs. The system is built on Hyperledger Fabric 2.2.3, and the chaincodes together with the client codes are written in Go 1.15.7. 

We divide all the PQ algorithms into three groups, namely level 1-2, level 3, level 4-5, which separately indicate low, medium and high security levels. The specified PQ algorithm names are summarized by groups in Table 1. In our experiments, we choose one PQ signature algorithm and one KEM algorithm in the same security group (i.e., the same line of Table 1).
<table class=MsoNormalTable border=1 cellspacing=0 cellpadding=0 width=311
 style='width:232.9pt;border-collapse:collapse;border:none;mso-border-alt:solid windowtext 1.0pt;
 mso-padding-alt:0cm 5.4pt 0cm 5.4pt;mso-border-insideh:1.0pt solid windowtext;
 mso-border-insidev:1.0pt solid windowtext'>
 <thead>
  <tr style='mso-yfti-irow:0;mso-yfti-firstrow:yes;height:12.0pt'>
   <td width=46 style='width:34.45pt;border:solid windowtext 1.0pt;border-left:
   none;padding:0cm 5.4pt 0cm 5.4pt;height:12.0pt'>
   <p class=tablecolhead style='margin-left:3.8pt;mso-para-margin-left:-.11gd;
   text-indent:-4.8pt;mso-char-indent-count:-.6'><span lang=EN-US
   style='color:black;mso-themecolor:text1'>Levels<o:p></o:p></span></p>
   </td>
   <td width=132 style='width:99.25pt;border:solid windowtext 1.0pt;border-left:
   none;mso-border-left-alt:solid windowtext 1.0pt;padding:0cm 5.4pt 0cm 5.4pt;
   height:12.0pt'>
   <p class=tablecolhead><span lang=EN-US style='color:black;mso-themecolor:
   text1'>Sig. Alg. Names<o:p></o:p></span></p>
   </td>
   <td width=132 style='width:99.2pt;border-top:solid windowtext 1.0pt;
   border-left:none;border-bottom:solid windowtext 1.0pt;border-right:none;
   mso-border-left-alt:solid windowtext 1.0pt;padding:0cm 5.4pt 0cm 5.4pt;
   height:12.0pt'>
   <p class=tablecolhead><span lang=EN-US style='color:black;mso-themecolor:
   text1'>KEM Alg. Names<o:p></o:p></span></p>
   </td>
  </tr>
 </thead>
 <tr style='mso-yfti-irow:1;height:19.0pt'>
  <td width=46 style='width:34.45pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-top-alt:solid windowtext 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid windowtext 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>Level<o:p></o:p></span></p>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>1~2</span><span
  lang=EN-US style='font-size:7.5pt;font-family:"Courier New";color:black;
  mso-themecolor:text1'><o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.25pt;border-top:none;border-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-alt:solid windowtext 1.0pt;mso-border-bottom-alt:solid windowtext .5pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(1).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Dilithium2-AES<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(2).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Falcon-512<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(3).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>picnic3_L1<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(4).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-128f-robust<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(5).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-128f-simple<o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.2pt;border:none;border-bottom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Kyber512-90s<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>NTRU-HPS-2048-509<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>LightSaber-KEM<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>BIKE1-L1-FO<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>FrodoKEM-640-SHAKE<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>HQC-128<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>ntrulpr653<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>sntrup653<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SIKE-p434<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SIKE-p503<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style='mso-yfti-irow:2;height:22.15pt'>
  <td width=46 style='width:34.45pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowtext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid windowtext 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:22.15pt'>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>Level<o:p></o:p></span></p>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>3</span><span
  lang=EN-US style='font-size:7.5pt;font-family:"Courier New";color:black;
  mso-themecolor:text1'><o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.25pt;border-top:none;border-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-style-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:22.15pt'>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(6).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Dilithium3-AES<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(7).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>picnic3_L3<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(8).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-192f-robust<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(9).<span style='font:7.0pt "Times New Roman"'>&nbsp;&nbsp;&nbsp;&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-192f-simple<o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.2pt;border:none;border-bottom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid windowtext 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:22.15pt'>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Kyber768-90s<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>NTRU-HPS-2048-677<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Saber-KEM<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>BIKE1-L3-FO<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>FrodoKEM-976-SHAKE<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>HQC-192<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>ntrulpr761<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>sntrup761<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SIKE-p610<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style='mso-yfti-irow:3;mso-yfti-lastrow:yes;height:20.05pt'>
  <td width=46 style='width:34.45pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:20.05pt'>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>Level<o:p></o:p></span></p>
  <p class=tablecopy align=center style='text-align:center'><span lang=EN-US
  style='font-size:7.5pt;color:black;mso-themecolor:text1'>4~5<o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.25pt;border-top:none;border-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid windowtext 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:20.05pt'>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(10).<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Dilithium5-AES<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(11).<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Falcon-1024<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(12).<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>picnic3_L5<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(13).<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-256f-robust<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:17.0pt;text-indent:-17.0pt;mso-list:
  l0 level1 lfo2'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  mso-fareast-font-family:"Times New Roman";color:black;mso-themecolor:text1'><span
  style='mso-list:Ignore'>(14).<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SPHINCS+-SHA256-256f-simple<o:p></o:p></span></p>
  </td>
  <td width=132 valign=top style='width:99.2pt;border:none;border-bottom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid windowtext 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:20.05pt'>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>Kyber1024-90s<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>NTRU-HPS-4096-821<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>FireSaber-KEM<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>FrodoKEM-1344-SHAKE<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>HQC-256<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>ntrulpr857<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>sntrup857<o:p></o:p></span></p>
  <p class=tablecopy style='margin-left:5.65pt;text-indent:-5.65pt;mso-list:
  l1 level1 lfo1'><![if !supportLists]><span lang=EN-US style='font-size:7.5pt;
  font-family:"Calibri",sans-serif;mso-fareast-font-family:Calibri;color:black;
  mso-themecolor:text1'><span style='mso-list:Ignore'>&#9642;<span style='font:7.0pt "Times New Roman"'>&nbsp;
  </span></span></span><![endif]><span lang=EN-US style='font-size:7.5pt;
  color:black;mso-themecolor:text1'>SIKE-p751<o:p></o:p></span></p>
  </td>
 </tr>
</table>
# References
[1] Liboqs, https://github.com/open-quantum-safe/liboqs

[2] Liboqs-go, https://github.com/open-quantum-safe/liboqs-go
