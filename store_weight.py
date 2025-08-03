# To store weight in a database. Should include date, weight.

import sqlite3

def create_sql_table():
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

def add_weight_entry(connection, weight_entry):
    insert_statement = ''' INSERT INTO weight_table(id, date, weight)
                            VALUES(?,?,?,?) '''
    
    cursor = connection.cursor()
    cursor.execute(insert_statement, weight_entry)
    connection.commit()

def main():

    weight_entries = [
        (0, "2025-08-01", 180.0),
        (0, "2025-08-02", 180.0),
        (0, "2025-08-03", 180.0),
        (0, "2025-08-04", 180.0),
        (0, "2025-08-05", 180.0)
    ]

    for weight_entry in weight_entries:
        add_weight_entry(weight_entry)

if __name__ == '__main__':
    main()