package address

import (
	"testing"
)

var valid_addresses = []string{
	"xrb_38nm8t5rimw6h6j7wyokbs8jiygzs7baoha4pqzhfw1k79npyr1km8w6y7r8",
	"xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
	"xrb_3arg3asgtigae3xckabaaewkx3bzsh7nwz7jkmjos79ihyaxwphhm6qgjps4",
	"xrb_3pczxuorp48td8645bs3m6c3xotxd3idskrenmi65rbrga5zmkemzhwkaznh",
	"xrb_3hd4ezdgsp15iemx7h81in7xz5tpxi43b6b41zn3qmwiuypankocw3awes5k",
	"xrb_1anrzcuwe64rwxzcco8dkhpyxpi8kd7zsjc1oeimpc3ppca4mrjtwnqposrs",
}

var invalid_addresses = []string{
	"xrb_38nm8t5rimw6h6j7wyokbs8jiygzs7baoha4pqzhfw1k79npyr1km8w6y7r7",
	"xrc_38nm8t5rimw6h6j7wyokbs8jiygzs7baoha4pqzhfw1k79npyr1km8w6y7r8",
}

func TestValidateAddress(t *testing.T) {
	for _, addr := range valid_addresses {
		if !ValidateAddress(addr) {
			t.Errorf("Valid address did not validate")
		}
	}

	for _, addr := range invalid_addresses {
		if ValidateAddress(addr) {
			t.Errorf("Invalid address was validated")
		}
	}

}

func BenchmarkGenerateAddress(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pub, _ := GenerateKey()
		PubKeyToAddress(pub)
	}
}
