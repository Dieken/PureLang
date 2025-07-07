# PureLang - A pure language with C skin and Lisp heart

## Why

To prove we can achieve programmer friendly advanced language with simplest core syntax,
Purelang syntax can be filled in a postcard just like Smalltalk language.

1. No keywords but still familiar syntax of C language family.
2. Expression-based language, all statements are expressions.
3. Type is similar with normal data variable
4. Macro syntax is isomorphic to normal syntax.
5. Support customized operators like Scala language.
6. Language implementation is isomorphic with normal program.

Code example:

``` scala
std := import("std");               # import() is a function

Int := std.types.Integer;           # type is similar to var
Double := std.types.Double;
String := std.types.String;
Void := std.types.Void;


a: Int = 3 + 4;                     # explicit type declaration and assignment

a = `+`(5, 6);                      # assignment and explicit call of operator

a = if(a > 10, 100, -100);          # if() is a special function

a = if(a > 10,
       { a + a },
       { a * a });                  # block is also expression


b := std.Map(1 -> 2, 3 -> 4);       # type inference, and customized operator "->" constructs a pair


a = 1.negative() + "hello".length();    # function on literal

f := (i: Int = 10, j: Int = 20) Int { i + j };              # multi-method

f := (i: Double = 1.0, j: Double = 2.0) Double { i + j };    # multi-method

a = f();

a = f(10);

a = f(j = 30, i = 20);

a = f(100, 200);

d := f(5.0, 6.0);

f = `+`;                            # assignment clears all previous multi-method

# Product type, just like `record class` in Scala
User = std.types.ProductType(name: String = "Jack", age: Int = 20);

greet := (user: User = 1) Void { println(user.name + fmt(" with age {user.age}")) };

you = User(name = "You", age = 30); # type is callable as constructor
greet(you);
you.greet();                        # unified call syntax
```

## Build

``` sh
go install github.com/pointlander/peg@latest

go generate && go build
```

## Run

``` sh
./PureLang some.pure

```
