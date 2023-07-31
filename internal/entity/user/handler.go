package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (u *userService) Create(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer func() {
		_ = r.Body.Close()
	}()

	data, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	var usr user
	err = json.Unmarshal(data, &usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	tx := u.conn.Create(&usr)

	fmt.Println(tx.Error)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, _ = w.Write([]byte(id))
}
