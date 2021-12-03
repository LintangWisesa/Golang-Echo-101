package main

func CanonicalizeMdn(mdn string) string {
	if mdn[0:1] == "+" {
		mdn = mdn[1:]
	} else if mdn[0:1] == "0" {
		mdn = "62" + mdn[1:]
	} else if !(mdn[0:2] == "62") {
		mdn = "62" + mdn
	}
	return mdn
}