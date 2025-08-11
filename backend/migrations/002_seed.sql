-- 002_seed.sql
-- SEED DATA DUMMY BANYAK: users, shops, categories (+sub), products (+images, +product_categories)

BEGIN;

-- =========================================
-- 1) USERS: beberapa seller + 1 buyer demo
-- =========================================
-- password_hash diset 'x' (placeholder); untuk login, gunakan akun yang kamu buat via /auth/register
INSERT INTO users (name, email, role, password_hash, status)
VALUES
  ('Buyer Demo',   'buyer@demo.local',  'buyer',  'x', 'active'),
  ('Seller One',   'seller1@demo.local','seller', 'x', 'active'),
  ('Seller Two',   'seller2@demo.local','seller', 'x', 'active'),
  ('Seller Three', 'seller3@demo.local','seller', 'x', 'active'),
  ('Seller Four',  'seller4@demo.local','seller', 'x', 'active'),
  ('Seller Five',  'seller5@demo.local','seller', 'x', 'active')
ON CONFLICT (email) DO NOTHING;

-- Ambil daftar seller untuk referensi shops
-- (Optional, hanya untuk dokumentasi)
-- SELECT id, email FROM users WHERE role='seller' ORDER BY id;

-- =========================================
-- 2) CATEGORIES: top-level + sub (banyak)
-- =========================================
-- Top-level
INSERT INTO categories (parent_id, name, slug, sort_order)
VALUES
  (NULL,'Sembako','sembako',1),
  (NULL,'Elektronik','elektronik',2),
  (NULL,'Fashion','fashion',3),
  (NULL,'Kesehatan','kesehatan',4),
  (NULL,'Olahraga','olahraga',5),
  (NULL,'Rumah Tangga','rumah-tangga',6),
  (NULL,'Ibu & Bayi','ibu-bayi',7),
  (NULL,'Hobi','hobi',8),
  (NULL,'Pertanian','pertanian',9),
  (NULL,'Otomotif','otomotif',10),
  (NULL,'Perlengkapan Kantor','perkantoran',11),
  (NULL,'Kecantikan','kecantikan',12)
ON CONFLICT (slug) DO NOTHING;

-- Sub‑categories: untuk setiap parent, buat 3 sub (total banyak)
WITH parents AS (
  SELECT id, slug FROM categories WHERE parent_id IS NULL ORDER BY id
)
INSERT INTO categories (parent_id, name, slug, sort_order)
SELECT p.id,
       CONCAT(initcap(replace(p.slug,'-',' ')),' Sub ', i) AS name,
       p.slug || '-sub-' || i AS slug,
       i
FROM parents p
CROSS JOIN generate_series(1,3) AS i
ON CONFLICT (slug) DO NOTHING;

-- =========================================
-- 3) SHOPS: puluhan toko tersebar ke para seller
-- =========================================
-- Tiap seller akan punya 6 toko: total 30 toko (kalau 5 seller)
WITH sellers AS (
  SELECT id AS seller_id FROM users WHERE role='seller' ORDER BY id
),
expanded AS (
  SELECT s.seller_id, gs.n
  FROM sellers s
  CROSS JOIN LATERAL generate_series(1,6) AS gs(n)
)
INSERT INTO shops (owner_id, name, slug, status)
SELECT seller_id,
       'Toko ' || seller_id || '-' || n AS name,
       'toko-' || seller_id || '-' || n AS slug,
       'active'
FROM expanded
ON CONFLICT (slug) DO NOTHING;

-- =========================================
-- 4) PRODUCTS: banyak produk untuk setiap shop
-- =========================================
-- Aturan:
--  - Setiap shop dibuat 12 produk
--  - Harga acak 10.000 - 1.000.000 (kelipatan ~10.000)
--  - Stok acak 10 - 200
--  - Slug unik per produk
--  - Status 'active'
--  - attributes kosong {}
WITH all_shops AS (
  SELECT id AS shop_id FROM shops ORDER BY id
),
gen AS (
  SELECT s.shop_id, gs.n
  FROM all_shops s
  CROSS JOIN LATERAL generate_series(1,12) AS gs(n)
),
ins AS (
  INSERT INTO products (shop_id, title, slug, description, price, stock, sku, status, attributes)
  SELECT
    g.shop_id,
    'Produk ' || g.shop_id || '-' || g.n AS title,
    'produk-' || g.shop_id || '-' || g.n AS slug,
    'Deskripsi produk ' || g.shop_id || '-' || g.n AS description,
    -- harga: 10.000 s/d 1.000.000
    ( (FLOOR(random()*90)::int + 10) * 10000 )::numeric(12,2) AS price,
    (FLOOR(random()*191)::int + 10) AS stock, -- 10..200
    'SKU-' || g.shop_id || '-' || g.n AS sku,
    'active' AS status,
    '{}'::jsonb AS attributes
  FROM gen g
  ON CONFLICT (slug) DO NOTHING
  RETURNING id
)
SELECT 1; -- dummy select agar CTE 'ins' dieksekusi

-- =========================================
-- 5) PRODUCT IMAGES: beri 1 gambar utama untuk setiap produk yang belum punya
-- =========================================
INSERT INTO product_images (product_id, url, is_primary, sort_order)
SELECT p.id,
       'https://picsum.photos/seed/prod' || p.id || '/600/600' AS url,
       TRUE,
       0
FROM products p
LEFT JOIN product_images i
  ON i.product_id = p.id AND i.is_primary = TRUE
WHERE i.id IS NULL;

-- =========================================
-- 6) PRODUCT CATEGORIES: assign 1 sub‑category random per produk (jika belum ada)
-- =========================================
INSERT INTO product_categories (product_id, category_id)
SELECT p.id, c.id
FROM products p
LEFT JOIN product_categories pc ON pc.product_id = p.id
CROSS JOIN LATERAL (
  SELECT id FROM categories
  WHERE parent_id IS NOT NULL
  ORDER BY random()
  LIMIT 1
) AS c
WHERE pc.product_id IS NULL;

COMMIT;

-- ====== VERIFIKASI CEPAT (opsional)
-- SELECT COUNT(*) AS users FROM users;
-- SELECT COUNT(*) AS shops FROM shops;
-- SELECT COUNT(*) AS categories FROM categories;
-- SELECT COUNT(*) AS products FROM products;
-- SELECT COUNT(*) AS images FROM product_images;
-- SELECT COUNT(*) AS prod_cats FROM product_categories;
