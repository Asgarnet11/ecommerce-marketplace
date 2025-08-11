-- 003_add_rest.sql — tabel pelengkap sesuai ERD

-- ===== addresses =====
CREATE TABLE IF NOT EXISTS addresses (
  id          BIGSERIAL PRIMARY KEY,
  user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  label       VARCHAR(80),
  recipient   VARCHAR(120) NOT NULL,
  phone       VARCHAR(32)  NOT NULL,
  line1       VARCHAR(200) NOT NULL,
  city        VARCHAR(120) NOT NULL,
  province    VARCHAR(120) NOT NULL,
  postal_code VARCHAR(16)  NOT NULL,
  notes       VARCHAR(200),
  is_default  BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE INDEX IF NOT EXISTS idx_addresses_user ON addresses(user_id);
CREATE UNIQUE INDEX IF NOT EXISTS uniq_default_address_per_user
  ON addresses(user_id) WHERE is_default = TRUE;

-- ===== carts & cart_items =====
CREATE TABLE IF NOT EXISTS carts (
  id         BIGSERIAL PRIMARY KEY,
  user_id    BIGINT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  session_id UUID UNIQUE
);

CREATE TABLE IF NOT EXISTS cart_items (
  id             BIGSERIAL PRIMARY KEY,
  cart_id        BIGINT NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
  product_id     BIGINT NOT NULL REFERENCES products(id),
  qty            INT    NOT NULL CHECK (qty > 0),
  price_snapshot NUMERIC(12,2) NOT NULL CHECK (price_snapshot >= 0),
  CONSTRAINT uniq_cart_product UNIQUE (cart_id, product_id)
);
CREATE INDEX IF NOT EXISTS idx_citems_cart    ON cart_items(cart_id);
CREATE INDEX IF NOT EXISTS idx_citems_product ON cart_items(product_id);

-- ===== orders =====
CREATE TABLE IF NOT EXISTS orders (
  id                 BIGSERIAL PRIMARY KEY,
  code               VARCHAR(40)  NOT NULL UNIQUE,
  checkout_group_id  UUID,
  user_id            BIGINT NOT NULL REFERENCES users(id),
  shop_id            BIGINT NOT NULL REFERENCES shops(id),
  address_snapshot   JSONB  NOT NULL,
  status             TEXT   NOT NULL DEFAULT 'pending'
                       CHECK (status IN ('pending','paid','processing','shipped','delivered','cancelled','refunded')),
  payment_status     TEXT   NOT NULL DEFAULT 'unpaid'
                       CHECK (payment_status IN ('unpaid','pending','paid','failed','refunded')),
  subtotal           NUMERIC(12,2) NOT NULL CHECK (subtotal >= 0),
  shipping_cost      NUMERIC(12,2) NOT NULL DEFAULT 0 CHECK (shipping_cost >= 0),
  total              NUMERIC(12,2) NOT NULL CHECK (total >= 0),
  note               VARCHAR(200),
  created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_orders_user   ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_shop   ON orders(shop_id);
CREATE INDEX IF NOT EXISTS idx_orders_group  ON orders(checkout_group_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status, payment_status);

-- ===== order_items =====
CREATE TABLE IF NOT EXISTS order_items (
  id             BIGSERIAL PRIMARY KEY,
  order_id       BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  product_id     BIGINT NOT NULL REFERENCES products(id),
  title_snapshot VARCHAR(180) NOT NULL,
  price          NUMERIC(12,2) NOT NULL CHECK (price >= 0),
  qty            INT NOT NULL CHECK (qty > 0)
);
CREATE INDEX IF NOT EXISTS idx_oit_order ON order_items(order_id);

-- ===== payments =====
CREATE TABLE IF NOT EXISTS payments (
  id           BIGSERIAL PRIMARY KEY,
  order_id     BIGINT NOT NULL UNIQUE REFERENCES orders(id) ON DELETE CASCADE,
  provider     TEXT   NOT NULL,        -- contoh: "dummy","midtrans","xendit"
  method       TEXT,
  amount       NUMERIC(12,2) NOT NULL CHECK (amount >= 0),
  status       TEXT   NOT NULL DEFAULT 'pending'
                 CHECK (status IN ('pending','paid','failed','refunded')),
  provider_ref TEXT,
  payload      JSONB  NOT NULL DEFAULT '{}'::jsonb,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);

-- ===== shipments =====
CREATE TABLE IF NOT EXISTS shipments (
  id            BIGSERIAL PRIMARY KEY,
  order_id      BIGINT NOT NULL UNIQUE REFERENCES orders(id) ON DELETE CASCADE,
  courier       TEXT   NOT NULL,       -- contoh: "JNE","SiCepat","Dummy"
  service       TEXT,
  tracking_code VARCHAR(80),
  status        TEXT   NOT NULL DEFAULT 'pending'
                 CHECK (status IN ('pending','packed','shipped','in_transit','delivered','returned','cancelled')),
  history       JSONB  NOT NULL DEFAULT '[]'::jsonb,
  shipped_at    TIMESTAMPTZ,
  delivered_at  TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_shipments_status ON shipments(status);
