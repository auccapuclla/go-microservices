package resthandlers

import (
	"encoding/json"
	"io"
	"microservices/api/restutil"
	"microservices/pb"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{
		authSvcClient: authSvcClient,
	}
}

func (h *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	user.Created = time.Now().Unix()
	user.Updated = user.Created
	user.Id = primitive.NewObjectID().Hex()
	resp, err := h.authSvcClient.SignUp(r.Context(), user)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandlers) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	vars := mux.Vars(r)
	user.Id = vars["id"]

	user.Updated = time.Now().Unix()
	respUser, err := h.authSvcClient.UpdateUser(r.Context(), user)
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, respUser)
}

func (h *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	respUser, err := h.authSvcClient.GetUser(r.Context(), &pb.GetUserRequest{Id: id})
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, respUser)
}

func (h *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := h.authSvcClient.DeleteUser(r.Context(), &pb.GetUserRequest{Id: id})
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, nil)
}

func (h *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := h.authSvcClient.ListUsers(r.Context(), &pb.ListUsersRequest{})
	if err != nil {
		restutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	users := make([]*pb.User, 0)
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		users = append(users, user)
	}
	restutil.WriteAsJson(w, http.StatusOK, users)
}
