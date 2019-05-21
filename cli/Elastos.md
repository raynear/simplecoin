https://github.com/elastos
https://github.com/elastos/Elastos/wiki/A-Non-Developer-Guide-to-Elastos
https://github.com/elastos/Elastos/wiki/A-Developer-Guide-to-Elastos
https://github.com/elastos/Elastos/wiki/Interact-with-an-ELA-node
https://github.com/elastos/Elastos.ELA.Client
https://github.com/elastos/Elastos.ELA
https://github.com/elastos/Elastos.ELA.SPV <= 컴파일 안됨
https://github.com/elastos/Elastos.ELA/blob/master/docs/Elastos_Wallet_Node_API_CN.md
https://steemit.com/coinkorea/@collector999/elastos-developer-news-testnet-launch?sort=trending <= 동영상


https://blockchain.elastos.org/blocks <= explorer
https://wallet.elastos.org/ <= 지갑 생성 가능 - 왜 가능한지 잘 모르겠음

Bitcoin에 기생하는 방식으로 Consensus를 얻습니다.
정확한 구조는 찾지 못했습니다.

Elastos 는 책같은 어떤 컨텐츠의 소유권을 거래하는 인터넷 환경을 목표로 하는 것으로 생각 됩니다.

코인데몬은 Elastos.ELA 라는 것으로 동작하며
Elastos.ELA.Client를 컴파일 하면 cli가 생성됩니다.
Wallet은 Elastos.ELA.SPV 이지만 현재 컴파일 되지 않습니다.
원래 cli에서 됐었는데 그 부분은 막혔습니다.
Elastos.ELA를 실행하고 curl로 restAPI를 던져주면 됩니다.
json rpc 2.0이 아닌 restapi를 사용합니다.
Elastos.ELA와 Elastos.ELA.Client는 config 파일을 필요로 합니다.

Chain DB는 실행한 위치에 Chain/ 아래 생깁니다.

testnet을 설정하는 부분이 이상합니다.
config.json 파일을 수정해야 하는데 ActiveNet을 MainNet으로 하고 peer들을 testnet peer로 설정해야 합니다.
