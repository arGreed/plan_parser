create or replace function querry_tables(
	p_query text
)
returns table(
	schema_name text,
	table_name text
)
language plpgsql as $function$
declare
	--? Перечень полученных запросов
	queries text[];
	--? Рассматриваемый запрос
	single_query text;
	--? Перечень временных таблиц
	tmp_tables text[] := '{}';
	--? План
	plan jsonb;
	--? Созданные таблицы
	created_temp text;
begin
	--? Парсинг полученного перечня запросов на список запросов
	queries := string_to_array(regexp_replace(p_query, ';\s*$', ''), ';');

	--? Проход по перечню запросов
	foreach single_query in array queries loop
		begin
			if lower(single_query) ~ '^\s*(create|drop|alter)'
				then
					-- Обработка DDL с извлечением временных таблиц
					if lower(single_query) ~ '^create\s+temporary\s+table\s+(\w+)'
						then
							created_temp := (regexp_match(single_query, '^create\s+temporary\s+table\s+(\w+)', 'i'))[1];
							tmp_tables := tmp_tables || created_temp;
					end if;
					execute single_query;
				else
					execute 'explain (verbose, format json) ' || single_query into plan;

					return query
					select distinct
						coalesce(elem->>'Schema', 'pg_temp'),
						elem->>'Relation Name'
					from jsonb_path_query(plan, '$.**') as elem
					where
						elem ? 'Relation Name';
			end if;
			exception when others then
			continue;
		end;
	end loop;
		
	--? Возврат временных таблиц
	if tmp_tables <> '{}'
		then
			return query
				select
					'pg_temp',
					unnest(tmp_tables);
	end if;
end;
$function$;