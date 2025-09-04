-- +goose Up
-- +goose StatementBegin

-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_admin BOOLEAN DEFAULT FALSE
);

-- Фильмы/Сериалы
CREATE TABLE films (
    id SERIAL PRIMARY KEY,
    title_ru VARCHAR(255) NOT NULL,
    title_original VARCHAR(255),
    description TEXT,
    year INTEGER,
    duration_minutes INTEGER,
    country VARCHAR(100),
    director VARCHAR(255),
    budget BIGINT,
    revenue BIGINT,
    age_rating VARCHAR(10),
    poster_url TEXT,
    backdrop_url TEXT,
    trailer_url TEXT,
    kinopoisk_id INTEGER UNIQUE,
    imdb_rating DECIMAL(3, 1),
    kinopoisk_rating DECIMAL(3, 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    type VARCHAR(10) DEFAULT 'film' CHECK (type IN ('film', 'series'))
);

-- Жанры
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    slug VARCHAR(50) UNIQUE NOT NULL
);

-- Связь фильмов и жанров
CREATE TABLE film_genres (
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genres (id) ON DELETE CASCADE,
    PRIMARY KEY (film_id, genre_id)
);

-- Актеры и съемочная группа
CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    name_ru VARCHAR(255) NOT NULL,
    name_original VARCHAR(255),
    biography TEXT,
    birth_date DATE,
    death_date DATE,
    photo_url TEXT,
    place_of_birth VARCHAR(255)
);

-- Роли в фильмах
CREATE TABLE film_persons (
    id SERIAL PRIMARY KEY,
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    person_id INTEGER REFERENCES persons (id) ON DELETE CASCADE,
    role_type VARCHAR(20) NOT NULL,
    role_name VARCHAR(255),
    "order" INTEGER DEFAULT 0
);

-- Отзывы и рецензии
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    title VARCHAR(255),
    content TEXT NOT NULL,
    rating INTEGER CHECK (rating BETWEEN 1 AND 10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_approved BOOLEAN DEFAULT FALSE,
    likes_count INTEGER DEFAULT 0
);

-- Оценки пользователей
CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (film_id, user_id)
);

-- Избранное
CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, film_id)
);

-- Серии для сериалов
CREATE TABLE episodes (
    id SERIAL PRIMARY KEY,
    film_id INTEGER REFERENCES films (id) ON DELETE CASCADE,
    season_number INTEGER NOT NULL,
    episode_number INTEGER NOT NULL,
    title_ru VARCHAR(255),
    title_original VARCHAR(255),
    description TEXT,
    duration_minutes INTEGER,
    release_date DATE,
    preview_url TEXT,
    UNIQUE (
        film_id,
        season_number,
        episode_number
    )
);

-- Лайки отзывов
CREATE TABLE review_likes (
    id SERIAL PRIMARY KEY,
    review_id INTEGER REFERENCES reviews (id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    is_like BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (review_id, user_id)
);

-- Индексы для оптимизации
CREATE INDEX idx_films_year ON films (year);

CREATE INDEX idx_films_rating ON films (kinopoisk_rating);

CREATE INDEX idx_films_title ON films (title_ru);

CREATE INDEX idx_reviews_film_id ON reviews (film_id);

CREATE INDEX idx_reviews_user_id ON reviews (user_id);

CREATE INDEX idx_ratings_film_id ON ratings (film_id);

CREATE INDEX idx_film_persons_film_id ON film_persons (film_id);

CREATE INDEX idx_film_persons_person_id ON film_persons (person_id);

CREATE INDEX idx_film_genres_film_id ON film_genres (film_id);

-- Добавляем основные жанры
INSERT INTO
    genres (name, slug)
VALUES ('драма', 'drama'),
    ('комедия', 'comedy'),
    ('боевик', 'action'),
    ('фантастика', 'sci-fi'),
    ('ужасы', 'horror'),
    ('триллер', 'thriller'),
    ('мелодрама', 'melodrama'),
    ('приключения', 'adventure'),
    ('детектив', 'detective'),
    ('фэнтези', 'fantasy');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS review_likes;

DROP TABLE IF EXISTS episodes;

DROP TABLE IF EXISTS favorites;

DROP TABLE IF EXISTS ratings;

DROP TABLE IF EXISTS reviews;

DROP TABLE IF EXISTS film_persons;

DROP TABLE IF EXISTS persons;

DROP TABLE IF EXISTS film_genres;

DROP TABLE IF EXISTS genres;

DROP TABLE IF EXISTS films;

DROP TABLE IF EXISTS users;

-- +goose StatementEnd