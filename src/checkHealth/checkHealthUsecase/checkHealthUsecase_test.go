package checkHealthUsecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCheckHealthRepository adalah mock untuk CheckHealthRepository
type MockCheckHealthRepository struct {
	mock.Mock
}

func (m *MockCheckHealthRepository) RetrieveVersion() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestGetVersion(t *testing.T) {
	// Membuat instance MockCheckHealthRepository
	mockRepo := new(MockCheckHealthRepository)

	// Memberikan behavior yang diharapkan saat fungsi RetrieveVersion dipanggil
	expectedVersion := "1.0.0"
	mockRepo.On("RetrieveVersion").Return(expectedVersion, nil)

	// Membuat instance checkHealthUsecase dengan mockRepo
	usecase := NewCheckHealthUsecase(mockRepo)

	// Memanggil fungsi GetVersion pada usecase
	version, err := usecase.GetVersion()

	// Memverifikasi bahwa fungsi RetrieveVersion pada mockRepo dipanggil sekali
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error yang diharapkan
	assert.NoError(t, err)

	// Memastikan versi yang dikembalikan sesuai dengan yang diharapkan
	assert.Equal(t, expectedVersion, version)
}

func TestGetVersionError(t *testing.T) {
	// Membuat instance MockCheckHealthRepository
	mockRepo := new(MockCheckHealthRepository)

	// Memberikan behavior yang diharapkan saat fungsi RetrieveVersion dipanggil dengan error
	expectedError := errors.New("mock error")
	mockRepo.On("RetrieveVersion").Return("", expectedError)

	// Membuat instance checkHealthUsecase dengan mockRepo
	usecase := NewCheckHealthUsecase(mockRepo)

	// Memanggil fungsi GetVersion pada usecase
	version, err := usecase.GetVersion()

	// Memverifikasi bahwa fungsi RetrieveVersion pada mockRepo dipanggil sekali
	mockRepo.AssertExpectations(t)

	// Memastikan error yang dikembalikan sesuai dengan yang diharapkan
	assert.EqualError(t, err, expectedError.Error())

	// Memastikan versi yang dikembalikan kosong
	assert.Empty(t, version)
}
