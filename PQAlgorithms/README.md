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

MIME-Version: 1.0
Content-Type: multipart/related; boundary="----=_NextPart_01D783CD.817287A0"

此文档为“单个文件网页”，也称为“Web 档案”文件。如果您看到此消息，但是您的浏览器或编辑器不支持“Web 档案”文件。请下载支持“Web 档案”的浏览器。

------=_NextPart_01D783CD.817287A0
Content-Location: file:///C:/B36127B1/file8817.htm
Content-Transfer-Encoding: quoted-printable
Content-Type: text/html; charset="gb2312"

<html xmlns:v=3D"urn:schemas-microsoft-com:vml"
xmlns:o=3D"urn:schemas-microsoft-com:office:office"
xmlns:w=3D"urn:schemas-microsoft-com:office:word"
xmlns:m=3D"http://schemas.microsoft.com/office/2004/12/omml"
xmlns=3D"http://www.w3.org/TR/REC-html40">

<head>
<meta http-equiv=3DContent-Type content=3D"text/html; charset=3Dgb2312">
<meta name=3DProgId content=3DWord.Document>
<meta name=3DGenerator content=3D"Microsoft Word 15">
<meta name=3DOriginator content=3D"Microsoft Word 15">
<link rel=3DFile-List href=3D"file8817.files/filelist.xml">
<!--[if gte mso 9]><xml>
 <o:DocumentProperties>
  <o:Author>dell</o:Author>
  <o:Template>Normal</o:Template>
  <o:LastAuthor>dell</o:LastAuthor>
  <o:Revision>2</o:Revision>
  <o:TotalTime>2</o:TotalTime>
  <o:Created>2021-07-28T08:27:00Z</o:Created>
  <o:LastSaved>2021-07-28T08:27:00Z</o:LastSaved>
  <o:Pages>1</o:Pages>
  <o:Words>79</o:Words>
  <o:Characters>453</o:Characters>
  <o:Lines>3</o:Lines>
  <o:Paragraphs>1</o:Paragraphs>
  <o:CharactersWithSpaces>531</o:CharactersWithSpaces>
  <o:Version>16.00</o:Version>
 </o:DocumentProperties>
 <o:OfficeDocumentSettings>
  <o:AllowPNG/>
 </o:OfficeDocumentSettings>
</xml><![endif]-->
<link rel=3DthemeData href=3D"file8817.files/themedata.thmx">
<link rel=3DcolorSchemeMapping href=3D"file8817.files/colorschememapping.xm=
l">
<!--[if gte mso 9]><xml>
 <w:WordDocument>
  <w:SpellingState>Clean</w:SpellingState>
  <w:GrammarState>Clean</w:GrammarState>
  <w:TrackMoves>false</w:TrackMoves>
  <w:TrackFormatting/>
  <w:PunctuationKerning/>
  <w:DrawingGridVerticalSpacing>7.8 磅</w:DrawingGridVerticalSpacing>
  <w:DisplayHorizontalDrawingGridEvery>0</w:DisplayHorizontalDrawingGridEve=
ry>
  <w:DisplayVerticalDrawingGridEvery>2</w:DisplayVerticalDrawingGridEvery>
  <w:ValidateAgainstSchemas/>
  <w:SaveIfXMLInvalid>false</w:SaveIfXMLInvalid>
  <w:IgnoreMixedContent>false</w:IgnoreMixedContent>
  <w:AlwaysShowPlaceholderText>false</w:AlwaysShowPlaceholderText>
  <w:DoNotPromoteQF/>
  <w:LidThemeOther>EN-US</w:LidThemeOther>
  <w:LidThemeAsian>ZH-CN</w:LidThemeAsian>
  <w:LidThemeComplexScript>X-NONE</w:LidThemeComplexScript>
  <w:Compatibility>
   <w:SpaceForUL/>
   <w:BalanceSingleByteDoubleByteWidth/>
   <w:DoNotLeaveBackslashAlone/>
   <w:ULTrailSpace/>
   <w:DoNotExpandShiftReturn/>
   <w:AdjustLineHeightInTable/>
   <w:BreakWrappedTables/>
   <w:SnapToGridInCell/>
   <w:WrapTextWithPunct/>
   <w:UseAsianBreakRules/>
   <w:UseWord2010TableStyleRules/>
   <w:DontGrowAutofit/>
   <w:SplitPgBreakAndParaMark/>
   <w:EnableOpenTypeKerning/>
   <w:DontFlipMirrorIndents/>
   <w:OverrideTableStyleHps/>
   <w:UseFELayout/>
  </w:Compatibility>
  <m:mathPr>
   <m:mathFont m:val=3D"Cambria Math"/>
   <m:brkBin m:val=3D"before"/>
   <m:brkBinSub m:val=3D"&#45;-"/>
   <m:smallFrac m:val=3D"off"/>
   <m:dispDef/>
   <m:lMargin m:val=3D"0"/>
   <m:rMargin m:val=3D"0"/>
   <m:defJc m:val=3D"centerGroup"/>
   <m:wrapIndent m:val=3D"1440"/>
   <m:intLim m:val=3D"subSup"/>
   <m:naryLim m:val=3D"undOvr"/>
  </m:mathPr></w:WordDocument>
</xml><![endif]--><!--[if gte mso 9]><xml>
 <w:LatentStyles DefLockedState=3D"false" DefUnhideWhenUsed=3D"false"
  DefSemiHidden=3D"false" DefQFormat=3D"false" DefPriority=3D"99"
  LatentStyleCount=3D"376">
  <w:LsdException Locked=3D"false" Priority=3D"0" QFormat=3D"true" Name=3D"=
Normal"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" QFormat=3D"true" Name=3D"=
heading 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 7"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 8"/>
  <w:LsdException Locked=3D"false" Priority=3D"9" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"heading 9"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 6"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 7"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 8"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index 9"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 7"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 8"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"toc 9"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Normal Indent"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"footnote text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"annotation text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"header"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"footer"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"index heading"/>
  <w:LsdException Locked=3D"false" Priority=3D"35" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"caption"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"table of figures"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"envelope address"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"envelope return"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"footnote reference"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"annotation reference"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"line number"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"page number"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"endnote reference"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"endnote text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"table of authorities"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"macro"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"toa heading"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Bullet"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Number"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Bullet 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Bullet 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Bullet 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Bullet 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Number 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Number 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Number 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Number 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"10" QFormat=3D"true" Name=3D=
"Title"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Closing"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Signature"/>
  <w:LsdException Locked=3D"false" Priority=3D"1" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"Default Paragraph Font"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text Indent"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Continue"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Continue 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Continue 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Continue 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"List Continue 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Message Header"/>
  <w:LsdException Locked=3D"false" Priority=3D"11" QFormat=3D"true" Name=3D=
"Subtitle"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Salutation"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Date"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text First Indent"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text First Indent 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Note Heading"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text Indent 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Body Text Indent 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Block Text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Hyperlink"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"FollowedHyperlink"/>
  <w:LsdException Locked=3D"false" Priority=3D"22" QFormat=3D"true" Name=3D=
"Strong"/>
  <w:LsdException Locked=3D"false" Priority=3D"20" QFormat=3D"true" Name=3D=
"Emphasis"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Document Map"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Plain Text"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"E-mail Signature"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Top of Form"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Bottom of Form"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Normal (Web)"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Acronym"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Address"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Cite"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Code"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Definition"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Keyboard"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Preformatted"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Sample"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Typewriter"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"HTML Variable"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"annotation subject"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"No List"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Outline List 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Outline List 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Outline List 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Simple 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Simple 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Simple 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Classic 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Classic 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Classic 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Classic 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Colorful 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Colorful 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Colorful 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Columns 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Columns 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Columns 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Columns 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Columns 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 6"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 7"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Grid 8"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 4"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 5"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 6"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 7"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table List 8"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table 3D effects 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table 3D effects 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table 3D effects 3"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Contemporary"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Elegant"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Professional"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Subtle 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Subtle 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Web 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Table Web 2"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Balloon Text"/>
  <w:LsdException Locked=3D"false" Priority=3D"0" Name=3D"Table Grid"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" Name=3D"Placeholder =
Text"/>
  <w:LsdException Locked=3D"false" Priority=3D"1" QFormat=3D"true" Name=3D"=
No Spacing"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" Name=3D"Revision"/>
  <w:LsdException Locked=3D"false" Priority=3D"34" QFormat=3D"true"
   Name=3D"List Paragraph"/>
  <w:LsdException Locked=3D"false" Priority=3D"29" QFormat=3D"true" Name=3D=
"Quote"/>
  <w:LsdException Locked=3D"false" Priority=3D"30" QFormat=3D"true"
   Name=3D"Intense Quote"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"60" Name=3D"Light Shading Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"61" Name=3D"Light List Accen=
t 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"62" Name=3D"Light Grid Accen=
t 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"63" Name=3D"Medium Shading 1=
 Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"64" Name=3D"Medium Shading 2=
 Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"65" Name=3D"Medium List 1 Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"66" Name=3D"Medium List 2 Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"67" Name=3D"Medium Grid 1 Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"68" Name=3D"Medium Grid 2 Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"69" Name=3D"Medium Grid 3 Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"70" Name=3D"Dark List Accent=
 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"71" Name=3D"Colorful Shading=
 Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"72" Name=3D"Colorful List Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"73" Name=3D"Colorful Grid Ac=
cent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"19" QFormat=3D"true"
   Name=3D"Subtle Emphasis"/>
  <w:LsdException Locked=3D"false" Priority=3D"21" QFormat=3D"true"
   Name=3D"Intense Emphasis"/>
  <w:LsdException Locked=3D"false" Priority=3D"31" QFormat=3D"true"
   Name=3D"Subtle Reference"/>
  <w:LsdException Locked=3D"false" Priority=3D"32" QFormat=3D"true"
   Name=3D"Intense Reference"/>
  <w:LsdException Locked=3D"false" Priority=3D"33" QFormat=3D"true" Name=3D=
"Book Title"/>
  <w:LsdException Locked=3D"false" Priority=3D"37" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" Name=3D"Bibliography"/>
  <w:LsdException Locked=3D"false" Priority=3D"39" SemiHidden=3D"true"
   UnhideWhenUsed=3D"true" QFormat=3D"true" Name=3D"TOC Heading"/>
  <w:LsdException Locked=3D"false" Priority=3D"41" Name=3D"Plain Table 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"42" Name=3D"Plain Table 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"43" Name=3D"Plain Table 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"44" Name=3D"Plain Table 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"45" Name=3D"Plain Table 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"40" Name=3D"Grid Table Light=
"/>
  <w:LsdException Locked=3D"false" Priority=3D"46" Name=3D"Grid Table 1 Lig=
ht"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k"/>
  <w:LsdException Locked=3D"false" Priority=3D"51" Name=3D"Grid Table 6 Col=
orful"/>
  <w:LsdException Locked=3D"false" Priority=3D"52" Name=3D"Grid Table 7 Col=
orful"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"Grid Table 1 Light Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"Grid Table 2 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"Grid Table 3 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"Grid Table 4 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"Grid Table 5 Dar=
k Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"Grid Table 6 Colorful Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"Grid Table 7 Colorful Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"46" Name=3D"List Table 1 Lig=
ht"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k"/>
  <w:LsdException Locked=3D"false" Priority=3D"51" Name=3D"List Table 6 Col=
orful"/>
  <w:LsdException Locked=3D"false" Priority=3D"52" Name=3D"List Table 7 Col=
orful"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 1"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 2"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 3"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 4"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 5"/>
  <w:LsdException Locked=3D"false" Priority=3D"46"
   Name=3D"List Table 1 Light Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"47" Name=3D"List Table 2 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"48" Name=3D"List Table 3 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"49" Name=3D"List Table 4 Acc=
ent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"50" Name=3D"List Table 5 Dar=
k Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"51"
   Name=3D"List Table 6 Colorful Accent 6"/>
  <w:LsdException Locked=3D"false" Priority=3D"52"
   Name=3D"List Table 7 Colorful Accent 6"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Mention"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Smart Hyperlink"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Hashtag"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Unresolved Mention"/>
  <w:LsdException Locked=3D"false" SemiHidden=3D"true" UnhideWhenUsed=3D"tr=
ue"
   Name=3D"Smart Link"/>
 </w:LatentStyles>
</xml><![endif]-->
<style>
<!--
 /* Font Definitions */
 @font-face
	{font-family:宋体;
	panose-1:2 1 6 0 3 1 1 1 1 1;
	mso-font-alt:SimSun;
	mso-font-charset:134;
	mso-generic-font-family:auto;
	mso-font-pitch:variable;
	mso-font-signature:3 680460288 22 0 262145 0;}
@font-face
	{font-family:"Cambria Math";
	panose-1:2 4 5 3 5 4 6 3 2 4;
	mso-font-charset:0;
	mso-generic-font-family:roman;
	mso-font-pitch:variable;
	mso-font-signature:-536869121 1107305727 33554432 0 415 0;}
@font-face
	{font-family:Calibri;
	panose-1:2 15 5 2 2 2 4 3 2 4;
	mso-font-charset:0;
	mso-generic-font-family:swiss;
	mso-font-pitch:variable;
	mso-font-signature:-469750017 -1073732485 9 0 511 0;}
@font-face
	{font-family:"Linux Libertine";
	mso-font-alt:"Times New Roman";
	mso-font-charset:0;
	mso-generic-font-family:auto;
	mso-font-pitch:variable;
	mso-font-signature:-536868097 1375790587 33554464 0 447 0;}
@font-face
	{font-family:"\@宋体";
	panose-1:2 1 6 0 3 1 1 1 1 1;
	mso-font-charset:134;
	mso-generic-font-family:auto;
	mso-font-pitch:variable;
	mso-font-signature:3 680460288 22 0 262145 0;}
 /* Style Definitions */
 p.MsoNormal, li.MsoNormal, div.MsoNormal
	{mso-style-unhide:no;
	mso-style-qformat:yes;
	mso-style-parent:"";
	margin:0cm;
	text-align:justify;
	text-justify:inter-ideograph;
	line-height:110%;
	mso-pagination:widow-orphan;
	font-size:9.0pt;
	mso-bidi-font-size:11.0pt;
	font-family:"Linux Libertine";
	mso-fareast-font-family:Calibri;
	mso-fareast-theme-font:minor-latin;
	mso-bidi-font-family:"Times New Roman";
	mso-bidi-theme-font:minor-bidi;
	mso-fareast-language:EN-US;}
p.MsoHeader, li.MsoHeader, div.MsoHeader
	{mso-style-priority:99;
	mso-style-link:"页眉 字符";
	margin:0cm;
	text-align:center;
	line-height:110%;
	mso-pagination:widow-orphan;
	tab-stops:center 207.65pt right 415.3pt;
	layout-grid-mode:char;
	border:none;
	mso-border-bottom-alt:solid windowtext .75pt;
	padding:0cm;
	mso-padding-alt:0cm 0cm 1.0pt 0cm;
	font-size:9.0pt;
	font-family:"Linux Libertine";
	mso-fareast-font-family:Calibri;
	mso-fareast-theme-font:minor-latin;
	mso-bidi-font-family:"Times New Roman";
	mso-bidi-theme-font:minor-bidi;
	mso-fareast-language:EN-US;}
p.MsoFooter, li.MsoFooter, div.MsoFooter
	{mso-style-priority:99;
	mso-style-link:"页脚 字符";
	margin:0cm;
	line-height:110%;
	mso-pagination:widow-orphan;
	tab-stops:center 207.65pt right 415.3pt;
	layout-grid-mode:char;
	font-size:9.0pt;
	font-family:"Linux Libertine";
	mso-fareast-font-family:Calibri;
	mso-fareast-theme-font:minor-latin;
	mso-bidi-font-family:"Times New Roman";
	mso-bidi-theme-font:minor-bidi;
	mso-fareast-language:EN-US;}
span.a
	{mso-style-name:"页眉 字符";
	mso-style-priority:99;
	mso-style-unhide:no;
	mso-style-locked:yes;
	mso-style-link:页眉;
	mso-ansi-font-size:9.0pt;
	mso-bidi-font-size:9.0pt;}
span.a0
	{mso-style-name:"页脚 字符";
	mso-style-priority:99;
	mso-style-unhide:no;
	mso-style-locked:yes;
	mso-style-link:页脚;
	mso-ansi-font-size:9.0pt;
	mso-bidi-font-size:9.0pt;}
p.tablecolsubhead, li.tablecolsubhead, div.tablecolsubhead
	{mso-style-name:"table col subhead";
	mso-style-unhide:no;
	margin:0cm;
	text-align:center;
	mso-pagination:widow-orphan;
	font-size:7.5pt;
	font-family:"Times New Roman",serif;
	mso-fareast-font-family:宋体;
	mso-fareast-language:EN-US;
	font-weight:bold;
	font-style:italic;}
p.tablecopy, li.tablecopy, div.tablecopy
	{mso-style-name:"table copy";
	mso-style-unhide:no;
	mso-style-parent:"";
	margin:0cm;
	text-align:justify;
	text-justify:inter-ideograph;
	mso-pagination:widow-orphan;
	font-size:8.0pt;
	font-family:"Times New Roman",serif;
	mso-fareast-font-family:宋体;
	mso-fareast-language:EN-US;
	mso-no-proof:yes;}
span.SpellE
	{mso-style-name:"";
	mso-spl-e:yes;}
.MsoChpDefault
	{mso-style-type:export-only;
	mso-default-props:yes;
	font-family:"Calibri",sans-serif;
	mso-bidi-font-family:"Times New Roman";
	mso-bidi-theme-font:minor-bidi;}
 /* Page Definitions */
 @page
	{mso-page-border-surround-header:no;
	mso-page-border-surround-footer:no;
	mso-footnote-separator:url("file8817.files/header.htm") fs;
	mso-footnote-continuation-separator:url("file8817.files/header.htm") fcs;
	mso-endnote-separator:url("file8817.files/header.htm") es;
	mso-endnote-continuation-separator:url("file8817.files/header.htm") ecs;}
@page WordSection1
	{size:595.3pt 841.9pt;
	margin:72.0pt 90.0pt 72.0pt 90.0pt;
	mso-header-margin:42.55pt;
	mso-footer-margin:49.6pt;
	mso-paper-source:0;
	layout-grid:15.6pt;}
div.WordSection1
	{page:WordSection1;}
-->
</style>
<!--[if gte mso 10]>
<style>
 /* Style Definitions */
 table.MsoNormalTable
	{mso-style-name:普通表格;
	mso-tstyle-rowband-size:0;
	mso-tstyle-colband-size:0;
	mso-style-noshow:yes;
	mso-style-priority:99;
	mso-style-parent:"";
	mso-padding-alt:0cm 5.4pt 0cm 5.4pt;
	mso-para-margin:0cm;
	mso-pagination:widow-orphan;
	font-size:10.5pt;
	mso-bidi-font-size:11.0pt;
	font-family:"Calibri",sans-serif;
	mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;
	mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin;
	mso-bidi-font-family:"Times New Roman";
	mso-bidi-theme-font:minor-bidi;
	mso-font-kerning:1.0pt;}
table.MsoTableGrid
	{mso-style-name:网格型;
	mso-tstyle-rowband-size:0;
	mso-tstyle-colband-size:0;
	mso-style-unhide:no;
	border:solid windowtext 1.0pt;
	mso-border-alt:solid windowtext .5pt;
	mso-padding-alt:0cm 5.4pt 0cm 5.4pt;
	mso-border-insideh:.5pt solid windowtext;
	mso-border-insidev:.5pt solid windowtext;
	mso-para-margin:0cm;
	mso-pagination:widow-orphan;
	font-size:10.0pt;
	font-family:"Times New Roman",serif;
	mso-fareast-font-family:"Times New Roman";
	mso-fareast-language:EN-US;}
</style>
<![endif]--><!--[if gte mso 9]><xml>
 <o:shapedefaults v:ext=3D"edit" spidmax=3D"2049"/>
</xml><![endif]--><!--[if gte mso 9]><xml>
 <o:shapelayout v:ext=3D"edit">
  <o:idmap v:ext=3D"edit" data=3D"1"/>
 </o:shapelayout></xml><![endif]-->
</head>

<body lang=3DZH-CN style=3D'tab-interval:21.0pt;word-wrap:break-word;text-j=
ustify-trim:
punctuation'>

<div class=3DWordSection1 style=3D'layout-grid:15.6pt'>

<div align=3Dcenter>

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

</div>

<p class=3DMsoNormal><span lang=3DEN-US><o:p>&nbsp;</o:p></span></p>

</div>

</body>

</html>

------=_NextPart_01D783CD.817287A0
Content-Location: file:///C:/B36127B1/file8817.files/themedata.thmx
Content-Transfer-Encoding: base64
Content-Type: application/vnd.ms-officetheme

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

------=_NextPart_01D783CD.817287A0
Content-Location: file:///C:/B36127B1/file8817.files/colorschememapping.xml
Content-Transfer-Encoding: quoted-printable
Content-Type: text/xml

<?xml version=3D"1.0" encoding=3D"UTF-8" standalone=3D"yes"?>
<a:clrMap xmlns:a=3D"http://schemas.openxmlformats.org/drawingml/2006/main"=
 bg1=3D"lt1" tx1=3D"dk1" bg2=3D"lt2" tx2=3D"dk2" accent1=3D"accent1" accent=
2=3D"accent2" accent3=3D"accent3" accent4=3D"accent4" accent5=3D"accent5" a=
ccent6=3D"accent6" hlink=3D"hlink" folHlink=3D"folHlink"/>
------=_NextPart_01D783CD.817287A0
Content-Location: file:///C:/B36127B1/file8817.files/header.htm
Content-Transfer-Encoding: quoted-printable
Content-Type: text/html; charset="gb2312"

<html xmlns:v=3D"urn:schemas-microsoft-com:vml"
xmlns:o=3D"urn:schemas-microsoft-com:office:office"
xmlns:w=3D"urn:schemas-microsoft-com:office:word"
xmlns:m=3D"http://schemas.microsoft.com/office/2004/12/omml"
xmlns=3D"http://www.w3.org/TR/REC-html40">

<head>
<meta http-equiv=3DContent-Type content=3D"text/html; charset=3Dgb2312">
<meta name=3DProgId content=3DWord.Document>
<meta name=3DGenerator content=3D"Microsoft Word 15">
<meta name=3DOriginator content=3D"Microsoft Word 15">
<link id=3DMain-File rel=3DMain-File href=3D"../file8817.htm">
<![if IE]>
<base href=3D"file:///C:/B36127B1/file8817.files/header.htm"
id=3D"webarch_temp_base_tag">
<![endif]><!--[if gte mso 9]><xml>
 <o:shapedefaults v:ext=3D"edit" spidmax=3D"2049"/>
</xml><![endif]-->
</head>

<body lang=3DZH-CN>

<div style=3D'mso-element:footnote-separator' id=3Dfs>

<p class=3DMsoNormal style=3D'line-height:normal'><span lang=3DEN-US><span
style=3D'mso-special-character:footnote-separator'><![if !supportFootnotes]>

<hr align=3Dleft size=3D1 width=3D"33%">

<![endif]></span></span></p>

</div>

<div style=3D'mso-element:footnote-continuation-separator' id=3Dfcs>

<p class=3DMsoNormal style=3D'line-height:normal'><span lang=3DEN-US><span
style=3D'mso-special-character:footnote-continuation-separator'><![if !supp=
ortFootnotes]>

<hr align=3Dleft size=3D1>

<![endif]></span></span></p>

</div>

<div style=3D'mso-element:endnote-separator' id=3Des>

<p class=3DMsoNormal style=3D'line-height:normal'><span lang=3DEN-US><span
style=3D'mso-special-character:footnote-separator'><![if !supportFootnotes]>

<hr align=3Dleft size=3D1 width=3D"33%">

<![endif]></span></span></p>

</div>

<div style=3D'mso-element:endnote-continuation-separator' id=3Decs>

<p class=3DMsoNormal style=3D'line-height:normal'><span lang=3DEN-US><span
style=3D'mso-special-character:footnote-continuation-separator'><![if !supp=
ortFootnotes]>

<hr align=3Dleft size=3D1>

<![endif]></span></span></p>

</div>

</body>

</html>

------=_NextPart_01D783CD.817287A0
Content-Location: file:///C:/B36127B1/file8817.files/filelist.xml
Content-Transfer-Encoding: quoted-printable
Content-Type: text/xml; charset="utf-8"

<xml xmlns:o=3D"urn:schemas-microsoft-com:office:office">
 <o:MainFile HRef=3D"../file8817.htm"/>
 <o:File HRef=3D"themedata.thmx"/>
 <o:File HRef=3D"colorschememapping.xml"/>
 <o:File HRef=3D"header.htm"/>
 <o:File HRef=3D"filelist.xml"/>
</xml>
------=_NextPart_01D783CD.817287A0--
