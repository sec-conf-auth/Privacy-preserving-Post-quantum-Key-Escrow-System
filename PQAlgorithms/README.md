To avoid the Shor’s attack, researchers have developed PQ public-key algorithms relying on different mathematics problems other than the IF, DL or ECDL problems. Based on the mathematics problems, the PQ public-key algorithms can be classified into different families introduced as follows.
2.3.1 PQ Digital Signature Algorithms. In the current round of NIST call, there are six PQ digital signature algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. This kind of cryptographic schemes are based on lattices, which are sets of points in n-dimensional spaces with a periodic structure. Lattice-based security schemes rely on presumed hardness of basic lattice problems like the Shortest Vector Problem (SVP) (i.e., find the shortest non-zero vector within a lattice), the Closest Vector Problem (CVP) (i.e., find a lattice vector that minimizes the distance from another target lattice) and the Shortest Independent Vectors Problem (SIVP) derived from SVP and CVP. In the current round, FALCON and CRYSTALS-DILITHIUM are lattice-based candidate algorithms and separately based on the Short Integer Solution problem (SIS) [39] over NTRU lattices and “Fiat-Shamir with Aborts” technique [40]. Lattice-based algorithms are promising quantum-resistant solutions with relatively efficient implementations, balanced key/signature sizes and strong security properties.

●	Hash-based cryptosystems. The security of these schemes depends on the security of underlying hash function instead of on the hardness of a mathematical problem. Currently, Picnic is a hash-based algorithm and makes use of hash-based zero-knowledge proof technique [41], while SPHINCS+ is a stateless hash-based signature algorithm built on Merkle tree. Normally, hash-based schemes have very small key sizes but suffer from large signature sizes.

●	Multivariate-based cryptosystems. Another family of problems that are used by some NIST PQ signature algorithms is related to solving multivariate quadratic equations over finite fields which is an NP-hard problem. Multivariate PQ schemes often lead to excessive key sizes, and the NIST multivariate-based signature algorithms are Rainbow and GeMSS.
As shown in Table 1 attached as Appendix A, we summarize the details of all the six PQ digital signature algorithms. By using different key sizes and parameter sets, all the algorithms can achieve different NIST security levels (i.e., bits-of-security level), which is defined as the effort required by a classical computer to perform a brute-force attack on a given-length cryptographic key. Normally NIST security levels 1~5 approximately imply 128/160/192/224/256-bits-of security levels.
2.3.2 PQ KEM Algorithms. In the current round of NIST call, there are nine PQ KEM algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. The lattice problems that NIST PQ KEM algorithms depend on are Learning With Errors (LWE) [42], Ring Learning With Errors (Ring-LWE) [43] and Module Learning with Rounding (MLWR) problem [44] with different lattices. And the list of NIST PQ KEM algorithms includes KYBER, NTRU, SABER, FrodoKEM and NTRU Prime, which have similar advantages just like lattice-based PQ signature algorithms.

●	Code-based cryptosystems. This family of algorithms are essentially based on the theory that supports error-correction codes. In the current round of NIST call, Classic McEliece is an example of code-based cryptosystem that dates back from 1970s and whose security is based on the syndrome decoding problem [45]. Classic McEliece provides fast encapsulation and relatively fast decapsulation, but is burdened with huge sizes of public and private key sizes. To ease this issue, the PQ KEM algorithms BIKE and HQC based on different codes have been proposed and selected as the alternate candidate algorithms in the current round of NIST.

●	Supersingular Elliptic Curve Isogeny (SECI) cryptosystems. The SECI cryptosystems are based on the isogeny protocol for ordinary elliptic curves and enhanced to withstand the quantum attack detailed in [46]. The only one isogeny-based PQ KEM algorithm passed to the current round of NIST call is SIKE, which is founded on pseudo-random walks in supersingular isogeny graphs.

We also summarize the details of all the nine PQ KEM algorithms in Table 2 attached as Appendix A. Just like the PQ signature algorithms, all the PQ KEM algorithms can achieve different NIST security levels by using different key lengths and related parameters.

<table class=3DMsoTableGrid border=3D1 cellspacing=3D0 cellpadding=3D0 widt=
h=3D608
 style=3D'width:456.2pt;border-collapse:collapse;border:none;mso-border-alt=
:solid windowtext 1.0pt;
 mso-yfti-tbllook:1184;mso-padding-alt:0cm 5.4pt 0cm 5.4pt'>
 <tr style=3D'mso-yfti-irow:0;mso-yfti-firstrow:yes;height:17.3pt'>
  <td width=3D57 valign=3Dtop style=3D'width:42.55pt;border:solid windowtex=
t 1.0pt;
  border-left:none;mso-border-top-alt:solid windowtext 1.0pt;mso-border-bot=
tom-alt:
  solid windowtext 1.0pt;mso-border-right-alt:solid windowtext .5pt;padding:
  0cm 5.4pt 0cm 5.4pt;height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Alg</span><span lang=3DEN-US style=3D'font-size:8.0pt;mso-fareast=
-language:
  ZH-CN;font-style:normal'>.</span><span lang=3DEN-US style=3D'font-size:8.=
0pt;
  font-style:normal'><o:p></o:p></span></p>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>types<o:p></o:p></span></p>
  </td>
  <td width=3D101 style=3D'width:75.6pt;border:solid windowtext 1.0pt;borde=
r-left:
  none;mso-border-top-alt:solid windowtext 1.0pt;mso-border-bottom-alt:soli=
d windowtext 1.0pt;
  mso-border-right-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;co=
lor:black;
  mso-themecolor:text1;font-style:normal'>Algorithms names</span><span
  lang=3DEN-US style=3D'font-size:8.0pt;mso-fareast-font-family:宋体;mso-fa=
reast-theme-font:
  minor-fareast;color:black;mso-themecolor:text1;mso-fareast-language:ZH-CN;
  font-style:normal'><o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border:solid windowtex=
t 1.0pt;
  border-left:none;mso-border-left-alt:solid windowtext .5pt;mso-border-top=
-alt:
  1.0pt;mso-border-left-alt:.5pt;mso-border-bottom-alt:1.0pt;mso-border-rig=
ht-alt:
  .5pt;mso-border-color-alt:windowtext;mso-border-style-alt:solid;padding:0=
cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Alg.<o:p></o:p></span></p>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>subtypes<o:p></o:p></span></p>
  </td>
  <td width=3D94 style=3D'width:70.85pt;border:solid windowtext 1.0pt;borde=
r-left:
  none;mso-border-left-alt:solid windowtext .5pt;mso-border-top-alt:1.0pt;
  mso-border-left-alt:.5pt;mso-border-bottom-alt:1.0pt;mso-border-right-alt:
  .5pt;mso-border-color-alt:windowtext;mso-border-style-alt:solid;padding:0=
cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Claimed NIST security level<o:p></o:p></span></p>
  </td>
  <td width=3D95 style=3D'width:70.9pt;border:solid windowtext 1.0pt;border=
-left:
  none;mso-border-left-alt:solid windowtext .5pt;mso-border-top-alt:1.0pt;
  mso-border-left-alt:.5pt;mso-border-bottom-alt:1.0pt;mso-border-right-alt:
  .5pt;mso-border-color-alt:windowtext;mso-border-style-alt:solid;padding:0=
cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Public key size (bytes)<o:p></o:p></span></p>
  </td>
  <td width=3D76 style=3D'width:2.0cm;border:solid windowtext 1.0pt;border-=
left:
  none;mso-border-left-alt:solid windowtext .5pt;mso-border-top-alt:1.0pt;
  mso-border-left-alt:.5pt;mso-border-bottom-alt:1.0pt;mso-border-right-alt:
  .5pt;mso-border-color-alt:windowtext;mso-border-style-alt:solid;padding:0=
cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Private key size (bytes)<o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border-top:solid windo=
wtext 1.0pt;
  border-left:none;border-bottom:solid windowtext 1.0pt;border-right:none;
  mso-border-left-alt:solid windowtext .5pt;padding:0cm 5.4pt 0cm 5.4pt;
  height:17.3pt'>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>Signature<o:p></o:p></span></p>
  <p class=3Dtablecolsubhead><span lang=3DEN-US style=3D'font-size:8.0pt;fo=
nt-style:
  normal'>size (bytes)<o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:1;height:10.5pt'>
  <td width=3D57 rowspan=3D6 style=3D'width:42.55pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-top-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext 1.0pt;mso-border-right-alt:solid w=
indowtext .5pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:10.5pt'>
  <p class=3Dtablecopy align=3Dcenter style=3D'text-align:center'><span lan=
g=3DEN-US
  style=3D'font-family:"Linux Libertine";color:black;mso-color-alt:windowte=
xt'>Digital
  signature</span><span lang=3DEN-US style=3D'font-family:"Linux Libertine"=
'><o:p></o:p></span></p>
  </td>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-top-alt:solid window=
text 1.0pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext .5pt;
  background:#DBE5F1;mso-background-themecolor:accent1;mso-background-theme=
tint:
  51;padding:0cm 5.4pt 0cm 5.4pt;height:10.5pt'>
  <p class=3Dtablecopy align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-family:"Linux Libertine";color:black;mso-themecolor:text1'>=
CRYSTALS-DILITHIUM<o:p></o:p></span></p>
  </td>
  <td width=3D98 style=3D'width:73.25pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-top-alt:solid windowtext =
1.0pt;
  background:#DBE5F1;mso-background-themecolor:accent1;mso-background-theme=
tint:
  51;padding:0cm 5.4pt 0cm 5.4pt;height:10.5pt'>
  <p class=3DMsoNormal><span lang=3DEN-US style=3D'font-size:8.0pt;line-hei=
ght:110%;
  mso-bidi-font-family:"Times New Roman";color:black;mso-color-alt:windowte=
xt;
  mso-no-proof:yes'>Lattice-based</span><span lang=3DEN-US style=3D'font-si=
ze:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";mso-no-proof:yes'=
><o:p></o:p></span></p>
  </td>
  <td width=3D94 style=3D'width:70.85pt;border-top:none;border-left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-top-alt:solid windowtext =
1.0pt;
  background:#DBE5F1;mso-background-themecolor:accent1;mso-background-theme=
tint:
  51;padding:0cm 5.4pt 0cm 5.4pt;height:10.5pt'>
  <p class=3DMsoNormal><span lang=3DEN-US style=3D'font-size:8.0pt;line-hei=
ght:110%;
  mso-bidi-font-family:"Times New Roman";color:black;mso-color-alt:windowte=
xt;
  mso-no-proof:yes'>2, 3, 5</span><span lang=3DEN-US style=3D'font-size:8.0=
pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";mso-no-proof:yes'=
><o:p></o:p></span></p>
  </td>
  <td width=3D95 style=3D'width:70.9pt;border-top:none;border-left:none;bor=
der-bottom:
  solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;mso-border-top=
-alt:
  solid windowtext 1.0pt;mso-border-left-alt:solid windowtext .5pt;mso-bord=
er-alt:
  solid windowtext .5pt;mso-border-top-alt:solid windowtext 1.0pt;backgroun=
d:
  #DBE5F1;mso-background-themecolor:accent1;mso-background-themetint:51;
  padding:0cm 5.4pt 0cm 5.4pt;height:10.5pt'>
  <p class=3DMsoNormal><span lang=3DEN-US style=3D'font-size:8.0pt;line-hei=
ght:110%;
  mso-bidi-font-family:"Times New Roman";color:black;mso-color-alt:windowte=
xt;
  mso-no-proof:yes'>1312</span><span lang=3DEN-US style=3D'font-size:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";color:black;
  mso-color-alt:windowtext'>~2592</span><span lang=3DEN-US style=3D'font-si=
ze:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";mso-no-proof:yes'=
><o:p></o:p></span></p>
  </td>
  <td width=3D76 style=3D'width:2.0cm;border-top:none;border-left:none;bord=
er-bottom:
  solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;mso-border-top=
-alt:
  solid windowtext 1.0pt;mso-border-left-alt:solid windowtext .5pt;mso-bord=
er-alt:
  solid windowtext .5pt;mso-border-top-alt:solid windowtext 1.0pt;backgroun=
d:
  #DBE5F1;mso-background-themecolor:accent1;mso-background-themetint:51;
  padding:0cm 5.4pt 0cm 5.4pt;height:10.5pt'>
  <p class=3DMsoNormal><span lang=3DEN-US style=3D'font-size:8.0pt;line-hei=
ght:110%;
  mso-bidi-font-family:"Times New Roman";color:black;mso-color-alt:windowte=
xt;
  mso-no-proof:yes'>2528</span><span lang=3DEN-US style=3D'font-size:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";color:black;
  mso-color-alt:windowtext'>~4864</span><span lang=3DEN-US style=3D'font-si=
ze:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";mso-no-proof:yes'=
><o:p></o:p></span></p>
  </td>
  <td width=3D88 style=3D'width:66.35pt;border:none;border-bottom:solid win=
dowtext 1.0pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext .5pt;
  mso-border-top-alt:solid windowtext 1.0pt;mso-border-left-alt:solid windo=
wtext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;background:#DBE5F1;mso-backgr=
ound-themecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:10=
.5pt'>
  <p class=3DMsoNormal><span lang=3DEN-US style=3D'font-size:8.0pt;line-hei=
ght:110%;
  mso-bidi-font-family:"Times New Roman";color:black;mso-color-alt:windowte=
xt;
  mso-no-proof:yes'>2420</span><span lang=3DEN-US style=3D'font-size:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";color:black;
  mso-color-alt:windowtext'>~4595</span><span lang=3DEN-US style=3D'font-si=
ze:8.0pt;
  line-height:110%;mso-bidi-font-family:"Times New Roman";mso-no-proof:yes'=
><o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:2;height:11.05pt'>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowt=
ext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext .5pt;
  background:#DBE5F1;mso-background-themecolor:accent1;mso-background-theme=
tint:
  51;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-themecolor:text1'>FALCON<o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Lattice-based</span><span lang=3DEN=
-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D94 valign=3Dtop style=3D'width:70.85pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1, 5</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D95 valign=3Dtop style=3D'width:70.9pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>897~1793</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D76 valign=3Dtop style=3D'width:2.0cm;border-top:none;border-l=
eft:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1281~2305</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border:none;border-bot=
tom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;background:#DBE5F1;mso-backgr=
ound-themecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>690~1330</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:3;height:11.05pt'>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowt=
ext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext .5pt;
  background:#DBE5F1;mso-background-themecolor:accent1;mso-background-theme=
tint:
  51;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Rainbow</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Multivariate-based</span><span
  lang=3DEN-US style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-fami=
ly:"Times New Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D94 valign=3Dtop style=3D'width:70.85pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1, 3, 5</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D95 valign=3Dtop style=3D'width:70.9pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>60192~1930600</span><span lang=3DEN=
-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D76 valign=3Dtop style=3D'width:2.0cm;border-top:none;border-l=
eft:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:#DBE5F1;mso-background-th=
emecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>64~1408736</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border:none;border-bot=
tom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;background:#DBE5F1;mso-backgr=
ound-themecolor:
  accent1;mso-background-themetint:51;padding:0cm 5.4pt 0cm 5.4pt;height:11=
.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>66~212</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:4;height:11.05pt'>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowt=
ext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext .5pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Picnic</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Hash-based</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D94 valign=3Dtop style=3D'width:70.85pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1, 3, 5</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D95 valign=3Dtop style=3D'width:70.9pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>33~65</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D76 valign=3Dtop style=3D'width:2.0cm;border-top:none;border-l=
eft:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>49~97</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border:none;border-bot=
tom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;background:white;mso-backgrou=
nd-themecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>34036~209510</span><span lang=3DEN-=
US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:5;height:11.05pt'>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowt=
ext .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;mso-border-right-alt:solid wi=
ndowtext .5pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>SPHINCS+</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Hash-based</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D94 valign=3Dtop style=3D'width:70.85pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1, 3, 5</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D95 valign=3Dtop style=3D'width:70.9pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>32~64</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D76 valign=3Dtop style=3D'width:2.0cm;border-top:none;border-l=
eft:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;background:white;mso-background-them=
ecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>64~128</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border:none;border-bot=
tom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-bottom-alt:solid windowtext .5pt;background:white;mso-backgrou=
nd-themecolor:
  background1;padding:0cm 5.4pt 0cm 5.4pt;height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>7856~49856</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
 </tr>
 <tr style=3D'mso-yfti-irow:6;mso-yfti-lastrow:yes;height:11.05pt'>
  <td width=3D101 valign=3Dtop style=3D'width:75.6pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-top-alt:solid windowt=
ext .5pt;
  mso-border-bottom-alt:solid windowtext 1.0pt;mso-border-right-alt:solid w=
indowtext .5pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span class=
=3DSpellE><span
  lang=3DEN-US style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-fami=
ly:"Times New Roman";
  color:black;mso-color-alt:windowtext'>GeMSS</span></span><span lang=3DEN-=
US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D98 valign=3Dtop style=3D'width:73.25pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-bottom-alt:solid windowte=
xt 1.0pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>Multivariate-based</span><span
  lang=3DEN-US style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-fami=
ly:"Times New Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D94 valign=3Dtop style=3D'width:70.85pt;border-top:none;border=
-left:
  none;border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1=
.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-bottom-alt:solid windowte=
xt 1.0pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>1, 3, 5</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D95 valign=3Dtop style=3D'width:70.9pt;border-top:none;border-=
left:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-bottom-alt:solid windowte=
xt 1.0pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>352190~3135590</span><span lang=3DE=
N-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D76 valign=3Dtop style=3D'width:2.0cm;border-top:none;border-l=
eft:none;
  border-bottom:solid windowtext 1.0pt;border-right:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  mso-border-alt:solid windowtext .5pt;mso-border-bottom-alt:solid windowte=
xt 1.0pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>128~256</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
  <td width=3D88 valign=3Dtop style=3D'width:66.35pt;border:none;border-bot=
tom:solid windowtext 1.0pt;
  mso-border-top-alt:solid windowtext .5pt;mso-border-left-alt:solid window=
text .5pt;
  background:white;mso-background-themecolor:background1;padding:0cm 5.4pt =
0cm 5.4pt;
  height:11.05pt'>
  <p class=3DMsoNormal align=3Dleft style=3D'text-align:left'><span lang=3D=
EN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman";
  color:black;mso-color-alt:windowtext'>258~600</span><span lang=3DEN-US
  style=3D'font-size:8.0pt;line-height:110%;mso-bidi-font-family:"Times New=
 Roman"'><o:p></o:p></span></p>
  </td>
 </tr>
</table>
