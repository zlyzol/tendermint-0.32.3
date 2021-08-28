# Evidence Reactor

## Channels

[#1503](https://github.com/zlyzol/tendermint-0.32.3/issues/1503)

Sending invalid evidence will result in stopping the peer.

Sending incorrectly encoded data or data exceeding `maxMsgSize` will result
in stopping the peer.
