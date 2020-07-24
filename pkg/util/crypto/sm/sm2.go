/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  sm2
 * @Version: 1.0.0
 * @Date: 2020/7/24 16:10
 */

package sm

import (
	"crypto/elliptic"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"math/big"
)

func FromECDSAPub(pub *sm2.PublicKey) []byte {

	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}

func FromECDSA(priv *sm2.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return math.PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}

func SignData(key *sm2.PrivateKey, digest []byte) (r, s, pub *big.Int, err error) {

	h := sm3.New()
	h.Write(digest)
	hash := h.Sum(nil)

	r, s, err = sm2.Sm2Sign(key, hash, default_uid)

	if err != nil {
		return
	}

	pb := FromECDSAPub(&key.PublicKey)

	pub = new(big.Int).SetBytes(pb[1:])

	return

}
