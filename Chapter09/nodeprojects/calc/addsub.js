
function add(i,j){
    return i+j;
}

function sub(i,j){
    return i-j;
}

function formatnumbers(Obj){
    return "First number: " + Obj.first + " second number: " + Obj.second;
}

module.exports={
    Add: add,
    Sub: sub,
    FormatNumbers: formatnumbers
}