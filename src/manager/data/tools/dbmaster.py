import sqlite3, os, time, uuid

MIGRATIONS_PATH = os.environ.get("ATTOCI_DB_MIGRATIONS_PATH")
DATABASES_PATH = os.environ.get("ATTOCI_DB_DATABASES_PATH")
JSON_PATH = os.environ.get("ATTOCI_DB_JSON_PATH")

MIGRATIONS_TABLE = "__migrations_history__"

def create_db(db_path:str):
    with sqlite3.connect(db_path) as con:
        con.executescript(f"CREATE TABLE {MIGRATIONS_TABLE} (name TEXT, applied_at_epoch FLOAT);")
        con.commit()

def migrate_db(dbname:str, reset:bool = False):
    print(f"Migrating '{dbname}'...")
    db_path = os.sep.join([DATABASES_PATH, f"{dbname}.db"])
    exists = os.path.isfile(db_path)
    if reset and exists:
        os.remove(db_path)
    if reset or not exists:
        create_db(db_path)
    db_migrations_path = os.sep.join([MIGRATIONS_PATH, dbname])
    for migration in sorted(os.listdir(db_migrations_path)):
        with sqlite3.connect(db_path) as con:
            exists = con.execute(f"SELECT COUNT(*) FROM {MIGRATIONS_TABLE} WHERE name = ?;", (migration,)).fetchone()[0] > 0
            if exists:
                print(f"Omitting existing migration '{migration}'")
                continue
            print(f"Applying migration '{migration}'...")
            with open(os.sep.join([db_migrations_path, migration]), 'r') as fi:
                statements = [f"{i.strip()};" for i in fi.read().split(";") if i.strip() != ""]
                for statement in statements:
                    print(statement)
                    con.executescript(statement)
                con.execute(f"INSERT INTO {MIGRATIONS_TABLE} (name, applied_at_epoch) VALUES (?, ?);", (migration, time.time()))
            con.commit()

def import_json_data(dbname:str, table_name:str, data_field:str, folder:str):
    db_path = os.sep.join([DATABASES_PATH, f"{dbname}.db"])
    json_folder = os.sep.join([JSON_PATH, folder])
    query = f"INSERT INTO {table_name}(id, {data_field}) VALUES (?, ?);"
    with sqlite3.connect(db_path) as con:
        for file in [os.sep.join([json_folder, f]) for f in os.listdir(json_folder)]:
            with open(file, 'r') as fi:
                json_data = fi.read()
                con.execute(query, (str(uuid.uuid4()), json_data))
        con.commit()

