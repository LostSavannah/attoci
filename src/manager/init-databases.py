from data.tools import dbmaster

dbmaster.migrate_db("workflows", True)
dbmaster.migrate_db("queues", True)
dbmaster.import_json_data("workflows", "workflows", "data", "workflows")