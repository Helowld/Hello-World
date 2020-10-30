-- When you use func as SQL, you need to write as declare -> begin -> end
declare
begin
  -- dbms_output.put_line displays the string inside quotation
  dbms_output.put_line('Hello World');
end;
/

-- You need to set serveroutput on if you want to see output from console, for more info, visit https://docs.oracle.com/database/121/LNPLS/toc.htm
-- OUTPUT:
-- Hello, World!
