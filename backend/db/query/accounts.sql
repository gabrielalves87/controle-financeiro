-- name: CreateAccount :one
INSERT INTO accounts(
    user_id,
    categories_id,
    title,
    type,
    description,
    value,
    date
)VALUES(
    $1,$2,$3,$4,$5,$6,$7
) RETURNING *;

-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value
FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: GetAccountsGraph :one
SELECT COUNT(*)
FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT 
a.id,
a.user_id,
a.title,
a.type,
a.description,
a.value,
a.date,
a.created_at,
c.title as category_title
FROM
  accounts a
LEFT JOIN categories c ON c.id = a.categories_id
WHERE
  a.user_id = @user_id
AND
  a.type = @type
AND 
  LOWER(a.title) LIKE CONCAT('%', LOWER(@title::text), '%')
AND 
  LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%') 
AND 
  a.categories_id = COALESCE(sqlc.narg('categories_id'),a.categories_id)
AND 
  a.date = COALESCE(sqlc.narg('date'),a.date);

-- name: UpdateAccount :one
UPDATE accounts
SET title = $2, description = $3, value = $4
WHERE id = $1
RETURNING * ;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
