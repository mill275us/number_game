Number Game Fun (ngfX)

Objective
=========
To continue a process of weaving in different programming  concepts along
with a longer-term,  multi-week project theme allowing you  to apply your
newfound programming skills and pose further opportunities for learning.

In addition to programming concepts and  understanding, a big part of the
ngfX projects will be in helping you to develop better time and attention
management  skills. Programming  is NOT  some hoop  you can  memorize and
regurgitate  your  way  to  succes. Indeed,  it  requires  understanding,
experimentation, playing, focus,  and more than a little  patience as you
dabble in new things and ENDURE them over the long haul.

Background
==========
We  will  be  stepping  through  a series  of  projects  exploring  basic
programming concepts applied in the theme of implementing a game. As this
is a first exposure to many, we will not be pursuing anything exceedingly
involved. Instead, we will explore  the implementation of a simple number
game, complete with board/playfield. Here's an example:

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 3 |   |   |   |   |   |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 4 |   |   |   |   |   |
---+---+---+---+---+---+
 4 |   |   |   |   |   |
---+---+---+---+---+---+

The premise is to determine which blank squares within contain a value (a
square will either  be populated or it will be  ultimately blank). We use
the header/heading  values off to  the far left and  top to inform  us of
information in the associated row or column (respectively). The resulting
intersection gives us two sources of information we can work with to  get
us to our eventual solution.

For instance, taking a closer look at the middle row of data:

---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+

^^^ ^^^^^^^^^^^^^^^^^^^ - data cells
 |
 +--- header information

 What we notice  here is an indication  that in this row,  exactly ONE of
 the data cells is populated. The  other four are blank. Alone, this clue
 isn't too useful, so  we have to utilize the column  headings as well to
 help us on our way.

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
    ^^^
     |
     +-- at most 2 (consecutive) populated

A common strategy is to identify  the larger consecutive values to try to
"mark off" undeniable  possibilities. For example, looking  at the bottom
two rows,  both containing  4 consecutive  values, we  know that  of that
5-cell row, EXACTLY 4 cells (and 4 consecutive cells) are occupied.

While we don't know  which cells are populated and which  are not, we CAN
make a simplification  that reveals which cells  ARE populated regardless
of the possibilities.

In our  case, with  a row  of FIVE  data cells  and knowing  that exactly
FOUR  consecutive  cells  are  populated, that  actually  leaves  us  TWO
possibilities, both of which are shown in the table below:

---+---+---+---+---+---+
 4 | X | X | X | X |   | <-- first possibility (marked with 'X')
---+---+---+---+---+---+
 4 |   | Y | Y | Y | Y | <-- second possibility (marked with 'Y')
---+---+---+---+---+---+

        ^^^^^^^^^^^ - cells in common in both

Right? We may not  yet know what the actuality is, we  CAN notice that in
all available possibilities, there are some populated cells that occur in
BOTH scenarios.  This allows  us to  make a  determination (using  '*' to
denote a populated cell):

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 3 |   |   |   |   |   |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+

We could do similar to the row with  3 in it. Here we actually have THREE
possibilities:

---+---+---+---+---+---+
 3 | X | X | X |   |   | <-- first possibility (marked with 'X')
---+---+---+---+---+---+
 3 |   | Y | Y | Y |   | <-- second possibility (marked with 'Y')
---+---+---+---+---+---+
 3 |   |   | Z | Z | Z | <-- third possibility (marked with 'Z')
---+---+---+---+---+---+
            ^^^
             |
             +-- cell in common in all scenarios

So, we can mark up the table as appropriate:

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 3 |   |   | * |   |   |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+

The remaining two rows are a  bit challenging at present, without further
information being determined.

Looking  to the  top  heading, notice  there  is a  column  of 4.  Again,
CONSECUTIVE (meaning exactly next to each  other, no gaps). Notice how we
have THREE of the FOUR populated  cells in that column i determined. This
should be easy (and simultaneously satisfies one of the rows of 1):

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   |   |   |   |
---+---+---+---+---+---+
 3 |   |   | * |   |   |
---+---+---+---+---+---+
 1 |   |   | * |   |   | <-- blazam
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
            ^^^
             |
             +-- this column is complete

You might also notice that we've  solve the "2" column immediately to the
right of the "4" column (and we've had it solved for some time).

It might  be intuitive to also  mark the determinedly blank  cells, so we
know that  no further considerations  have to be  made. We'll use  '.' to
denote a cell that has been determined to be blank.

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   | . | . |   |
---+---+---+---+---+---+
 3 |   |   | * | . |   |
---+---+---+---+---+---+
 1 | . | . | * | . | . | <-- this row is also solved
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
                ^^^
                 |
                 +-- as is this column

You may also notice that the row of  3 can now be solved (since we know a
cell that connot be populated):

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   |   | . | . |   |
---+---+---+---+---+---+
 3 | * | * | * | . | . | <-- this row now solved
---+---+---+---+---+---+
 1 | . | . | * | . | . |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+
 4 |   | * | * | * |   |
---+---+---+---+---+---+

What next? How  about that right-most column of 2?  Since both cells have
to be  exactly next  to each other,  we should notice  there is  only one
place that fits that requirement. Hey, another column solved!

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 |   | . | . | . | . |
---+---+---+---+---+---+
 3 | * | * | * | . | . |
---+---+---+---+---+---+
 1 | . | . | * | . | . |
---+---+---+---+---+---+
 4 |   | * | * | * | * | <-- hey, this row is also solved
---+---+---+---+---+---+
 4 |   | * | * | * | * | <-- as is this row
---+---+---+---+---+---+
        ^^^         ^^^
         |           |
         |           +-- column solved!
         |
         +---- when we solved the row of 3, we had also got this

NOTE: the "1 2" column indicates  that THREE total cells are populated, a
group of ONE, and a group of TWO.  There must be a break between them, as
is seen on the solved column.

Finally, we have that sole remaining column solved, by similar conditions
as we've already experienced. So the final solved puzzle is:

   |   | 1 |   |   |   |
   | 2 | 2 | 4 | 2 | 2 |
---+---+---+---+---+---+
 1 | * | . | . | . | . |
---+---+---+---+---+---+
 3 | * | * | * | . | . |
---+---+---+---+---+---+
 1 | . | . | * | . | . |
---+---+---+---+---+---+
 4 | . | * | * | * | * |
---+---+---+---+---+---+
 4 | . | * | * | * | * |
---+---+---+---+---+---+

Boards can be larger, containing many groups of various sizes.

Program
=======
Functions  and structs  continue to  aid us  in the  organization of  our
program logic.

Take notice  that, through proper,  strategy deployment of  functions, we
can  further optimize  out some  other instances  of redundancy-  getting
closer to that ideal of ONE copy  of logic that many different aspects of
the program can utilize.

Program output is no longer fixed; each  time you run the program, due to
the incorporation of randomness into our code, the output will differ.

There will be two main modes of operation, a "regular" (default) and  the
"celldata" mode that can be useful for debugging purposes.  Running  your
program as normal should produce output like the following (but again, as
we are now randomly generating data, each successive run should result in
an entirely different board makeup):

<cli>
lab46:~/src/spring2021/cprog/ngfX$ ./ngfX
            |  2 |  1 |  1 |  1 |  2 |
            |    |  3 |  3 |    |    |
            |    |    |    |    |    |
------------+----+----+----+----+----+
  3         |    |    |    |    |    |
------------+----+----+----+----+----+
  1         |    |    |    |    |    |
------------+----+----+----+----+----+
  2         |    |    |    |    |    |
------------+----+----+----+----+----+
  2   1     |    |    |    |    |    |
------------+----+----+----+----+----+
  4         |    |    |    |    |    |
------------+----+----+----+----+----+
lab46:~/src/spring2021/cprog/ngfX$
</cli>

By setting a shell variable by the name of CELLDATA to the string "true",
program execution should also display the contents of the board cells:

<cli>
lab46:~/src/spring2021/cprog/ngfX$ CELLDATA="true" ./ngfX
            |  2 |  1 |  1 |  1 |  2 |
            |    |  3 |  3 |    |    |
            |    |    |    |    |    |
------------+----+----+----+----+----+
  3         |  1 |  1 |  1 |    |    |
------------+----+----+----+----+----+
  1         |  1 |    |    |    |    |
------------+----+----+----+----+----+
  2         |    |  1 |  1 |    |    |
------------+----+----+----+----+----+
  2   1     |    |  1 |  1 |    |  1 |
------------+----+----+----+----+----+
  4         |    |  1 |  1 |  1 |  1 |
------------+----+----+----+----+----+
lab46:~/src/spring2021/cprog/ngfX$
</cli>

The "make check" target in the Makefile has been updated to try BOTH  the
modes of operation.

Example code / Tutorial process
===============================
To assist you in your development  efforts, I have provided fully working
code to  a DIFFERENT  game, that  of tic-tac-toe, at  a similar  state of
development, and  fully commented, consistently indented,  and serving to
provide an application of necessary concepts in a manner you can observe,
discover, and understand, so that you can apply what needs to be done for
your required program you are to submit.

Be  sure  to  utilize  the  class  discord  to  ask  questions  and  gain
clarification  on  the  various  concepts and  details  surrounding  this
project.

Please  do  make   sure  you  read,  explore,  and   understand  what  is
happening in  the provided  example; I have  presented many  comments and
scenarios that  are directly applicable  and useful to your  overall ngfX
implementation, and problem solving skillset.

Strategy
========
A basic outline of actions you can take that can lead to maximal success:

  * use 'grabit' on lab46 to obtain project resources
  * add/commit/push to semester repository
  * pull/update to access on your pi development system
  * take a look around:
    * read this file
    * view the example code
  * explore the related content:
    * read/understand the comments in the example code
    * tinker with things to gain better understanding
    * make sure you understand what output the code produces
    * ask questions on things that aren't clear
  * once you are familiar with what you need to do,  and the  level  that
    this project calls for (you're implementing the number game to  about
    the same level of functionality as the example code implements it for
    tic tac toe), you can begin implementing your program (in ngfX.c)
  * be sure to comment as you go along! Explain your reasoning, even if a
    later change requires changes.
  * commit and push changes as you implement pieces, or complete a coding
    session, so that you have multiple snapshots to call upon should  you
    need to.
  * compile and run your code to ensure it is doing what you'd like it to
    do, or debug and troubleshoot problems to get it progressively closer
    to that desired state.
    * running 'make' will compile the code, placing any compiler messages
      into the 'errors' file (which you can view with cat/less/nano)
    * running 'make debug' directly shows you the compiler messages as it
      compiles
    * running 'make check' tries to show both your program's output along
      with that of the reference implementation. Unless there is a  level
      of randomness or other input opportunities of deviation, the output
      of your code should match that of the reference implementation.