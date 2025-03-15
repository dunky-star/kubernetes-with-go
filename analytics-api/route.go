package main

func (s *MuxServer) routes() {
	s.gorilla.HandleFunc("/user", s.createUser).Methods("POST")
	s.gorilla.HandleFunc("/users", s.getUsers).Methods("GET")
	s.gorilla.HandleFunc("/user", s.getUser).Methods("GET")
	s.gorilla.HandleFunc("/user/{id}", s.updateUser).Methods("PUT")
	s.gorilla.HandleFunc("/user/{id}", s.deleteUser).Methods("DELETE")
}
