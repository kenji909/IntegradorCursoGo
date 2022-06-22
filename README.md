# IntegradorCursoGo

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
