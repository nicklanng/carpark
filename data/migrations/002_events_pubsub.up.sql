CREATE OR REPLACE FUNCTION new_event_notify() RETURNS TRIGGER AS $$
  DECLARE
    payload varchar;
  BEGIN
    payload := CAST(NEW.seq AS text) || ',' || CAST(NEW.type AS text) || ',' || CAST(encode(NEW.data::bytea, 'hex') AS text);
    PERFORM pg_notify('new_event', payload);
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER event_insert AFTER INSERT ON events
FOR EACH ROW EXECUTE PROCEDURE new_event_notify();
