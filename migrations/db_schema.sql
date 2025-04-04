CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    profile JSONB NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE posts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    content JSONB NOT NULL,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE comments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    text text NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    post_id uuid REFERENCES posts(id) ON DELETE CASCADE,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    parent_id uuid REFERENCES comments(id) ON DELETE CASCADE,
    deleted_at TIMESTAMPTZ
);

-- Indexes for faster lookups
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_parent ON comments(parent);
    