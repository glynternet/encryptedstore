package http

import (
	"encoding/json"
	"github.com/glynternet/encryptedstore/pkg/storage"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type server struct {
	store storage.Encrypted
	logger *log.Logger
}

type storeRequestBody struct {
	Id []byte `json:"id"`
	Payload []byte `json:"payload"`
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

	k, err := s.store.Store(sr.Id, sr.Payload)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "storing store request payload"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(k)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "writing key response"))
	}
}

type retrieveRequestBody struct {
	Id []byte `json:"id"`
	Key []byte `json:"key"`
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

	P, err := s.store.Retrieve(rr.Id, rr.Key)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "storing retrieve request payload"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(P)
	if err != nil {
		s.logger.Print(errors.Wrap(err, "writing payload response"))
	}
}