# What is ECMAScript
- Scripting language standard and spec
    - Javascript
    - Jscript
    - ActionScript
- ES6 is the most recent version of JavaScript, a very significant update.
- Also called ES2015
- On the subject of compatibility:
    - A lot of ways to go for some browsers
    - Transpilers such as `Babel` can be used to compile ES6 code to ES5

# What is JavaScript?
It is a programming language and has these characteristics.
### Weakly Typed
> No explicit type assignment to variables. 

> Dynamic changes to a variable's data type is possible
### Object Oriented 
> Data can be organized in logical objects

> Primitive and reference data types possible.
### Versatile
> Can run in browser or directly on a PC/Server

> A lot of tasks can be performed 
## Core Syntax
### Variables and Scopes
```js
// Declaring variables
var n = 'Name';
var age = 28;
// Funtions
function summarizeUser(userName, userAge){
    return 'Name is ' + userName + ', age is ' + userAge;
}
// Here variable scope must be kept in mind. Which means that the variable userName cannot be used outside "summarizeUser"
```
```js
// Another way of declaring variables is 
let age = 10;

// Continuing with the core JS 

// The scoping works a bit different her compared to var
const value = 'Will Never Change';
```
### Functions 
```js
// Anonymous Function
const summary = function (userName) {
    // We store this function inside the variable summary essentially making it a named function
    return 'Name = ' + userName;
}
// Arrow Function
const summaryArrow = (username) => return 'Name : ' + userName; 
const add = (a, b) => a+b;
const sub = (a, b) => {return a+b;}
const addOne = a => a+1;
const add12 = () => 1+2;
// These are some ways of using single statement arrow functions

// Below is how to use arrow function with multiple statements allowing us to remove the "function" keyword
const summary = (userName) => {
    username = 'Modded' + username;
    return 'Name = ' + username;
}

```
### Objects
```js
// Example of an object
const person = {
    name : 'Suteerth',
    age : 20,
    // Below wont work when .greet() is called because "this" refers to global run time scope when it is called on person.geet()
    greet : () =>  console.log('greetings from ' + this.name);
    // Aliter is
    greet: function () { 
        console.log('greetings from ' + this.name);
    },
    greet_2() { 
        // Preferred way
        console.log('2nd function');
    },
}
console.log(person);
person.greet();
// Output >> { name: 'Suteerth', age: 20 }
```
### Arrays
```js
// Declaring arrays
const arr = ['String', 45];
const nums = [1,2,3,4];
const strs = ['String A', 'String B', 'String C'];
// Iterating through the array
// Using for of
for(let num of nums){
    console.log(num);
}
// Traditional way
for (let index = 0; index < strs.length; index++) {
    console.log(strs[index]);
}
```
There are a lot in built methods applicable on the array manipulating them, getting a subset of them, going through them etc.

Map is used quite a lot to transform the values returning a new array 
```js
const nums = [1,2,3,4];
console.log(nums.map((num) => {return num + 1;}));
// Output >> [ 2, 3, 4, 5 ]
```
## Important
> **Objects** and **Arrays** are reference types.
>
> Which means that we can sort of edit the arrays after declaring them as const
>
> Reference types only store pointers to the objects/entity stored in the memory. Because the memory address is not changed when we "add" an element, it doesn't violate the const property.
```js
// Here is a demonstration
const arr = [1,2,3];
console.log(arr);
arr.push(4);
console.log(arr);
/* 
Output 
>> [ 1, 2, 3 ] 
>> [ 1, 2, 3, 4 ]
*/
```
## Spread and Rest Operator
### Spread Operator
In order to explain this, consider this snippet.
```js
const arr = [1,2,3];
arr = [4,...];
// This is also how you add a new element while creating a new array
// We get an eror 
TypeError: Assignment to constant variable.
```
Now, why did we get this error?
What we do here is, we add 4 and add copy of arr to arr (this time changing the object reference)

It will be clear with this example where I demonstrate methods of copying an array
```js
const arr = [1,2,3];
let copied_arr = arr.silce();
// This won't work 
let array_in_array = [arr]; // [ [1,2,3] ]
let spread_arr = [...arr]; // Creates a copy 
```
### Rest Operator
It is sort of opposite of spread operator. This will be clear with the snippets given below.
```js
const toArray = (arg1, arg2, arg3) => {
    return [arg1, arg2, arg3];
};
// This will work but only for 3 arguments, what if we want to add 4?
const toArrayRest =(...args) => {
    return args;
};
// This can work for n number of arguments
```
> It looks like the spread operator. The placement of the operator is what creates the difference. Keep that in mind.

## Destructuring
Now this allows us to be more efficient with our functions.
```js
const person = {
    name : 'Suteerth',
    age : 20,
    greet() {
        return 'I am ' + this.name;
    }
};
const printName = (person) => {
    console.log(person.name);
}
```
This works but the entire person object is sent to the function where we only want to work with the name property
```js
const printNameDestructured = ({name}) => {
    console.log(name);
}
```
What this does is, only the name property is retained and stored inside the name variable which can then be used instead of accessing the entire object.

It can happen in this way as well
```js
// The variable names need to be the same as property names
const {name, age} = person;
console.log(name, age); 
```
Our outpute will be name, age with their respective object values but here we just destructured the object for getting those values

Destructuring is not just limited to objects, arrays can also be destructured. Consider the snippet given below.
```js
const array = [1,2,3];
// Here, the variable name can be anything and based on position, the values are assigned
const [element1, element2] = array;
console.log(element1, element2);
// Going out of bound will simply keep these variables "undefined"
```
## String Literals
It's a different way of writing strings.

Instead of using double or single quotation marks:
```js
'A String'
```
or
```js
"Another string"
```
you can use backticks **( ` )**
```js
`Another way of writing strings`
```
With that syntax, you can dynamically add data into a string like this:
```js
const name = "Suteerth";
const age = 20;
console.log(`My name is ${name} and I am ${age} years old.`);
```
This is of course shorter and easier to read than the "old" way of concatenating strings.

## Async and Promises
Asynchronous code is a piece of code which doesn't execute immediately such as a timed log.
```js
// Asynchronous code
setTimeout(() => {
    console.log('Timer done!');
}, 2000);
// Synchronous code
console.log('Hello');
console.log('World');
```
Now 'Hello World' will be printed first and 'Timer Done' after some delay.

Code execution is NOT blocked until async code is executed.

Instead the callback function is recognized and be executed in the future. Move onto the next line and then execute the sync code.

Now dependencies can be cumbersome to deal with.

```js
const fetchData = callback => {
    setTimeout(() => {
        callback('Done!');
    }, 1500);
};
// callback is an argument which is a function which receives a text as ITS argument
setTimeout(() => {
    console.log('Time is done');
    fetchData(text => {
        console.log(text);
    })
    // The text will be the text passed by the callback 
}, 2000);
```
### Promises
The usage is mostly through packages but it can be done like this.
```js
const fetchData = () => {
    const promise = new Promise({resolve, reject} => {
        setTimeout(() => {
            resolve('Done!');
        }, 1500);
    }); 
    return promise;
};
setTimeout(() => {
    console.log('Time is done');
    // then is callable only on a promise
    fetchData().then(text => {
        console.log(text);
    });
}, 2000);
```
```js
// Nested callbacks
setTimeout(() => {
    console.log('Time is done');
    // then is callable only on a promise
    fetchData().then(text => {
        console.log(text);
        return fetchData();
        // Returning something in a thenblock converts it into a promise which is instantly resolved
    }).then(text2 => {
        console.log(text2);
    });
}, 2000);
```
```js
fetch(api).then((res) => res.json()).then((data) => console.log(data));
```
This is a bit complex and can be done like this
```js
async function getData(){
    const res = await fetch(api);
    const data = await res.json();
    console.log(data);
}
// JS waits only inside the async function but the rest of the code execution takes place 
```

# Optional Chaining
```js
function getTotalReviewCount(book){
    const goodReads = book.reviews.goodreads.reviewsCount;
    const librarything = book.reviews.librarything?.reviewsCount ?? 0; 
    // We wont try to read reviewsCount, in case 
}
```
## Filter method on array
```js
const longBooks = books.filter((book) => books.pages > 500);

const longBooks = books.filter((book) => books.pages > 500).filter((book) => books.hasMovieAdaptation);

const longBooks = books.filter((book) => books.genre.includes("adventure")).map((book) => book.title);
```

## Reduce method
It is called reduce because it reduces the array to a value.
```js
// The starter value is the init value of acc or accumulator 
// Every iteration, pages get added to acc for all the elements in book array
const pages = books.reduce((acc, book) => acc + book.pages , 0)
```

## Sort method
This mutilates the data
```js
const sorted = arr.sort((a,b) => a - b); // Ascending
const sorted = arr.sort((a,b) => b - a);
// a is first value, b is second
// If the result is negative, a < b, ascending order
// If the result is positive, a > b descending order
const sorted = arr.slice().sort((a,b) => b - a);

// Practical example
const sortByBooksDesc = books.slice().sort((a,b) => b.pages - a.pages);
```

# Add elements without changing original

```js
const newBook = {
    id:2,
    title:"Harry Potter",
    author:"JK Rowling",
}
// How to add a new object
const booksAfterAdd = [...books, newBook];
// How to delete
const booksAfterDelete = booksAfterAdd.filter((book) => book.id !== 3);
// How to update, change ONE property while keeping the rest same (the exmample here does THAT)
const booksAfterUpdate = booksAfterDelete.map((book) => book.id === 1 ? {...book, pages: 1210} : book);
```
# Some Extra Stuff
