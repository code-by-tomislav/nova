package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Student

func (b *Bouncer) GETStudent(w http.ResponseWriter, r *http.Request) {
	studentId := chi.URLParam(r, "id")
	student, err := b.m.ReadStudent(studentId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching student %s", studentId)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	} else if student == nil {
		errMsg := fmt.Sprintf("student %s not found", studentId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (b *Bouncer) GETStudents(w http.ResponseWriter, r *http.Request) {
	students, err := b.m.ListStudents()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching students: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

// Member

func (b *Bouncer) GETMember(w http.ResponseWriter, r *http.Request) {
	memberId := chi.URLParam(r, "id")
	member, err := b.m.ReadMember(memberId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching member %s: %v", memberId, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	} else if member == nil {
		errMsg := fmt.Sprintf("member %s not found", memberId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (b *Bouncer) GETMembers(w http.ResponseWriter, r *http.Request) {
	members, err := b.m.ListMember()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching members: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(members)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

// Employee

func (b *Bouncer) GETEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := chi.URLParam(r, "id")
	employee, err := b.m.ReadEmployee(employeeId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching employee %s: %v", employeeId, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	} else if employee == nil {
		errMsg := fmt.Sprintf("employee %s not found", employeeId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (b *Bouncer) GETEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := b.m.ListEmployees()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errMsg := fmt.Sprintf("error fetching employees: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	}

	json, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}
