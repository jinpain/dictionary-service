-- name: CreateInsuranceCompany :one
INSERT INTO
    insurance_companies (
        type,
        name,
        full_name,
        address,
        phone,
        email,
        website,
        ogrn,
        okpo,
        okato,
        created_ts,
        created_by
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    ) RETURNING id;