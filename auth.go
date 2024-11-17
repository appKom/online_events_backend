package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func authenticateUser(userID string, token string, db *pgxpool.Pool) (*User, error) {
	req, err := http.NewRequest("GET", "https://old.online.ntnu.no/api/v1/profile/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch profile: status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profileData map[string]interface{}
	if err := json.Unmarshal(body, &profileData); err != nil {
		return nil, err
	}

	profileUserID, ok := profileData["id"].(string)
	if !ok || profileUserID != userID {
		return nil, fmt.Errorf("user ID mismatch or missing in profile data")
	}

	var user User
	err = db.QueryRow(context.Background(), `
        SELECT id, grade, email, image, interest_group FROM users WHERE id = $1
    `, userID).Scan(&user.ID, &user.Grade, &user.Email, &user.Image, &user.InterestGroup)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
