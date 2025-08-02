# To store weight in a database. Should include date, weight.

import sqlite3

database = 'weight.db'
create_table_sql = """
CREATE TABLE IF NOT EXISTS products (
    id TEXT PRIMARY KEY,
    date TEXT NOT NULL, 
    weight REAL NOT NULL-- Stores weight
);
"""

with sqlite3.connect(database) as connection:
    cursor = connection.cursor()
    cursor.execute(create_table_sql)
    connection.commit()
