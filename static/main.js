var p = document.getElementById("test")
const baseURL = "http://localhost:8000/"

submitButton = document.getElementById("submit")
submitButton.onclick = function() {
    var xhr = new XMLHttpRequest();
    var fullURL = document.getElementById("fullURL");
    var shortURL = document.getElementById("shortURL");
    var data = JSON.stringify({ "fullURL": fullURL.value, "shortcut": shortURL.value }); 
    xhr.open('POST', 'http://localhost:8000', true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status = 200) {
                var path = JSON.parse(xhr.responseText)
                p.innerHTML = ('Success! Here\'s your link: ' + '<a href="' + baseURL + path + '">' + baseURL + path + '</a>')
            } else {
                console.log("The server returned an error.")
            }
        }
      };
    console.log("Sending data: " + data)
    xhr.send(data)
}