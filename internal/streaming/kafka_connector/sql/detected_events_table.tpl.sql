CREATE TABLE IF NOT EXISTS OM_DETECTED_EVENTS
WITH (
    KAFKA_TOPIC = {{ .Topic | squote }},
    KEY_FORMAT = 'JSON_SR',
    VALUE_FORMAT = 'JSON_SR',
    PARTITIONS = {{ .Partitions }}
) AS
SELECT
    ID AS KEY1,
    SOURCE AS KEY2,
    AS_VALUE(ID) AS ID,
    EARLIEST_BY_OFFSET(TYPE) AS TYPE,
    AS_VALUE(SOURCE) AS SOURCE,
    EARLIEST_BY_OFFSET(SUBJECT) AS SUBJECT,
    EARLIEST_BY_OFFSET(TIME) AS TIME,
    EARLIEST_BY_OFFSET(DATA) AS DATA,
    COUNT(ID) as ID_COUNT
FROM OM_EVENTS
WINDOW TUMBLING (
    SIZE {{ .Retention }} DAYS, 
    RETENTION {{ .Retention }} DAYS
)
GROUP BY 
    ID, 
    SOURCE;