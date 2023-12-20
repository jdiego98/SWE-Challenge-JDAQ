package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jdiego98/SWE-Challenge-JDAQ/api/models"
)

// Esta funcion maneja las peticiones de busqueda
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// Se verifica que no falten las credenciales
	if username == "" || password == "" {
		http.Error(w, "Credenciales no proporcionadas", http.StatusUnauthorized)
		return
	}

	// Se construye el body request segun la peticion
	searchBody, err := constructSearchBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Se realiza la peticion al api de zicsearch
	results, err := SearchInZinc(searchBody, username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// constructSearchBody se encarga de construir el query para zincsearch
func constructSearchBody(r *http.Request) ([]byte, error) {
	from := r.URL.Query().Get("From")
	to := r.URL.Query().Get("To")
	subject := r.URL.Query().Get("Subject")

	var must []map[string]interface{}

	if subject != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]string{"Subject": subject},
		})
	}
	if from != "" {
		must = append(must, map[string]interface{}{
			"wildcard": map[string]string{"From": from + "*"},
		})
	}
	if to != "" {
		must = append(must, map[string]interface{}{
			"wildcard": map[string]string{"To": to + "*"},
		})
	}

	searchQuery := map[string]interface{}{
		"search_type": "querystring",
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}

	return json.Marshal(searchQuery)
}

func SearchInZinc(searchBody []byte, username, password string) ([]models.Email, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	url := "http://localhost:4080/es/test_emails/_search"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(searchBody))
	if err != nil {
		return nil, err
	}

	// Se codifican las credenciales y agregan a los headers
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Set("Authorization", "Basic "+encodedCreds)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResults models.ZincSearchResponse
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return nil, err
	}

	var emails []models.Email
	for _, hit := range searchResults.Hits.Hits {

		// bodyString, err := json.Marshal(hit.Source.Body)

		if err != nil {
			return nil, err
		}

		// Se mapean los resultados al modelo Email
		email := models.Email{
			Date:      hit.Source.Date,
			From:      hit.Source.From,
			Subject:   hit.Source.Subject,
			To:        hit.Source.To,
			MessageID: hit.Source.MessageID,
			// Body:      string(bodyString),
			Body: hit.Source.Body,
		}

		emails = append(emails, email)
	}

	return emails, nil
}
