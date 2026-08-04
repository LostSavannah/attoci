CREATE TABLE queues(
    id TEXT PRIMARY KEY,
    queue_name TEXT,
    data TEXT,
    status TEXT,
    created_at INT,
    reserved_at INT,
    completed_at INT
);