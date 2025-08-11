-- 004_seed_address.sql
INSERT INTO addresses (user_id, label, recipient, phone, line1, city, province, postal_code, notes, is_default)
VALUES
  (2, 'Rumah', 'Buyer Demo', '0812000000', 'Jl. Melati 1', 'Makassar', 'Sulawesi Selatan', '90111', 'Patokan masjid', true)
ON CONFLICT DO NOTHING;
