# To store weight in a database. Should include date, weight.

import sqlite3

connection = sqlite3.connect('weight.db')

cursor = connection.cursor()

create_table_sql = """
CREATE TABLE IF NOT EXISTS products (
    id TEXT PRIMARY KEY,
    date TEXT, 
    weight INTEGER -- Stores weight
);
"""
cursor.execute(create_table_sql)

cursor.close()