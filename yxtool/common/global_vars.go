package common

var (
	COMMIT_ID   = ""
	BRANCH      = ""
	LAST_AUTHOR = ""
	publicKey   = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCkcV8oxWJVkH87y8ePTEGxXGiw
PsE/WYO7ze3z88yYtT+5ktsPF/ZKas6Jhlo1mblmHjewYeVu8j8gl6OmDxUQ4EIu
4ZYze5pqkZ2TjJI0RyoK6Ahc4Rf3bxPz3iHrggBVNhBEvSUDkJbm/HdvvU3MSL1R
pQK+UICW9uQR5fOJ5wIDAQAB
-----END PUBLIC KEY-----`
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCkcV8oxWJVkH87y8ePTEGxXGiwPsE/WYO7ze3z88yYtT+5ktsP
F/ZKas6Jhlo1mblmHjewYeVu8j8gl6OmDxUQ4EIu4ZYze5pqkZ2TjJI0RyoK6Ahc
4Rf3bxPz3iHrggBVNhBEvSUDkJbm/HdvvU3MSL1RpQK+UICW9uQR5fOJ5wIDAQAB
AoGAIII5Kbo7wmaagJu8i9urcQCvwYfmdDuiC1KepAeift32Bx29KLK4AeYzhoSE
+Ou29zUqFTiFPvWs/Sh9NartouNL0ZuLxQLHtKwsoU4Hf+Np2yXScsZJz+hFUWoG
VCwWmHBklT7UrHzogWbLKuUsPJhqs2SGTPvQ5mjium/eToECQQDSpBRiMgBVS/V/
d85Gu/BsVDcgz4/hFai8eqtbrDUJUV3X+oOMwNv1ZXcGDUpCqVg1N1F3ZoANERMW
mgC7tUoRAkEAx9qMeyzc8v9aM2Vx6mAlcnjokjxYfD5Avvu6z0niDUwAegs129zv
sixVV01sNdGMaBtOC27YqM1yIvlODRBcdwJAfy8pS50rTFVf90ZMm7MQ2SgEXWBW
XK7qQlvDKcgSAf1KL4WIdq1gYoyQUGwn4ZZEt2UeYHzYZotfvi7gsHYnMQJBAK1Y
/pL0le1b2dRMwKNHpBOHlksJdsLFiIlfdWjJL/7SlnNFeADtwSbtV1SzPhHsYXoH
fQljqx+3m7GERtQEBf8CQQCn5V1DizTjMFinqCueII0/hn242IC4g1lME60KBk5n
Ha2a55wjVEvnVg/rzS02srsUuKOvQSQZGSLT3j/Ay8YW
-----END RSA PRIVATE KEY-----`
	V3_ONLY_UIN_OWNERUIN = "v3_only_uin_owneruin"
)

func SetPrivateKey(ptKey string) {
	privateKey = ptKey
}

func GetPrivateKey() string {
	return privateKey
}

func SetPublicKey(pubKey string) {
	publicKey = pubKey
}

func GetPublicKey() string {
	return publicKey
}
