# Random Number Algorithm
BytomDice Random Number Algorithm

The Algorithm we are using for BytomDice.com random number is easy to understand. 

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



