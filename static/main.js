var p = document.getElementById("test")
p.innerText = "This is some text"

submitButton = document.getElementById("submit")
submitButton.onclick = function() {
    var xhr = new XMLHttpRequest();

    var fullURL = document.getElementById("fullURL");
    var shortURL = document.getElementById("shortURL");
    var data = JSON.stringify({ "fullURL": fullURL.value, "shortcut": shortURL.value }); 

    xhr.open('POST', 'http://localhost:8000', true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(data)

    console.log("Sending data: " + data)
}