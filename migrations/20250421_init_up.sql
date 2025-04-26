CREATE TABLE "products" (
                           "id" SERIAL PRIMARY KEY,
                           "name" varchar(255) NOT NULL,
                           "description" varchar(255) NOT NULL,
                           "sku" varchar(255) UNIQUE NOT NULL,
                           "price" numeric(18,2),
                           "quantity_in_stock" int,
                           "reorder_level" int
);
