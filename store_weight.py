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

def get_weight_data(database, data_length):
    
    with sqlite3.connect(database) as connection:
        cursor = connection.cursor()
        cursor.execute(f'SELECT weight FROM weight_table ORDER BY id DESC LIMIT {data_length}')
        weekly_weight = cursor.fetchall()
        
        # This look like: 
        # [(184.0,), (183.0,), (182.0,), (181.0,), (180.0,), (184.0,), (183.0,)]
        # Need to unpack it
        return [weight for (weight,) in weekly_weight]

def calculate_avg_weight(weight_data):
    """ Calculate average for list of weight values. """
    total_week_sum = sum(weight_data)
    avg_weekkly_weight = total_week_sum / len(weight_data)
    
    return round(avg_weekkly_weight, 2)

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
    
    week_data = get_weight_data(database,7) # For a week, 7 days

    avg_weekly_weight = calculate_avg_weight(week_data)

    print(avg_weekly_weight)

if __name__ == '__main__':
    main()