package aoc2019

const day01myInput = `109067
75007
66030
93682
83818
108891
139958
129246
80272
119897
112804
69495
95884
85402
148361
75986
120063
127683
146962
76907
61414
98452
134330
53858
82662
143258
82801
60279
131782
105989
102464
96563
71172
113731
90645
94830
133247
110149
54792
134863
125919
145490
69836
108808
87954
148957
110182
126668
148024
96915
117727
147378
75967
91915
60130
85331
66800
103419
72627
72687
61606
113160
107082
110793
61589
105005
73952
65705
117243
140944
117091
113482
91379
148185
113853
119822
78179
85407
119886
109230
68783
63914
51101
93549
53361
127984
106315
54997
138941
81075
120272
120307
98414
115245
105649
89793
88421
121104
97084
56928
`

const day02myInput = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,5,19,23,1,6,23,27,1,27,10,31,1,31,5,35,2,10,35,39,1,9,39,43,1,43,5,47,1,47,6,51,2,51,6,55,1,13,55,59,2,6,59,63,1,63,5,67,2,10,67,71,1,9,71,75,1,75,13,79,1,10,79,83,2,83,13,87,1,87,6,91,1,5,91,95,2,95,9,99,1,5,99,103,1,103,6,107,2,107,13,111,1,111,10,115,2,10,115,119,1,9,119,123,1,123,9,127,1,13,127,131,2,10,131,135,1,135,5,139,1,2,139,143,1,143,5,0,99,2,0,14,0`

const day03myInput = `R1006,U867,R355,D335,L332,U787,L938,U987,L234,U940,R393,D966,R57,U900,R619,D693,L606,U272,L45,D772,R786,U766,R860,U956,L346,D526,R536,D882,L156,U822,L247,D279,R515,U467,R208,D659,R489,D295,R18,D863,L360,D28,R674,U203,L276,U518,L936,D673,L501,D414,L635,U497,R402,D530,L589,D247,L140,U697,R626,D130,L109,D169,L316,D2,R547,D305,L480,U871,R551,D48,L710,D655,R562,D395,L925,D349,L795,U308,L861,D265,R88,U116,L719,D204,R995,D197,R167,U786,R459,U978,L506,D232,L37,U530,L808,D75,R79,D469,L851,D152,R793,D362,L293,D760,L422,U516,L400,D275,L498,U877,R202,D302,L89,U924,L55,U161,L945,D578,R861,U853,R358,D427,L776,U571,R670,D29,R52,D116,R879,U359,R493,D872,L567,U382,R804,D168,R316,D376,R711,U627,R36,D241,R876,U407,L481,D360,R679,U561,L947,U449,R232,U176,R677,U850,R165,D257,R683,D666,L31,U237,L906,U810,R198,U890,L462,D928,R915,D778,L689,U271,L486,D918,L995,U61,R748,U516,L80,D109,L328,U649,L784,D546,R584,D751,L543,U217,L976,D400,L795,U332,R453,U399,L761,U823,R142,U532,R133,U255,R722,D726,L862,D845,L813,U981,R272,D800,L918,D712,R259,U972,R323,D214,R694,D629,L817,D814,L741,U111,L678,D627,L743,D509,R195,U875,R46,D344,L361,D102,L656,U897,L841,U124,L95,D770,L785,U767,L504,D309,L955,D142,L401,U914,R117,D897,R715,D117,R675,U248,R182,D725,L751,U562,R385,D120,R730,U185,L842,D446,L432,U640,R994,D482,R576,U915,R645,U109,R77,D983,L327,D209,R686,D486,R566,D406,R130,D759,R441,U895,R597,U443,L773,D704,R219,U222,R244,D11,L723,U804,L264,D121,L81,D454,R279,D642,L773,D653,R127,D199,R119,U509,L530
L1003,D933,L419,D202,L972,U621,L211,U729,R799,U680,R925,U991,L167,U800,R198,U214,R386,D385,R117,D354,L914,D992,L519,U797,L28,D125,R163,D894,R390,D421,L75,D577,L596,U95,L403,U524,L160,D39,R209,D373,L464,U622,L824,D750,L857,U658,L109,D188,R357,D295,L227,U904,L268,U814,L483,U897,R785,U194,R865,U300,L787,U812,L321,D637,R761,U560,R800,U281,R472,D283,L490,D629,L207,D589,L310,D980,R613,U129,R668,U261,R82,D594,R627,D210,L865,U184,R387,U995,R497,U68,L776,U657,R559,D38,R981,D485,L196,D934,R313,D128,R269,D225,L32,U677,R425,U728,L665,D997,R271,D847,R715,U300,L896,D481,L30,U310,L793,D600,L219,D944,R197,D945,L564,D603,L225,U413,L900,U876,R281,D26,R449,D506,L846,D979,L817,D794,R309,D841,R735,U11,R373,U530,R74,D534,L668,U185,L972,D436,L377,D164,L88,U249,L8,D427,R711,D530,L850,D921,L644,U804,L388,U500,L813,D223,L572,U246,R309,U241,R185,D48,L545,U678,L344,D964,L772,D985,L178,U686,R937,U821,R214,D346,L648,D824,L943,D651,R98,D225,R832,D883,L814,D894,L995,D975,R440,D502,L177,D320,R675,U5,R188,D866,R974,U169,R432,D627,L424,D5,L273,U184,R657,U830,R681,U610,R170,U106,L726,D861,L257,D381,L918,D607,L820,D757,R556,D621,R21,U510,L575,D545,L590,D302,R446,D225,L164,D817,L520,D204,L33,U272,L648,D478,R945,U369,L924,D932,R46,D584,R630,U592,R613,U136,R253,D343,L983,U328,L442,D311,L258,U173,L574,U658,R283,D927,L247,D37,R36,D61,L692,U663,L207,U48,L114,U511,L924,U229,L221,D504,R345,U51,R464,D516,L115,D311,L71,D418,R378,D173,R154,U436,L403,D871,L765,D803,R630,U108,L79,U625,R77,U176,R911`

const day04myInput = `271973-785961`

const day05myInput = `3,225,1,225,6,6,1100,1,238,225,104,0,1,191,196,224,1001,224,-85,224,4,224,1002,223,8,223,1001,224,4,224,1,223,224,223,1101,45,50,225,1102,61,82,225,101,44,39,224,101,-105,224,224,4,224,102,8,223,223,101,5,224,224,1,224,223,223,102,14,187,224,101,-784,224,224,4,224,102,8,223,223,101,7,224,224,1,224,223,223,1001,184,31,224,1001,224,-118,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1102,91,18,225,2,35,110,224,101,-810,224,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1101,76,71,224,1001,224,-147,224,4,224,102,8,223,223,101,2,224,224,1,224,223,223,1101,7,16,225,1102,71,76,224,101,-5396,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1101,72,87,225,1101,56,77,225,1102,70,31,225,1102,29,15,225,1002,158,14,224,1001,224,-224,224,4,224,102,8,223,223,101,1,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1007,226,226,224,1002,223,2,223,1006,224,329,1001,223,1,223,8,226,677,224,1002,223,2,223,1005,224,344,1001,223,1,223,107,226,677,224,1002,223,2,223,1006,224,359,1001,223,1,223,8,677,677,224,1002,223,2,223,1005,224,374,1001,223,1,223,1108,226,226,224,1002,223,2,223,1005,224,389,1001,223,1,223,7,677,226,224,1002,223,2,223,1005,224,404,101,1,223,223,7,226,226,224,102,2,223,223,1006,224,419,1001,223,1,223,1108,226,677,224,102,2,223,223,1005,224,434,1001,223,1,223,1107,226,226,224,1002,223,2,223,1006,224,449,1001,223,1,223,1007,677,677,224,102,2,223,223,1006,224,464,1001,223,1,223,107,226,226,224,1002,223,2,223,1005,224,479,101,1,223,223,1107,677,226,224,1002,223,2,223,1005,224,494,1001,223,1,223,1008,677,677,224,102,2,223,223,1005,224,509,101,1,223,223,107,677,677,224,102,2,223,223,1005,224,524,1001,223,1,223,1108,677,226,224,1002,223,2,223,1005,224,539,1001,223,1,223,7,226,677,224,102,2,223,223,1006,224,554,1001,223,1,223,8,677,226,224,1002,223,2,223,1006,224,569,101,1,223,223,108,226,226,224,1002,223,2,223,1006,224,584,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,599,101,1,223,223,1008,226,226,224,102,2,223,223,1005,224,614,1001,223,1,223,1007,226,677,224,1002,223,2,223,1006,224,629,1001,223,1,223,108,677,226,224,102,2,223,223,1005,224,644,101,1,223,223,1008,226,677,224,1002,223,2,223,1005,224,659,101,1,223,223,108,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226`

const day06sampleInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
const day06myInput = `NKB)PZS
KBG)9JH
PZS)KZD
8J1)GT7
7HT)9ZM
RV2)14G
RM1)YG7
36C)BPH
BY7)5RG
PYW)RWT
VHV)DCN
8GP)3DJ
PWQ)KNW
WRQ)R4M
PYK)WQ7
C47)RSZ
KVF)879
M8B)TRX
RL1)N46
H7D)5CD
FZM)T6P
99X)DN8
1M7)ZWZ
2XZ)C9G
K9Q)CLK
KVF)WGY
S2F)3XK
LRC)DHQ
3VV)549
XFJ)YV2
764)7PZ
RV2)2G9
CGP)CZ5
PZP)WLC
NKB)ZLV
6YS)WHG
547)HYD
CPP)GNM
LNZ)T4D
WLC)55F
GNP)5XG
RT1)TP3
223)51G
GDY)2DF
CDN)9CS
JVW)N4N
4C4)FYN
C27)3NK
R1Z)ZD9
2RV)4GR
YR3)H8J
67C)46J
RQP)35J
T3G)FTR
QHT)8JG
JML)H46
74P)KTH
SC9)CDN
633)WRD
6WJ)NZQ
JQS)QHT
FHW)QL9
347)XLW
H25)5QX
M6Y)MTC
F1D)XZN
NJ6)D26
NHJ)B87
8H5)FF9
DNL)9K9
JM7)GNN
BRD)JGW
WGH)VM1
T4F)33Y
MC6)7M9
LKQ)N5L
WGY)Q4F
XL4)SD8
N7W)J59
TDX)8N3
M9S)64V
NNT)WPY
C9G)RSM
J8M)WDS
5X4)346
KGG)TZR
4LY)JST
5MM)MCX
PG8)6S4
G4R)NNT
FYY)SFS
Q1B)MLQ
N9F)B7F
B5F)452
57Y)TQ3
NCC)2ZZ
XQ2)18R
G7V)7SQ
93T)SR1
T1X)9WB
46W)GDY
Q74)RP1
B1F)GV4
KR9)Q1J
ZZZ)4NL
7PZ)NSK
G5Q)KCQ
6K7)2W6
GRL)VTL
RWY)G7B
LB4)DT4
QZW)51L
NV5)D8H
49G)4NQ
XZC)4RB
Y9S)7XV
4GT)N9M
39K)7X1
W83)JQR
547)KLC
4GT)2LM
S4F)RN9
BBL)47N
JYJ)W2D
F16)3H9
K5F)1XC
JPY)6RY
9TW)5X1
9QS)P5L
CZ7)PZ6
R3X)M6Y
JQ6)CN3
R3M)PTR
X6Y)8RW
YB8)X6L
HZQ)GRL
P5J)DNS
7C1)SX8
JF5)72W
MYG)XZC
BRD)5XK
QJV)P4N
SKC)7BF
67C)PXD
LCC)7R1
BMC)BCT
WW3)145
W4D)VT3
8L2)W4W
M85)GQ4
P77)HZ2
DN1)1P2
493)6R8
2K4)NZL
NZ5)NV5
SBZ)W7Q
SMT)DP7
L5H)BMC
W7Q)K7X
3QT)GPV
ZDQ)KQ6
8LG)PNS
3XG)3XT
KFQ)X3K
3QT)D9X
PL5)WH6
PYX)HNG
D7G)VC5
NM2)39K
RRN)LZ5
342)V9B
1F5)S8L
WHG)LS5
3D3)XSD
1J9)MNR
LNH)LBK
MZR)L5P
M3T)333
X4N)7VX
SCH)VJP
5XJ)CMS
CBY)DQD
PWQ)VZB
1H9)8H5
MFN)C29
PLY)VQ3
GQP)T85
K7X)891
JZD)HLQ
FMM)D9Z
P85)T1Y
NWJ)764
L5P)VWM
MVS)9TW
VQ6)P3D
CXV)L2K
46J)3PL
SDY)VQ6
8LT)77D
8GR)XBF
17Z)NRH
T9T)3KD
D9Z)49G
BFV)HMX
XBF)L28
BDT)LWN
6R8)B8J
G1V)9VK
V9B)VSQ
WKG)4GD
879)ZHH
SFS)VR7
B2H)88T
8HB)XMH
YSN)Z5G
G8R)HP2
KHT)57D
RS2)74P
4Q9)KG8
G85)33L
FJH)CSD
5Z8)JQS
WPX)6XQ
LNZ)5GG
PVS)Y28
253)5Z8
KMP)8ZL
16H)GXP
47N)JQ6
SX8)PHF
WM1)JPY
TRS)8KG
W7Q)S83
QWL)ZXH
SQ4)KJ7
J2S)5ZY
51G)CGP
HZ2)VL2
5PL)WGR
QTY)LYQ
8TJ)2BP
MJK)RW8
P4N)F7R
KPS)JJJ
KYX)7S3
VTL)4QW
2LM)BS2
YDH)M1H
VL2)5MH
99P)B3C
33Y)QG3
RWT)K4F
M4J)8FL
QL5)RT3
8QT)RV2
R6Y)MFN
ZFT)KSJ
CSD)W8J
WH6)PZP
7S3)KZX
5LK)QC5
VT3)9XG
6H4)F6S
VR7)W4D
43J)4X6
638)389
PVT)QWL
GG4)BFL
D87)SC9
KVS)NN1
81W)NP7
Z4K)RM1
5XK)KND
CZ5)GXC
WCZ)NHP
ZBC)5PL
RSM)DQG
L9Q)4LY
QQY)5BK
1Z6)JKF
6XQ)ZPT
T57)HHX
DB4)8L2
DNS)M85
T4V)XD4
VY3)9NQ
59V)F1D
PZ6)HXZ
P5L)F49
XMF)LNH
472)FBN
8G2)372
959)93T
97T)H13
2W6)4R9
D26)LKQ
34D)F6Q
NZ5)THG
D36)HX3
Z9P)QKJ
G1X)V8D
7WF)MVS
N6J)Z1L
Y3Q)JZY
VJP)9PF
ZNV)BZ6
V8D)DKJ
R7K)GHD
6D7)GNP
BHN)F24
Y32)YSW
MJD)T4Z
54M)H8Y
NWJ)5XT
LQ2)99P
LR8)D2R
JZD)ZWG
LDC)R9L
FF9)CMM
NWL)Q5R
11B)Q2Q
Q1J)V81
YSW)JDJ
KYX)S5R
6SK)6Y4
S8L)1H9
DCB)142
6RW)995
777)J13
QN7)83F
5NL)SC3
DTY)15Z
Q85)R3T
XZN)YH4
XB9)6RB
GV4)5RV
6F3)V9P
CN3)TLT
2BP)9CW
Q8R)SYP
Z5V)GX4
7ND)2XZ
8PS)7D7
1R9)T4F
ZWZ)QL4
KPM)Z32
3XK)1F5
87K)FNM
TFQ)L5J
9QQ)K5F
KCQ)717
KQV)16H
PJ9)MW1
35J)ZMQ
LFD)CR6
XGK)9QQ
ZWG)K2F
DXR)31S
46R)XZM
C29)T57
8J2)66C
83F)SSF
ZHN)JKT
N46)ZNV
7SQ)P2N
XK3)839
ZH8)7LG
D45)PMJ
QNN)QGN
NSJ)54M
F8H)KJD
COM)1Z6
3T1)H1Z
31D)XGK
3R5)3CF
YB2)KPS
JJJ)57Y
RPZ)W43
VV1)M8D
7WF)5XR
BG5)GBM
RT3)C18
TSX)31B
46C)RRN
J3X)PFZ
2FG)MPX
W4W)2V1
JWZ)Q8R
8CF)DN1
MLQ)81P
DMN)SM7
D1D)8TJ
XVW)RWY
KQ6)Y91
M3P)HY7
VDR)1ZS
SXD)J9M
Q51)7C4
Z1X)3N1
PXD)QVX
T81)942
7C4)2Y3
DMN)6YS
KMD)T28
15H)KSF
2DF)FJH
86S)R8G
WRD)YB8
G58)7C1
6WR)C23
BK9)LMM
HP2)2BY
FZ3)H2C
TJC)BBL
9YQ)L9N
4VR)DZ4
PP1)P65
YG7)BZH
QKJ)QB7
R3G)FZ3
XF7)D45
47N)PL5
PVT)KHT
HXZ)X6Y
LK2)QBF
6WJ)7ND
7D5)KLK
346)R24
551)Z4K
CZN)QN7
BGP)FBY
4X6)LCX
31B)DQF
N9M)S2S
DP7)8M2
4TC)GJT
GJZ)JHD
QHB)W8V
88T)N47
452)S8Q
V2B)SCH
VYJ)N75
4VJ)XRS
VM9)1VD
HX3)D78
SZ4)Q74
TLQ)G8R
29Y)YSN
NGK)14N
B4J)XF2
NWY)F8H
47X)D1D
QF5)QYX
K2F)G81
31S)MF8
L19)SJN
KTH)LLF
FMT)JF5
XW3)Y32
ZB7)CW4
5BF)RCX
WF8)756
ZXK)36C
3KD)VRR
MJB)HRY
F8T)7LQ
5XG)9SV
T34)H3W
1WQ)BZ7
TS8)8BC
8LF)69Q
615)R9C
SB8)W8P
612)BGP
9SV)M1T
SFS)PJ9
FYN)66G
1LF)YQ5
WLV)JM7
6XL)Z84
CT9)SZ4
2V1)V2B
8PS)6K7
9ZM)3VM
JYP)6TT
MPX)ZB7
9BM)M9S
CMM)KNC
X27)DMN
SR1)ZSF
59V)ZSD
JKF)DB4
B3V)X6K
RG4)DMB
9Z5)TZ8
T28)BS7
T8M)KGD
CQR)JZD
5QD)2M3
K8F)L19
XF2)GNL
MP5)12X
MFK)MJD
BGR)HCD
GPV)F32
2Y3)RBD
GQ4)F2S
342)G5Q
2BR)81W
6S4)46W
ZLV)TRS
WSD)9ZN
1TS)9MZ
VW9)YOU
XWM)M7F
VRR)G68
JFZ)X27
32L)HSW
D5V)M3T
HQF)G8W
N5L)MFK
27Z)NWL
H2C)M8B
HVP)B5F
DN8)P85
RP1)C5Y
8KG)75C
3M4)XL4
3B2)WCZ
X8Z)Y9D
NZQ)HCV
142)Y3Q
14G)3GJ
X8L)SKC
RTJ)JX8
ZFD)G8P
G8W)3T5
F7K)SL3
9VK)BFV
PHF)55C
RND)TVX
WSS)4GT
MHR)G35
8KQ)1Q8
JMF)G85
R24)H67
CZ7)65P
NDQ)Z93
QG3)L4J
QC5)GVX
711)2ZQ
Z32)BY7
T1Y)547
WK8)RFD
KGZ)2FG
51G)551
H8Y)GNW
GT7)XMF
JLZ)G9H
RZH)QBB
TZX)X54
6HT)N4W
FNM)FFD
PCG)SZD
4GN)SBT
BBL)F6Z
CWC)XJ9
W9D)NZ5
KNW)SSV
D2R)KWC
3H9)FCY
145)B8Z
TYT)3FH
6W9)BGR
5MH)L9Q
2CD)VYJ
6V2)JZ2
QBF)J4X
XH1)X8R
NN1)ZNF
ZX9)DHP
M7F)V12
TVX)NJC
4LQ)YY7
3VM)R7K
KCP)MYG
YH4)8LF
TP3)4MC
CN9)MRC
PJ9)NK2
RBW)VP8
XLW)T68
QL9)QNN
DQF)LW6
GHD)MRY
S53)638
839)X8L
7WG)KV4
33L)ZBC
LTY)HPV
T4D)PVT
WMQ)WRQ
DJ3)347
3RJ)1KL
QVX)SXD
LQW)JLZ
6YQ)XB9
L2K)Y9K
TH1)9PW
ZMB)6HF
15K)9YQ
1JD)4TJ
35P)PG8
H63)1HW
TGV)6QV
TZR)6D7
VM1)6MW
7LG)WVQ
8LX)MDS
4L5)DCB
Q5R)1LR
G8P)RPZ
8FX)ZQ2
WST)NMF
452)L6D
SL3)Z5V
BFK)K3W
KGD)PLY
8M2)DXW
JKZ)47X
VV5)HLF
8MM)QQ2
XSD)CZN
7XV)D8Q
MRD)HZQ
N29)DP8
MS4)35L
G8H)RD3
QL6)7Z8
SZ6)RH5
9HW)H9D
SJN)LLP
KSR)27Z
H46)G7V
K9L)P46
QDL)633
66C)2Z9
PP1)PYW
W8V)RYK
HL5)NGK
18R)Y9S
DL9)RJK
LWN)KYX
HJT)RL1
BLY)342
6TP)LJ2
ZG8)SKB
25Y)KHV
K8X)Y95
1L5)3D3
DWW)T3G
1P2)BMB
XRS)NJ6
K4F)M4J
P2N)KRT
4N8)57V
WLS)3R5
RCX)9HW
7K8)T34
D9X)MZR
9ZN)4X7
RZ4)LQ2
JKT)JGF
LW4)7HT
D8Q)281
8BC)ZDQ
SJN)BRD
5NL)RX9
3G7)D87
9K9)BR9
7X1)T4V
65P)V3M
845)1L5
8P7)XPF
G8H)WLS
SRB)2P2
QB7)G8H
DKJ)LB4
3DJ)D36
8JG)ZX9
RMX)SLV
RQC)QZW
58D)JYJ
THG)GG4
5CD)YQN
Y28)FWP
78H)8LP
XDF)KMD
LYT)CFR
281)2M2
2N1)W9D
6R7)6SK
42D)8NV
KRT)6YQ
RP8)W83
SBT)DNL
MMY)TLQ
GNM)R3M
8PV)YDH
6PL)5Y8
8RW)11B
HG2)Q1B
XYR)GGV
K4J)35X
RSZ)Z56
56R)P9X
P9X)96H
H63)4QK
XMH)WMQ
3VP)6R7
F32)MP5
ZMB)PLP
B7H)26Y
TLT)KFQ
VJ5)223
5ZP)8KQ
GXP)T8M
6TT)LH9
XSW)LFD
RBD)612
KQW)V3W
MNR)9FZ
W8J)ZHS
4PT)ZHN
MHG)8YQ
J9M)B25
KGG)845
Y81)SAN
3NK)RDX
7Z8)SG7
4NQ)PP1
RYK)5LK
VWM)7SC
KXZ)1J9
KSY)YJM
KSJ)Q51
ZQ2)8LG
F24)6WJ
Q4F)KQX
BPH)C27
2M2)JML
6QV)QTY
F3Z)2NF
JRS)TH1
RH5)34D
GNV)493
QMY)Y5G
2NF)R34
N4N)JWZ
LH9)TSX
SJG)8Z2
NRH)H63
YGZ)6L9
NP7)7L6
RY7)MS4
55F)HW4
GCP)8LT
YQH)QJV
GH5)T76
RVC)R3X
PRG)B1F
8N3)631
J59)WFN
JLH)15H
QNN)6H4
2K9)LK2
S8Q)T81
DHP)SBZ
69Q)SW1
J7G)KGG
GBM)D4T
35X)GWC
LQW)5ZP
9PW)QMY
YY7)VV1
SD8)XDF
TNH)9YX
3CF)X8J
1S8)Q85
YJM)LQW
SVZ)2JH
FFD)D9L
CN9)M3P
KG8)JLH
YTC)F17
JHD)LCC
372)LR8
PZP)8J2
5Y8)HL5
DCN)MJK
HNG)QNX
NHP)SB9
4DQ)GYD
84T)3G5
5RV)LT2
GXC)ZG8
8ZL)G4R
Q82)XN5
HQ8)KMF
F2S)777
L9Q)K4J
N7Y)KCP
M6Y)JVW
7LM)Y99
418)5MM
12X)LYT
1KL)J8M
JQR)3V7
VQ3)K11
TQ3)TZP
2M3)G1P
C23)DWW
2G9)RVC
GJT)KW6
Z32)BLY
K11)J9Y
Z84)4Z4
V3M)G1V
BMT)5NL
DQF)78C
6R7)WSS
LLF)QF5
9HP)X4N
P46)J12
FRL)YV3
KND)L5H
SSF)WLV
M8B)1S8
YDZ)RT8
B87)TJC
WQ7)PYX
CW3)87K
45Q)MQ5
VH4)3JY
TFB)S53
KQX)GFM
H3W)BNH
KGD)MJB
995)8GR
ZSD)B2H
75C)XF7
W29)3VP
2GY)GTN
SY3)XRB
Y5G)VYP
JGW)R7Y
DP2)PVS
999)6W9
XD4)7GJ
TSW)C1Y
R34)MMY
5BK)9SL
51L)G1X
VC5)9P5
8J1)K9L
WFD)15K
LCX)NSJ
ZD9)Y1J
38S)3VV
4HD)NPY
79C)1M7
5LK)8G2
D9L)LCZ
QD5)84Z
RL1)VY3
BBH)6QK
BGD)XYR
N47)KBG
S83)DJ3
4CC)NKB
BPX)TZX
HLF)8GW
3KB)KMP
LJ2)LPV
3H3)RQC
KGT)615
KMF)DP2
BZ7)JYP
GX4)MD5
GJZ)WKG
MTC)TSW
7D7)P5J
HGR)ZZZ
4MC)86S
RJW)35P
W8P)XH1
VZB)G1J
JRC)YTC
6MW)KSR
QMY)45Q
F6S)3VZ
631)YB2
VWN)MC6
YV3)B8K
Q76)1HB
PN4)XW3
NJC)ZFT
4XQ)KSY
MRY)9KV
S8L)BL8
BR9)711
R8G)J1S
57V)4HD
68Q)W6C
HHL)CRM
RWV)VH4
9KC)NHG
F6Q)3KB
NN4)C9K
5X1)FHW
6Y4)171
7M9)9RW
5RG)N29
LYQ)5X4
PFZ)QDL
V3W)2F6
YQ5)ZH8
9FM)L64
H8J)BLH
1Q8)7K8
7SC)6TS
Y29)F8T
RFD)VM9
N75)2K4
CW4)51C
524)MRD
WFN)Y29
LMM)JRC
8NV)YNF
R7Y)XVW
HRY)DNT
C6Z)RS2
7S3)D3M
T6P)XFJ
SC3)Q7P
JGF)H82
SB8)9MF
8LD)6HT
KNC)XQ2
GTN)G9M
PCD)6XS
MZR)RG4
GDY)FXJ
KHV)BG5
5XT)CZ7
SBL)GJB
X3K)YDZ
77D)7CB
2JH)VP3
4FC)SGM
PG2)HDH
H25)Z3K
9XG)GJZ
891)DR5
BR8)22Y
XPF)L83
NK2)JMF
717)7CL
B3C)PCG
R4M)HHL
BLH)KVF
4GR)K8X
DXW)58D
4R9)84T
LS5)6F3
Z56)WM1
J9M)LRC
P3D)DXR
NZQ)4XQ
549)TYT
2BY)68Q
ZXH)S9G
BS2)WPX
4JP)JG9
RDX)WW3
LRC)199
QZW)8P7
RLP)DQV
JLZ)8GP
FXJ)S9H
N4W)524
9WB)S64
142)PN4
R34)FZM
L9N)9Z5
D1W)S5B
C9K)SZ6
B25)C47
9PF)RND
FZM)Z9P
L64)8QT
4Z4)CBY
XFL)Y5Y
W2D)VW9
V6D)N9F
SYP)5TW
S5R)KPM
QYX)T1X
MW1)2RV
DZ4)6TP
P64)P73
PZ6)JRN
JRC)CNR
RYF)Z2Z
MF8)FDB
2ZZ)XL3
LBK)B7H
15Z)17Z
3W8)K8F
BZ6)RJW
G1J)CBR
2N9)SVF
NSK)253
78X)1MZ
3TH)BGD
15K)2K9
BS7)CGB
BGR)XSW
4GD)CPP
RX9)WK8
RJK)VWN
9RW)N6J
J4X)4C4
35L)1WQ
HZ2)QL6
RND)VJ5
2Z9)76C
Z2Z)YB4
JNP)8MM
Z1L)SRB
3VZ)NHJ
8GW)56R
C1Y)8LX
4QW)XK3
G68)SQ4
Y9D)HQF
T85)38S
Y9K)9FM
RD3)RWV
F17)B8P
HMG)1JD
T8M)46C
JC8)HJT
D8H)F6K
55C)3BM
TZ8)VDR
FBY)BDT
CNR)4JP
GMJ)CW3
G9M)TGV
G81)ZSQ
C5Y)MHR
6RB)99X
9JH)F16
G5F)VFV
T1L)2N1
L6D)9QS
LT2)97T
CGB)RS9
HW4)WSD
22Y)7LM
PKQ)CN9
LPV)H25
LLP)SJ1
LCZ)SB8
CR6)DFJ
VV1)RLP
S9G)PWQ
9FZ)3TH
X9R)4R4
6TS)RFG
GNL)RZ4
HCV)8CF
KZD)3W8
DR5)F3Z
1R2)8FX
G5W)MHG
TL3)TDX
HPV)1R9
NMF)67C
LW6)N7W
LB9)KQV
DMB)25Y
9MF)B9K
DP8)43J
ZHS)1LW
Q51)HQ8
H9D)SJG
RW8)PKQ
6XS)FYY
YV2)WGH
KLK)9BM
TZP)PYK
M41)4TC
GFM)418
CW4)HGF
K3W)T9T
H1Z)GBK
VFV)XXH
Q7R)6V2
L5J)TL3
LYT)29Y
HLQ)HG2
9YX)HMG
NJ6)X1P
CRM)RYF
8GW)D1W
6L9)G58
DT4)ZMB
QWL)BMT
9SV)WST
ZPT)LB9
3G5)W92
199)F7K
WGR)794
F6Z)HGR
NW9)ZFD
V4S)Y81
XRB)BFK
5XR)TS8
6HT)B4J
6XQ)SY3
6HF)PRG
78C)D5V
4X7)FRL
GYD)JNP
2F6)QD5
T68)GH5
Z93)P64
G35)4Q9
5ZY)32L
M8D)46R
WVQ)J7G
8FL)Q82
KWC)K9Q
1ZS)WF8
84Z)T1L
ZSQ)MMQ
MC6)78H
C6Z)J3X
Y91)RBW
WPY)Q7R
GVX)VHV
CW3)W29
TRX)KL8
1LR)3B2
XL4)LNZ
FDB)79C
X6K)S4F
HDH)S2F
Y1J)CQR
5ZP)GMJ
QQ2)7WF
YJM)GQP
DQV)RQP
KLC)Z1X
Q7P)R1Z
PLP)2GY
7VX)7WG
G1X)VV5
H3G)KQW
FCY)RY7
333)QHB
VV5)8PS
X6L)TNH
QGN)NW9
YNF)QL5
XDF)CXV
F6K)9KC
G7B)YKY
Y99)4CC
J1S)SMT
4R4)NCC
Z3K)BHN
CBR)4VJ
L4J)N2C
1L5)G5W
3JY)KGT
1LW)4L5
XZM)5XJ
MQ5)M41
7BF)3QT
756)9HP
GNW)1LF
CLK)JC8
SGM)V4S
BCT)XFL
B8Z)1T4
2V1)BK9
FWP)PP8
L83)4C5
C18)KH1
HHX)KGZ
1XC)6XL
QBB)SBL
GHD)LTY
5XJ)2TJ
4CC)BPX
6YQ)4LQ
HGF)3RJ
6QK)QQY
VP8)31D
V9B)GNV
YR3)472
4C5)YGZ
389)ZXK
3BM)QWB
WH7)JKZ
9NQ)42D
S2S)X8Z
S9H)4N8
R9L)T9D
3T1)JRS
W92)4DQ
8K7)2N9
8LP)8LD
5GG)3T1
M1T)WH7
GBK)8PV
X54)3H3
7C1)QZ8
KQ6)KVS
7L6)NWJ
51L)R3G
DHQ)3M4
DNT)YQH
HCD)6PL
3N1)3G7
YKY)3XG
B8J)J2S
P73)278
9P5)TFQ
Z5G)PCD
3PL)B3V
X8R)NWY
S5B)YR3
CFR)LMD
SKB)DK6
BFL)78X
CWQ)NDQ
9LQ)N7Y
JST)6RW
RN9)RZH
V9P)V6D
4TJ)8J1
9KC)NM2
9CW)CWC
KTH)1R2
WDS)4PT
JZ2)RTJ
1HW)4VR
942)LDC
278)FMT
794)KR9
84T)H7D
CMS)G5F
V81)1TS
V12)2BR
TRS)BBH
8Z2)8K7
JRN)CT9
B8P)SVZ
KSF)RP8
QNX)SDY
LMD)DTY
9MZ)BR8
4RB)8HB
X8J)5BF
KZX)9LQ
JX8)DL9
2P2)H3G
H67)C6Z
66G)59V
KL8)HVP
D3M)TFB
HY7)FMM
SW1)RMX
GJB)NN4
B7F)4GN
XXH)ZCP
XJ9)XWM
9QQ)X9R
RS9)2CD
G1P)4FC
QL4)RT1
7LM)CWQ
5TW)Q76
9LQ)P77
3XT)2T6
RFG)PYJ
JZY)D7G
KJ7)LW4
VSQ)999
SVF)WFD
51C)959
NZL)JFZ
81P)27W
P64)6WR
D9Z)KXZ
X1P)TXH
ZSF)5QD
HYD)R6Y
ZNF)7D5
J12)PG2
BL8)GCP
`

const day07myInput = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`

const day08myInput = `222222222222222222222021022222222202002222222202222222220222222222222221222222202211220222102122202220212222222221222222202122210222222022222222222212222222222222222222222120022222222222122222222222222222220222222222222221222222202211220222102222222220202222222221222222212222220222222222222222222202222222222222222222222221122222222212202222222222222222221220222221222220222222212222220222222022212222212202222221222222200122220222222122222222222212222222222222222222222021122222222202212222222212222222220220222220222222222222202200220022012122222220222202222220222222202222222222222122222222222220222222222222222222222120122222222222122222222212222222220221222220222222222222212211220222122022212222202202222220222222220222221222222122222222222220222222222222222222222122022222222222022222222202222222221222222221222212222222222201221022012222202222212222222220222222211222221222222122222222222222222222222222222222222220222222222212212222222202222222222220222220222222222222222200221122012222222222222202222220222222200022201222222122222222222221222222222222222222222222122222222202122222222202222222221222222222222210222122212212220222102121202222202222222220222222222022112222222222222222222211222222222222222222222221222222222212112222222212222222220221222220222220222022202201222222102220222221222212222221212222211222121222222222222222222202222222222222222222222021022222222222022222222212222222220222222220222210220022212200222122002220212222202202222220002222211122011222222022222222222200222222222222222222222221222222222222212222222212222222222220222220222211220022222221222022012222212222202202222220022222222222202222222122222222222221222222222222222222222020122222222212012022222202222222220221222221022200222122212202220022102120212222202202222221002222211122211222222222222222222210222222222222222222222021222222222212012112222222222022221222222221122201222122222220220122012022222220202222222222212222202122100222222122222222222202222222222222222222222021122222222222112022222212222122222221222221022212222022202211220222112122202221222202222220022222211222201222222122222222222211222222222222222222222221122222222202122022222222222122220221222222222202220122202222220022112021202221202202222221012222200222201221222022222222222201222222222222222222222121122222222212102222222112222222221221222222222220221222212220220122112122222220222202222222102222222222102222222022222222222212222222222222222222222121222222222212212112222022222021220221222220022221222122212202220122022022212222202212222221202222210222201222222222222222222201222222222222222222222222022222222222022012222112222120220220222221022211220022202221222222122020222220202202222222202222202122201220222222222222222210222222222222222222222121122222222202102212222202222021221222222220022221221222202212221022012021222221202202222220102222220022221222222022222222222202222222222222222222222022122222222212222122222002222020222222222220022202220022222201220122222222202220202222222221222222222222211220222222222222222222222222222222222222222020122222202202202002222002122022220221222222022201221222202202221022212122202220212222222222022222210122102202222022222222222222222222222222222222222221122222222222112022222212222022222221222220022212220122202220220022212221212221212222222220012222212022102221222022222222222221222222222222222222222020022222222222112222012002022021220220222222222220221122212201221222212220212021212212222221122222220122022211222022222222222222222222222222222222222021222222212212002112002122122122222222222222222201221122202200220122102220202222222222222220212222212022221222222122222212222211222222222222222222222221222222222222002112022122222120220221222222222211222022222201220222012220202122202212222221222222202022122201222022222212222210222222222222222222222022222222222212202222202212222021221221222220122202220122202201221222102220202020212202222220102222212122001201222222222222222222222202222222222222222220122222202202112202212002022121221222222221122221220022212200221222102021202220202202222221012222201222210210222122222212222202222222222222222222222220222222202212022122212222122120220220222220122222221222212212221122102020212221002202222222102222220122020212222122222222222222222222220222222222222121022222212222022102222202122120222221222220122200221122222220220022202121222020002212222222212222222022011221222222222212222200222212221222222222222121122222222202022212122022222021221220222221122212222022212220221122012021202020102202222222022222210122022210222122222212222221222202221222222222222021120222222222212122012212222022221222222220222202220122222211221122022221202221102212222221002222201022220221222122222222222210222212221222222222122021221222212222212022222122122122222220222222122222221222212200220122122120212220022202222221022222201122212210222222222202222221222202220222222222122021121222212212202212202122222221220222222222122211121022212220222022202020202221022222222222102222202022011202222222222202222222222222221222222222122020220222212222022122002112222120220221222221022201021222202201220022202121212120022212222222012222210222021200222022222222222210222202222222222222222221121212212202002002002022222220220220222221122221020122222210220222002020202220002202222220122222222220120210222222222222222220222222221222222222222021022222212212102212012222222021222201222221122201120122202202221022002221202221002202222221012222211121011210222022222222222202222202221222222222122022021222202222022012022202022020221200222221122221221222212212220022202122212220122202222220212222210220012221222022222222222210222212220222222222122220222222212212222212102102022022220210222222222210120022202221220222202220202222112222222221112222200221221202222122222212222202222212221222222222122021122202222202012102212022122221221200222222122221022222222222221222202122202220022202222221122222221220220222222122222212222221222222220222222222022122220202212212112222122112122222220211222222022222021122202221220122202121212121002112222022212222221121110200222122222222222211222202220222222222122221022212212202122112222222222220222202222222122202222022212202220122022020222021012012222221002222211022012222222122222222222202222222222222222222222102122212212222112212022112022122222200222220222211220222202211021122122221222222012022122121102222201122110212222022222222222201222222220222222222122100120002202222202102202120122221221202222220022220120022212202111220022120212222102112122022102222220022122211222222222222222212222202221222222222121110221202202202002102022110022222220222222020022221020022202211110021112222212022222002222022122222220021121220222022222222222201222202221222222222120102120202212222102112002220222021220200222122222202121022222212112120102221212122122102122220012222201120110212222222222202222220222222222222222222220220020202222212102202112022222020222211222121022210220022212201222221022121222021022102022020122222200120001202222222222212222212222222222222222222021010021012202212212212022110022022220221222020122210221222202202020121022121222122122222022222122222200020111212222022222202222222222212222222222222121000121222212212112212212222022121220212222220022202222222222201020022002221222221212202222220112222210221210220222022222222222220222202221222220222221212222022212212002202222210022120222210202220122221222122202222210221102122202222002212122122212222200022212212222022222212222211222202220222221222120201020122212222102212022121022020222222202020222201022122202210102221102120202220222212022222222222221120220211222022222212222210222212220222220222222020021202212222202112002122122020221211212220122222022122212210102022112221202222102212222222222222220222221221222122222202222222222212222222220222120120022222212212002212112122122220220221222122222221020222212222012220002221222222212222222021222222202220122202222222202212222211222212221222221222122221021022202222212212202021222222220221220222122211120122202211221020020122212220202202122022102222221221212212222122222202222200222212222222222222021122221002212222222212212202122221220221212120222201122022222211100020221222202221222022222121022222212201200200212222212222222202222212222222221222121000121022212222212002002022122121220201211120022222220122212212010021001021202121012222222122022222201102212211212122202212222212222212221222222222020111022102202212212212202200222221222222201021022220222120212201111020212120202021022112102121222222200102022212212222222212222220222202222222220222120010022122222202122022102112222122220202201020222202221222211210011120200121212222112222102221002222202212022221222022212212222211222202220222222222120122222122212202102022102111122220221220210220022211221122212200010222012021222021022202202021012222200022201210212022212222222200222202221222220222221220020022202222122102102212122222220222221022222221021022212211010221112221222022002022222020202222202102212200212022212222222221222222222222222222021221020122202202212102102200022120220202201122222201220020212210121221000022202122002012002121122222200110100221222122212202222222222222221222221222121022120222202202201022002120222122220211221221222221121022210221201020211120212022102012002121112222201212022220202222212202222221222222221222221222222111122222222212210002112220222221220202212222222212121120210221112222220221212020022202021120212222220010010211202222202222222201222212222222220222020220021102212202121122122021122122222221212020022221221020202210102020210220222020212222001222122222202112201202222022212202222221222202121222220222220100122120202202112112202100222022220200220021222222221020221220202121222222222022112112210220102222211220200220222222222212222212222222221222220222022101120112202222002202012001222122222202212122222212021221202201211121100120202022212222111122202222212200012211202122212222222201222212120222220222121222021000212222201002222221121221221222200121022221122222201221220121101021222020202022200021212222200011211210202222202222222221222212220222220222021021221002112222122012012210221120222211211122022210221122212202112120110222202020202102220022222222200112010221202112202202222212222222121222220222222111222002002202002202222122220221221020221021222212222020220210011221122221202120212002100120202222212120202212212122202202222221222202220222220222222100121120002222102200102200021221220111221221121210221020212202010222100022202222222122022121000222221102001201212022202202222222222212122222220222121020021222002202122202022110222022220101200222221200020020201201121021102120202020112212111121120222222111111212222122212222022211222202121202202222122002022211222212000102202002120021221212200221222222221121212222010121110020212021112212221120222222211101112210202012222202122212222202122012202222021210021212022202201012002220120220222210200020120212220120210200002220102222222122022122010220212222211122220210222022222212222222222222120002202222221220020010202212220201002100220021220000221221120221022122202210101021212221222020112012201120020222222110220212202222212202022222222222221212202222220102020010002222222120102001220220220100201022122210120221212202200120022120212220102012211020111222222200021200212002222202222201222202221202210222120122122121112212020002102121220121221002200020222220122122212201211220212220212121202012012220210222202122201222202212202202120201222222122012220222020122020212212212120201022012222222211112202001022222021021200221212221001122212120222012210122022222212200222222212122222222121221222202120112201222020111020101112202021001202010020022200002221100121201022220201211211222102121212220222212021120000222100211202212212012202222121221222222021212201222021100220021002222021000102211122022201210221110220201222022200220120120100021212121002012210122110220222122222220202112212222022212222212220022200222221020120211202202010020202102221120201112212221022201120020210202200122112222222122212102002122222221001221021200202212222212222200222222121212210222222120021220022212211101022022221121221010212002121212222121210212101122010002212022112122020120120221210200001201222102202212220001222212222102202222021011120012112222121210212110022021200221222200121201121020211221020022012212202222012112201022102222222101012200202102222212222101222212121022201222220001021120102212112000002102021020220200201021220200121020210202201221111120212120122122201120201222121110100202212202212222120112222222220202221222220110222011002202221021222111222222210202201010221202020222200212012022021200222022102102011222000222222021022202222022222202120002222202120202201202020000120212022212110020222222022121211212200000022210021120210221102222121121202122012022120121210222221112001201212102002212022101222212222012220202121102221220022212101221112212022021200201121000120221222022201210111022101121212220122122011121000220220012220221212012112222222102222212122102211222022011221012012222021011022020020020202200020210020221222121211201002020110022222221022002221021222221211011002210212222002222221121122202020212200222020222022010022222220110002102020020210020211201022202122120212211010120112122212221222012212122002222021220102202212212022202022001022202220212211202020012022112122202012221202221022122201102122022220212220222221212220220210020202022012102000120122222100200000211110112222202021012022222120202222212220001222020022202210010012002221220211020102020200201121121220011121022101022222020212122020120210221101212211200221002212212120102222212221112212222022012121211002202012000202220022022202101212100122201020120222002221221201211212222222112201222202220212010020212222102102202020221222222020002221202222020121201202212110000012202021220201220000102220222022022212212210121121021202122012202101121110222220120201222011022122222121102122112022102212212220201120020202202101212012110121022221002121001220202020022222222222022000020202120212112110020121220120211100211221222202202122220222102220102220212020021210001012222112210012022122222202110121112021210021122211211010220022011202120222122222122111220221200100212122002222212121201022212021212221202122211210211002212210001002222221020222200020020210221021122220200111020001120212220122022202021221222101101022200020222202202022100122222121202211212020200202111102212220222022210122121200201212000100212021121220110120220102111202020102222222122220222011102120222111102220102122201222022021012201222021200220212002212110010212200122021212022220020110222121022212020010221021021212022202122022120220221121020012221211222110022020112002102022202202222021122102011002212002020022112222222220111012101012210120120212220000021022200202222102022212222200211001110020201220212211010121110012000220202202212122220112010202202022000112221020121222000102220201210121122201122001021011121222222102222021220110212220121112200100112202022020122112000021112221112220210022100022212210222022001021122202022012110002220221220222020212222111021222220212200201121212220100210212201200212201222120012100212110111102010100022220012200121100101020112102200021110020021101012112202100222112101021000011111220022100100221021210022212011010011000221211102`

const day09myInput = `1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1102,1,3,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1101,0,493,1024,1102,1,38,1015,1101,20,0,1011,1101,0,509,1026,1101,0,32,1018,1101,0,333,1022,1102,1,0,1020,1101,326,0,1023,1101,0,33,1010,1101,21,0,1016,1101,25,0,1004,1102,28,1,1008,1102,1,506,1027,1102,488,1,1025,1101,0,27,1013,1101,1,0,1021,1101,0,34,1019,1101,607,0,1028,1102,1,23,1003,1102,26,1,1007,1102,29,1,1009,1101,31,0,1000,1102,37,1,1012,1101,30,0,1005,1101,602,0,1029,1101,36,0,1002,1102,1,22,1001,1102,1,35,1014,1102,24,1,1006,1102,39,1,1017,109,4,21102,40,1,6,1008,1010,40,63,1005,63,203,4,187,1106,0,207,1001,64,1,64,1002,64,2,64,109,13,1206,3,221,4,213,1106,0,225,1001,64,1,64,1002,64,2,64,109,-5,1208,-9,22,63,1005,63,241,1106,0,247,4,231,1001,64,1,64,1002,64,2,64,109,-5,21107,41,40,3,1005,1010,263,1106,0,269,4,253,1001,64,1,64,1002,64,2,64,109,-1,1202,3,1,63,1008,63,29,63,1005,63,295,4,275,1001,64,1,64,1106,0,295,1002,64,2,64,109,16,21108,42,42,-8,1005,1014,313,4,301,1105,1,317,1001,64,1,64,1002,64,2,64,109,-4,2105,1,5,1001,64,1,64,1105,1,335,4,323,1002,64,2,64,109,-5,1207,-4,28,63,1005,63,355,1001,64,1,64,1105,1,357,4,341,1002,64,2,64,109,2,21102,43,1,-1,1008,1014,45,63,1005,63,377,1106,0,383,4,363,1001,64,1,64,1002,64,2,64,109,-10,1208,-3,36,63,1005,63,401,4,389,1106,0,405,1001,64,1,64,1002,64,2,64,109,6,21107,44,45,1,1005,1012,423,4,411,1105,1,427,1001,64,1,64,1002,64,2,64,109,4,21101,45,0,3,1008,1018,45,63,1005,63,453,4,433,1001,64,1,64,1105,1,453,1002,64,2,64,109,-23,2101,0,10,63,1008,63,36,63,1005,63,475,4,459,1106,0,479,1001,64,1,64,1002,64,2,64,109,26,2105,1,6,4,485,1105,1,497,1001,64,1,64,1002,64,2,64,109,4,2106,0,5,1105,1,515,4,503,1001,64,1,64,1002,64,2,64,109,-25,1201,10,0,63,1008,63,26,63,1005,63,537,4,521,1105,1,541,1001,64,1,64,1002,64,2,64,109,18,21101,46,0,-1,1008,1014,43,63,1005,63,565,1001,64,1,64,1106,0,567,4,547,1002,64,2,64,109,-6,1201,-4,0,63,1008,63,33,63,1005,63,587,1105,1,593,4,573,1001,64,1,64,1002,64,2,64,109,22,2106,0,-3,4,599,1105,1,611,1001,64,1,64,1002,64,2,64,109,-28,2102,1,-2,63,1008,63,22,63,1005,63,633,4,617,1105,1,637,1001,64,1,64,1002,64,2,64,109,-1,21108,47,44,9,1005,1011,653,1105,1,659,4,643,1001,64,1,64,1002,64,2,64,109,10,2107,24,-8,63,1005,63,681,4,665,1001,64,1,64,1105,1,681,1002,64,2,64,109,-11,2107,31,4,63,1005,63,697,1106,0,703,4,687,1001,64,1,64,1002,64,2,64,109,8,2101,0,-8,63,1008,63,23,63,1005,63,727,1001,64,1,64,1105,1,729,4,709,1002,64,2,64,109,-16,2108,21,10,63,1005,63,749,1001,64,1,64,1106,0,751,4,735,1002,64,2,64,109,17,2108,36,-8,63,1005,63,769,4,757,1105,1,773,1001,64,1,64,1002,64,2,64,109,-10,1207,1,23,63,1005,63,791,4,779,1105,1,795,1001,64,1,64,1002,64,2,64,109,-3,2102,1,6,63,1008,63,22,63,1005,63,815,1106,0,821,4,801,1001,64,1,64,1002,64,2,64,109,16,1205,7,837,1001,64,1,64,1105,1,839,4,827,1002,64,2,64,109,-5,1202,0,1,63,1008,63,30,63,1005,63,863,1001,64,1,64,1106,0,865,4,845,1002,64,2,64,109,4,1205,9,883,4,871,1001,64,1,64,1106,0,883,1002,64,2,64,109,16,1206,-7,899,1001,64,1,64,1106,0,901,4,889,4,64,99,21102,1,27,1,21101,915,0,0,1105,1,922,21201,1,47633,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21102,942,1,0,1105,1,922,22102,1,1,-1,21201,-2,-3,1,21101,957,0,0,1106,0,922,22201,1,-1,-2,1105,1,968,22101,0,-2,-2,109,-3,2106,0,0`

const day10sampleInput = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`
const day10myInput = `.##.#.#....#.#.#..##..#.#.
#.##.#..#.####.##....##.#.
###.##.##.#.#...#..###....
####.##..###.#.#...####..#
..#####..#.#.#..#######..#
.###..##..###.####.#######
.##..##.###..##.##.....###
#..#..###..##.#...#..####.
....#.#...##.##....#.#..##
..#.#.###.####..##.###.#.#
.#..##.#####.##.####..#.#.
#..##.#.#.###.#..##.##....
#.#.##.#.##.##......###.#.
#####...###.####..#.##....
.#####.#.#..#.##.#.#...###
.#..#.##.#.#.##.#....###.#
.......###.#....##.....###
#..#####.#..#..##..##.#.##
##.#.###..######.###..#..#
#.#....####.##.###....####
..#.#.#.########.....#.#.#
.##.#.#..#...###.####..##.
##...###....#.##.##..#....
..##.##.##.#######..#...#.
.###..#.#..#...###..###.#.
#..#..#######..#.#..#..#.#
`
