package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jorelosorio/go-signature/signature"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Signature",
		Usage:       "sign",
		Description: "A command-line tool playground to encode, sign and decode data",
		Version:     "1.0.0",
		Compiled:    time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jorge Osorio",
				Email: "jorelosorio@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "create-keys",
				Aliases: []string{"cks"},
				Usage:   "Creates a new pair of private and public keys",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "export-path",
						Aliases:  []string{"ep"},
						Usage:    "Exports keys as .pem files in the specified path",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					if outputPath := c.String("export-path"); outputPath != "" {
						asymmetricKey := signature.NewAsymmetricKey()

						asymmetricKey.ExportPrivateKeyToPem(outputPath)
						asymmetricKey.ExportPublicKeyToPem(outputPath)
					}

					return nil
				},
			},
			{
				Name:    "encode-message",
				Aliases: []string{"emsg"},
				Usage:   "Creates a new message, sign it and encode it using -base64 encoding-",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "sender",
						Aliases:  []string{"s"},
						Usage:    "Who sends the message?",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "payload",
						Aliases: []string{"p"},
						Usage:   "Message payload",
					},
					&cli.StringFlag{
						Name:    "payload-path",
						Aliases: []string{"pp"},
						Usage:   "Message payload from a file",
					},
					&cli.StringFlag{
						Name:     "private-key-path",
						Aliases:  []string{"prkp"},
						Usage:    "The private key path to sign the message",
						Required: true,
					},
					// TODO: Include an option to export data to a file
				},
				Action: func(c *cli.Context) error {
					if sender, payload, payloadPath, prkPath := c.String("sender"), c.String("payload"), c.String("payload-path"), c.String("private-key-path"); sender != "" && prkPath != "" && (payload != "" || payloadPath != "") {
						asymmetricKey := signature.AsymmetricKey{}
						asymmetricKey.ImportPrivateKey(prkPath)

						payloadData := []byte(payload)
						if payloadPath != "" {
							payloadData = signature.ReadFile(payloadPath)
						}

						message := &signature.Message{Sender: sender, Payload: payloadData}

						signatureByte, encodedContainer := signature.PackAndSignMessage(message, &asymmetricKey)
						fmt.Printf("Signature\n==========\n%s\n\n", signature.EncodeBase64(signatureByte))
						fmt.Printf("Container data\n==========\n%s\n", signature.EncodeBase64(encodedContainer))
					}

					return nil
				},
			},
			{
				Name:    "decode-message",
				Aliases: []string{"dmsg"},
				Usage:   "Decode a message, verifies the signature and print out the message content",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "base64-message",
						Aliases:  []string{"b64msg"},
						Usage:    "Encoded message in Base64 format",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "public-key-path",
						Aliases:  []string{"pkp"},
						Usage:    "The public key path to verify the message",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					if base64EncodedMessage, pkPath := c.String("base64-message"), c.String("public-key-path"); base64EncodedMessage != "" && pkPath != "" {
						asymmetricKey := signature.AsymmetricKey{}
						asymmetricKey.ImportPublicKey(pkPath)

						decodedMessage := signature.DecodeBase64(base64EncodedMessage)
						messageContainer := signature.DecodeContainer(decodedMessage)

						if isAuthentic := signature.IsAuthentic(messageContainer, &asymmetricKey); isAuthentic {
							fmt.Println("The message is authentic!")
							fmt.Println(messageContainer.Message)
							os.Exit(0)
						}

						fmt.Println("The message is not authentic!")
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
