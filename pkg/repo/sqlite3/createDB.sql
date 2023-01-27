-- таблица состояний модуль ключ значение
-- INSERT OR REPLACE INTO app_state (module, key, value) values ("utm", "host", "localhost");
-- select value from app_state where module = 'utm' and key = 'host';
CREATE TABLE if not exists app_state (
  module TEXT,
  key TEXT,
  value TEXT,
  UNIQUE(module, key)
);
