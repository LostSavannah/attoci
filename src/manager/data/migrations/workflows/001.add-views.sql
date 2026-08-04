CREATE VIEW workflow_cron_triggers AS
SELECT 
    workflows.id as workflow_id, 
    value as cron 
FROM 
    workflows, json_each(data, '$.triggers.crons');

CREATE VIEW workflow_event_triggers AS
SELECT 
    workflows.id as workflow_id, 
    value as cron 
FROM 
    workflows, json_each(data, '$.triggers.crons');