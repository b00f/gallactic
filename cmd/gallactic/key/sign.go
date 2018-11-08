package key

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gallactic/gallactic/cmd"
	"github.com/gallactic/gallactic/crypto"
	"github.com/gallactic/gallactic/keystore/key"
	"github.com/jawher/mow.cli"
)

//Sign signs the message with the private key and returns the signature hash
func Sign() func(c *cli.Cmd) {
	return func(c *cli.Cmd) {
		messageFile := c.String(cli.StringOpt{
			Name: "f file",
			Desc: "Message file path to read the file and sign the message inside",
		})
		message := c.String(cli.StringOpt{
			Name: "m message",
			Desc: "Text message to sign",
		})
		privateKey := c.String(cli.StringOpt{
			Name: "p privateKey",
			Desc: "Private key to sign the message",
		})
		keyFile := c.String(cli.StringOpt{
			Name: "k keyfile",
			Desc: "Path to the encrypted key file",
		})
		keyFileAuth := c.String(cli.StringOpt{
			Name: "a auth",
			Desc: "Key file's passphrase",
		})

		c.Spec = "[-f=<message file>] | [-m=<message to sign>]" +
			" [-p=<private key>] | [-k=<path to the key file>] [-a=<key file's passphrase>]"
		c.LongDesc = "Signing a message "
		c.Before = func() { fmt.Println(title) }
		c.Action = func() {
			var msg []byte
			var err error
			//extract the message to be signed
			if *message != "" {
				msg = []byte(*message)
			} else if *messageFile != "" {
				msg, err = ioutil.ReadFile(*messageFile)
				if err != nil {
					log.Fatalf("Can't read message file: %v", err)
				}
			}
			var signature crypto.Signature
			var pv crypto.PrivateKey
			//Sign the message with the private key
			if *privateKey != "" {
				pv, err = crypto.PrivateKeyFromString(*privateKey)
				if err != nil {
					log.Fatalf("Could not obtain privateKey: %v", err)
				}
				signature, err = pv.Sign(msg)
				if err != nil {
					log.Fatalf("Error in signing: %v", err)
				}
			} else if *keyFile != "" {
				var passphrase string
				if *keyFileAuth == "" {
					passphrase = cmd.PromptPassphrase("Passphrase: ", false)
				} else {
					passphrase = *keyFileAuth
				}

				kj, err := key.DecryptKeyFile(*keyFile, passphrase)
				if err != nil {
					log.Fatalf("Could not decrypt file: %v", err)
				}
				pv = kj.PrivateKey()
				signature, err = pv.Sign(msg)
			}

			//display the signature
			fmt.Println("Signature: ", signature)
		}
	}
}