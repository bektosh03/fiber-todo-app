package postgres

import (
	"github.com/bektosh/fiber-app/models/responses"
	"github.com/lib/pq"
	"log"
)

func (p *Postgres) CheckUserEmail(email string) (bool, error) {
	var exists bool

	err := p.db.Get(
		&exists,
		`SELECT count(1) = 1 FROM users WHERE email = $1`,
		email,
	)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (p *Postgres) CreateUser(user responses.User, passwordHash string) (responses.User, error) {
	var (
		txErr   error
		pqArray interface{}
	)
	tx, err := p.db.Beginx()
	if err != nil {
		return responses.User{}, err
	}

	if user.CompanyID == "" {
		pqArray = pq.Array([]string{})
	} else {
		pqArray = pq.Array(user.CompanyID)
	}
	_, err = tx.Exec(
		`INSERT INTO users_auth (id, refresh_token, email, password) VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.RefreshToken,
		user.Email,
		passwordHash,
	)
	if err != nil {
		txErr = tx.Rollback()
		if txErr != nil {
			log.Fatalln("unable to rollback:", txErr)
		}
		return responses.User{}, err
	}

	_, err = tx.Exec(
		`INSERT INTO users
    			(user_id, email, first_name, last_name, username, age, gender, profile_photo, created_at, company_id)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, $9)`,
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Username,
		nullInt(int(user.Age)),
		nullString(user.Gender),
		nullString(user.ProfilePhoto),
		pqArray,
	)
	if err != nil {
		txErr = tx.Rollback()
		if txErr != nil {
			log.Fatalln("unable to rollback:", txErr)
		}
		return responses.User{}, err
	}

	txErr = tx.Commit()
	if txErr != nil {
		log.Fatalln("unable to commit:", txErr)
	}

	return user, nil
}

//func (p *Postgres) GetCompanyUsers(page, limit int, companyID, name string) ([]responses.User, int, error) {
//	var (
//		rows         *sqlx.Rows
//		err          error
//		profilePhoto sql.NullString
//		offset       = (page - 1) * limit
//		count        int
//	)
//	if name == "" {
//		err = p.db.Get(
//			&count,
//			`SELECT count(*) FROM users
//				WHERE company_id = $1`,
//			companyID,
//		)
//		if err != nil {
//			return nil, 0, err
//		}
//		rows, err = p.db.Queryx(`
//			SELECT user_id,
//				   first_name,
//				   last_name,
//				   profile_photo
//			FROM users
//			WHERE company_id = $1
//			LIMIT $2 OFFSET $3`,
//			companyID, limit, offset,
//		)
//	} else {
//		err = p.db.Get(
//			&count,
//			`SELECT count(*) FROM users
//				WHERE company_id = $1`,
//			companyID,
//		)
//		if err != nil {
//			return nil, 0, err
//		}
//		rows, err = p.db.Queryx(`
//			SELECT user_id,
//				   first_name,
//				   last_name,
//				   profile_photo
//			FROM users
//			WHERE company_id = $1
//			LIMIT $2 OFFSET $3`,
//			companyID, limit, offset,
//		)
//	}
//}
