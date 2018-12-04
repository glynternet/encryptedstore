package http

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/glynternet/encryptedstore/pkg/storage"
	"github.com/pkg/errors"
)

type server struct {
	store  storage.Encrypted
	logger *log.Logger
}

// NewEncrypterServeMux creates a ServeMux that uses an encrypted storage backend
func NewEncrypterServeMux(logger *log.Logger) *http.ServeMux {
	s := server{
		logger: logger,
	}
	m := http.NewServeMux()
	m.Handle("/store", http.HandlerFunc(s.Store))
	m.Handle("/retrieve", http.HandlerFunc(s.Retrieve))
	return m
}

type storeRequestBody struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

func (s *server) Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var sr storeRequestBody
	err := json.NewDecoder(r.Body).Decode(&sr)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "decoding store request body"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	k, err := s.store.Store([]byte(sr.ID), []byte(sr.Payload))
	if err != nil {
		s.logger.Print(errors.Wrap(err, "storing store request payload"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(k)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(encoded))
	if err != nil {
		s.logger.Print(errors.Wrap(err, "writing key response"))
	}
}

type retrieveRequestBody struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

func (s *server) Retrieve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var rr retrieveRequestBody
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "decoding retrieve request body"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(rr.Key))
	if err != nil {
		s.logger.Print(errors.Wrap(err, "base64 decoding retrieve key"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	p, err := s.store.Retrieve([]byte(rr.ID), decoded)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "retrieving retrieve request payload"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(p)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "writing payload response"))
	}
}
