
// hx-get="/datetime-option" hx-trigger="click"
//         hx-target="#date-time-option" hx-swap="outerHTML"
function getDataDateTime(){

        console.log('add class')
    var element = document.getElementById("date-time-option")
    element.innerHTML = "hello"
    element.classList.toggle('hidden')
    // htmx.ajax('GET','/datetime-option',{target:'#date-time-option',swap:'outerHTML'})
    
}

function getDefaultData(){
    htmx.ajax('GET','/option-title',{target:'#dropdown-datetime',swap:'outerHTML'})
}
function getConent(page){
    htmx.ajax('GET',`/data?component=${page}`,{target:'#content',swap:'outerHTML'})
}