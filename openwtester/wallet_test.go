package openwtester

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

var appID = "assets-adapter"

func InitWalletManager() *openw.WalletManager {

	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = "."
	tc.EnableBlockScan = false
	return openw.NewWalletManager(tc)
}
func getPassword() string {
	return "12345678"
}
func Test_a(t *testing.T) {
	wm := InitWalletManager()

	var accountID = "Dd4gAqptJ1dDw3qt8Q4CG6b8RnZeYzP5hVcfihDyJekK"
	var address = "terra1vqgd4h30z5wzrr7f6danmvqtw349yjr5wqxrhx"
	account, err := wm.GetAssetsAccountInfo(appID, "", accountID)

	wrapper, err := wm.NewWalletWrapper(appID, account.WalletID)
	if err != nil {
		t.Error(err)
		return
	}
	err = wrapper.UnlockWallet(getPassword(), 5*time.Second)
	if err != nil {
		t.Error("get HDKey from wallet wrapper failed, err=", err)
		return
	}
	key, err := wrapper.HDKey()
	if err != nil {
		t.Error("get HDKey from wallet wrapper failed, err=", err)
		return
	}
	addressObj, err := wm.GetAddress(appID, account.WalletID, accountID, address)
	if err != nil {
		t.Error(err)
		return
	}

	childKey, _ := key.DerivedKeyWithPath(addressObj.HDPath, owcrypt.ECC_CURVE_SECP256K1)
	keyBytes, err := childKey.GetPrivateKeyBytes()
	pubkeyBytes := childKey.GetPublicKey().GetPublicKeyBytes()
	log.Info("hex public key:", hex.EncodeToString(pubkeyBytes))
	log.Info("hex private key:", hex.EncodeToString(keyBytes))
}
