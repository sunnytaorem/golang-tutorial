-- Count rows for all expected loadcell tags across selected days.
-- Pattern covered: Leg1..4, ChordA..C, Motor1..8 => 96 tags per day.
WITH day_windows AS (
  SELECT
    DATE '2025-06-29' AS day,
    TIMESTAMP '2025-06-29 00:00:00' AS start_ts,
    TIMESTAMP '2025-06-30 00:00:00' AS end_ts
  UNION ALL
  SELECT
    DATE '2025-11-09' AS day,
    TIMESTAMP '2025-11-09 00:00:00' AS start_ts,
    TIMESTAMP '2025-11-10 00:00:00' AS end_ts
),
all_tags AS (
  WITH legs AS (
    SELECT leg_num
    FROM (VALUES (1), (2), (3), (4)) AS l(leg_num)
  ),
  chords AS (
    SELECT chord
    FROM (VALUES ('A'), ('B'), ('C')) AS c(chord)
  ),
  motors AS (
    SELECT motor_num
    FROM (VALUES (1), (2), (3), (4), (5), (6), (7), (8)) AS m(motor_num)
  )
  SELECT
    'Leg' || leg_num || '_Chord' || chord || '_Motor' || motor_num || '_Loadcell_ActLoad' AS tagname
  FROM legs
  CROSS JOIN chords
  CROSS JOIN motors
)
SELECT
  d.day,
  t.tagname,
  COUNT(v.tagname) AS row_count
FROM day_windows d
CROSS JOIN all_tags t
LEFT JOIN timeseries.v_raw v
  ON v.tagname = t.tagname
 AND v.timestamp_utc >= d.start_ts
 AND v.timestamp_utc < d.end_ts
GROUP BY d.day, t.tagname
ORDER BY d.day, t.tagname;

-- -- Optional sanity check: quick sample from raw table.
-- SELECT *
-- FROM timeseries.raw
-- WHERE timestamp_utc >= TIMESTAMP '2025-06-29 00:00:00'
--   AND timestamp_utc < TIMESTAMP '2025-06-30 00:00:00'
-- ORDER BY timestamp_utc DESC
-- LIMIT 50;