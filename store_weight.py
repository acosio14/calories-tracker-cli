# To store weight in a database. Should include date, weight.

import sqlite3

database = 'weight.db'
create_table_sql = """
CREATE TABLE IF NOT EXISTS weight_table (
    id TEXT PRIMARY KEY,
    date TEXT NOT NULL, 
    weight REAL NOT NULL-- Stores weight
);
"""

with sqlite3.connect(database) as connection:
    cursor = connection.cursor()
    cursor.execute(create_table_sql)
    connection.commit()

def add_weight(conn, weight_entry):
    insert_statement = ''' INSERT INTO weight_table(id, date, weight)
                            VALUES(?,?,?,?) '''
    
    cursor = conn.cursor()

    cursor.execute(insert_statement, weight_entry)

    conn.commit()

    return cursor.lastrowid