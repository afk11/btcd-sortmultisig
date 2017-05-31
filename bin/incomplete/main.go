package main

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"log"
	"fmt"
	"github.com/afk11/sortmultisig/sortutil"
)

func main() {
	//txHex := "0100000001951aefe7968498e74fc5fc52d81009d95a01eb2eaae67cc9fcef61b68ebbc0b800000000fd5c01004730440220516b0f747d126b12cdfb891239b3dfc547a719175edfe0309503e76d6f1ea27602205ff41618e5019dcd748ae51b7e4aa43b3d0acb0d3ecff4c05aa43ad87a1e2df90147304402207335ff2de2b32168ec9ea26752bfe1c34b53ad15cc2904cb570e62a4cc851db2022045575bce2cf374ab0d7091003a418c20bc05fcea52f632762b6691c1aedcd378014cc95241043e49ec68abcf030dfc8ec7dfcb388b17fed99134d5f910c87e947f0cc86a1cf9c29b27ddbd0443b6d40fc5cb35eb13fcb41bf0e4f63d74bea2576e8db07dd1864104e38fa9a9dfa216d45d90cfca8ca2566f2d9aa9c1846e6dd6ab0756c07262abe1c7f8a60ff2357bc2ea9fb597bfbcf4d3e4fe605a294171dc32421578125136e0410473682ed776e9d0afee6cd52f4a4b20ad458956ef5567d5747853b689bb44a6a90736b515aa63bc5703c7d1a5662b7e2421b9436597fd8bf8da216f2b1cba01cc53aeffffffff01010000000000000017a914221154d32a82ae83f9e75431feae77a37af771a68700000000"

	// 011, so sigs 0, 1, not 2
	//txHex := "0100000001951aefe7968498e74fc5fc52d81009d95a01eb2eaae67cc9fcef61b68ebbc0b800000000fd5c01004730440220516b0f747d126b12cdfb891239b3dfc547a719175edfe0309503e76d6f1ea27602205ff41618e5019dcd748ae51b7e4aa43b3d0acb0d3ecff4c05aa43ad87a1e2df90147304402207335ff2de2b32168ec9ea26752bfe1c34b53ad15cc2904cb570e62a4cc851db2022045575bce2cf374ab0d7091003a418c20bc05fcea52f632762b6691c1aedcd378014cc95241043e49ec68abcf030dfc8ec7dfcb388b17fed99134d5f910c87e947f0cc86a1cf9c29b27ddbd0443b6d40fc5cb35eb13fcb41bf0e4f63d74bea2576e8db07dd1864104e38fa9a9dfa216d45d90cfca8ca2566f2d9aa9c1846e6dd6ab0756c07262abe1c7f8a60ff2357bc2ea9fb597bfbcf4d3e4fe605a294171dc32421578125136e0410473682ed776e9d0afee6cd52f4a4b20ad458956ef5567d5747853b689bb44a6a90736b515aa63bc5703c7d1a5662b7e2421b9436597fd8bf8da216f2b1cba01cc53aeffffffff01010000000000000017a914221154d32a82ae83f9e75431feae77a37af771a68700000000"
	// 100 so sigs 2, not 0, 1
	txHex := "0100000001951aefe7968498e74fc5fc52d81009d95a01eb2eaae67cc9fcef61b68ebbc0b800000000fd14010047304402202478513599ca49bec6de0751d36fdc47ffe697a25d3ada64694b08040782983d02201915074336baf3d9bca7f7751aa81605c2a8013c221c88e1fefc76ea1cd68d74014cc95241043e49ec68abcf030dfc8ec7dfcb388b17fed99134d5f910c87e947f0cc86a1cf9c29b27ddbd0443b6d40fc5cb35eb13fcb41bf0e4f63d74bea2576e8db07dd1864104e38fa9a9dfa216d45d90cfca8ca2566f2d9aa9c1846e6dd6ab0756c07262abe1c7f8a60ff2357bc2ea9fb597bfbcf4d3e4fe605a294171dc32421578125136e0410473682ed776e9d0afee6cd52f4a4b20ad458956ef5567d5747853b689bb44a6a90736b515aa63bc5703c7d1a5662b7e2421b9436597fd8bf8da216f2b1cba01cc53aeffffffff01010000000000000017a914221154d32a82ae83f9e75431feae77a37af771a68700000000"
	txBytes, _ := hex.DecodeString(txHex)
	utilTx, err := btcutil.NewTxFromBytes(txBytes)
	if err != nil {
		panic(err)
	}
	tx := utilTx.MsgTx()

	scriptPubKey, _ := hex.DecodeString("a914221154d32a82ae83f9e75431feae77a37af771a687")

	flags := txscript.ScriptVerifySigPushOnly | txscript.ScriptBip16 | txscript.ScriptVerifyCleanStack

	engine, err := txscript.NewEngine(scriptPubKey, tx, 0, flags, nil, nil, 0)
	if err != nil {
		panic(err)
	}
	ops, err := engine.ExecuteSignOp(true)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", ops)
	var complete string = "no"
	path, pos, err := ops.IsComplete(0)
	if err != nil {
		panic(err)
	}
	if path && pos {
		complete = "yes"
	}
	log.Printf("Op 1 is complete? %s\n", complete)

	keys, sigs , err:= ops.GetIncompleteOps(0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(keys); i++ {
		sig, ok := sigs[i]
		key, _ := sortutil.FormatPublicKey(keys[i].Key, keys[i].Format)
		fmt.Printf("idx %d (for %s) \n", i, hex.EncodeToString(key))
		if ok {
			fmt.Printf("  - %s \n", hex.EncodeToString(sig.Serialize()))
		} else {
			fmt.Printf("  - null \n")
		}
	}
}
