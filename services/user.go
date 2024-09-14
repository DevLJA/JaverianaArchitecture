package services

import (
	"CircuitBreaker-LauraJoya/dtos"
	"CircuitBreaker-LauraJoya/utils/circuit_braker"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserService interface {
	GetCompleteInfoClientWithCircuitBreaker(idUser string) (*dtos.Client, error)
	GetCompleteInfoClient(idUser string) (*dtos.Client, error)
	GetBasicInfoUser(idUser string) (*dtos.User, error)
}

type userServiceImp struct {
	client *http.Client
}

func NewUserService() UserService {
	return userServiceImp{
		client: &http.Client{},
	}
}

func (u userServiceImp) GetCompleteInfoClientWithCircuitBreaker(idUser string) (*dtos.Client, error) {
	switch circuit_braker.GetInstance().State {
	case circuit_braker.StateOpen:
		if time.Since(circuit_braker.GetInstance().LastFailureTime) > circuit_braker.GetInstance().Timeout {
			circuit_braker.SetValueState(circuit_braker.StateHalfOpen)
			return nil, errors.New("Soft error: circuit breaker is HALF-OPEN")
		} else {
			return nil, errors.New("Soft error: circuit breaker is OPEN")
		}
	case circuit_braker.StateHalfOpen:
		result, err := u.GetCompleteInfoClient(idUser)
		if err != nil {
			circuit_braker.SetValueState(circuit_braker.StateOpen)
			circuit_braker.SetValueLastFailureTime(time.Now())
			return result, err
		}
		circuit_braker.SetValueState(circuit_braker.StateClosed)
		circuit_braker.SetValueFailureCount(0)
		return result, nil
	case circuit_braker.StateClosed:
		result, err := u.GetCompleteInfoClient(idUser)
		if err != nil {
			circuit_braker.SetValueFailureCount(circuit_braker.GetInstance().FailureCount + 1)
			if circuit_braker.GetInstance().FailureCount >= circuit_braker.GetInstance().FailureThreshold {
				circuit_braker.SetValueState(circuit_braker.StateOpen)
				circuit_braker.SetValueLastFailureTime(time.Now())
				return nil, errors.New("Soft error: circuit breaker is OPEN")
			}
			return result, err
		}
		circuit_braker.SetValueFailureCount(0)
		return result, nil
	}
	return nil, nil
}

func (u userServiceImp) GetCompleteInfoClient(idUser string) (*dtos.Client, error) {
	//Here the users-api service is consumed
	responseEmail, err := u.client.Get(fmt.Sprintf("http://localhost:8080/user/%s", idUser))
	defer responseEmail.Body.Close()
	if err != nil || responseEmail.StatusCode != http.StatusOK {
		return nil, errors.New("users-api was consumed and threw an error")
	}
	email, err := io.ReadAll(responseEmail.Body)

	responseAPI := &struct {
		Data dtos.User `json:"data"`
	}{}

	if err := json.Unmarshal(email, &responseAPI); err != nil {
		return nil, err
	}

	// Create response
	client := dtos.Client{
		User: responseAPI.Data,
		StatusClient: dtos.StatusClient{
			Status: "active",
			Points: 23456543,
		},
	}
	return &client, nil
}

func (u userServiceImp) GetBasicInfoUser(idUser string) (*dtos.User, error) {
	if idUser == "123" {
		return &dtos.User{
			Id:    idUser,
			Name:  "Laura Joya",
			Age:   27,
			Email: "lauracja.123@gmail.com",
		}, nil
	} else {
		return nil, errors.ErrUnsupported
	}
}
