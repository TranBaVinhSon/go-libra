package main

import (
	"context"
	"fmt"

	"github.com/the729/go-libra/client"
)

const (
	//defaultServer    = "ac.testnet.libra.org:8000"
	defaultServer    = "http://hk2.wutj.info:38080"
	trustedPeersFile = `
	[19f93314cbe8c0925a4492eb2f2f197ce6e11717449c218f50e043e37fa7a5f3]
	ns = "f6c7c31b68157b839cfc662cd9330ea5d27ed03d0b1ec9b2970c05fc66cd80d0"
	ni = "f0799c6e2b843066d5a23f42f775dce008fe92b9a0e8e9bf1208e83f5cec883a"
	c = "0397615aa5cc4cedadaa870511f381423e914012f5e341f474f3f608e0224e1d"
	
	[f9770caa0be0c0ad427f204c22a2c2d7b22ee373a1b9cf6fd768fbf48a079554]
	ns = "d844e3eb78ae751157af04c844c80bf55bc2bd7d2d7feea9b22b6ca833e6ac30"
	ni = "3f64b2233f8eb629f4b00724099895b819dd23e8986fdb5774bf9e176c9c8f26"
	c = "6d9e7d0d6c1c0acf7884ab8f8258f9f9dca74ca20653422c09652e348cf66b78"
	
	[4d78ab90b759ecacafe4e687c5db9cc2936a7a29c84aa8be777f54db519d756b]
	ns = "b0c0773494d1e87dd15e34a2636dfcaed771b01353227c69ab71c175a5ed437a"
	ni = "8383c5dbe5bc2888534e063c626084e58ba57661c62aa557022b807ee9838c02"
	c = "cec7be0f4808b68823ffe3db481564568a37f91ccdaa3bfc5d0d31b664e53695"
	
	[3b7c756cce9ad7d801b078a08ee91df5f8122e44011b4fdf6d6c20c016823b8f]
	ns = "d00667dcedc2c8359a261df5844408fe17fb768c33e211835f98893cfa304c91"
	ni = "d7eae84059a32ed1e2723909f8e05a3fe4ff96e3e496c8cb64917cd14e73ee3d"
	c = "ef84908705e82f835d665a94aa39990b101de671935a108c7d979fec91cbac10"
	
	[9102bd7b1ad7e8f31023c500371cc7d2971758b450cfa89c003efb3ab192a4b8]
	ns = "5f5ecda9576edd942ed22aa4735939092161445177cd456fd087c7bc1d6de403"
	ni = "b5eb9a2e5814c66df6c01a1dc94252a4ae6733e93a58187c5eb48d1f53be0b28"
	c = "576e91b04632683a11c3be3dc47a19f9f0a31ae947211f59c5fe02dfa2d07d68"
	
	[6687e9a6e6c3de0dc4f7e91eacbc676a228a9c0c46450bbccbd1072780bfcb30]
	ns = "dbe17f5fce01ad52dc78574d63124c108102291a042e332b4c82cf30c62ddca6"
	ni = "5c538501145dbc5e24b822edb1d0543ee63f00a022d2bd7b3e72f5468b12e82a"
	c = "62767f364921a4446a8d141298362497ab74b6cf697f2aa49eb2916b41a1b1cc"
	
	[c28b953590c58117ae8431456ea28f67c2f9e1733078b208e1a7bd5bd4081e9e]
	ns = "865a2c72cee4ede366bfdc391420885e4dc49f924506778d214dc3cf6d09c8c4"
	ni = "727bef5a30048db7000bbc54e6eb8dfe74b1fbac4fc627c8b3c78153e2e19805"
	c = "b62aa0cde6584521a9c5ba81fa0fe659eba8987356db81fc6f54e812ed0c2437"
	
	[4995559c4844b7e4101c486035329402a8ba4976c1be23080bac5bddf6a60f0d]
	ns = "e1a61336ee0a122e1c86e228dfc5067f24941b76a34da400af4e3a4da8da04a6"
	ni = "18ce05c58c4109cb123a9bfdf57b541b84814e7504f1aaeb09d23e05d20fb155"
	c = "d29b1294c223ec10749df7fee80c2ddf39c1fb32f1a89ae028e835b0502c19a0"
	
	[dfb9c683d1788857e961160f28d4c9c79b23f042c80f770f37f0f93ee5fa6a96]
	ns = "246ca919a3b39c95110e3bee891136ab087a9b3b9e84fa90cbf8f19c8abe62e3"
	ni = "8aa297d686dd2444de86ea3a68353d74af74b9659990d06ccaf4344e2b629b33"
	c = "3ca1400fb865befa8a21c58e90fc636ef2f84993a8396cb0e10008f876a00afd"
	
	[26873decd9330065988b0acf5027662b5097fb50913ae2a2652b50a9869df4fb]
	ns = "1512ea6e18f7c3069372d6883220aadb8fa525f9b6bb4d0eb473fc94a19ca87b"
	ni = "d19114154b3cbb76434d5fb5f50d215eab6ddf2b67907850118c8b3525896b11"
	c = "ada40b3c2039b0fe06e31f88c13505acc3d0731d9e6a09410f463a5c02ce5ef7"
`
)

func main() {
	c, err := client.New(defaultServer, trustedPeersFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	addrStr := "18b553473df736e5e363e7214bd624735ca66ac22a7048e3295c9b9b9adfc26a"
	addr := client.MustToAddress(addrStr)

	provenState, err := c.QueryAccountState(context.TODO(), addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	if provenState.IsNil() {
		fmt.Printf("Account %s does not exist at version %d.", addrStr, provenState.GetVersion())
		return
	}

	provenResource, err := c.GetLibraCoinResourceFromAccountBlob(provenState.GetAccountBlob())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Balance (microLibra): %d", provenResource.GetBalance())
	fmt.Printf("Sequence Number: %d", provenResource.GetSequenceNumber())
	fmt.Printf("SentEventsCount: %d", provenResource.GetSentEvents().Count)
	fmt.Printf("    Key: %x", provenResource.GetSentEvents().Key)
	fmt.Printf("ReceivedEventsCount: %d", provenResource.GetReceivedEvents().Count)
	fmt.Printf("    Key: %x", provenResource.GetReceivedEvents().Key)
	fmt.Printf("DelegatedWithdrawalCapability: %v", provenResource.GetDelegatedWithdrawalCapability())
	fmt.Println("OK")

}