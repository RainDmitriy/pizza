CREATE TABLE items(
  ItemId SERIAL PRIMARY KEY REFERENCES items(ItemId),
  Title TEXT UNIQUE,
  Types integer[],
  Rating INT
);

CREATE TABLE props(
  ItemId SERIAL PRIMARY KEY REFERENCES items(ItemId),
  Spicy BOOL,
  Meat BOOL,
  Vegetarian BOOL,
  Grilled BOOL,
  Closed BOOL
);

CREATE TABLE prices(
  ItemId SERIAL PRIMARY KEY REFERENCES items(ItemId),
  forSmall INT,
  forMiddle INT,
  forBig INT
);

CREATE TABLE sizes(
  ItemId SERIAL PRIMARY KEY REFERENCES items(ItemId),
  small INT,
  middle INT,
  big INT
);

CREATE TABLE image(
  ItemId SERIAL PRIMARY KEY REFERENCES items(ItemId),
  Thin TEXT,
  Traditional TEXT
);

CREATE TABLE cart(
  CartId SERIAL PRIMARY KEY,
  ItemId INT REFERENCES items(ItemId),
  SelectedType INT,
  SelectedSize INT,
  Quantity INT
);

CREATE TABLE orders(
  OrderId SERIAL PRIMARY KEY,
  CartItems JSON,
  TotalPrice INT
);