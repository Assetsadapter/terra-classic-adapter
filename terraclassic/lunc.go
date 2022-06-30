package terraclassic

import "github.com/blocktree/openwallet/v2/log"
import "github.com/Assetsadapter/terra-adapter/terra"

const (
	Symbol = "LUNC"
)

type WalletManager struct {
	*terra.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = terra.NewWalletManager()
	wm.Config = terra.NewConfig(Symbol, "lunc master key")
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "terra-classic"
}
