package validate

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/futurenda/google-auth-id-token-verifier"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	command        = kingpin.Command("validate", "Start token validator (reads stdin)")
	identifier     = command.Flag("id", "Perforce identifier (user or email)").Required().String()
	clientIP       = command.Flag("client-ip", "Perforce client IP").Short('c').IP()
	googleClientID = kingpin.Flag("google-client-id", "Google Client ID").Required().String()
)

func FullCommand() string {
	return command.FullCommand()
}

func RunValidate() error {
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buf.Write(scanner.Bytes())
		buf.WriteByte('\n')
		fields := bytes.Fields(scanner.Bytes())
		if len(fields) != 1 {
			continue
		}

		token := string(fields[0])
		log.Debugf("ID: %s", *identifier)
		log.Debugf("Token: %s", token)

		verifier := googleAuthIDTokenVerifier.Verifier{}
		err := verifier.VerifyIDToken(token, []string{
			*googleClientID,
		})
		if err != nil {
			log.Debugf("Token verified failed: %v", err)
			return err
		}

		claimSet, err := googleAuthIDTokenVerifier.Decode(token)
		if err != nil {
			return err
		}

		if claimSet.Email != *identifier {
			return fmt.Errorf("User did not match claimSet Email: (%s != %s)", *identifier, claimSet.Email)
		}
		log.Debugf("User authenicated: %s", claimSet.Email)
		return nil
	}
	if *clientIP != nil {
		log.Debugf("client ip: %s", clientIP.String())
		if clientIP.Equal(net.IPv4(127, 0, 0, 1)) {
			return nil
		}
	}
	return fmt.Errorf("Invalid input: %s", string(buf.Bytes()))
}
