/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  SM2KeyGenOpts
 * @Version: 1.0.0
 * @Date: 2020/7/16 15:11
 */

package keystore

// ECDSAP256KeyGenOpts contains options for ECDSA key generation with curve P-256.
type SM2KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM2KeyGenOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM2KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}
