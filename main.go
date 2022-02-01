package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
)

// run this as ./main <emoji string wanted> like ./main "🐨🧋✨"


var (
	emojiCoder = encoding{set: emoji, dataLen: addrBytesLen}
	textCoder  = encoding{set: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), dataLen: addrBytesLen}
	emoji = []rune(`🦆🐣🐤🐥🎟🪐🤖🌥🎯🦋🟪🛷📸🤾🔞🎀🔈🪦🦢➖🥊⚧🌲🪗😚🎌🌡🐁👨🎗🌝⛄💨🖕🧨📏🦒🎨🔦🧆🚊🗑🏅🛸🤿🐶🚂🚀🦨🪲🍵🤶🎥💟🥷🦟🐞😳🥦📎🏕🐅🏇🕚🤫🕧🌀😫🤣🕘🥱🐠📣🌙🏒💱🎏🟡🚅🏝🔧🧃😹🅰🕤🥼🚍🫂🦩🖨😺🌗🌭📢🖐💎🚣🌉🤽🦻🤦🧱🌵💦📳🔚⚾😡😻🤠🐦⛰🕕🐂💫📿🌚🤺⏰😰📬🛄🤴➰🍰🙋🍇🔹🏷👱🦀🚬🏣💇⚫🥑💃🚰🦶🐟😒📩🐹💋🐵🥣📴⛱💅🤑🦎🦜🤐👌🚧🏈💗🎋🩰🛌🏹😩🏺🏏🗽🫀🎅🧘🎁🪥😤📘🥇🛻🤌⏭🕺💻⛪🚦🟨❌💽🙆💡🪒🔛🍾🦉🖲💬🚡👯⚡🧁🏉🌪💁🌎🦫💔🧭🏬🍒🧦🍩🐖😞⏯🛶⛓🥳◽⏺🚜🥈🦭👆🤲🏩🛺⚪🌇🧶👻🪙➗🚆🍤👣📪🤳🔫🏟🧗🪄🐎🚮🚹🎾🕜👒🔨🌊🧎💲🌸⌚🍘🖋⏸🍣🌤📖💐🪓🧳🪰😾👪👸💩♿🚝🦮🍈🔘🚋🔱🛬🦄🏑👘🩸💪👥👋🛏🎚🏰🩲🔩🔰🎪🐔🌒🧓🧯🔗👟🦍🎍🤹🥽🟤🏯📇🟥🛀🕗🔴🪶🔄🧤🌴🚨🌑🥮👗🐍🔖🕞⏱📮🪟🐘🍌🕰🪴📲🐌🔶📰🍴🏜😪🍔⛳🦽🏂👫🔉✊📯👈🐴📡🎱🐽🏌🕹😄🥌🪆✋🤞🤧⛏🎐🚓🥥🐃🚉😶🥖🛑🚛🔇💿🍿🛰🐊📋🪁🛂🚚🚙🏭🧫👵🆖🏍🌽🧲⏫🔳😔🥒🤷🪱⬜👤🧷🅾🏊🎡🧼🧧🧄🍉🍪👚🔎🐛🍅🎽🤍🆚🌦🚔🛅😓🦘👢🥠🎠🪡👇🍷🤱📕🧥🎰👉🪣🥟🐐🔵🥐💧🦠👛🥗🎿🛣🕳🧔🪠🍆🚩💥🎦🚐🏨💌🧢🙂✅📚🏦🦬🕑👖🦹🧑🦖😌🐸🍐🧿😿🎇🥫✨👝🧖🍗🎼🕡🚈🟩⚓🪚❔🦇🔻🚃💠🟧🐒🚠😛🗣🦊🐱💷🦂🟫🔬👧🫖📟🍼🟦🧋🏁🚳🆗🧩🚿🫑🚺🤭🐿🚎📁💤🏤🤎📄👜🐜🌅🔁🔃😀💘🔐🚘🕶🕣🚼💾🥧📧❕🔼🗄🌼🪞⏪🕢📔👺🕒👶🏠💕🥓🎉🦙🛤🪘📥🧏🫁🦑🥁🩴🔕🌫🥰🦾📼🥯⛹🐕🛳🕟🦯🙀📜🍸🤥💆🕓🔓🍲🥴🐬🛃🧣👦🐨🪨😱🛩👓👹🩹⛸🛋🪜⛴🦁🏛🚾🦚👎🧙🥉🦛🧮🤤🎴🕊🍜🧪🤵🏓💯🍯🚫🧬💮🟣🛁😉🗨🦺🤏🎆😊🤗🚑🌺🦈🕔😬🦏🗾🚪😥🫒🦵😷🥭🚽🔌😝🎤🐚🪂🆘🍺🍀🛼🦤💖🍦🕸🆙⛑💴🐙❎🥻🍕🥡🚢👽😑🥜🔺🐉⬛🐢🛎🍻🅱⌛🥺🏚🐯🧞🔋⚽📀🍂🪝🍮🪖🐪👿🏵🥙💰🗺🧺➿🌓🗡🧾🐫🌱📉👙🌕🚥🤩🍨🦓🌧🥶📆🏮🧜🐗🌠🧟🚯🤙🐭🚏🚕🦡🔲🚲📂🌟📵📻🪵👲🎢🧚🔟👩🚻😎🧐😨📍🎞😯❗🥎🐄🦦🍱🎃🖼⛎🪔🪃🌰🗞🆕👍🤒💼❓🥬💈🙊🛗🎄📺🥚🍓🥤💸🤼🌜⭕🤰🗯🏸🎊⛩🧒🗜🦅🤮🔂📝🩳😈🥿🍎🚴🧹🥢🍡🎖👳😭☔🆓🎙📑🍑🦪🏗🥨🐮🐝🥘🌋🔆🌿👭🚗🔜🍟👐🚱🍢🏐🌶🖊🦥🃏🚒🥀🖇🖱🕠🥅📨🏞📅🗳🐻🧸🦗🥾🙉🌐🛢😐😁🎷🦌😣🏪🚶🌘🚭💄🥍🥕😏🍭📒💳🗓🤓🚞💂🙁🏴🍛🦣🔷🙅👔🆎💝🎣🌏👠😃🎭🤝📗🏙📈🗿🤪🍍🍫🏆👊🕥🔥😟🐆🤬🪅📽🐺⛔🌾🦔💢🤡🌷🌆🙏📌💉😆🗝👾🖍🐩👃🀄🦴🧵🗂🛫🕦🚁🦸🍝🥃🆒🕝⛈🌯🧠🤚🛕🆔🧡👬😅🔙🦞🎈🧕🏖📤🕙😴💓⛲🛥🐡💣🕛🌻🎮🎹😂👏🪤🦐🛴🧀🚌🔮💍🥸🏡🚵🐑⛷🛍🌛🍊🍹⏮🛹🔊📙🏎👡🫓🍋🐏🎳👮⏬🦃🐋🐾⛅🪧😖💊🥩😘👰🏥🏢🆑📊🍚👑🧴🌁🟢🧊⛽👂🧽🏃💶🍄🙎👞⏹🛠🧝🫐🏔👕⏳🎩🍠⭐🥲😵💹🕵🚷📫📓🪢🥞🩺🧛🙌🔭🚄🙍📛🕷😢🔝🎓💚💞😦🧍🔍😽⛵🐲🗒🎎🐓💀👀📐📠🦧🎫🍥🔒😧🗼👼🌳🏘📶🗃🏄💛😇🛡🙇🌄🧻😲😗🚤🌬👷🏧🤟🕯🎂⏲🥂💙🧰🐼🌌📃💵😮🔏👴🐷🌨➕🤘🍬🛖👅🫔🌂🦼🧂🍁☕🩱🧈🥄🔪◾🚇🥔📱🦝🕖🥋💺🥏📦🤨🕐💭💑🍞🌖⏩🍽🟠🥛🐰🔀🍃📷🗻🤢😕🐇🔸🪕🏋🦷🏫👁🌔🔔🎵👄🍏🦿🎲🔅🐳🚟😠🌮🖖🤛🅿⛺🌩🎬🏳😍🐀😜🙈📭🫕🎺📞🧅🖌🪀🎑🌹🌞💜🌈🪑🙃😋🖥😼🪛🍙💏🚖🥵🤔🙄🤜🎻🌃🍶🥝🎧🤯🔑🖤🛵😸📹🚸🛒🥪🎛🌍🐈🍖🏀🤕🎸🍧🎶🐧🦕🧇💒🎒🕴🤸🍳😙🧉🪳`)

	wantPref = os.Args[1] // change this to a string like "🐨🧋✨" if running as a script
)

const (
	addrBytesLen           = 24
	checksumLen            = 4
	textAddressPrefixChar  = 'Q'
	emojiAddressPrefixChar = '🦆'
	//versionChar               = '0'
	runs = 100_000_000
)

func main() {
	for i := 0; i < runs; i++ {
		pub, priv , err := MakeKeyPair()
		if err != nil {
			fmt.Println(err)
			return
		}
		b := addChecksum(
			DoubleShasumBytes([]byte(pub))[:20], // Truncation is fine. SHA256 is designed to be comparable to a random oracle and the US Government itself is okay with it: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf page 32 (section 7)
		)
		//s :=
		if strings.Contains(emojiCoder.Encode(b), wantPref) {
			fmt.Println("\n\nGot a match! As emoji:", KeyToAddress(pub).Emoji, "As text:", KeyToAddress(pub).Text)
			fmt.Println("Public key PEM (save to ~/.config/duckcoin/pubkey.pem):")
			fmt.Println("\n-----BEGIN DUCKCOIN (ECDSA) PUBLIC KEY-----\n" + pub + "\n-----END DUCKCOIN (ECDSA) PUBLIC KEY-----\n")
			fmt.Println("Private key: (save to ~/.config/duckcoin/privkey.pem):")
			fmt.Println("\n-----BEGIN DUCKCOIN (ECDSA) PRIVATE KEY-----\n"+priv+"\n-----END DUCKCOIN (ECDSA) PRIVATE KEY-----")
		}
		if i % 100_000 == 0 && i != 0 {
			fmt.Println("at", i, 100*float64(i)/float64(runs), "% of", runs, "runs")
		}
	}
}


// KeyToAddress derives a Duckcoin Address from a Duckcoin Public Key
func KeyToAddress(key string) Address {
	return BytesToAddress(sliceToAddrBytes(addChecksum(
		DoubleShasumBytes([]byte(key))[:20], // Truncation is fine. SHA256 is designed to be comparable to a random oracle and the US Government itself is okay with it: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf page 32 (section 7)
	)))
}

func sliceToAddrBytes(addrSlice []byte) [addrBytesLen]byte {
	var arr [addrBytesLen]byte
	copy(arr[:], addrSlice)
	return arr
}

func addChecksum(data []byte) []byte {
	dataCopy := make([]byte, len(data), cap(data))
	copy(dataCopy, data) // don't modify original data

	hash := DoubleShasumBytes(data)
	return append(dataCopy, hash[:checksumLen]...)
}

func BytesToAddress(b [addrBytesLen]byte) Address {
	addr := Address{
		bytes: b,
	}
	addr.Text = string(textAddressPrefixChar) + textCoder.Encode(addr.bytes[:]) // len(base64(20 + 4 bytes)) + len("q" + versionChar) = 24 * 4/3 + 2 = 34 len addrs
	addr.Emoji = string(emojiAddressPrefixChar) + emojiCoder.Encode(addr.bytes[:])
	return addr
}

// DoubleShasumBytes returns sha256(sha256(record))
func DoubleShasumBytes(record []byte) []byte {
	h := sha256.New()
	h.Write(record)
	hashed := h.Sum(nil)
	h.Reset()
	h.Write(hashed)
	hashed = h.Sum(nil)
	return hashed
}

type Address struct {
	Emoji string `json:",omitempty"`
	Text  string `json:",omitempty"`
	bytes [addrBytesLen]byte
}

// MakeKeyPair creates a new public and private key pair
func MakeKeyPair() (pub, priv string, err error) {
	pubkeyCurve := elliptic.P256()                              // see http://golang.org/pkg/crypto/elliptic/#P256
	privkey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	if err != nil {
		return "", "", err
	}
	pubkey := &privkey.PublicKey
	pub, err = publicKeytoDuck(pubkey)
	if err != nil {
		return "", "", err
	}
	priv, err = privateKeytoDuck(privkey)
	if err != nil {
		return "", "", err
	}
	return pub, priv, nil
}


// publicKeytoDuck serializes public keys to a base64 string
func publicKeytoDuck(pubkey *ecdsa.PublicKey) (string, error) {
	marshalled, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(marshalled), nil
}

// privateKeytoDuck serializes private keys to a base64 string
func privateKeytoDuck(privkey *ecdsa.PrivateKey) (string, error) {
	marshalled, err := x509.MarshalECPrivateKey(privkey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(marshalled), nil
}


type encoding struct {
	set []rune
	// dataLen is the length of the byte data used. In duckcoin, this is always 24.
	dataLen int
}

func (e *encoding) Encode(data []byte) string {
	convertedBase := toBase(new(big.Int).SetBytes(data), "", e.set)
	// repeat emoji[0] is to normalize result length. 0 because that char has zero value in the set
	return strings.Repeat(string(e.set[0]), e.EncodedLen()-len([]rune(convertedBase))) + convertedBase
}

func (e *encoding) EncodedLen() int {
	return int(math.Ceil(
		float64(e.dataLen) * math.Log2(256) / math.Log2(float64(len(e.set))),
	))
}

func toBase(num *big.Int, buf string, set []rune) string {
	base := int64(len(set))
	div, rem := new(big.Int), new(big.Int)
	div.QuoRem(num, big.NewInt(base), rem)
	if div.Cmp(big.NewInt(0)) != 0 {
		buf += toBase(div, buf, set)
	}
	return buf + string(set[rem.Uint64()])
}