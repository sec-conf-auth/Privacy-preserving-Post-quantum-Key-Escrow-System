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
<table class=3DMsoNormalTable border=3D1 cellspacing=3D0 cellpadding=3D0 wi=
dth=3D315
 style=3D'width:235.9pt;border-collapse:collapse;border:none;mso-border-alt=
:solid windowtext 1.0pt;
 mso-padding-alt:0cm 5.4pt 0cm 5.4pt;mso-border-insideh:1.0pt solid windowt=
ext;
 mso-border-insidev:1.0pt solid windowtext'>
 <thead>
  <tr style=3D'mso-yfti-irow:0;mso-yfti-firstrow:yes;height:12.0pt'>
   <td width=3D72 style=3D'width:53.65pt;border:solid windowtext 1.0pt;bord=
er-left:
   none;padding:0cm 5.4pt 0cm 5.4pt;height:12.0pt'>
   <p class=3Dtablecolhead style=3D'margin-left:3.8pt;mso-para-margin-left:=
-.11gd;
   text-indent:-4.8pt;mso-char-indent-count:-.6'><span class=3DSpellE><span
   lang=3DEN-US style=3D'color:black;mso-themecolor:text1'>Chaincodes</span=
></span><span
   lang=3DEN-US style=3D'color:black;mso-themecolor:text1;mso-fareast-langu=
age:
   ZH-CN'>/<o:p></o:p></span></p>
   <p class=3Dtablecolhead style=3D'margin-left:3.8pt;mso-para-margin-left:=
-.11gd;
   text-indent:-4.8pt;mso-char-indent-count:-.6'><span lang=3DEN-US
   style=3D'color:black;mso-themecolor:text1;mso-fareast-language:ZH-CN'>Cl=
ient</span><span
   lang=3DEN-US style=3D'color:black;mso-themecolor:text1'> codes<o:p></o:p=
></span></p>
   </td>
   <td width=3D57 style=3D'width:42.55pt;border:solid windowtext 1.0pt;bord=
er-left:
   none;mso-border-left-alt:solid windowtext 1.0pt;padding:0cm 5.4pt 0cm 5.=
4pt;
   height:12.0pt'>
   <p class=3Dtablecolhead><span lang=3DEN-US style=3D'color:black;mso-them=
ecolor:
   text1'>APIs<o:p></o:p></span></p>
   </td>
   <td width=3D186 style=3D'width:139.7pt;border-top:solid windowtext 1.0pt;
   border-left:none;border-bottom:solid windowtext 1.0pt;border-right:none;
   mso-border-left-alt:solid windowtext 1.0pt;padding:0cm 5.4pt 0cm 5.4pt;
   height:12.0pt'>
   <p class=3Dtablecolhead align=3Dleft style=3D'text-align:left'><span lan=
g=3DEN-US
   style=3D'color:black;mso-themecolor:text1'>Descriptions<o:p></o:p></span=
></p>
   </td>
  </tr>
 </thead>
 <tr style=3D'mso-yfti-irow:1;height:19.0pt'>
  <td width=3D72 rowspan=3D6 style=3D'width:53.65pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-top-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'font-family:"Courier New";color:black;mso-themecolor:text1'>PqCa=
/<o:p></o:p></span></p>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'font-family:"Courier New";color:black;mso-themecolor:text1'>PqCa=
_Cli.go<o:p></o:p></span></p>
  </td>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext 1.0pt;
  mso-border-alt:solid windowtext 1.0pt;mso-border-bottom-alt:solid windowt=
ext .5pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Gen<b>_<br>
  </b>CertCA<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Generate
  the key pair and certificate (self-signed) for the CA<o:p></o:p></span></=
p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:2;height:19.0pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Gen<b>_<br>
  </b>CertID<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Generate
  the key pair and ID certificate (signed by the CA) for each user<o:p></o:=
p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:3;height:19.0pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Gen<b>_<br>
  </b>CertKE<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Generate
  the KEM key pair and KE certificate for each user<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:4;height:19.0pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Query<b>_</b><br>
  Cert<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Query
  one certificate based on a given serial number<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:5;height:19.0pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Verify<b>_</b><br>
  Cert<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Verify
  one certificate and its parent certificates in the certificate chain<o:p>=
</o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:6;height:19.0pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:19.0pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Revoke<b>_</b><br>
  Cert<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:19.0pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Revoke
  one certificate based on a given serial number<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:7;height:22.15pt'>
  <td width=3D72 rowspan=3D4 style=3D'width:53.65pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:22.15pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'font-family:"Courier New";color:black;mso-themecolor:text1'>PqUs=
er/<o:p></o:p></span></p>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><a
  name=3D"_Hlk74209996"><span lang=3DEN-US style=3D'font-family:"Courier Ne=
w";
  color:black;mso-themecolor:text1'>PqUser_<br>
  Cli</span></a><span lang=3DEN-US style=3D'font-family:"Courier New";color=
:black;
  mso-themecolor:text1'>.go<o:p></o:p></span></p>
  </td>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:22.15pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Query<b>_</b><br>
  Cert<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:22.15pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Invoke
  <span class=3DSpellE>chaincode</span> </span><span class=3DSpellE><span
  lang=3DEN-US style=3D'font-size:8.0pt;line-height:110%;font-family:"Couri=
er New";
  color:black;mso-themecolor:text1'>PqCa</span></span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'> and
  query one certificate with a given serial number<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:8;height:20.6pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:20.6pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Verify<b>_</b><br>
  Cert<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:20.6pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Invoke
  <span class=3DSpellE>chaincode</span> </span><span class=3DSpellE><span
  lang=3DEN-US style=3D'font-size:8.0pt;line-height:110%;font-family:"Couri=
er New";
  color:black;mso-themecolor:text1'>PqCa</span></span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>
  and verify one certificate (together with the related certificate chain)<=
o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:9;height:20.45pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:.5pt;mso-border-left-alt:1.0pt;mso-border-bottom-alt:.=
5pt;
  mso-border-right-alt:1.0pt;mso-border-color-alt:windowtext;mso-border-sty=
le-alt:
  solid;padding:0cm 5.4pt 0cm 5.4pt;height:20.45pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Encap<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:20.45pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Perform
  the encapsulation operation and upload encapsulated shared secret<o:p></o=
:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:10;mso-yfti-lastrow:yes;height:20.05pt'>
  <td width=3D57 style=3D'width:42.55pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:20.05pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'color:black;mso-themecolor:text1'>Decap<o:p></o:p></span></p>
  </td>
  <td width=3D186 style=3D'width:139.7pt;border:none;border-bottom:solid wi=
ndowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text 1.0pt;
  padding:0cm 5.4pt 0cm 5.4pt;height:20.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;color:black;mso-themecolor:text=
1'>Download
  the encapsulated share secret and perform decapsulation operation<o:p></o=
:p></span></p>
  </td>
 </tr>
</table>

# References
[1] Liboqs, https://github.com/open-quantum-safe/liboqs

[2] Liboqs-go, https://github.com/open-quantum-safe/liboqs-go
