var tf = document.getElementById("input")
var result = document.getElementById("result")
var btn = document.getElementById("button")

tf.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
        event.preventDefault();
        sendRequest(tf.value)
    }
});

btn.addEventListener("click",() => {
    console.log("clicked")
    sendRequest(tf.value)
})

function sendRequest(value) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url: value })
    };
    fetch('/shrink', requestOptions)
        .then(response => response.json())
        .then(data => appendChild(data));
}

function appendChild(data) {
    console.log(data)
    result.innerHTML = '';
    if (!data["is_success"]) {
        var p = document.createElement('p');
        p.textContent = data["err"];
        result.appendChild(p);
        result.style.display = "block"
    } else{
        var a = document.createElement('a');
        console.log(data["full_url"])
        a.href = data["hex_value"];
        a.textContent = data["full_url"];
        result.appendChild(a);
        result.style.display = "block"
    }
}