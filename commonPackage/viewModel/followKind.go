package viewModel

import (
	"commonPackage/model/account"
)

type FollowKind struct {
	KindId   int64
	KindName string
	CONTENT  []account.Lctb_userInfo
}
