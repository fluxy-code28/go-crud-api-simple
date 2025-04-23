package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// struct adalah kumpulan definisi variabel dan atau fungsi, yang dibungkus sebagai
// tipe data baru dengan nama tertentu. Bisa memiliki berbagai tipe data(variabelnya)
// Dari struct kita bisa membuat variabel baru yang memiliki atribut sesuai skema struct
// Ada yang memanggil variabel tersebut dengan istilah objet atau variabel object.
// Type itu sendiri bisa dibilang alias, jadi User adalah alias dari struct
// inisialisasi object struct bisa dengan var s1 User atau var s1 = User{}
// `json:"id"` disebut tag struct, fungsinya mapping nama field ke JSON Key

var db *sql.DB

// deklarasi variabel db dengan tipe *sql.DB, Tipe *sql.DB adalah pointer ke struct sql.DB. Pointer digunakan agar kita bisa mengakses dan memodifikasi objek database di memori secara efisien tanpa harus membuat salinan dari objek tersebut.

func main() {
	var err error
	db, err = sql.Open("mysql", "alduraimron:MobilePubg122@tcp(127.0.0.1:3306)/go_crud_api")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	// inisialisasi router

	router.HandleFunc("/users", getUsers).Methods("GET") //menghandel fungsi GET ke path /users dan memanggil fungsi getUsers

	log.Fatal(http.ListenAndServe(":8000", router))
	//memulai server di port 8000 dan router sebagai HTTP handler, log.Fatal() akan memberhentikan dan mencetak error apbila ada error
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
