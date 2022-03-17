package main

import (
	"fmt"
	"os"
	"signatures-playground/helpers"
	"signatures-playground/structs"
	pb "signatures-playground/structspb"
	"signatures-playground/utilities"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Signatures Playground",
		Usage:       "sp",
		Description: "A command-line tool playground to encode, sign and decode data",
		Version:     "1.0",
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
						asymmetricKey := structs.NewAsymmetricKey()

						asymmetricKey.ExportPrivateKeyToPem(outputPath)
						asymmetricKey.ExportPublicKeyToPem(outputPath)
					}

					return nil
				},
			},
			{
				Name:    "pack-message",
				Aliases: []string{"pmsg"},
				Usage:   "Creates a new message, sign it and pack it using base64 encoding",
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
				},
				Action: func(c *cli.Context) error {
					if sender, payload, payloadPath, prkPath := c.String("sender"), c.String("payload"), c.String("payload-path"), c.String("private-key-path"); sender != "" && prkPath != "" && (payload != "" || payloadPath != "") {
						asymmetricKey := structs.AsymmetricKey{}
						asymmetricKey.ImportPrivateKey(prkPath)

						payloadData := []byte(payload)
						if payloadPath != "" {
							payloadData = utilities.ReadFile(payloadPath)
						}

						message := &pb.Message{Sender: sender, Payload: payloadData}

						signature, encodedContainer := helpers.PackAndSignMessage(message, &asymmetricKey)
						fmt.Printf("Signature\n==========\n%s\n\n", helpers.EncodeBase64(signature))
						fmt.Printf("Container data\n==========\n%s\n", helpers.EncodeBase64(encodedContainer))
					}

					return nil
				},
			},
			{
				Name:    "unpack-message",
				Aliases: []string{"umsg"},
				Usage:   "Unpack a message, verifies the signature and print out the message content",
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
						asymmetricKey := structs.AsymmetricKey{}
						asymmetricKey.ImportPublicKey(pkPath)

						decodedMessage := helpers.DecodeBase64(base64EncodedMessage)
						messageContainer := helpers.DecodeContainer(decodedMessage)

						if isAuthentic := helpers.IsAuthentic(messageContainer, &asymmetricKey); isAuthentic {
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
