# A extend of ctf-wiki

> 知识点太多了回忆不过来了😫😫😫,各位多多来添砖加瓦ORZ
>
> 按照wiki的路线做一些扩展,放一些好的学习资料和连接上去
>
> 脚本(一把梭哈的轮子)在网上应该都能找到

# In the frist:Using sagemath in a grace way

- 搞一个Windows termial（很Duang很漂亮）

- 搞一个基于wsl2的linux子系统（个人偏爱ubuntu）

- apt 安装 sagemath（linux自己折腾）

  
  
  ![](https://gitee.com/ljahum/images/raw/master/img/20211128223016.png)

然后，调整颜色让他好看一点（也可以不改）

![](https://gitee.com/ljahum/images/raw/master/img/20211129113709.png)

调好了可以在terminal中用`code .`用vscode联动编程

也可以用上流的jupyter，我闲的无聊搭了个sage的jupyter笔记本给大伙玩23333

http://139.224.231.64:8888/

密码123321

## stream crypto

### 随机数生成器

#### mt19937

大多数随机数算法的实现方式

三年前的攻击方法：https://github.com/tna0y/Python-random-module-cracker

#### LCG线性同余生成器

攻击方法：推公式

example：https://github.com/fghcvjk/HGAME2021/blob/master/crypto/%E5%A4%BA%E5%AE%9D%E5%A4%A7%E5%86%92%E9%99%A92.py

#### 线性反馈移位寄存器 - LFSR

说白了就是解一个方程，如果有位数缺失可以想办法去爆破

sagemath中有快速解决求运算符为异或的矩阵的方法

<img src="https://gitee.com/ljahum/images/raw/master/img/20211128222422.png" style="zoom:50%;" />

异或在数学上被称为GF有限域上的加法，对矩阵指定类型，sage自动将该矩阵之后的运算限制为我们规定的域下的运算（指定运算方式）

模运算和多项式也是可以的



## block-cipher

### 比特反转

最最最基本的一种方式，把iv拿给别人改也是挺蠢的

### oracle padding attack

主流的blockcipher有aes，des等，看完了wiki应该会有一个基本了解：

- 对massage进行分组，填充
- 放入加密器中加密，加密的规则由秘钥决定
- 讲密文输出，输出后的密文和明文长度应该保持一样

wiki上面值给出了加密器（黑盒）中的分析和破解

但现在流通的blockcipher的黑盒已经十分优秀了，所以这里建议从填充、encrypto mode上的漏洞下手

填充漏洞最经典的就是ECB 填充和CBC填充

ecb填充是利用往下一个块泄露单个未知字节来进行诸位爆破（嘛，没有找到例题，画一个好了）

![](https://gitee.com/ljahum/images/raw/master/img/SmartSelect_20211129-115937_Samsung%20Notes.jpg)



cbc填充例题，和ecb基本思想是一样的，在一些地方进行了小小的修改

https://github.com/wm-team/WMCTF2020-WriteUp/blob/master/WMCTF%202020%E5%AE%98%E6%96%B9WriteUp.md#game

### cipherMode

wiki上面已经把基本模式讲得比较详细了，

https://ctf-wiki.org/crypto/blockcipher/mode/cfb/

cfb oracle padding的攻击方式

https://ljahum.top/bytectf2021/#document-for-justdecrypt

其他奇奇怪怪的ciphermode的攻击

- aes-ocb2模式的攻击方式

https://dawn-whisper.hack.best/2021/04/04/Wp_for_%E7%BA%A2%E6%98%8E%E8%B0%B7_crypto/#babyFogery

https://eprint.iacr.org/2018/1090.pdf

>  嘛，原理有点复杂，做个积累就行



### DFA 差分隐私分析

> 这个玩意只碰到过一次，碰到了只能去极限找资料，然后wp八成是跑国外👴的poc跑出来的

github一搜就有，找到能用能玩的就行

古董级玩具

https://github.com/Daeinar/dfa-aes

对应题

https://github.com/ljahum/crypto-challenges/tree/main/ciscn2021/oddaes

新一点的(没用过)

https://github.com/SideChannelMarvels/JeanGrey

https://github.com/guojuntang/sm4_dfa

## RSA

> 这里我们先总结非交互情况下的rsa攻击方式
>
> wiki上已经给出了基本的破解方式这里直接从coppersmith开始



### 格相关前置知识

格学习资料

>  这里给出了基本的格和格相关，重要的还是后面的应用
>
> 格相关的题目脚本更多的是对题宝具形式，找到轮子一把梭哈。。。
>
> 毕竟复现论文拿来出中型赛题过于奢侈了

by HIT阮师傅 ，博客格相关入门内容翻译质量非常高（高于一些学者的翻译），但仍有非常模棱两可的部分

https://www.ruanx.net/coppersmith/

> 背包算法wiki上给的比较详细了，可以配合阮师傅的博客构造一下格子

除wiki以外的加强版copper脚本

https://github.com/defund/coppersmith

### 类维纳题型 （私钥d较小）

wiki给出了<!-- $d<N^{0.29}$ --> <img style="transform: translateY(0.1em);filter: invert(100%)" src="https://render.githubusercontent.com/render/math?math=d%3CN%5E%7B0.29%7D">的范围的形式

某次比赛给出了扩展形式，本文章翻译得比较好，证明平板也比较高质量

这个题把d范围扩展到了
<!-- $d<N^{5\over 14}$ --> <img style="transform: translateY(0.1em);filter: invert(100%)" src="https://render.githubusercontent.com/render/math?math=d%3CN%5E%7B5%5Cover%2014%7D">

https://blog.csdn.net/jcbx_/article/details/109306542

> 嘛，核心点嫖到脚本就完事
>
> 理解层面看懂怎么从公式构造到格子就行
>
> 呃呃呃,好像这个wiki已经收录了,看来wiki并没有停止运维....

### Common Private Exponent

一个2019年没人玩，2021年被打烂的知识点，构造格子的经典案例

https://github.com/SycloverSecurity/SCTF2020/tree/master/Crypto/RSA



### 共模攻击plus

https://github.com/De1ta-team/De1CTF2020/tree/master/writeup/crypto/easyRSA

有别与正常的公模，这玩意只给了一组公钥，脚本给出了2组公钥分解N的格子构造方式

四组的:

https://huangx607087.online/2021/03/01/LatticeNotes6/#toc-heading-6 

这个博客里面还有很多其他总结得非常好的东西，维纳包括维纳扩展的总结



### Known High Bits Message Attack AND Coppersmith’s short-pad attack etc.

这玩意纯纯的屯脚本了

推荐题目:https://blog.csdn.net/zippo1234/article/details/109409929

环境：https://github.com/CTFTraining/qwb_2019_crypto_copperstudy

### RSA 选择明密文攻击

wiki给出了三种攻击方式

- 泄露解密后的最低位byte
- 泄露解密后的最低位bit(奇偶性)
- 泄露解密后的最低位bit(奇偶性),但每次会变,这时算法变为给密文乘上$(2^{-i})^e\pmod{n}$,也是一种很好的思路

扩展:

- 高位泄露

这种情况就不能用乘$2^i$的方式来泄露不同位置的明文了,因为超过n后会被模掉

影响明文的数据,所以要用二分法的思想去搞

例题:

https://ljahum.top/%E5%A4%8D%E7%9B%982/#%E8%93%9D%E5%B8%BDber-final

- mini版[Bleichenbacher Attack](https://crypto.stackexchange.com/questions/12688/can-you-explain-bleichenbachers-cca-attack-on-pkcs1-v1-5)

南邮👴的练习题,本质上也是二分法的思想

src https://github.com/fghcvjk/NCTF2020/tree/master/crypto/Oracle

wp https://ctf.njupt.edu.cn/562.html#Oracle

### 数学/算法题

#### 数论题:

- 东华杯2020

https://github.com/ljahum/crypto-challenges/tree/main/2020%E5%BD%92%E6%A1%A3/%E4%B8%8A%E6%B5%B7%E5%A4%A7%E5%AD%A6%E7%94%9F2020/baby_rsa

- 东华杯2021（每年都来恶心人了属于是）

https://blog.csdn.net/weixin_56678592/article/details/121194962

- gkctf2021

https://blog.csdn.net/weixin_51867782/article/details/118573717

- 祥云杯2020 more_calc(银行赞助的抢钱大赛)

下面的连接没代码,附一个

```python
import gmpy2
from Crypto.Util.number import *

flag = b"flag{xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx}"

p = getStrongPrime(2048)
for i in range(1, (p+1)//2):
    s += pow(i, p-2, p)
s = s % p
q = gmpy2.next_prime(s)
n = p*q
e = 0x10001
c = pow(bytes_to_long(flag), e, n)
print(p)
print(c)
#27405107041753266489145388621858169511872996622765267064868542117269875531364939896671662734188734825462948115530667205007939029215517180761866791579330410449202307248373229224662232822180397215721163369151115019770596528704719472424551024516928606584975793350814943997731939996459959720826025110179216477709373849945411483731524831284895024319654509286305913312306154387754998813276562173335189450448233216133842189148761197948559529960144453513191372254902031168755165124218783504740834442379363311489108732216051566953498279198537794620521800773917228002402970358087033504897205021881295154046656335865303621793069
#350559186837488832821747843236518135605207376031858002274245004287622649330215113818719954185397072838014144973032329600905419861908678328971318153205085007743269253957395282420325663132161022100365481003745940818974280988045034204540385744572806102552420428326265541925346702843693366991753468220300070888651732502520797002707248604275755144713421649971492440442052470723153111156457558558362147002004646136522011344261017461901953583462467622428810167107079281190209731251995976003352201766861887320739990258601550606005388872967825179626176714503475557883810543445555390014562686801894528311600623156984829864743222963877167099892926717479789226681810584894066635076755996423203380493776130488170859798745677727810528672150350333480506424506676127108526488370011099147698875070043925524217837379654168009179798131378352623177947753192948012574831777413729910050668759007704596447625484384743880766558428224371417726480372362810572395522725083798926133468409600491925317437998458582723897120786458219630275616949619564099733542766297770682044561605344090394777570973725211713076201846942438883897078408067779325471589907041186423781580046903588316958615443196819133852367565049467076710376395085898875495653237178198379421129086523

```

https://blog.csdn.net/cccchhhh6819/article/details/110302200

> 嘛,前年祥云杯还有几个其他的小数学题

https://www.anquanke.com/post/id/223383#h3-22

#### 算法(剪枝)

> 这种吧,就是换着法子恶心人,换个姿势又是一个新题,对对应的新问题又要重新安排剪枝的条件
>
> 大比赛一般不会出，小比赛用这个拖时间

1.[GKCTF 2021]XOR

https://blog.csdn.net/m0_57291352/article/details/119974008

2.Mini-L CTF 2021-asr 

SRC:

https://github.com/traLk/Mini-L-CTF-2021/blob/main/Challenges/Crypto/asr/task.py

solve：

https://d33b4t0.com/H4n53r-TEAM/#asr-dbt-done

etc. buu上记得有一些类似的，算是一种近期被挖掘出来的玩法吧。。。。



## 离散对数



> 离散对数即离散也对数
>
> 主要体现可以理解为在咱的运算模式下只有加法没有乘法
>
> 可以用加法实现了快速乘法却没有快速除法
>
> 这里整理了较为常见的离散对数玩法，ECC加密和DSA签名



## DSA相关

### 常规DSA



基本攻击要点wiki已经讲过了，整个例题

- 东华2020

https://github.com/ljahum/crypto-challenges/tree/main/2020%E5%BD%92%E6%A1%A3/%E4%B8%8A%E6%B5%B7%E5%A4%A7%E5%AD%A6%E7%94%9F2020/baby_dsa

- 狮吼2020 好像找不到推导的公式了，这个题是刚开始学密码的时候碰到的，现在看起来真简单啊。。。。。

https://github.com/ljahum/crypto-challenges/tree/main/2020%E5%BD%92%E6%A1%A3/%E5%98%B6%E5%90%BC

wp https://blog.csdn.net/jcbx_/article/details/110944601

### ECDSA

阿巴阿巴，瞎几把搜搜到了UESTC👴的ppthttps://blog.arpe1s.xyz/files/DSA.pdf

里面由关于LLL的部分主要是为了解决bias nonce attck格子的构造：k低位泄露的攻击方式

https://github.com/daedalus/BreakingECDSAwithLLL 这个仓库给出了对于ECDSA bias nonce attck 的攻击脚本和攻击范围（范围在论文里面）

ECDSA和DSA核心签名的方式是一样的，只是计算离散对数数值的方式不一样，应该可以作用于常规DSA💀

- 对于weak ecdsa的一些讨论

里面代码不能直接用，要改一下

https://blog.trailofbits.com/2020/06/11/ecdsa-handle-with-care/















