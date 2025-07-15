package app_test

import (
	"bytes"
	"errors"
	"log"
	"sauce-service/src/app"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Mock ConfigLoader qui garde en mémoire si LoadEnv a été appelée
type mockConfigLoader struct {
	loaded bool
}

func (m *mockConfigLoader) LoadEnv() {
	m.loaded = true
}

// Mock DBConnection avec contrôle sur Close() et GetDB()
type mockDBConnection struct {
	closed      bool
	closeErr    error
	returnedDB  *gorm.DB
}

func (m *mockDBConnection) Close() error {
	m.closed = true
	return m.closeErr
}

func (m *mockDBConnection) GetDB() *gorm.DB {
	return m.returnedDB
}

// Mock DBConnector qui simule la connexion
type mockDBConnector struct {
	connectCalled bool
	returnError   bool
	closeErr     error
}

func (m *mockDBConnector) Connect() (app.DBConnection, error) {
	m.connectCalled = true
	if m.returnError {
		return nil, errors.New("connection failed")
	}
	return &mockDBConnection{closeErr: m.closeErr}, nil
}

// Mock Router qui crée un *gin.Engine et garde trace de l'appel
type mockRouter struct {
	setupCalled bool
	engine      *gin.Engine
}

func (m *mockRouter) Setup(db *gorm.DB) *gin.Engine {
	m.setupCalled = true
	if m.engine == nil {
		m.engine = gin.New()
	}
	return m.engine
}

// Mock Server qui simule le démarrage
type mockServer struct {
	startCalled bool
	returnError bool
	engine      *gin.Engine
}

func (m *mockServer) Start(e *gin.Engine) error {
	m.startCalled = true
	m.engine = e
	if m.returnError {
		return errors.New("server failed")
	}
	return nil
}

// Test succès complet
func TestRunApp_Success(t *testing.T) {
	config := &mockConfigLoader{}
	dbConnector := &mockDBConnector{}
	router := &mockRouter{}
	server := &mockServer{}
	err := app.RunApp(config, dbConnector, router, server)
	assert.NoError(t, err, "RunApp should not return error")
	assert.True(t, config.loaded, "LoadEnv should be called")
	assert.True(t, dbConnector.connectCalled, "Connect should be called")
	assert.True(t, router.setupCalled, "Setup should be called")
	assert.True(t, server.startCalled, "Start should be called")
	assert.Equal(t, router.engine, server.engine, "Server should receive the engine from Router")
}

// Test échec de connexion à la DB
func TestRunApp_DBConnectionFails(t *testing.T) {
	config := &mockConfigLoader{}
	dbConnector := &mockDBConnector{returnError: true}
	router := &mockRouter{}
	server := &mockServer{}
	err := app.RunApp(config, dbConnector, router, server)
	assert.Error(t, err)
	assert.EqualError(t, err, "connection failed")
	assert.True(t, config.loaded, "LoadEnv should be called")
	assert.True(t, dbConnector.connectCalled, "Connect should be called")
	assert.False(t, router.setupCalled, "Setup should NOT be called on DB error")
	assert.False(t, server.startCalled, "Start should NOT be called on DB error")
}

// Test échec au démarrage du serveur
func TestRunApp_ServerStartFails(t *testing.T) {
	config := &mockConfigLoader{}
	dbConnector := &mockDBConnector{}
	router := &mockRouter{}
	server := &mockServer{returnError: true}
	err := app.RunApp(config, dbConnector, router, server)
	assert.Error(t, err)
	assert.EqualError(t, err, "server failed")
	assert.True(t, config.loaded, "LoadEnv should be called")
	assert.True(t, dbConnector.connectCalled, "Connect should be called")
	assert.True(t, router.setupCalled, "Setup should be called")
	assert.True(t, server.startCalled, "Start should be called")
}

// Test que la fermeture DB est appelée et que l'erreur est loggée
func TestRunApp_DBCloseErrorIsLogged(t *testing.T) {
	config := &mockConfigLoader{}
	dbConnector := &mockDBConnector{closeErr: errors.New("close failed")}
	router := &mockRouter{}
	server := &mockServer{}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)
	err := app.RunApp(config, dbConnector, router, server)
	assert.NoError(t, err, "RunApp should not return error even if Close fails")
	logged := buf.String()
	assert.Contains(t, logged, "Erreur fermeture DB : close failed", "Error on DB close should be logged")
}
