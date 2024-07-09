package api_controller

import (
	"encoding/json"
	"github.com/ternaryinvalid/todo1/internal/app/domain/todo"
	"github.com/ternaryinvalid/todo1/internal/app/domain/user"
	"log"
	"net/http"
)

func (ctr *Controller) SignUP(w http.ResponseWriter, r *http.Request) {
	var request user.TodoUserCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	ctx := r.Context()

	token, err := ctr.service.CreateUser(ctx, request)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		log.Println(err)

		http.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}
}

func (ctr *Controller) CreateTODO(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var task todo.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	task.UserID = userID

	ctx := r.Context()

	id, err := ctr.service.CreateTODO(ctx, task)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error creating task", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]int{"task_id": id})
	if err != nil {
		log.Println(err)

		http.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}
}

func (ctr *Controller) DeleteTODO(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var deletionPayload todo.DeleteTaskRequest

	err := json.NewDecoder(r.Body).Decode(&deletionPayload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	deletionPayload.UserID = userID

	ctx := r.Context()

	err = ctr.service.DeleteTODO(ctx, deletionPayload)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error deleting task", http.StatusInternalServerError)

		return
	}
}

func (ctr *Controller) GetAllTODO(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	ctx := r.Context()

	tasks, err := ctr.service.GetAllTODO(ctx, userID)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error creating task", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string][]todo.Task{"tasks": tasks})
	if err != nil {
		log.Println(err)

		http.Error(w, "Error getting task", http.StatusInternalServerError)

		return
	}
}

func (ctr *Controller) Done(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var req todo.DoneTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	req.UserID = userID

	ctx := r.Context()

	err = ctr.service.Done(ctx, req)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error done task", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
}
