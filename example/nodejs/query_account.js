// Use following require statement when using npm.
// const { libra } = require("gopherjs-libra");
const { libra } = require("../../gopherjs-libra/gopherjs-libra.js");

const fromHexString = hexString =>
    new Uint8Array(hexString.match(/.{1,2}/g).map(byte => parseInt(byte, 16)));

const toHexString = bytes =>
    bytes.reduce((str, byte) => str + byte.toString(16).padStart(2, '0'), '');

const defaultServer = "http://hk2.wutj.info:38080",
    waypoint = "0:bf7e1eef81af68cc6b4801c3739da6029c778a72e67118a8adf0dd759f188908";

var addrStr = "18b553473df736e5e363e7214bd624735ca66ac22a7048e3295c9b9b9adfc26a"
var addr = fromHexString(addrStr)

var client = libra.client(defaultServer, waypoint)
client.queryAccountState(addr)
    .then(r => {
        if (r.isNil()) {
            throw "Account " + addrStr + " does not exist at version " + r.getVersion() + "."
        }
        return r.getAccountBlob()
    })
    .then(r => r.getLibraAccountResource())
    .then(r => {
        console.log("Address: ", addrStr)
        console.log("Balance (microLibra): %d", r.getBalance())
        console.log("Sequence Number: ", r.getSequenceNumber())
        console.log("Sent Events: ", r.getSentEvents())
        console.log("Received Events: ", r.getReceivedEvents())
        console.log("Delegated withdrawal capability: ", r.getDelegatedWithdrawalCapability())
        console.log("Event generator: ", r.getEventGenerator())
        console.log("Proven at ledger version / time: ", r.getLedgerInfo().getVersion(), r.getLedgerInfo().getTimestampUsec() / 1000000)
    })
    .catch(e => {
        console.log("Error: ", e)
    })
