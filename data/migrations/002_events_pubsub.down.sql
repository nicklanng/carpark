DROP TRIGGER IF EXISTS event_insert ON events CASCADE;

DROP FUNCTION IF EXISTS new_event_notify;
