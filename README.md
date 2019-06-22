# Random Number Algorithm (English Version)

Random Number Algorithm

The Algorithm is easy to understand. 

A blockchain's new block is in the future and unpredictable, therefore using that block hash and combine it with server seed and TX Index we can have a unique Seed, then using Mersenne Twister alg to generate a random number.

It is random in terms of unpredictable, but the result is repeatable which means you always get the same result using
the same input.

Seed = Server Seed + Deposit TX + Deposite TX Index

* Server Seed:
  - c9656775fd0f9d40dcde80fd597cc7e863144cb4fccbfc1e47de1ff4849ab054
* Deposit TX
  - b31e198874365efc77338439e16c9b92c9b0c75a49c82c61c4e4423b71139996
* Deposite TX Index:
  - 0
  
### Example (Game id 8888: https://bytomdice.com/fairgame/8888)
### Run it online: https://repl.it/@Satoshigod/RandomNumberAlgorithm
### BytomDice Simulator: https://bytomdice.com/simulator


# 区块链随机数算法

在区块链随机数出现之前，几千年来中心化系统实现的随机数存在不公平的可能性，篡改概率，修改随机数很普遍，也无法证明或者证伪。

并且中心化随机数存在不可重复性，也就是事后无法验证公平。

区块链随机数解决了这个几千年来没法解决的问题，因为下一个区块的哈希值（TX ID）是无法预测的，结合交易在区块的位置（Index）和一个提前上链的时效为24小时的种子哈希，可以形成一个独有的种子组合。（24小时候种子也会公开，提前存上比原链的哈希值可证明种子没有被篡改）

独特种子公式： 种子 = 服务器种子 + 交易哈希 + 交易哈希所在区块位置（排名，0，1，2等）

有了不可预测的独特种子后通过梅森旋转（随机数算法中最常用的一种算法）624轮打乱后可以得到一个区块链随机数
通过验证工具 https://bytomdice.com/simulator 可以用来验证过往任何一笔交易所生成的随机数。

这里所说的随机是指区块链随机数不可预测，生成后可验证其公平性（可重复生成）。

验证工具所用的交易是（游戏ID 8888） https://bytomdice.com/fairgame/8888

* 服务器种子:
  - c9656775fd0f9d40dcde80fd597cc7e863144cb4fccbfc1e47de1ff4849ab054
* 交易哈希：
  - b31e198874365efc77338439e16c9b92c9b0c75a49c82c61c4e4423b71139996
* 交易哈希排名：
  - 0

随机数模拟器（验证工具）页面： https://bytomdice.com/simulator
