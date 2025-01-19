CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY, -- ID pengguna (string)
    email VARCHAR(255) UNIQUE NOT NULL, -- Email pengguna (unik dan wajib)
    password VARCHAR(255) NOT NULL, -- Password pengguna (wajib)
    name VARCHAR(100), -- Nama pengguna (opsional, panjang 2-100 karakter)
    image_uri VARCHAR(255), -- URI gambar profil pengguna (opsional)
    preference VARCHAR(50) CHECK (preference IN ('CARDIO', 'WEIGHT')), -- Preferensi pengguna (enum)
    weight_unit VARCHAR(50) CHECK (weight_unit IN ('KG', 'LBS')), -- Unit berat (enum)
    height_unit VARCHAR(50) CHECK (height_unit IN ('CM', 'INCH')), -- Unit tinggi (enum)
    weight SMALLINT CHECK (weight >= 10 AND weight <= 1000), -- Berat badan (wajib, 10-1000)
    height SMALLINT CHECK (height >= 10 AND height <= 1000), -- Tinggi badan (wajib, 10-1000)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu pembuatan
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Waktu pembaruan
);