package user

import (
	"context"

	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

// TODO: GetUsers(w http.ResponseWriter, r *http.Request)
// TODO: CreateUser(w http.ResponseWriter, r *http.Request)
// TODO: GetUser(w http.ResponseWriter, r *http.Request)
// TODO: DeleteUser(w http.ResponseWriter, r *http.Request)
// TODO: UpdateUser(w http.ResponseWriter, r *http.Request)

type Service interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	// GetByLogin(ctx context.Context, username, email string) (model.User, error)
	// GetById(ctx context.Context, userId uuid.UUID) (model.User, error)
}

type handler struct {
	service Service
	logger  *zap.Logger
}

func NewHandler(service Service, logger *zap.Logger) *handler {
	return &handler{service: service, logger: logger}
}

// func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")
// 	userId, err := uuid.Parse(id)
// 	if err != nil {
// 		h.logger.Error("invalid id", zap.Error(err))
// 		http.Error(w, "invalid id", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.service.GetById(r.Context(), userId)
// 	if err != nil {
// 		h.logger.Error("failed get user", zap.Error(err))
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)

// 	if err = json.NewEncoder(w).Encode(dto.ToUserResponse(&user)); err != nil {
// 		h.logger.Error("failed create response", zap.String("email", user.Email), zap.Error(err))
// 	}
// }

// func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var input dto.UserRequest

// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		h.logger.Error("failed decode request", zap.Error(err))
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.service.CreateUser(r.Context(), input.Username, input.Email, input.Password)
// 	if err != nil {
// 		h.logger.Error("failed create user", zap.Error(err))
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	if err = json.NewEncoder(w).Encode(dto.ToUserResponse(&user)); err != nil {
// 		h.logger.Error("failed create response", zap.String("email", input.Email), zap.Error(err))
// 	}
// }
