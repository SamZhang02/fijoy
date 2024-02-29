package handler

import (
	"database/sql"
	"encoding/json"
	"fijoy/internal/gen/postgres/railway/public/model"
	"fmt"
	"net/http"
	"time"

	. "fijoy/internal/gen/postgres/railway/public/table"

	. "github.com/go-jet/jet/v2/postgres"

	"github.com/go-playground/validator"
	"github.com/nrednav/cuid2"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type workspaceHandler struct {
	tokenAuth *jwtauth.JWTAuth
	db        *sql.DB
}

func NewWorkspaceHandler(r *chi.Mux, tokenAuth *jwtauth.JWTAuth, db *sql.DB) {
	handler := workspaceHandler{tokenAuth, db}

	r.Route("/v1/workspaces", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/", handler.createWorkspace)
		r.Get("/", handler.getWorkspaces)
		r.Get("/{workspaceId}", handler.getWorkspace)
	})
}

type createFijoyWorkspace struct {
	Name      string `json:"Name" validate:"required,min=1"`
	Namespace string `json:"Namespace" validate:"required,min=1"`
}

func (wh workspaceHandler) createWorkspace(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["user_id"].(string)

	var createWorkspace createFijoyWorkspace

	err := json.NewDecoder(r.Body).Decode(&createWorkspace)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(createWorkspace)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := wh.db.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, "Failed to begin transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	workspace := model.FijoyWorkspace{
		Name:      createWorkspace.Name,
		Namespace: createWorkspace.Namespace,
		ID:        "workspace_" + cuid2.Generate(),
		CreatedAt: time.Now(),
	}

	insertWorkspaceStmt := FijoyWorkspace.INSERT(FijoyWorkspace.AllColumns).MODEL(workspace)

	_, err = insertWorkspaceStmt.ExecContext(r.Context(), tx)
	if err != nil {
		http.Error(w, "Failed to create workspace: "+err.Error(), http.StatusInternalServerError)
		return
	}

	workspaceUser := model.FijoyWorkspaceUser{
		WorkspaceID: workspace.ID,
		UserID:      userId,
		Role:        "owner",
	}

	insertWorkspaceUserStmt := FijoyWorkspaceUser.INSERT(FijoyWorkspaceUser.AllColumns).MODEL(workspaceUser)

	_, err = insertWorkspaceUserStmt.ExecContext(r.Context(), tx)
	if err != nil {
		http.Error(w, "Failed to create workspace user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workspace)
}

func (wh workspaceHandler) getWorkspaces(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["user_id"].(string)

	stmt := SELECT(FijoyWorkspace.AllColumns).
		FROM(FijoyWorkspaceUser.
			INNER_JOIN(FijoyUser, FijoyWorkspaceUser.UserID.EQ(FijoyUser.ID)).
			INNER_JOIN(FijoyWorkspace, FijoyWorkspaceUser.WorkspaceID.EQ(FijoyWorkspace.ID))).WHERE(FijoyUser.ID.EQ(String(userId)))

	dest := []*model.FijoyWorkspace{}

	err := stmt.QueryContext(r.Context(), wh.db, &dest)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to get workspaces: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dest)
}

func (wh workspaceHandler) getWorkspace(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["user_id"].(string)
	workspaceId := chi.URLParam(r, "workspaceId")

	isNamespace := r.URL.Query().Get("namespace")

	var stmt SelectStatement

	if isNamespace == "true" {
		stmt = SELECT(FijoyWorkspace.AllColumns).
			FROM(FijoyWorkspaceUser.
				INNER_JOIN(FijoyUser, FijoyWorkspaceUser.UserID.EQ(FijoyUser.ID)).
				INNER_JOIN(FijoyWorkspace, FijoyWorkspaceUser.WorkspaceID.EQ(FijoyWorkspace.ID))).
			WHERE(AND(FijoyUser.ID.EQ(String(userId)), FijoyWorkspace.Namespace.EQ(String(workspaceId))))
	} else {
		stmt = SELECT(FijoyWorkspace.AllColumns).
			FROM(FijoyWorkspaceUser.
				INNER_JOIN(FijoyUser, FijoyWorkspaceUser.UserID.EQ(FijoyUser.ID)).
				INNER_JOIN(FijoyWorkspace, FijoyWorkspaceUser.WorkspaceID.EQ(FijoyWorkspace.ID))).
			WHERE(AND(FijoyUser.ID.EQ(String(userId)), FijoyWorkspace.ID.EQ(String(workspaceId))))
	}

	var dest struct {
		model.FijoyWorkspace
	}

	err := stmt.QueryContext(r.Context(), wh.db, &dest)
	if err != nil {
		http.Error(w, "Failed to get workspace: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dest)
}
