package proto

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func DialServer(addr string) (*grpc.ClientConn, error) {
	// load certs
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Error().Err(err).Msg("failed to read ca.crt")
		return nil, err
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		err := errors.New("failed to append ca.crt to cert pool")
		log.Error().Err(err).Msg("internal error")
		return nil, err
	}

	tlsCreds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})

	// setup grpc client
	creds := grpc.WithTransportCredentials(tlsCreds)
	conn, err := grpc.Dial(addr, creds)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial grpc server")
		return nil, err
	}

	return conn, nil
}
