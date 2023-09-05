package product_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"
	"go-learn/library/meta"
	"math"
	"strings"

	"github.com/google/uuid"
)

func (c *_ProductRepoImp) Create(pr *entities.Product) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	queryInsert := `INSERT INTO products (id, title, description, points, quantity, image, type, banner, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = tx.Exec(queryInsert, pr.ID, pr.Title, pr.Description, pr.Points, pr.Qty, pr.Image, pr.Type, pr.Banner, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	queryInfo := `INSERT INTO product_info (id, product_id, info, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	newId, _ := uuid.NewUUID()
	_, err = tx.Exec(queryInfo, newId, pr.ID, pr.Info, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_ProductRepoImp) Detail(id uuid.UUID) (*entities.Product, error) {
	query := `
		SELECT 
			p.id, 
			p.title, 
			p.description, 
			p.points, 
			p.quantity, 
			COALESCE(ph.rating, 0) AS rating,
			p.image, 
			p.type, 
			p.banner, 
			pi.info, 
			p.created_at, 
			p.updated_at 
		FROM 
			products p 
		JOIN 
			product_info pi 
		ON 
			p.id = pi.product_id 
		LEFT JOIN (
				SELECT product_id, AVG(rating) AS rating
				FROM product_history
				GROUP BY product_id
				) ph ON p.id = ph.product_id
			
		WHERE 
			p.id = $1`

	var object entities.Product

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.Title,
		&object.Description,
		&object.Points,
		&object.Qty,
		&object.Rating,
		&object.Image,
		&object.Type,
		&object.Banner,
		&object.Info,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	object.Rating = c.RoundRatingToNearestHalf(object.Rating)

	return &object, nil
}

func (c *_ProductRepoImp) GetAll(m *meta.Metadata) ([]entities.Product, error) {
	q, err := meta.ParseMetaData(m, c)
	if err != nil {
		return nil, err
	}
	stmt := `SELECT p.id, p.title, p.description, p.points, p.quantity, p.image, p.type, p.banner, p.created_at, p.updated_at, 
	COALESCE(ph.rating, 0) AS rating
	FROM products p
	LEFT JOIN (
		SELECT product_id, AVG(rating) AS rating
		FROM product_history
		GROUP BY product_id
		) ph ON p.id = ph.product_id
	`
	queries := QueryStatement(stmt)
	var (
		searchBy = q.SearchBy
		order    = q.OrderBy
	)
	if len(q.Search) > 2 {
		if len(q.SearchBy) != 0 {
			queries = queries.Where(searchBy, like, q.Search)
		}
	}
	if q.DateEnd.Valid && q.DateFrom.Valid {
		queries = queries.Where(order, between, q.DateFrom.Time.Local(), q.DateEnd.Time.Local())
	}

	query, _, args := queries.Order(order, direction(strings.ToUpper(q.OrderDirection))).
		Offset(q.Offset).
		Limit(q.Limit).Build()

	rows, err := c.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]entities.Product, 0)
	for rows.Next() {
		var f entities.Product
		if err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Description,
			&f.Points,
			&f.Qty,
			&f.Image,
			&f.Type,
			&f.Banner,
			&f.CreatedAt,
			&f.UpdatedAt,
			&f.Rating,
		); err != nil {
			return nil, err
		}

		f.Rating = c.RoundRatingToNearestHalf(f.Rating)

		collections = append(collections, f)
	}

	return collections, nil
}

func (c *_ProductRepoImp) Update(pr *entities.Product) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	query := `
		UPDATE products
		SET
			title = $2,
			description = $3,
			points = $4,
			quantity = $5,
			image = $6,
			type = $7,
			banner = $8,
			created_at = $9,
			updated_at = $10
		WHERE
			id = $1
	`

	_, err = tx.Exec(query, pr.ID, pr.Title, pr.Description, pr.Points, pr.Qty, pr.Image, pr.Type, pr.Banner, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	queryInfo := `
			UPDATE product_info 
			SET 
				info = $2, 
				updated_at = $3
			WHERE 
				product_id = $1`
	_, err = tx.Exec(queryInfo, pr.ID, pr.Info, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_ProductRepoImp) Delete(id uuid.UUID) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	query := `DELETE FROM products WHERE id = $1`

	_, err = tx.Exec(query, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}

	query2 := `DELETE FROM product_info WHERE product_id = $1`

	_, err = tx.Exec(query2, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}
	return nil
}

func (c *_ProductRepoImp) Redeem(rq *entities.RedeemRequired) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	updateQuery := `
		UPDATE products
		SET
			quantity = $2
		WHERE
			id = $1
	`

	_, err = tx.Exec(updateQuery, rq.ProductID, rq.QtyAfter)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	insertQuery := `INSERT INTO product_history (id, product_id, email_user, quantity, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(insertQuery, rq.ID, rq.ProductID, rq.Email, rq.QtyReq, rq.CreatedAt, rq.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil

}

func (c *_ProductRepoImp) Rating(rq *entities.RatingRequired) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	checkQuery := `SELECT COUNT(*) FROM product_history WHERE email_user = $1 AND product_id = $2 AND rating is NULL`
	var count int
	err = tx.QueryRow(checkQuery, rq.Email, rq.ProductID).Scan(&count)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	if count == 0 {
		return entities.ErrNeverRedeem
	}

	checkRating := `SELECT COUNT(*) FROM product_history WHERE email_user = $1 AND product_id = $2 AND rating is NOT NULL`
	var countRating int
	err = tx.QueryRow(checkRating, rq.Email, rq.ProductID).Scan(&countRating)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	if countRating == 1 {
		return entities.ErrAlreadyRating
	}

	insertQuery := `INSERT INTO product_history (id, product_id, email_user,rating, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(insertQuery, rq.ID, rq.ProductID, rq.Email, rq.Rating, rq.CreatedAt, rq.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_ProductRepoImp) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}

func (c *_ProductRepoImp) RoundRatingToNearestHalf(rating float64) float64 {
	rounded := math.Round(rating*2) / 2
	return rounded
}
