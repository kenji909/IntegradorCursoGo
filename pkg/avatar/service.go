/*
CONSIGNA:

Dada una información personal, como una dirección de correo electrónico, una dirección IP
o una clave pública, se debe desarrollar un package que permita generar un avatar único.

Imagine, por ejemplo, que está creando una nueva aplicación y desea que todos sus
usuarios tengan un avatar único y predeterminado. Para ello, debe desarrollar y publicar el
package que escribirá permitirá la generación de dichos avatares. Por ejemplo, GitHub
utiliza este enfoque y genera un idéntico para todos los usuarios nuevos que no tienen una
imagen de avatar seleccionada
--------------------------------------------------------------------------------------------
Logica que voy a usar:
1- obtener la información personal, en este caso usare el nombre de usario "UserName"
2- generar un hash con la información personal, de esta manera normalizo un string empleando
	la iterface "CryptoEcoder".
3- generar un avatar con el hash generado, creando una imagen "avatar.png"
	empleado la interfaz "ImageGenerator".
-------------------------------------------------------------------------------------------
Ejemplo de uso:

package main

import (
	"pkg/pkg/avatar"
	"fmt"
)

func main() {
	var username string
	println("Enter your username: ")
	fmt.Scanln(&username)
	avatar.AvatarGenerator().GenerateAndSaveAvatar(avatar.Information{UserName: username})
}
*/

package avatar

import (
	"log"
	"pkg/pkg/avatar/encoder"
	"pkg/pkg/avatar/images"
)

// cryptoEncoder is someone who can encode information in a cryptographic way
type CryptoEncoder interface {
	EncodeInformation(strInformation string) (EncodeInformation []byte, err error)
}

// imageGenerator is someone who can generate an images
type ImageGenerator interface {
	BuildAndSaveImage(EncodeInformation []byte) error
}

// service contains funcionalities related to avatar generation
type Service struct {
	encoder   CryptoEncoder
	generator ImageGenerator
}

// NewService creates a new service with the given encoder and generator
func AvatarGenerator() *Service {
	return &Service{
		encoder:   &encoder.CryptoEncoder{},
		generator: &images.ImageGenerator{},
	}
}

// contain information about the avatar
type Information struct {
	UserName string
}

// GenerateAndSaveAvatar generates and saves an avatar
func (s *Service) GenerateAndSaveAvatar(information Information) error {
	// defer for recovering from panic errors of the encoder
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// encode the information
	encodeBytes, err := s.encoder.EncodeInformation(information.UserName)
	if err != nil {
		return err
	}
	// defer for recovering from panic errors of the generator and save the image
	defer func() {
		if errImage := recover(); errImage != nil {
			log.Println(errImage)
		}
	}()
	// generate and save the image
	errImage := s.generator.BuildAndSaveImage(encodeBytes)
	if errImage != nil {
		return errImage
	}
	println("The avatar has been generated")
	return nil
}
