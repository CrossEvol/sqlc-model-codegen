# relation "`Category`" does not exist
```shell
$ sqlc generate 
# package sqliteDao
assets\queries\Category.sql:1:1: relation "`Category`" does not exist
assets\queries\Category.sql:8:1: relation "`Category`" does not exist
assets\queries\Category.sql:12:1: relation "`Category`" does not exist
assets\queries\Category.sql:17:1: relation "`Category`" does not exist
assets\queries\Category.sql:21:1: relation "`Category`" does not exist
assets\queries\Category.sql:25:1: relation "`Category`" does not exist
assets\queries\Category.sql:33:1: relation "`Category`" does not exist
```

`Category.sql`
```sqlite
-- name: GetCategory :one
SELECT *
FROM `Category`
WHERE id = ?
LIMIT 1;

-- name: GetCategories :many
SELECT *
FROM `Category`;

-- name: GetCategoriesByIds :many
SELECT *
FROM `Category`
WHERE id IN (sqlc.slice('ids'));

-- name: CountCategories :one
SELECT count(*)
FROM `Category`;

-- name: CreateCategory :execresult
INSERT INTO `Category` (`name`, `desc`, `created_at`, `updated_at`)
VALUES (?, ?, ?, ?);

-- name: UpdateCategory :execresult
UPDATE `Category`
SET `name`       = CASE WHEN @name IS NOT NULL THEN @name ELSE `name` END,
    `desc`       = CASE WHEN @desc IS NOT NULL THEN @desc ELSE `desc` END,
    `created_at` = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE `created_at` END,
    `updated_at` = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE
FROM `Category`
WHERE id = ?;

```

the table name in queries/ should be same as in the migrations/

# Wrong in types of UpdateXxxParams
`User.dto.go`
```go
func (dto *UpdateUserDTO)Map2UpdateUserParams() *sqliteDao.UpdateUserParams {
	Emailverified:= time.UnixMilli(dto.Emailverified)
		
	params := sqliteDao.UpdateUserParams{
        Name:  dto.Name  ,
        Password:  dto.Password  ,
        Email:  dto.Email  ,
        EmailVerified:  dto.EmailVerified  ,
        Image:  dto.Image  ,
        Role:  dto.Role  ,
        ID:  dto.ID  ,
        
    }
    return &params
}

```
the name `Emailverified` is not compatible with `EmailVerified`

`internal/database/sqliteDao/models.go`
```go
type User struct {
	ID            string     `db:"id" json:"id"`
	Name          *string    `db:"name" json:"name"`
	Password      *string    `db:"password" json:"password"`
	Email         *string    `db:"email" json:"email"`
	Emailverified *time.Time `db:"emailverified" json:"emailverified"`
	Image         *string    `db:"image" json:"image"`
	Role          string     `db:"role" json:"role"`
}
```

`internal/database/sqliteDao/User.sql.go`
```go
type UpdateUserParams struct {
	Name          interface{} `db:"name" json:"name"`
	Password      interface{} `db:"password" json:"password"`
	Email         interface{} `db:"email" json:"email"`
	EmailVerified interface{} `db:"emailVerified" json:"email_verified"`
	Image         interface{} `db:"image" json:"image"`
	Role          interface{} `db:"role" json:"role"`
	ID            string      `db:"id" json:"id"`
}
```

the wrong will happen in `codegen/collect_data.go`
```go
if target != nil && origin != nil {
    for _, targetFieldMeta := range target.FieldMetas {
        for _, originFieldMeta := range origin.FieldMetas {
            if targetFieldMeta.Name == originFieldMeta.Name && targetFieldMeta.Type != originFieldMeta.Type && targetFieldMeta.Type == "*ast.InterfaceType" {
                targetFieldMeta.Type = originFieldMeta.Type
                if strings.Index(targetFieldMeta.Type, "*") == -1 {
                    targetFieldMeta.Type = "*" + targetFieldMeta.Type
                }
            }
        }
    }
}
```
when compare the name, should use lowerCamelCase 
when define the column name in the table, should use the snake_case naming strategy