package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guicazaroto/learning-go/config"
	"github.com/guicazaroto/learning-go/model"
	"github.com/guicazaroto/learning-go/schemas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPing(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", PingHandler)

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCleanerHandler(t *testing.T) {
	t.Run("GetCleanerHandler", func(t *testing.T) {
		t.Run("Return Cleaners", func(t *testing.T) {
			resultCleanerMock := []schemas.Cleaner{
				{
					Id:     1,
					UserId: 1,
					UserInfos: schemas.User{
						Name:      "teste",
						Email:     "teste@teste.com",
						Password:  "cbcjdsb",
						Role:      "cleaner",
						Active:    true,
						ImagemUrl: "ima_url",
					},
					Telefone:       "1254785",
					CPF:            "115165516",
					DataNascimento: "2024-05-15T19:21:30+03:00",
					Cep:            "1248522",
					Logradouro:     "rua teste",
					Numero:         "10",
					Cidade:         "Teste",
					Uf:             "tt",
					Descricao:      "descricao teste",
				},
				{
					Id:     2,
					UserId: 2,
					UserInfos: schemas.User{
						Name:      "teste 2",
						Email:     "teste2@teste.com",
						Password:  "cbcjdsb",
						Role:      "cleaner",
						Active:    true,
						ImagemUrl: "ima_url",
					},
					Telefone:       "1254785",
					CPF:            "145878523",
					DataNascimento: "2024-05-15T19:21:30+03:00",
					Cep:            "1248522",
					Logradouro:     "rua teste 2",
					Numero:         "10",
					Cidade:         "Teste 2",
					Uf:             "tt",
					Descricao:      "descricao teste 2",
				},
			}

			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/search", GetCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleaners", mock.Anything).Return(resultCleanerMock)

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/search", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Data []model.CleanerResponse
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)
			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Len(t, response.Data, 2)
		})
		t.Run("Dont return Cleaners", func(t *testing.T) {
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/search", GetCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleaners", mock.Anything).Return([]schemas.Cleaner{})

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/search", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNotFound, resp.Code)
		})
	})

	t.Run("GetCleanerByIdHandler", func(t *testing.T) {
		t.Run("Return Cleaner by id", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "115165516",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			cleanerID := "1"
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/:id", GetCleanerByIdHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", cleanerID).Return(resultCleanerMock)

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/"+cleanerID, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Data model.CleanerResponse
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)
			assert.Equal(t, http.StatusOK, resp.Code)
			assert.NotNil(t, response.Data)
			assert.Equal(t, resultCleanerMock.Cidade, response.Data.Cidade)
			assert.Equal(t, resultCleanerMock.Descricao, response.Data.Descricao)
			assert.Equal(t, resultCleanerMock.UserInfos.Email, response.Data.Email)
			assert.Equal(t, resultCleanerMock.Id, response.Data.Id)
			assert.Equal(t, resultCleanerMock.UserInfos.ImagemUrl, response.Data.ImagemUrl)
			assert.Equal(t, resultCleanerMock.UserInfos.Name, response.Data.Name)
			assert.Equal(t, resultCleanerMock.Telefone, response.Data.Telefone)
			assert.Equal(t, resultCleanerMock.Uf, response.Data.Uf)
		})

		t.Run("Dont return Cleaner by id", func(t *testing.T) {
			cleanerID := "1"
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/:id", GetCleanerByIdHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", cleanerID).Return(nil)

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/"+cleanerID, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNotFound, resp.Code)
		})
	})

	t.Run("GetCleanerMeByIdHandler", func(t *testing.T) {
		t.Run("Return cleaner/me", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "115165516",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/me", GetCleanerMeByIdHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(resultCleanerMock)

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/me", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Data model.CleanerMeResponse
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)
			assert.Equal(t, http.StatusOK, resp.Code)
			assert.NotNil(t, response.Data)
			assert.Equal(t, resultCleanerMock.CPF, response.Data.CPF)
			assert.Equal(t, resultCleanerMock.Cep, response.Data.Cep)
			assert.Equal(t, resultCleanerMock.Cidade, response.Data.Cidade)
			assert.Equal(t, resultCleanerMock.DataNascimento, response.Data.DataNascimento)
			assert.Equal(t, resultCleanerMock.Descricao, response.Data.Descricao)
			assert.Equal(t, resultCleanerMock.UserInfos.Email, response.Data.Email)
			assert.Equal(t, resultCleanerMock.Id, response.Data.Id)
			assert.Equal(t, resultCleanerMock.UserInfos.ImagemUrl, response.Data.ImagemUrl)
			assert.Equal(t, resultCleanerMock.Logradouro, response.Data.Logradouro)
			assert.Equal(t, resultCleanerMock.UserInfos.Name, response.Data.Name)
			assert.Equal(t, resultCleanerMock.Numero, response.Data.Numero)
			assert.Equal(t, resultCleanerMock.Telefone, response.Data.Telefone)
			assert.Equal(t, resultCleanerMock.Uf, response.Data.Uf)

		})

		t.Run("Dont return cleaner/me", func(t *testing.T) {
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.GET("/cleaner/me", GetCleanerMeByIdHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(nil)

			req, _ := http.NewRequest(http.MethodGet, "/cleaner/me", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNotFound, resp.Code)

		})
	})

	t.Run("CreateCleanerHandler", func(t *testing.T) {
		logger = config.GetLogger("main")
		t.Run("Create a cleaner", func(t *testing.T) {
			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner", CreateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndCpf", request.Email, request.CPF).Return(int64(0))
			cleanerRepositoryMock.On("CreateCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Data model.CleanerResponse
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)
			assert.Equal(t, http.StatusCreated, resp.Code)
			assert.NotNil(t, response.Data)
			assert.Equal(t, request.Cidade, response.Data.Cidade)
			assert.Equal(t, request.Descricao, response.Data.Descricao)
			assert.Equal(t, request.Email, response.Data.Email)
			assert.Equal(t, request.Name, response.Data.Name)
			assert.Equal(t, request.Telefone, response.Data.Telefone)
			assert.Equal(t, request.Uf, response.Data.Uf)
		})

		t.Run("Request Body error", func(t *testing.T) {
			request := `{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}`
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner", CreateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndCpf", "teste@teste.com", "1458574596").Return(int64(0))
			cleanerRepositoryMock.On("CreateCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusBadRequest, resp.Code)
		})

		t.Run("Request body invalid", func(t *testing.T) {
			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner", CreateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndCpf", request.Email, request.CPF).Return(int64(0))
			cleanerRepositoryMock.On("CreateCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusBadRequest, resp.Code)
		})

		t.Run("Cleaner exists", func(t *testing.T) {
			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner", CreateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndCpf", request.Email, request.CPF).Return(int64(1))
			cleanerRepositoryMock.On("CreateCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusConflict, resp.Code)
		})

		t.Run("Create a cleaner error", func(t *testing.T) {
			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner", CreateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndCpf", request.Email, request.CPF).Return(int64(0))
			cleanerRepositoryMock.On("CreateCleaner", mock.Anything).Return(errors.New("erro ao cadastrar"))

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusInternalServerError, resp.Code)
		})
	})

	t.Run("UpdateCleanerHandler", func(t *testing.T) {
		t.Run("Update a cleaner", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.PUT("/cleaner", UpdateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(resultCleanerMock)
			cleanerRepositoryMock.On("SaveCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPut, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Data model.CleanerResponse
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Equal(t, request.Cidade, response.Data.Cidade)
			assert.Equal(t, request.Descricao, response.Data.Descricao)
			assert.Equal(t, request.Email, response.Data.Email)
			assert.Equal(t, request.Name, response.Data.Name)
			assert.Equal(t, request.Telefone, response.Data.Telefone)
			assert.Equal(t, request.Uf, response.Data.Uf)
		})

		t.Run("Cleaner not exists", func(t *testing.T) {
			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.PUT("/cleaner", UpdateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(nil)
			cleanerRepositoryMock.On("SaveCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPut, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNotFound, resp.Code)
		})

		t.Run("Request body erro", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			request := `{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}`
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.PUT("/cleaner", UpdateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(resultCleanerMock)
			cleanerRepositoryMock.On("SaveCleaner", mock.Anything).Return(nil)

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPut, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusBadRequest, resp.Code)
		})

		t.Run("Save cleaner erro", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			request := model.CleanerRequest{
				Telefone:       "1655167",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1254785",
				Logradouro:     "rua teste",
				Numero:         "125",
				Cidade:         "teste",
				Uf:             "tt",
				Descricao:      "teste descricao",
				Name:           "teste",
				Email:          "teste@teste.com",
				Password:       "zasdsd",
				Active:         true,
			}
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.PUT("/cleaner", UpdateCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerById", mock.Anything).Return(resultCleanerMock)
			cleanerRepositoryMock.On("SaveCleaner", mock.Anything).Return(errors.New("erro ao salvar"))

			body, _ := json.Marshal(request)
			req, _ := http.NewRequest(http.MethodPut, "/cleaner", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusInternalServerError, resp.Code)
		})
	})

	t.Run("DeleteCleanerHandler", func(t *testing.T) {
		t.Run("Delete a cleaner", func(t *testing.T) {
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.DELETE("/cleaner", DeleteCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("DeleteCleaner", mock.Anything).Return(nil)

			req, _ := http.NewRequest(http.MethodDelete, "/cleaner", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusOK, resp.Code)
		})

		t.Run("Delete clenaer erro", func(t *testing.T) {
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.DELETE("/cleaner", DeleteCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("DeleteCleaner", mock.Anything).Return(errors.New("erro ao deletar"))

			req, _ := http.NewRequest(http.MethodDelete, "/cleaner", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusInternalServerError, resp.Code)
		})
	})

	t.Run("LoginHandler", func(t *testing.T) {
		t.Run("Login", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			var creds struct {
				Email    string
				Password string
			}
			creds.Email = "teste@teste.com"
			creds.Password = "scdscjdnck"
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner/login", LoginCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndPassword", creds.Email, creds.Password).Return(resultCleanerMock)

			body, _ := json.Marshal(creds)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner/login", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			type Response struct {
				Token string
			}

			response := Response{}
			json.Unmarshal([]byte(resp.Body.String()), &response)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.NotNil(t, response.Token)
		})

		t.Run("Request invalid", func(t *testing.T) {
			resultCleanerMock := &schemas.Cleaner{
				Id:     1,
				UserId: 1,
				UserInfos: schemas.User{
					Name:      "teste",
					Email:     "teste@teste.com",
					Password:  "cbcjdsb",
					Role:      "cleaner",
					Active:    true,
					ImagemUrl: "ima_url",
				},
				Telefone:       "1254785",
				CPF:            "1458574596",
				DataNascimento: "2024-05-15T19:21:30+03:00",
				Cep:            "1248522",
				Logradouro:     "rua teste",
				Numero:         "10",
				Cidade:         "Teste",
				Uf:             "tt",
				Descricao:      "descricao teste",
			}

			creds := `{
				Email    string
				Password string
			}`
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner/login", LoginCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndPassword", mock.Anything, mock.Anything).Return(resultCleanerMock)

			body, _ := json.Marshal(creds)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner/login", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusBadRequest, resp.Code)
		})

		t.Run("Invalid creds", func(t *testing.T) {
			var creds struct {
				Email    string
				Password string
			}
			creds.Email = "teste@teste.com"
			creds.Password = "scdscjdnck"
			router := gin.Default()
			cleanerRepositoryMock := new(mockCleanerRepository)
			router.POST("/cleaner/login", LoginCleanerHandler(cleanerRepositoryMock))

			cleanerRepositoryMock.On("GetCleanerByEmailAndPassword", creds.Email, creds.Password).Return(nil)

			body, _ := json.Marshal(creds)
			req, _ := http.NewRequest(http.MethodPost, "/cleaner/login", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusUnauthorized, resp.Code)
		})
	})
}
