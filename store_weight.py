# To store weight in a database. Should include date, weight.

import sqlite3

def create_sql_table(database):
    create_table_sql = """
    CREATE TABLE IF NOT EXISTS weight_table (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        date TEXT NOT NULL, 
        weight INTEGER NOT NULL-- Stores weight
    );
    """
    with sqlite3.connect(database) as connection:
        cursor = connection.cursor()
        cursor.execute(create_table_sql)
        connection.commit()

def add_weight_entry(database, weight_entry):
    with sqlite3.connect(database) as connection:
        insert_statement = ''' INSERT INTO weight_table(date, weight)
                                VALUES(?,?) '''
        
        cursor = connection.cursor()
        cursor.execute(insert_statement, weight_entry)
        connection.commit()

def get_weekly_avg(database):
    with sqlite3.connect(database) as connection:
        cursor = connection.cursor()
        cursor.execute('SELECT weight FROM weight_table ORDER BY id DESC LIMIT 7')
        return cursor.fetchall()
        
def main():
    database = 'weight.db'

    create_sql_table(database)

    weight_entries = [
        ("2025-08-01", 180.0),
        ("2025-08-02", 181.0),
        ("2025-08-03", 182.0),
        ("2025-08-04", 183.0),
        ("2025-08-05", 184.0),
        ("2025-08-06", 180.0),
        ("2025-08-07", 181.0),
        ("2025-08-08", 182.0),
        ("2025-08-09", 183.0),
        ("2025-08-10", 184.0)
    ]
    
    for weight_entry in weight_entries:
        add_weight_entry(database, weight_entry)
    
    weight_list = []
    for entry in get_weekly_avg(database):
        weight_list.append(*entry)
    
    weekly_avg = round(sum(weight_list) / len(weight_list),2)

    print(weekly_avg)

if __name__ == '__main__':
    main()