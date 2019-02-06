package subject

import (
    "encoding/json"
    "net/http"
    "strconv"
    "team-academy/component"

    "github.com/gorilla/mux"
)

func CreateSubjectController(w http.ResponseWriter, r *http.Request) {
    var subject Subject
    err := json.NewDecoder(r.Body).Decode(&subject)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    err = CreateSubject(component.App.DB, subject)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    newSubject, err := GetSubjectByName(component.App.DB, subject.Name)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    subjectInfo, err := json.Marshal(newSubject)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Write([]byte(subjectInfo))
}

func FetchAllSubjectsController(w http.ResponseWriter, r *http.Request) {
    subjects, err := GetAllSubjects(component.App.DB)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    subjectsInfo, err := json.Marshal(subjects)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Write(subjectsInfo)
}

func FetchSubjectByIDController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["ID"]

    subjectID, err := strconv.Atoi(id)
    if err != nil {
        return
    }

    subject, err := GetSubjectByID(component.App.DB, subjectID)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    subjectInfo, err := json.Marshal(subject)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Write(subjectInfo)
}