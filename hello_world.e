-- To learn more: https://www.eiffel.org/doc/
--
--
--
-- 
-- Following format is often used in practical
note    
    description: "A Simple Hello World Program" 
    author: "Elizabeth W. Brown"

class 
    HELLO 

create    
    make

feature

    make
            -- Print a simple message.         
        do      
            io.put_string ("Hello, World!")
            io.put_new_line
        end

end -- class HELLO
