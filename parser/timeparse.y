%{
package parser

import (
    tr "github.com/mz-eco/timerange"
    "time"
)
%}


%union {
    stmt int
    tm   DateTime
    date Date
    time Time
    cb func (int) tr.Interval
    interval tr.Interval
}

%token<stmt>     num
%token<stmt>     big
%token<cb>       interval
%token<stmt>     next
%token<stmt>     stmt

%type<date> date
%type<time> time
%type<tm>   datetime
%type<interval> offset
%type<stmt> size

%start range
%%

split: ' ' | 'T'

range: datetime ',' datetime          {yylex.(*Lex).tr = makeRange($1,$3)}
     | datetime ',' offset            {yylex.(*Lex).tr = tr.RangeTo(makeTime($1,time.Now()), $3)}
     | '[' datetime ',' datetime ')'  {yylex.(*Lex).tr = makeRange($2,$4)}



size: num  {$$ = $1}
    | stmt {$$ = $1}
    | big  {$$ = $1}

offset: '+' size interval { $$ = $3( 1*$2)}
      | '-' size interval { $$ = $3(-1*$2)}

datetime: date split time { $$ = NewDateTime($1             ,$3    )}
        | date            { $$ = NewDateTime($1             ,Time{})}
        | time            { $$ = NewDateTime(NewDate(-1,0,0),$1    )}



date: num                  {$$ = NewDate(-1, 0,$1)}
    | num  '-' num         {$$ = NewDate(-1,$1,$3)}
    | stmt '-' num '-' num {$$ = NewDate($1,$3,$5)}
    | stmt                 {$$ = NewDate($1, 0, 0)}


time: num ':' num         {$$ = NewTime($1,$3,0 )}
    | num ':' num ':' num {$$ = NewTime($1,$3,$5)}

%%

