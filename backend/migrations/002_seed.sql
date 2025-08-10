-- 002_seed.sql
INSERT INTO users (name, email, role, password_hash) VALUES
('Seller Demo', 'seller@demo.local', 'seller', 'x'),
('Buyer Demo',  'buyer@demo.local',  'buyer',  'x');

INSERT INTO shops (owner_id, name, slug, status) VALUES
(1, 'Toko Demo', 'toko-demo', 'active');

INSERT INTO categories (name, slug, sort_order) VALUES
('Sembako', 'sembako', 1),
('Pertanian','pertanian',2),
('Elektronik','elektronik',3)
ON CONFLICT DO NOTHING;

INSERT INTO products (shop_id, title, slug, description, price, stock, sku, status) VALUES
(1, 'Beras Premium 5kg', 'beras-premium-5kg', 'Beras wangi kualitas premium', 75000, 120, 'SKU-001', 'active'),
(1, 'Minyak Goreng 2L', 'minyak-goreng-2l', 'Minyak goreng sawit 2 liter', 38000, 50, 'SKU-002', 'active');

INSERT INTO product_images (product_id, url, is_primary) VALUES
(1, 'https://picsum.photos/seed/beras/600/600', true),
(2, 'https://picsum.photos/seed/minyak/600/600', true);

INSERT INTO product_categories (product_id, category_id) VALUES
(1, 1), (2, 1);
