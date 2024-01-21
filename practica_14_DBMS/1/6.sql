UPDATE expense 
SET 
  creation_date = CURRENT_DATE,
  import = 100
WHERE 
  creation_date < TO_DATE('15/11/2010', 'DD/MM/YYYY');