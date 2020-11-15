// Hello World Program in Verilog HDL.
// Any program in Verilog starts with reserved word 'module' <module_name>.
module main;
  // initial block is another construct typically used to initialize signals nets and variables for simulation. 
  // This block gets executed only once after the simulation starts, at time=0 (0ns).
  initial
  begin
  // To display the signal value to the screen.
  $display("Hello, World!");
  $finish ;
  end
// End of Module main
endmodule

// For more info check this site http://www.asic-world.com/verilog/veritut.html
// For more info check this site https://www.javatpoint.com/verilog

// Output :
// Hello, World!
