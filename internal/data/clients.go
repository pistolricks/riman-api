package data

import (
	"context"

	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pistolricks/riman-api/internal/validator"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateUsername = errors.New("duplicate clientname")
)

type Client struct {
	ID                     int64          `json:"id"`
	CreatedAt              time.Time      `json:"created_at"`
	FirstName              string         `json:"first_name"`
	MiddleName             string         `json:"middle_name,omitempty"`
	LastName               string         `json:"last_name"`
	Suffix                 string         `json:"suffix,omitempty"`
	Email                  string         `json:"email"`
	Mobile                 string         `json:"mobile,omitempty"`
	Username               string         `json:"username"`
	RimanUserId            int64          `json:"riman_user_id"`
	Status                 bool           `json:"status"`
	OrganizationType       string         `json:"organization_type,omitempty"`
	SignupDate             string         `json:"signup_date"`
	AnniversaryDate        string         `json:"anniversary_date"`
	AccountType            string         `json:"account_type"`
	SponsorUsername        string         `json:"sponsor_username,omitempty"`
	MemberId               string         `json:"member_id"`
	Rank                   string         `json:"rank,omitempty"`
	EnrollmentDate         string         `json:"enrollment_date,omitempty"`
	PersonalOrdersVolume   float64        `json:"personal_orders_volume,omitempty"`
	PersonalClientsVolume  float64        `json:"personal_clients_volume,omitempty"`
	TotalPersonalVolume    float64        `json:"total_personal_volume,omitempty"`
	CurrentMonthSp         float64        `json:"current_month_sp,omitempty"`
	CurrentMonthBp         float64        `json:"current_month_bp,omitempty"`
	LastOrderDate          string         `json:"last_order_date,omitempty"`
	LastOrderId            int64          `json:"last_order_id,omitempty"`
	LastOrderSp            float64        `json:"last_order_sp,omitempty"`
	LastOrderBp            float64        `json:"last_order_bp,omitempty"`
	LifetimeSpend          float64        `json:"lifetime_spend,omitempty"`
	MostRecent12MonthSpend float64        `json:"most_recent_12_month_spend,omitempty"`
	Data                   any            `json:"data,omitempty"`
	PasswordHash           string         `json:"password_hash"`
	Token                  string         `json:"token,omitempty"`
	Password               secretPassword `json:"-"`
}

type secretPassword struct {
	plaintext *string
	hash      []byte
}

func (p *secretPassword) SetHash(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *secretPassword) MatchesPassword(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateUsername(v *validator.Validator, username string) {
	v.Check(username != "", "username", "must be provided")
	// v.Check(validator.Matches(username, validator.EmailRX), "username", "must be a valid email address")
}

func ValidateClient(v *validator.Validator, client *Client) {
	v.Check(client.FirstName != "", "name", "must be provided")
	v.Check(len(client.FirstName) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(client.LastName != "", "name", "must be provided")
	v.Check(len(client.LastName) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, client.Email)

	if client.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *client.Password.plaintext)
	}

	if client.Password.hash == nil {
		panic("missing password hash for client")
	}
}

type ClientModel struct {
	DB *sql.DB
}

func (m ClientModel) Insert(client *Client) error {
	query := `
        INSERT INTO clients (first_name, middle_name, last_name, suffix, email, mobile, username, riman_user_id, status, organization_type, signup_date, anniversary_date, account_type, sponsor_username, member_id, rank, enrollment_date, personal_orders_volume, personal_clients_volume, total_personal_volume, current_month_sp, current_month_bp, last_order_date, last_order_id, last_order_sp, last_order_bp, lifetime_spend, most_recent_12_month_spend, data, token ) 
        VALUES ( 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33 )
        RETURNING id, created_at`

	args := []any{client.FirstName, client.MiddleName, client.LastName, client.Suffix, client.Email, client.Mobile, client.Username, client.RimanUserId, client.Status, client.OrganizationType, client.SignupDate, client.AnniversaryDate, client.AccountType, client.SponsorUsername, client.MemberId, client.Rank, client.EnrollmentDate, client.PersonalOrdersVolume, client.PersonalClientsVolume, client.TotalPersonalVolume, client.CurrentMonthSp, client.CurrentMonthBp, client.LastOrderDate, client.LastOrderId, client.LastOrderSp, client.LastOrderBp, client.LifetimeSpend, client.MostRecent12MonthSpend, client.Data}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&client.ID, &client.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "clients_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m ClientModel) GetByUsername(username string) (*Client, error) {
	query := `
        SELECT id, created_at, first_name, middle_name, last_name, suffix, email, mobile, email, riman_user_id, status, organization_type, signup_date, anniversary_date, account_type, sponsor_username, member_id, rank, enrollment_date, personal_orders_volume, personal_clients_volume, total_personal_volume, current_month_sp, current_month_bp, last_order_date, last_order_id, last_order_sp, last_order_bp, lifetime_spend, most_recent_12_month_spend, data, token
        FROM clients
        WHERE username = $1`

	var client Client

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, username).Scan(&client.ID, &client.CreatedAt, &client.FirstName, &client.MiddleName, &client.LastName, &client.Suffix, &client.Email, &client.Mobile, &client.Username, &client.RimanUserId, &client.Status, &client.OrganizationType, &client.SignupDate, &client.AnniversaryDate, &client.AccountType, &client.SponsorUsername, &client.MemberId, &client.Rank, &client.EnrollmentDate, &client.PersonalOrdersVolume, &client.PersonalClientsVolume, &client.TotalPersonalVolume, &client.CurrentMonthSp, &client.CurrentMonthBp, &client.LastOrderDate, &client.LastOrderId, &client.LastOrderSp, &client.LastOrderBp, &client.LifetimeSpend, &client.MostRecent12MonthSpend, &client.Data, &client.Token)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &client, nil
}

func (m ClientModel) GetByEmail(email string) (*Client, error) {
	query := `
        SELECT id, created_at, first_name, middle_name, last_name, suffix, email, mobile, username, riman_user_id, status, organization_type, signup_date, anniversary_date, account_type, sponsor_username, member_id, rank, enrollment_date, personal_orders_volume, personal_clients_volume, total_personal_volume, current_month_sp, current_month_bp, last_order_date, last_order_id, last_order_sp, last_order_bp, lifetime_spend, most_recent_12_month_spend, data, token
        FROM clients
        WHERE email = $1`

	var client Client

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(&client.ID, &client.CreatedAt, &client.FirstName, &client.MiddleName, &client.LastName, &client.Suffix, &client.Email, &client.Mobile, &client.Username, &client.RimanUserId, &client.Status, &client.OrganizationType, &client.SignupDate, &client.AnniversaryDate, &client.AccountType, &client.SponsorUsername, &client.MemberId, &client.Rank, &client.EnrollmentDate, &client.PersonalOrdersVolume, &client.PersonalClientsVolume, &client.TotalPersonalVolume, &client.CurrentMonthSp, &client.CurrentMonthBp, &client.LastOrderDate, &client.LastOrderId, &client.LastOrderSp, &client.LastOrderBp, &client.LifetimeSpend, &client.MostRecent12MonthSpend, &client.Data, &client.Token)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &client, nil
}

func (m ClientModel) Update(client *Client) error {
	query := `
        UPDATE clients
               SET data = $1, first_name = $2, middle_name = $3, last_name = $4, suffix = $5, email = $6, mobile = $7, username = $8, riman_user_id = $9, status = $10, organization_type = $11, signup_date = $12, anniversary_date = $13, account_type = $14, sponsor_username = $15, member_id = $16, rank = $17, enrollment_date = $18, personal_orders_volume = $19, personal_clients_volume = $20, total_personal_volume = $21, current_month_sp = $22, current_month_bp = $23, last_order_date = $24, last_order_id = $25, last_order_sp = $26, last_order_bp = $27, lifetime_spend = $28, most_recent_12_month_spend = $29, token = $30
	WHERE id = $31
     RETURNING id`

	args := []any{
		client.Data,
		client.FirstName,
		client.MiddleName,
		client.LastName,
		client.Suffix,
		client.Email,
		client.Mobile,
		client.Username,
		client.RimanUserId,
		client.Status,
		client.OrganizationType,
		client.SignupDate,
		client.AnniversaryDate,
		client.AccountType,
		client.SponsorUsername,
		client.MemberId,
		client.Rank,
		client.EnrollmentDate,
		client.PersonalOrdersVolume,
		client.PersonalClientsVolume,
		client.TotalPersonalVolume,
		client.CurrentMonthSp,
		client.CurrentMonthBp,
		client.LastOrderDate,
		client.LastOrderId,
		client.LastOrderSp,
		client.LastOrderBp,
		client.LifetimeSpend,
		client.MostRecent12MonthSpend,
		client.Token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&client.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrEditConflict
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m ClientModel) UpdateToken(id int64, token string) error {
	query := `
        UPDATE clients
               SET token = $1
	WHERE id = $2
     RETURNING token`

	args := []any{token, id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&token)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "a lot of different ones"`:
			return ErrEditConflict
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m ClientModel) GetAll(username string, filters Filters) ([]*Client, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(),id, created_at, first_name, middle_name, last_name, suffix, email, mobile, username, riman_user_id, status, organization_type, signup_date, anniversary_date, account_type, sponsor_username, member_id, rank, enrollment_date, personal_orders_volume, personal_clients_volume, total_personal_volume, current_month_sp, current_month_bp, last_order_date, last_order_id, last_order_sp, last_order_bp, lifetime_spend, most_recent_12_month_spend, data, token
        FROM clients
        WHERE (to_tsvector('simple', username) @@ plainto_tsquery('simple', $1) OR $1 = '')      
        ORDER BY %s %s, id ASC
        LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{username, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	clients := []*Client{}

	for rows.Next() {
		var client Client

		err := rows.Scan(
			&totalRecords,
			&client.ID,
			&client.CreatedAt,
			&client.FirstName,
			&client.MiddleName,
			&client.LastName,
			&client.Suffix,
			&client.Email,
			&client.Mobile,
			&client.Username,
			&client.RimanUserId,
			&client.Status,
			&client.OrganizationType,
			&client.SignupDate,
			&client.AnniversaryDate,
			&client.AccountType,
			&client.SponsorUsername,
			&client.MemberId,
			&client.Rank,
			&client.EnrollmentDate,
			&client.PersonalOrdersVolume,
			&client.PersonalClientsVolume,
			&client.TotalPersonalVolume,
			&client.CurrentMonthSp,
			&client.CurrentMonthBp,
			&client.LastOrderDate,
			&client.LastOrderId,
			&client.LastOrderSp,
			&client.LastOrderBp,
			&client.LifetimeSpend,
			&client.MostRecent12MonthSpend,
			&client.Data,
			&client.Token,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		clients = append(clients, &client)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return clients, metadata, nil
}

func (m ClientModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM clients
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m PermissionModel) AddForClient(userID int64, clientID int64) error {
	query := `
        INSERT INTO clients_users (client_id, user_id)
    	VALUES ($1, $2)
		RETURNING client_id, user_id
	`

	args := []any{clientID, userID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&clientID, &userID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "clients_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return err
}
