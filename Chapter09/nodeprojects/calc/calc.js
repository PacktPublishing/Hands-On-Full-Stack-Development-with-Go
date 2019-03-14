var calc = require('./addsubgo.js');

//Call Add() then save result in the add variable
var add = calc.Add(2,3);

//Call Sub() then save result in the sub variable
var sub = calc.Sub(5,2);

//Call FormatWords then save the result in the fw variable
var fw = calc.FormatNumbers({
    first: 10,
    second: 20,
});

console.log(add);
console.log(sub);
console.log(fw);

